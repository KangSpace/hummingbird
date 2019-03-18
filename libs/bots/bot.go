// author: kango2gler@gmail.com
// date: 2018-06-12 23:53:36

// 机器人超类
//
// func:
//
package bots

// Bot 对象
type Bot interface {
	Send(interface{}) interface{}
}

//Bot 发送消息成功返回对象
type BotResponseMsg struct {
	result string
	msg    string
	data   string
}
