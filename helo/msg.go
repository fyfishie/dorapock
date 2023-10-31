/*
 * @Author: fyfishie
 * @Date: 2023-05-19:11
 * @LastEditors: fyfishie
 * @LastEditTime: 2023-05-21:19
 * @Description: :)
 * @email: muren.zhuang@outlook.com
 */
package helo

import "encoding/json"

type Message struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

func MakeMsg(status bool, msg string) Message {
	return Message{Status: status, Msg: msg}
}

func MakeMsgBytes(status bool, msg string) []byte {
	bs, _ := json.Marshal(Message{Status: status, Msg: msg})
	return bs
}

func ParseMsg(bs []byte) (Message, error) {
	m := Message{}
	err := json.Unmarshal(bs, &m)
	return m, err
}

func ParseMsgNoErr(bs []byte) Message {
	m := Message{}
	json.Unmarshal(bs, &m)
	return m
}
