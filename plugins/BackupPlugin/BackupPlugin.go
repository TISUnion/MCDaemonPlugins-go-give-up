package BackupPlugin

import (
	"MCDaemon-go/command"
	"MCDaemon-go/lib"
)

type BackupPlugin struct{}

func (hp BackupPlugin) Handle(c *command.Command, s lib.Server) {
	text := "使用 !!backup save [备份名称] 指定备份的名称"
}
