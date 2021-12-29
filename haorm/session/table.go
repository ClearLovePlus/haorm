package haosession

import (
	"errors"
	"fmt"
	haolog "gochen/haorm/log"
	"gochen/haorm/schema"
	"reflect"
	"strings"
)

func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		haolog.Error("Model is not set")
		panic("Model is not set")
	}
	return s.refTable
}

//创建表
func (s *Session) CreateTable() error {
	table := s.refTable
	if table == nil {
		haolog.Error("refTable must not be null")
		return errors.New("refTable must not be null")
	}
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)).Exec()
	return err
}
