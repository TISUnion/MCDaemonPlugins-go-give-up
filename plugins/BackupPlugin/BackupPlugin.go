package BackupPlugin

import (
	"MCDaemon-go/command"
	"MCDaemon-go/config"
	"MCDaemon-go/lib"
	"fmt"
	"os/exec"
	"runtime"
)

type BackupPlugin struct {
	backupName string
}

func (bp BackupPlugin) Handle(c *command.Command, s lib.Server) {
	server_name := config.Cfg.Section("MCDeamon").Key("server_name").String()
	server_path := config.Cfg.Section("MCDeamon").Key("server_path").String()
	serverfile := fmt.Sprintf("%s/%s", server_path, server_name)
	switch c.Argv[0] {
	case "save":
		bp.backupName = c.Argv[1]
		s.Execute("/save-all flush")
	case "saved":
		Copy(serverfile, "back-up/"+bp.backupName)
	case "save-compress":
		Copy(serverfile, "back-up/"+bp.backupName)
		if runtime.GOOS == "windows" {
			s.Tell(c.Player, "windows服务器不支持压缩功能")
		} else {
			cmd := exec.Command("tar", "zcvf", "back-up/"+bp.backupName+".tar.gz", "back-up/"+bp.backupName)
			if err := cmd.Run(); err != nil {
				s.Tell(c.Player, fmt.Sprint("压缩姬出问题了，因为", err))
			}
		}
	}
}
