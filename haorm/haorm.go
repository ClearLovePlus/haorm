package hao

import (
	"database/sql"
	"gochen/haorm/dialect"
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
	db      *sql.DB
	dbName  string
	dialect dialect.Dialect
}

type TxFunc func(*session.Session) (interface{}, error)

func NewEngine(driver, source string, idleNum, maxConnect int, idleTimeout, liveTimeout time.Duration, dbName string) (e *Engine, err error) {
	if dbName == "" {
		log.Error("dbName can not be null")
		return
	}
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("%s dialect cannot be null,should create the dialect belong to %s", driver, driver)
	}
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
	e = &Engine{
		db:      db,
		dbName:  dbName,
		dialect: dial,
	}
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
	return session.New(enginde.db, enginde.dialect, enginde.dbName)
}

func (engine *Engine) Transation(f TxFunc) (result interface{}, err error) {
	s := engine.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = s.Rollback()
			panic(p)
		} else if err != nil {
			_ = s.Rollback()
		} else {
			if err1 := s.Commit(); err1 != nil {
				_ = s.Rollback()
			}
		}
	}()
	return f(s)
}
