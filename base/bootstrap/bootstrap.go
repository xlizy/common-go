package bootstrap

import (
	"github.com/xlizy/common-go/base/config"
	"github.com/xlizy/common-go/components/nacos/v2"
	"github.com/xlizy/common-go/utils/snowflake"
	"github.com/xlizy/common-go/utils/zlog"
)

func Init() {
	zlog.InitLogger(config.GetLogCfg().Path)
	snowflake.Init(1, 1)
	zlog.Info("项目开始启动...")
	nacos.InitNacos()
}
