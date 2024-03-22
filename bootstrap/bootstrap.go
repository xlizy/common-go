package bootstrap

import (
	"github.com/xlizy/common-go/config"
	"github.com/xlizy/common-go/nacos/v2"
	"github.com/xlizy/common-go/snowflake"
	"github.com/xlizy/common-go/zlog"
)

func Init() {
	zlog.InitLogger(config.GetLogCfg().Path)
	snowflake.Init(1, 1)
	zlog.Info("项目开始启动...")
	nacos.InitNacos()
}
