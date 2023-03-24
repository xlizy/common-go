package bootstrap

import (
	"github.com/xlizy/common-go/config"
	"github.com/xlizy/common-go/nacos"
	"github.com/xlizy/common-go/zlog"
)

func Init() {
	zlog.InitLogger(config.GetLogCfg().Path)
	nacos.InitNacos()
}
