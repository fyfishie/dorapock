/*
 * @Author: fyfishie
 * @Date: 2023-05-21:19
 * @LastEditors: fyfishie
 * @LastEditTime: 2023-05-21:19
 * @Description: :)
 * @email: muren.zhuang@outlook.com
 */
package helo

import "encoding/json"

type MessageWithErrorCode struct {
	Status    bool   `json:"status"`
	Msg       string `json:"msg"`
	ErrorCode int    `json:"error_code"`
}

func MakeMsgWCBytes(status bool, msg string, errorCode int) []byte {
	bs, _ := json.Marshal(MessageWithErrorCode{Status: status, Msg: msg, ErrorCode: errorCode})
	return bs
}

func ParseMsgWC(bs []byte) (MessageWithErrorCode, error) {
	m := MessageWithErrorCode{}
	err := json.Unmarshal(bs, &m)
	return m, err
}
func ParseMsgWCNoErr(bs []byte) MessageWithErrorCode {
	m := MessageWithErrorCode{}
	json.Unmarshal(bs, &m)
	return m
}
