package plugin

import (
	"MCDaemon-go/command"
	"MCDaemon-go/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

type SDChatPlugin struct{}

func (hp *SDChatPlugin) Handle(c *command.Command, s lib.Server) {
	text := "!!SDChat all start 开启全局聊天模式\\n!!SDChat start 开启私聊模式（别的玩家看不见沙雕机器人给你发的信息）\\n!!SDChat stop 关闭聊天模式"
	if len(c.Argv) < 1 {
		s.Tell(text, c.Player)
		return
	}
	switch c.Argv[0] {
	case "all":
		command.Group.AddPlayer("SDChat-all", c.Player)
		s.Tell("开启全局聊天模式成功", c.Player)
	case "start":
		command.Group.AddPlayer("SDChat", c.Player)
		s.Tell("开启全局聊天模式成功", c.Player)
	case "stop":
		command.Group.DelPlayer("SDChat", c.Player)
		command.Group.DelPlayer("SDChat-all", c.Player)
		s.Tell("退出聊天模式成功", c.Player)
	case "say":
		s.Tell(chat(c.Argv[1], c.Player), c.Player)
	case "say-all":
		s.Say(chat(c.Argv[1], c.Player))
	default:
		s.Tell(text, c.Player)
	}
}

//封装JSON
func LightEncode(elememt interface{}) string {
	//拼接的结果字符串
	var s string
	//若为对象，则拼接字符串
	if LightJson, err := elememt.(map[string]interface{}); !err {
		s = string("{")
		for key, val := range LightJson {
			s += fmt.Sprintf("\"%s\":\"%s\",", key, LightEncode(val))
		}
		s += string("}")
	} else {
		jsonStr, err := json.Marshal(elememt)
		if err != nil {
			log.Fatal("Can't transform jsonString,Because ", err)
		}
		s = string(jsonStr)
	}
	return s
}

//http POST请求
func httpPost(data string) string {
	resp, err := http.Post("http://openapi.tuling123.com/openapi/api/v2",
		"application/x-www-form-urlencoded",
		strings.NewReader(data),
	)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}

//发出请求获取聊天回复
func chat(data string, player string) string {
	_requestMap := map[string]interface{}{
		"perception": map[string]interface{}{
			"inputText": map[string]interface{}{
				"text": data,
			},
		},
		"userInfo": map[string]interface{}{
			"apiKey":     "b0891402dce941e48394a6090e304b51",
			"userId":     player,
			"groupId":    10,
			"userIdName": player,
		},
	}
	return gjson.Get(httpPost(LightEncode(_requestMap)), "results.0.values.text").String()
}
