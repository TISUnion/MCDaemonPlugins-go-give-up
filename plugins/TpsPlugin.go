package plugin

import (
	"MCDaemon-go/command"
	"MCDaemon-go/lib"
	"strconv"
	"strings"
	"time"
)

type TpsPlugin string

func (hp TpsPlugin) Handle(c *command.Command, s lib.Server) {
	if second, ok := strconv.Atoi(c.Argv[0]); ok == nil {
		if second > 30 {
			second = 30
		}
		s.Execute("debug start")
		time.Sleep(time.Second * time.Duration(second))
		s.Execute("debug stop")
	} else if strings.Contains(c.Argv[0], "seconds") {
		s.Say(c.Argv[0])
	} else {
		text := "使用 !!tps [秒数] 指定获取多少秒内的tps"
		s.Tell(c.Player, text)
	}
}
