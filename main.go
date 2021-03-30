package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v3"
	"log"
	"net/http"
	"strings"
	"time"
)

type Info struct {
	MsgCnt float64 `json:"msg_cnt"`
	UID    string  `json:"uid"`
}

var maxWorker int
var minCount int64
var limit int64
var test_uid string
var redo bool
var client *elastic.Client
var wordLis = make([]string, 0, 120000)
var specialWords = make(map[string]int, 1000)  // map类型方便快速检索
var c *http.Client
var postURL = "https://172.16.100.8:11570/special/tag/upsert/data"
// 1244136556  1493776399  1119449737  1271536028  1492155360

func init() {
	// 解析命令参数
	flag.IntVar(&maxWorker, "worker", 8, "处理每个群组分词的go协程数量")
	flag.StringVar(&test_uid, "testUID", "", "测试模式下指定需要测试的群组ID，测试模式不会更新ES")
	flag.Int64Var(&minCount, "minCount", -1, "最少聊天数")
	flag.Int64Var(&limit, "limit", 100000, "聊天数区间")
	flag.BoolVar(&redo, "redo", false, "是否爬取已存在keywords的群组")
	flag.Parse()

	// 加载词库
	loadWords()

	// 初始化es
	var err error
	client, err = elastic.NewClient(elastic.SetURL("http://172.16.100.6:9200", "http://172.16.100.7:9200", "http://172.16.100.8:9200"))
	if err != nil {
		panic(err)
	}

	// 初始化跳过认证的客户端连接
	c = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

func main() {
	groupList, err := getGroups()
	if err != nil {
		log.Fatalln("获取群组ID失败", err)
	}
	if groupList == nil {
		log.Println("无群组ID")
		return
	}
	fmt.Printf("一共%d群组\n", len(groupList))
	for _, groupInfo := range groupList {
		fmt.Printf("%s群组%s共%d\n", groupInfo.EsID, groupInfo.UID, groupInfo.MsgCNT)
		t := time.Now()
		handleGroup(groupInfo)
		fmt.Printf("%s|%s|%d|花费%f秒\n", groupInfo.EsID, groupInfo.UID, groupInfo.MsgCNT, time.Since(t).Seconds())
		fmt.Println(strings.Repeat("#", 100))
		if test_uid != "" {
			break
		}
	}
}
