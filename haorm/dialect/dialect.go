package dialect

import (
	haolog "gochen/haorm/log"
	"reflect"
)

var dialectMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string, dbName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	_, ok := dialectMap[name]
	if ok {
		haolog.Warn("duplicate key,replace this?")
	}
	dialectMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	if !ok {
		haolog.Warn("this record is not exist!")
	}
	return
}
