package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type OrmConfig struct {
	Dsn             string        `yaml:"dsn"`
	MaxOpen         int           `yaml:"max-open"`
	MaxIdle         int           `yaml:"max-idle"`
	ConnMaxLifetime time.Duration `yaml:"conn-max-lifetime"`
}

type MultiOrmConfig struct {
	Multi map[string]*OrmConfig
}

var _primaryDb *gorm.DB
var _multiDb map[string]*gorm.DB

// InitOrm 初始化单数据源
func InitOrm(config OrmConfig) {
	var err error
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if db.Error != nil {
		panic(db.Error)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(config.MaxOpen)
	sqlDB.SetMaxIdleConns(config.MaxIdle)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	_primaryDb = db
}

// MultiOrm 初始化多数据源
func MultiOrm(config MultiOrmConfig) {
	multi := config.Multi
	if multi != nil {
		for key, _config := range multi {
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
	}
}

func Conn() *gorm.DB {
	return _primaryDb
}

func ConnMulti(key string) *gorm.DB {
	return _multiDb[key]
}
