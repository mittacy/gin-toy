package eMysql

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

// GetConnectByConf 获取gorm连接
func GetConnectByConf(conf MultipleConf) (*gorm.DB, error) {
	dsn := conf.Sources[0].DSN()

	db, err := gorm.Open(mysql.Open(dsn), gormConfig())
	if err != nil {
		return nil, err
	}

	if len(conf.Sources) > 1 || len(conf.Replicas) > 0 {
		sources, replicas := conf.ConfToGormDialector()

		dbResolverConfig := dbresolver.Config{
			Sources:  sources,
			Replicas: nil,
			Policy:   dbresolver.RandomPolicy{},
		}

		if len(replicas) > 0 {
			dbResolverConfig.Replicas = replicas
		}

		if err = db.Use(dbresolver.Register(dbResolverConfig)); err != nil {
			return nil, err
		}
	}

	return db, nil
}

var ErrNoInit = errors.New("eMysql: please initialize with eMysql.Init()")

// EGorm 在结构体引入组合并赋值ConfName，即可通过GDB()获取gorm连接
// Example
// type User struct {
// 	 EGorm
// }
//
// var user = User{EGorm{ConfName: "localhost"}}
//
// func (u *User) GetUser(id int64) error {
// 	 u.GDB().Where("id = ?", id).First()
// }
type EGorm struct {
	ConfName string
}

// GDB 获取DB连接
func (ctl *EGorm) GDB() *gorm.DB {
	return GetConnect(ctl.ConfName)
}

func GetConnect(name string) *gorm.DB {
	if !initFlag {
		panic(ErrNoInit)
	}

	if db, ok := connectPool[name]; ok {
		return db
	}

	conf, isExist := connectConf[name]
	if !isExist {
		panic("eGorm: " + name + "配置不存在, 请检查配置")
	}

	db, err := GetConnectByConf(conf)
	if err != nil {
		l.Error(context.Background(), fmt.Sprintf("eGorm: 连接数据库失败, conf: %+v, err: %+v", conf, err))
		return &gorm.DB{}
	}
	connectPool[name] = db
	return db
}

// gormConfig gorm连接配置
func gormConfig() *gorm.Config {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 是否禁用表名复数形式
		Logger:         l,
	}
}
