//钉钉机器人
package dingdingbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

type DingDingBot struct {
	WebhookUrl string
	Token      string
	Name       string
}

//发送消息返回对象
type DingDingReturnMsg struct {
	IsSuccess bool
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
}

//发送消息输入参数对象
//https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.0UsMvA&treeId=257&articleId=105735&docType=1#s2
type DingDingMessage interface {
	//获取map
	GetValues() []byte
}

//文本消息
//https://open-doc.dingtalk.com/docs/api.htm?spm=a219a.7386797.0.0.UaCPAW&source=search&apiId=37246
type DingDingTextMessage struct {
	Text      string   //文本
	AtMobiles []string //@用户的手机号
	isAtAll   bool     //是否@所有人
}

//Link类型消息
//https://open-doc.dingtalk.com/docs/api.htm?spm=a219a.7386797.0.0.UaCPAW&source=search&apiId=37246
type DingDingLinkMessage struct {
	Title      string //标题
	Text       string //文本
	PicUrl     string //图片地址
	MessageUrl string //消息地址
}

//MarkDown 文本消息
//https://open-doc.dingtalk.com/docs/api.htm?spm=a219a.7386797.0.0.UaCPAW&source=search&apiId=37246
type DingDingMarkDownMessage struct {
	Title     string   //标题
	Text      string   //文本
	AtMobiles []string //@用户的手机号
	isAtAll   bool     //是否@所有人
}

//ActionCard消息btn参数
type DingDingActionCardBtnType struct {
	Title     string //按钮标题
	ActionURL string //点击按钮的操作
}

//整体跳转ActionCard类型
//https://open-doc.dingtalk.com/docs/api.htm?spm=a219a.7386797.0.0.UaCPAW&source=search&apiId=37246
type DingDingActionCardMessage struct {
	Title          string                      //标题
	Text           string                      //文本,markdown格式的消息
	HideAvatar     int                         //0-正常发消息者头像,1-隐藏发消息者头像
	BtnOrientation int                         //0-按钮竖直排列，1-按钮横向排列
	SingleTitle    string                      //单个按钮的方案。(设置此项和singleURL后btns无效。)
	SingleURL      string                      //点击singleTitle按钮触发的URL
	Btns           []DingDingActionCardBtnType //按钮的信息：title-按钮方案，actionURL-点击按钮触发的URL
}

//FeedCard类型Item
type DingDingFeedCardItem struct {
	Title      string //标题
	MessageURL string //
	PicURL     string //点击按钮的操作
}

//FeedCard类型
type DingDingFeedCardMessage struct {
	Links []DingDingFeedCardItem
}

//文本消息获取参数
func (msg *DingDingTextMessage) GetValues() []byte {
	var values = make(map[string]interface{}, 3)
	var textMap = make(map[string]string, 1)
	textMap["content"] = msg.Text
	var atMap = make(map[string]interface{}, 2)
	if len(msg.AtMobiles) > 0 {
		var at = ""
		for i, v := range msg.AtMobiles {
			at += v
			if i < len(msg.AtMobiles) {
				at += ","
			}
		}
		atMap["atMobiles"] = at
	}
	var isAtAll = false
	if msg.isAtAll {
		isAtAll = true
	}
	atMap["isAtAll"] = isAtAll
	values["msgtype"] = "text"
	values["text"] = textMap
	values["at"] = atMap
	v, _ := json.Marshal(values)
	return v
}

//Link消息获取参数
func (msg *DingDingLinkMessage) GetValues() []byte {
	var values = make(map[string]interface{}, 2)
	var linkMap = make(map[string]string, 4)
	linkMap["text"] = msg.Text
	linkMap["title"] = msg.Title
	linkMap["picUrl"] = msg.PicUrl
	linkMap["messageUrl"] = msg.MessageUrl
	values["msgtype"] = "link"
	values["link"] = linkMap
	v, _ := json.Marshal(values)
	return v
}

//MarkDown消息获取参数
func (msg *DingDingMarkDownMessage) GetValues() []byte {
	var values = make(map[string]interface{}, 3)
	var markdownMap = make(map[string]string, 2)
	markdownMap["text"] = msg.Text
	markdownMap["title"] = msg.Title
	var atMap = make(map[string]interface{}, 2)
	if len(msg.AtMobiles) > 0 {
		var at = ""
		for i, v := range msg.AtMobiles {
			at += v
			if i < len(msg.AtMobiles) {
				at += ","
			}
		}
		atMap["atMobiles"] = at
	}
	var isAtAll = false
	if msg.isAtAll {
		isAtAll = true
	}
	atMap["isAtAll"] = isAtAll
	values["msgtype"] = "markdown"
	values["markdown"] = markdownMap
	values["at"] = atMap
	v, _ := json.Marshal(values)
	return v
}

//ActionCard消息获取参数
func (msg *DingDingActionCardMessage) GetValues() []byte {
	var values = make(map[string]interface{}, 2)
	var actionCardMap = make(map[string]interface{}, 6)
	actionCardMap["text"] = msg.Text
	actionCardMap["title"] = msg.Title
	actionCardMap["hideAvatar"] = msg.HideAvatar
	actionCardMap["btnOrientation"] = msg.BtnOrientation
	if msg.SingleTitle != "" && msg.SingleURL != "" {
		actionCardMap["singleTitle"] = msg.SingleTitle
		actionCardMap["singleURL"] = msg.SingleURL
	} else if msg.Btns != nil && len(msg.Btns) > 0 {
		var btns []map[string]string
		for _, v := range msg.Btns {
			var btn = make(map[string]string, 2)
			btn["title"] = v.Title
			btn["actionURL"] = v.ActionURL
			btns = append(btns, btn)
		}
		actionCardMap["btns"] = btns
	}
	values["msgtype"] = "actionCard"
	values["actionCard"] = actionCardMap
	v, _ := json.Marshal(values)
	return v
}

//FeedCard消息获取参数
func (msg *DingDingFeedCardMessage) GetValues() []byte {
	var values = make(map[string]interface{}, 2)
	var feedCardMap = make(map[string]interface{}, 1)
	var linkArray []map[string]string
	for _, v := range msg.Links {
		var linkMap = make(map[string]string, 3)
		linkMap["title"] = v.Title
		linkMap["messageURL"] = v.MessageURL
		linkMap["picURL"] = v.PicURL
		linkArray = append(linkArray, linkMap)
	}
	feedCardMap["links"] = linkArray
	values["msgtype"] = "feedCard"
	values["feedCard"] = feedCardMap
	v, _ := json.Marshal(values)
	return v
}

//发送消息
func (bot *DingDingBot) Send(data DingDingMessage) (*DingDingReturnMsg, error) {
	var values io.Reader
	var resp *http.Response
	var err error
	if data != nil {
		values = bytes.NewReader(data.GetValues())
	}
	if resp, err = http.Post(bot.WebhookUrl, "application/json", values); err != nil {
		return nil, err
	}
	var byteBody []byte
	if byteBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		var ddReturnMsg DingDingReturnMsg
		if err := json.Unmarshal(byteBody, &ddReturnMsg); err != nil {
			return nil, err
		} else {
			ddReturnMsg.IsSuccess = ddReturnMsg.Errcode == 0
			return &ddReturnMsg, nil
		}
	}
	return nil, errors.New("http response status: " + resp.Status + " body:" + string(byteBody))
}
