package hao

import (
	"database/sql"
	log "gochen/haorm/log"
	session "gochen/haorm/session"
	"time"
)

var defalutIdleNum int = 1
var defalutMaxConns int = 50
var defaultIdleTimeout time.Duration = 10
var defaultLiveTimeout time.Duration = 100

//自定义orm的主要入口，用于配置各种数据库
type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string, idleNum, maxConnect int, idleTimeout, liveTimeout time.Duration) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	//配置数据库连接池的最大空闲连接数
	if idleNum == 0 {
		db.SetMaxIdleConns(defalutIdleNum)
	} else {
		db.SetMaxIdleConns(idleNum)
	}
	//配置数据库连接池的最大连接数
	if maxConnect == 0 {
		db.SetMaxOpenConns(defalutMaxConns)
	} else {
		db.SetMaxOpenConns(maxConnect)
	}
	//配置空闲链接超时时间
	if idleTimeout == 0 {
		db.SetConnMaxIdleTime(defaultIdleTimeout)
	} else {
		db.SetConnMaxIdleTime(idleTimeout)
	}
	//配置最大连接存活时间
	if liveTimeout == 0 {
		db.SetConnMaxLifetime(defaultLiveTimeout)
	} else {
		db.SetConnMaxLifetime(liveTimeout)
	}
	e = &Engine{db: db}
	log.Info("connect database succes")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close databse")
	}
	log.Info("Close database success")
}

func (enginde *Engine) NewSession() *session.Session {
	return session.New(enginde.db)
}
