package dingdingbot

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
)

//发消息测试
func TestSend(t *testing.T) {
	var url = "1https://oapi.dingtalk.com/robot/send?access_token=468470912f6c1241b1d804292794523940a9763eed0b04bf379469ee2163a103"
	ddBot := &DingDingBot{url, url, "ddBot"}
	//一般文本消息
	//a := &DingDingTextMessage{Text: "hello"}
	//链接消息
	//a := &DingDingLinkMessage{"Hello Title", "Hello", "", "http://www.20dot.com"}
	//markdown消息
	//a := &DingDingMarkDownMessage{"杭州天气", "#### 杭州天气  \n > 9度，@1825718XXXX 西北风1级，空气良89，相对温度73%\n\n > ![screenshot](http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png)\n  > ###### 10点20分发布 [天气](http://www.thinkpage.cn/)", []string{"18500153794"}, true}
	//ActionCard跳转消息
	// a := &DingDingActionCardMessage{"乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
	// 	"![screenshot](@lADOpwk3K80C0M0FoA) \n\n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
	// 	1, 1, "阅读全文", "https://20dot.com", nil}
	// a := &DingDingActionCardMessage{"乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
	// 	"![screenshot](@lADOpwk3K80C0M0FoA) \n\n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
	// 	1, 0, "", "",
	// []DingDingActionCardBtnType{DingDingActionCardBtnType{"赞成", "https://baidu.com"}, DingDingActionCardBtnType{"反对", "https://baidu.com"}}}
	//FeedCard类型消息
	a := &DingDingFeedCardMessage{
		[]DingDingFeedCardItem{
			DingDingFeedCardItem{"时代的火车向前开", "https://mp.weixin.qq.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI", "https://www.dingtalk.com/"},
			DingDingFeedCardItem{"时代的火车向前开2", "https://mp.weixin.qq.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI", "https://www.dingtalk.com/"}}}
	returnMsg, err := ddBot.Send(a)
	log.Println("结果:", returnMsg)
	log.Println("err:", err)
}

func testS(t *testing.T) {
	var data = `{"errmsg":"param error","errcode":300001}`
	var e DingDingReturnMsg
	json.Unmarshal([]byte(data), &e)
	log.Printf("%+v", e)
}

func testSend(t *testing.T) {
	fmt.Println(reflect.TypeOf(DingDingTextMessage{}).String())
}
