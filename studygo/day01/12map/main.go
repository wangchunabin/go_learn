package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 80

	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	//判断map中某个键值是否存在
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	//map遍历
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//若只想遍历key
	for k := range scoreMap {
		fmt.Println(k)
	}

	//使用delete删除key
	delete(scoreMap, "小明")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	userInfo := map[string]string{
		"username": "wangdachui",
		"password": "123456",
	}
	fmt.Println(userInfo)

	fmt.Println("========================================")
	//按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano()) //初始化随机种子
	var scoreMap1 = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap1[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap1 {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap1[key])
	}

	fmt.Println("========================================")
}
