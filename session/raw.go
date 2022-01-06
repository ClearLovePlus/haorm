package haosession

import (
	"database/sql"
	"strings"
	"sync"

	"github.com/ClearLovePlus/haorm/clause"
	"github.com/ClearLovePlus/haorm/dialect"
	log "github.com/ClearLovePlus/haorm/log"
	"github.com/ClearLovePlus/haorm/schema"
)

var mu sync.Mutex

//和数据库交互的session对象
type Session struct {
	db     *sql.DB
	dbName string
	//sql语句
	sql      strings.Builder
	dialect  dialect.Dialect
	refTable *schema.Schema
	clause   clause.Clause
	sqlVars  []interface{}
	tx       *sql.Tx
}
type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

func New(db *sql.DB, dialect dialect.Dialect, dbName string) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
		dbName:  dbName,
	}
}
func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db

}
func (s *Session) Clear() {
	mu.Lock()
	defer mu.Unlock()
	s.sql.Reset()
	s.sqlVars = nil
	s.dbName = ""
	s.clause = clause.Clause{}
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

//执行一般的语句
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

//查询单条数据
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

//查询多条数据
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
