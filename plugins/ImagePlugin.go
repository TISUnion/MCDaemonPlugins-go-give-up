/**
 * 开启、关闭镜像
 * 前置插件： 备份插件backup
 */

package plugin

import (
	"MCDaemon-go/command"
	"MCDaemon-go/container"
	"MCDaemon-go/lib"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
)

type ImagePlugin struct{}

func (lp *ImagePlugin) Handle(c *command.Command, s lib.Server) {
	if len(c.Argv) == 0 {
		c.Argv = append(c.Argv, "help")
	}
	cor := container.GetInstance()
	switch c.Argv[0] {
	case "show":
		backupfiles, _ := filepath.Glob("back-up/*")
		//标记是否启动
		for k, v := range backupfiles {
			if cor.IsRuntime(v) {
				backupfiles[k] += "   已启动"
			} else {
				backupfiles[k] += "   未启动"
			}
		}
		text := "备份如下：\\n" + strings.Join(backupfiles, "\\n")
		s.Tell(c.Player, text)
	case "start":
		if len(c.Argv) == 1 {
			s.Tell(c.Player, "缺少启动的镜像名称")
		}
		cor = container.GetInstance()
		svr := s.Clone()
		//修改端口
		sercfg, _ := ini.Load("back-up/" + c.Argv[1] + "/server.properties")
		sercfg.Section("").NewKey("server-port", svr.GetPort())
		//启动
		cor.Add(c.Argv[1], "back-up/"+c.Argv[1], svr)
	case "stop":
		if len(c.Argv) == 1 {
			s.Tell(c.Player, "缺少停止的镜像名称")
		}
		cor = container.GetInstance()
		cor.Del(c.Argv[1])
	default:
		text := "!!image show 查看镜像\\n!!image start [镜像名称] 开启镜像 \\n!!image stop [镜像名称] 关闭镜像"
		s.Tell(c.Player, text)
	}
}

func (lp *ImagePlugin) Init(s lib.Server) {
}
