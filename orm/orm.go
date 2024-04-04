package orm

import (
	sqlDriverMySql "github.com/go-sql-driver/mysql"
	"github.com/xlizy/common-go/enums/common_error"
	"github.com/xlizy/common-go/response"
	"github.com/xlizy/common-go/zlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
	"time"
)

type RootConfig struct {
	Orm   ormConfig             `yaml:"orm"`
	Multi map[string]*ormConfig `yaml:"orm-multi"`
}

type ormConfig struct {
	Dsn             string        `yaml:"dsn"`
	MaxOpen         int           `yaml:"max-open"`
	MaxIdle         int           `yaml:"max-idle"`
	ConnMaxLifetime time.Duration `yaml:"conn-max-lifetime"`
}

var _primaryDb *gorm.DB
var _multiDb = make(map[string]*gorm.DB)

var wd = ""

type ormLoggerWriter struct {
	logger.Writer
}

func (w ormLoggerWriter) Printf(template string, args ...interface{}) {
	if len(args) > 0 {
		args[0] = strings.Replace(args[0].(string), wd+"/", "", 1)
	}
	template = strings.Replace(template, "\n", " ", -1)
	zlog.Info(template, args...)
}

func NewConfig() *RootConfig {
	return &RootConfig{}
}

// InitOrm 初始化数据源
func InitOrm(rc *RootConfig) {
	wd, _ = os.Getwd()
	newLogger := logger.New(
		&ormLoggerWriter{},
		logger.Config{
			SlowThreshold: 1 * time.Minute,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)
	if rc.Multi != nil && len(rc.Multi) > 0 {
		for key, _config := range rc.Multi {
			var err error
			_db, err := gorm.Open(mysql.Open(_config.Dsn), &gorm.Config{PrepareStmt: true, Logger: newLogger})
			if err != nil {
				zlog.Error("连接Mysql异常:{}", err)
				panic(err)
			}
			if _db.Error != nil {
				zlog.Error("连接Mysql异常:{}", _db.Error.Error())
				panic(_db.Error)
			}
			sqlDB, _ := _db.DB()
			sqlDB.SetMaxOpenConns(_config.MaxOpen)
			sqlDB.SetMaxIdleConns(_config.MaxIdle)
			sqlDB.SetConnMaxLifetime(_config.ConnMaxLifetime)
			if _primaryDb == nil {
				_primaryDb = _db
			}
			_multiDb[key] = _db
		}
	} else {
		var err error
		db, err := gorm.Open(mysql.Open(rc.Orm.Dsn), &gorm.Config{PrepareStmt: true, Logger: newLogger})
		if err != nil {
			zlog.Error("连接Mysql异常:{}", err)
			panic(err)
		}
		if db.Error != nil {
			zlog.Error("连接Mysql异常:{}", db.Error.Error())
			panic(db.Error)
		}
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(rc.Orm.MaxOpen)
		sqlDB.SetMaxIdleConns(rc.Orm.MaxIdle)
		sqlDB.SetConnMaxLifetime(rc.Orm.ConnMaxLifetime)
		_primaryDb = db
	}

}

func Conn() *gorm.DB {
	return _primaryDb
}

func ConnMulti(key string) *gorm.DB {
	return _multiDb[key]
}

func ErrHandler(err error) *response.Response {
	zlog.Error("sql exec error:{}", err.Error())
	if e, ok := err.(*sqlDriverMySql.MySQLError); ok {
		ce := common_error.SYSTEM_ERROR
		if e.Number == uint16(1008) {
			ce = common_error.MYSQL_ERR_1008
		} else if e.Number == uint16(1012) {
			ce = common_error.MYSQL_ERR_1012
		} else if e.Number == uint16(1020) {
			ce = common_error.MYSQL_ERR_1020
		} else if e.Number == uint16(1021) {
			ce = common_error.MYSQL_ERR_1021
		} else if e.Number == uint16(1022) {
			ce = common_error.MYSQL_ERR_1022
		} else if e.Number == uint16(1037) {
			ce = common_error.MYSQL_ERR_1037
		} else if e.Number == uint16(1044) {
			ce = common_error.MYSQL_ERR_1044
		} else if e.Number == uint16(1045) {
			ce = common_error.MYSQL_ERR_1045
		} else if e.Number == uint16(1048) {
			ce = common_error.MYSQL_ERR_1048
		} else if e.Number == uint16(1049) {
			ce = common_error.MYSQL_ERR_1049
		} else if e.Number == uint16(1054) {
			ce = common_error.MYSQL_ERR_1054
		} else if e.Number == uint16(1062) {
			ce = common_error.MYSQL_ERR_1062
		} else if e.Number == uint16(1065) {
			ce = common_error.MYSQL_ERR_1065
		} else if e.Number == uint16(1114) {
			ce = common_error.MYSQL_ERR_1114
		} else if e.Number == uint16(1130) {
			ce = common_error.MYSQL_ERR_1130
		} else if e.Number == uint16(1133) {
			ce = common_error.MYSQL_ERR_1133
		} else if e.Number == uint16(1141) {
			ce = common_error.MYSQL_ERR_1141
		} else if e.Number == uint16(1142) {
			ce = common_error.MYSQL_ERR_1142
		} else if e.Number == uint16(1143) {
			ce = common_error.MYSQL_ERR_1143
		} else if e.Number == uint16(1149) {
			ce = common_error.MYSQL_ERR_1149
		} else if e.Number == uint16(1169) {
			ce = common_error.MYSQL_ERR_1169
		} else if e.Number == uint16(1216) {
			ce = common_error.MYSQL_ERR_1216
		}
		return response.Error(ce, e.Message)
	}
	return response.Succ()
}
