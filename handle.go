package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"gopkg.in/olivere/elastic.v3"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
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

func updateTag(tagLis []string, info *MetaGroup) {
	if len(tagLis) == 0 {
		return
	}
	dataLis := make([]string, 0, len(tagLis))
	tm := strconv.Itoa(int(time.Now().Unix()))
	for _, wd := range tagLis {
		dd := map[string]string{
			"uid":       info.UID,
			"utype":     info.Ctype,
			"keyword":   wd,
			"storetime": tm,
		}
		_b, err := json.Marshal(dd)
		if err != nil {
			continue
		}
		dataLis = append(dataLis, string(_b))
	}

	t := map[string][]string{"tags": dataLis}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println("失败", err)
		return
	}
	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(b))
	if err != nil {
		log.Println("构造请求失败", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = c.Do(req)
	if err != nil {
		log.Println("请求失败", err)
		return
	}
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

	// 关键词和标签
	var valueLis []int               // 存放词频  [100, 100, 98...]
	m := make(map[int][]string, 80)  // 预开辟长度为80 {100: ["xx", "yy"], 98: []}
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

	// 对词频表进行从大到小排序 [105, 88, 88, 76......]
	sort.Slice(valueLis, func(i, j int) bool {
		return valueLis[i] > valueLis[j]
	})

	current := 0
	mm := make(map[string]int, 100)
	tagLis := make([]string, 0, 30)
	for _, v := range valueLis {
		if v == current { // 去重
			continue
		}
		if words, ok := m[v]; ok {
			current = v
			for _, wd := range words {
				if _, ok = specialWords[wd]; ok{
					tagLis = append(tagLis, wd)
				}
				if len(mm) < 100 {
					mm[wd] = v
				}
			}
		}
	}
	b, _ := json.Marshal(mm)
	fmt.Println("词频")
	fmt.Println(len(mm), string(b))
	fmt.Println("标签")
	fmt.Println(len(tagLis), tagLis)

	if test_uid != "" {
		return
	}

	// 更新ES群组关键词
	_, err := client.Update().Index("group_channel_entity").Type("info").Id(info.EsID).
		Doc(map[string]interface{}{"keywords": string(b)}).Do()
	client.Index().Index("tags").BodyJson(map[string]string{"name": "chen"})
	if err != nil {
		log.Println("更新es关键词失败", err)
	}

	// 更新ES tag
	updateTag(tagLis, info)
}
