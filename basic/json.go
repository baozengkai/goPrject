package main

import (
	"encoding/json"
	"fmt"
)

type jsonStruct struct {
	Number int    `json:"number"`
	Phone  string `json:"phone"`
}

func main() {
	// 1.结构体转换为json字符串
	var jStruct jsonStruct
	jStruct.Number = 1
	jStruct.Phone = "17702137349"
	fmt.Println(jStruct)

	jsonByte, _ := json.Marshal(jStruct)
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)

	// 2.json字符串转换为结构体
	var jsonStr2 string = `{"Number": 2, "phone": "123456"}`
	var jStruct2 jsonStruct
	err := json.Unmarshal([]byte(jsonStr2), &jStruct2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jStruct2.Number)
}
