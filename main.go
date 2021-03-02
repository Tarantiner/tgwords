package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v3"
	"log"
	"os"
	"strings"
	"time"
)

type Info struct {
	MsgCnt float64 `json:"msg_cnt"`
	UID    string  `json:"uid"`
}

var maxWorker int
var test_uid string
var client *elastic.Client
var wordLis []string

func init() {
	flag.IntVar(&maxWorker, "w", 10, "处理每个群组分词的go协程数量")
	flag.StringVar(&test_uid, "test_uid", "", "测试模式下指定需要测试的群组ID，测试模式不会更新ES")
	flag.Parse()
	var err error
	// 加载词库
	f, err := os.Open("sense.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dd := json.NewDecoder(f)
	err = dd.Decode(&wordLis)
	if err != nil {
		panic(err)
	}

	// 初始化es
	client, err = elastic.NewClient(elastic.SetURL("http://172.16.100.6:9200", "http://172.16.100.7:9200", "http://172.16.100.8:9200"))
	if err != nil {
		panic(err)
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
	for _, groupInfo := range groupList {
		fmt.Printf("%s群组%s共%d\n", groupInfo.EsID, groupInfo.UID, groupInfo.MsgCNT)
		t := time.Now()
		handleGroup(groupInfo)
		fmt.Printf("%s|%s|%d|花费%f秒\n", groupInfo.EsID, groupInfo.UID, groupInfo.MsgCNT, time.Since(t).Seconds())
		fmt.Println(strings.Repeat("#", 80))
		if test_uid != "" {
			break
		}
	}
}
