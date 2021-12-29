package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type mysql struct{}

// var _Dialect = (*mysql)(nil)

func init() {
	RegisterDialect("mysql", &mysql{})
}

func (m *mysql) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Int, reflect.Int32, reflect.Uint, reflect.Uint32:
		return "int"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Int8, reflect.Uint8:
		return "tinyint"
	case reflect.Float32, reflect.Float64:
		return "double"
	case reflect.String:
		return "varchar"
	case reflect.Array, reflect.Slice:
		return "text"
	case reflect.Bool:
		return "bool"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

func (m *mysql) TableExistSQL(tbName string, dbName string) (string, []interface{}) {
	args := []interface{}{dbName, tbName}
	return "select  TABLE_NAME  from  INFORMATION_SCHEMA . TABLES  where TABLE_SCHEMA = ? and  TABLE_NAME =?", args
}
