package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"lib"
	"net/http"
)

var WxPush = &cli.Command{
	Name:    "WxPush",
	Aliases: []string{"wx"},
	Usage:   "send a message to wx",
	Action:  WxPushAction,
}

func WxPushAction(c *cli.Context) error {
	if len(c.Args().Slice()) == 0 {
		println("no argv run go mod tidy")
		lib.Run("go mod tidy")
		return nil
	} else {
		// connect argvs to string
		message := ""
		for _, value := range c.Args().Slice() {
			message += value + " "
		}
		SendMessage(message)
	}

	return nil
}

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
	println("send message to wx success")
	defer res.Body.Close()
}
