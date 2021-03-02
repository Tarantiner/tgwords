package main

import (
	"fmt"
	"sort"
	"sync"
)

func handle(lis *[]string, mp sync.Map, li[]int) {
	for _, v := range li{
		mp.Range(func(key, value interface{}) bool {
			if value.(int) == v {
				*lis = append(*lis, key.(string))
				mp.Delete(key)
				return true
			}else{
				return false
			}
		})
	}
}

func main() {
	var mp sync.Map
	mp.Store("d", 55)
	mp.Store("a", 18)
	mp.Store("e", 23)
	mp.Store("b", 15)
	mp.Store("c", 23)

	var wordLis []string
	var valueLis []int
	m := make(map[int][]string)
	mp.Range(func(key, value interface{}) bool {
		vv := value.(int)
		if v,ok := m[vv]; ok{
			v = append(v, key.(string))
			m[vv] = v
		}else{
			v = []string{key.(string)}
			m[vv] = v
		}
		valueLis = append(valueLis, vv)
		return true
	})

	// 将值从大到小排序
	sort.Slice(valueLis, func(i, j int) bool {
		return valueLis[i] > valueLis[j]
	})


	fmt.Println(m)
	fmt.Println(wordLis)
	fmt.Println(valueLis)

	for _, v := range valueLis {
		fmt.Println(v)
	}

	// b, a, e, c, d
	//mp.Range(func(key, value interface{}) bool {
	//	fmt.Println(key, value)
	//	return true
	//})
	//fmt.Println(strings.Repeat("#", 50))
	//mp.Range(func(key, value interface{}) bool {
	//	if value.(int) < 24{
	//		mp.Delete(key)
	//	}
	//	return true
	//})
	//mp.Range(func(key, value interface{}) bool {
	//	fmt.Println(key, value)
	//	return true
	//})

	//li := []int{55, 23, 23}
	//lis := make([]string, 0, 3)
	//handle(&lis, mp, li)
	//mp.Range(func(key, value interface{}) bool {
	//	fmt.Println(key, value)
	//	return true
	//})
	//fmt.Println(lis)
}
