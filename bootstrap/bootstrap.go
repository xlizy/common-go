package bootstrap

import (
	"github.com/xlizy/common-go/config"
	"github.com/xlizy/common-go/nacos"
	"github.com/xlizy/common-go/zlog"
)

func Init() {
	zlog.InitLogger(config.GetLogCfg().Path)
	zlog.Info("项目开始启动...")
	nacos.InitNacos()
}
