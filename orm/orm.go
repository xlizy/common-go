package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type RootConfig struct {
	Orm   OrmConfig             `yaml:"orm"`
	Multi map[string]*OrmConfig `yaml:"orm-multi"`
}

type OrmConfig struct {
	Dsn             string        `yaml:"dsn"`
	MaxOpen         int           `yaml:"max-open"`
	MaxIdle         int           `yaml:"max-idle"`
	ConnMaxLifetime time.Duration `yaml:"conn-max-lifetime"`
}

var _primaryDb *gorm.DB
var _multiDb map[string]*gorm.DB

// InitOrm 初始化数据源
func InitOrm(rc RootConfig) {
	if rc.Multi != nil && len(rc.Multi) > 0 {
		for key, _config := range rc.Multi {
			var err error
			_db, err := gorm.Open(mysql.Open(_config.Dsn), &gorm.Config{})
			if err != nil {
				panic(err)
			}
			if _db.Error != nil {
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
		db, err := gorm.Open(mysql.Open(rc.Orm.Dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		if db.Error != nil {
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
