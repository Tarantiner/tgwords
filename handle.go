package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"gopkg.in/olivere/elastic.v3"
	"io"
	"log"
	"sort"
	"strings"
	"sync"
)

func handle(mp *sync.Map, msgString string) {
	// 每个go协程对负责聊天内容统计词频
	for _, word := range wordLis {
		count := strings.Count(msgString, word)
		if count <= 0 {
			continue
		}
		// 若属于特别敏感词库，则频数翻倍
		if _, ok := specialWords[word]; ok {
			count*=2
		}
		if n, ok := (*mp).Load(word); ok {
			if n, ok := n.(int); ok {
				(*mp).Store(word, n+count)
			}
		} else {
			(*mp).Store(word, count)
		}
	}
}

type ChatMsg struct {
	Msg string `json:"msg"`
}

func handleGroup(info *MetaGroup) {
	// 处理每个群组的聊天信息
	// 每1000条聊天信息，分发一个go处理，最多x个go同时处理
	msgChan := make(chan string, 1000)
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		defer close(msgChan)
		fsc := elastic.NewFetchSourceContext(true).Include("msg") //.Exclude("*.description")
		q := elastic.NewBoolQuery()
		q.Must(elastic.NewTermQuery("to", info.UID))
		q.Must(elastic.NewTermQuery("msgtype", "text"))
		scroll := client.Scroll("app_chatinfo_01", "app_chatinfo_02", "app_chatinfo_03").Type("info").Query(q).FetchSourceContext(fsc).Size(1000)
		for i := 0; i < 1000; i++ { // 表示最多解析100万条聊天记录
			results, err := scroll.Do()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

			var s string
			for _, hit := range results.Hits.Hits {
				var cm ChatMsg
				err := json.Unmarshal(*hit.Source, &cm)
				if err != nil {
					continue
				}
				s += cm.Msg
			}
			select {
			case msgChan <- s:
			case <-ctx.Done(): // 应该是在第二个协程报错时传递过来的吧
				return ctx.Err()
			}
		}
		return nil
	})

	var mp sync.Map

	for i := 0; i < maxWorker; i++ {
		g.Go(func() error {
			for msgString := range msgChan {
				handle(&mp, msgString)
				select {
				default:
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Println("处理群组聊天信息失败", err)
		return
	}

	// 搜索关键词和标签
	Top100 := make([]string, 0, 100) // 存放前100高频词
	Top20 := make([]string, 0, 20)   // 存放前20高频词
	var valueLis []int               // 存放词频
	m := make(map[int][]string, 80)  // 预开辟长度为80
	mp.Range(func(key, value interface{}) bool {
		vv := value.(int)
		if v, ok := m[vv]; ok {
			v = append(v, key.(string))
			m[vv] = v
		} else {
			v = []string{key.(string)}
			m[vv] = v
		}
		valueLis = append(valueLis, vv)
		return true
	})

	// 对词频表进行从大到小排序
	sort.Slice(valueLis, func(i, j int) bool {
		return valueLis[i] > valueLis[j]
	})

	// 得出前100和前20高频词
	current := 0
	for _, v := range valueLis {
		if v == current { // 去重
			continue
		}
		if words, ok := m[v]; ok {
			current = v
			if len(Top100)+len(words) > 100 {
				Top100 = append(Top100, words[:100-len(Top100)]...)
				break
			} else {
				Top100 = append(Top100, words...)
				if len(Top20) == 20 {
					continue
				} else if len(Top20)+len(words) > 20 {
					Top20 = append(Top20, words[:20-len(Top20)]...)
				} else {
					Top20 = append(Top20, words...)
				}
			}
		}
	}
	fmt.Println("搜索关键词")
	fmt.Println(Top100)
	fmt.Println("label标签")
	fmt.Println(Top20)

	if test_uid != "" {
		return
	}

	// 更新ES群组搜索关键词和标签
	_, err := client.Update().Index("group_channel_entity").Type("info").Id(info.EsID).Doc(map[string]interface{}{"search_keywords": strings.Join(Top100, ","), "label": strings.Join(Top20, ",")}).Do()
	if err != nil {
		log.Println("更新es失败", err)
	}
}
