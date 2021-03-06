package main

import (
	"encoding/json"
	"gopkg.in/olivere/elastic.v3"
	"io"
)

type MetaGroup struct {
	UID    string `json:"uid"`
	MsgCNT int    `json:"msg_cnt"`
	EsID   string
	Keywords  string `json:"keywords"`
	Ctype string `json:"comm_type"`
}

// groupIDList []string
func getGroups() (groupIDList []*MetaGroup, err error) {
	fsc := elastic.NewFetchSourceContext(true).Include("uid", "msg_cnt", "keywords", "comm_type") //.Exclude("*.description")
	var scroll *elastic.ScrollService
	if test_uid != "" {
		q := elastic.NewTermQuery("uid", test_uid)
		scroll = client.Scroll("group_channel_entity").Type("info").Query(q).FetchSourceContext(fsc).Size(1)
	} else {
		if minCount == -1 {
			scroll = client.Scroll("group_channel_entity").Type("info").FetchSourceContext(fsc).Sort("msg_cnt", true).Size(500)
		} else {
			q := elastic.NewRangeQuery("msg_cnt").Gt(minCount).Lte(minCount + limit)
			scroll = client.Scroll("group_channel_entity").Type("info").Query(q).FetchSourceContext(fsc).Sort("msg_cnt", true).Size(500)
		}
	}
	for {
		results, err := scroll.Do()
		if err == io.EOF {
			return groupIDList, nil // all results retrieved
		}
		if err != nil {
			return groupIDList, err // something went wrong
		}
		for _, hit := range results.Hits.Hits {
			var info MetaGroup
			err = json.Unmarshal(*hit.Source, &info)
			if err != nil {
				continue
			}
			if hit.Id != "" {
				if info.Keywords == "" || test_uid != "" || redo {
					info.EsID = hit.Id
					groupIDList = append(groupIDList, &info)
				}
			}
		}
	}
}
