package main

import (
	"encoding/json"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func loadSpecialWords() {
	db, err := sql.Open("mysql", "root:1qaz2wsx@tcp(172.16.100.8:3306)/yingguang?parseTime=true&charset=utf8")
	if err != nil {
		panic(fmt.Errorf("无法连接到mysql:%v", err))
	}
	defer db.Close()
	ret, err := db.Query("select tag from label")
	if err != nil {
		panic(fmt.Errorf("执行sql失败:%v", err))
	}
	defer ret.Close()
	for ret.Next() {
		var tag string
		err = ret.Scan(&tag)
		if err != nil {
			continue
		}
		if tag != "" {
			// 添加到特别敏感词库
			specialWords[tag] = 0
		}
		// 如果敏感词库没有收录，则添加到敏感词库
		var F bool
		for _, word := range wordLis {
			if word == tag {
				F = true
			}
		}
		if !F{
			wordLis = append(wordLis, tag)
		}
	}
}

func loadWords() {
	// 加载词库
	f, err := os.Open("sense.json")
	if err != nil {
		panic(fmt.Errorf("打开词库文件失败:%v", err))
	}
	defer f.Close()
	dd := json.NewDecoder(f)
	err = dd.Decode(&wordLis)
	if err != nil {
		panic(fmt.Errorf("解析词库文件失败:%v", err))
	}
	loadSpecialWords()
}
