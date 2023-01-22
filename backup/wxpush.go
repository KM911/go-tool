package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//	{
//	  "appToken":"AT_xxx",
//	  "content":"Wxpusher祝你中秋节快乐!",
//	  "summary":"消息摘要",//消息摘要，显示在微信聊天页面或者模版消息卡片上，限制长度100，可以不传，不传默认截取content前面的内容。
//	  "contentType":1,//内容类型 1表示文字  2表示html(只发送body标签内部的数据即可，不包括body标签) 3表示markdown
//	  "topicIds":[ //发送目标的topicId，是一个数组！！！，也就是群发，使用uids单发的时候， 可以不传。
//	      123
//	  ],
//	  "uids":[//发送目标的UID，是一个数组。注意uids和topicIds可以同时填写，也可以只填写一个。
//	      "UID_xxxx"
//	  ],
//	  "url":"https://wxpusher.zjiecode.com", //原文链接，可选参数
//	  "verifyPay":false //是否验证订阅时间，true表示只推送给付费订阅用户，false表示推送的时候，不验证付费，不验证用户订阅到期时间，用户订阅过期了，也能收到。
//	}
type Message struct {
	AppToken    string   `json:"appToken"`
	Content     string   `json:"content"`
	Summary     string   `json:"summary"`
	ContentType int      `json:"contentType"`
	TopicIds    []int    `json:"topicIds"`
	Uids        []string `json:"uids"`
	Url         string   `json:"url"`
	VerifyPay   bool     `json:"verifyPay"`
}

func SendMessage(ms string) {
	// 这里我只是为了给我自己使用 所以我不需要其他的参数了

	// 创建我们的json对象
	message := Message{"AT_bfuM6nXGqseMPQZXkaC9r7ce4jLUAY1S", ms, "特别通知", 1, []int{123}, []string{"UID_YqkKAyhXfS1Om9vRb4AwGq7qzfB9"}, "http://81.68.91.70/", false}
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}
	res, err := http.Post("https://wxpusher.zjiecode.com/api/send/message", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
}
func main() {
	SendMessage("服务器挂了 请立即查看")
}
