package main

import (
	"encoding/json"
	"fmt"
)

type IsRcContent interface {
	isRcContentData()
}

type RcText struct {
	Content string `json:"content"`
}

func (*RcText) isRcContentData() {}

type RcMsgContent struct {
	IsRcContent `json:",ASN.1"`
	Extra       RcMsgExtra `json:"extra"`
	User        RcMsgUser  `json:"user"`
}
type RcMsgExtra struct {
	MsgId uint64 `json:"msgId"`
}

type RcMsgUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func main() {
	msg := &RcMsgContent{
		IsRcContent: &RcText{
			Content: "haha",
		},
		User: RcMsgUser{
			Id: "1",
		},
	}

	jsonStr, err := json.Marshal(msg)

	fmt.Printf("%s,%v", jsonStr, err)
}
