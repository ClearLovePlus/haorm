package clause

import (
	"fmt"
	"strings"
)

type generator func(values ...interface{}) (string, []interface{})

var generatorMap map[Type]generator

func init() {
	generatorMap = make(map[Type]generator)
	generatorMap[INSERT] = _insert
	generatorMap[VALUES] = _values
	generatorMap[SELECT] = _select
	generatorMap[LIMIT] = _limit
	generatorMap[WHERE] = _where
	generatorMap[ORDERBY] = _orderBy
	generatorMap[UPDATE] = _update
	generatorMap[DELETE] = _delete
	generatorMap[COUNT] = _count
}

func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ",")
}

//拼接 insert into user (id,name,age) 这样子的语句
func _insert(value ...interface{}) (string, []interface{}) {
	tableName := value[0]
	fields := strings.Join(value[1].([]string), ",")
	return fmt.Sprintf("insert into %s (%v)", tableName, fields), []interface{}{}
}

//拼接 VALUES (?,?,?,?,?),(?,?,?,?)这样的语句
func _values(values ...interface{}) (string, []interface{}) {
	var bindStr string
	var sql strings.Builder
	var vars []interface{}
	sql.WriteString("VALUES ")
	for i, value := range values {
		v := value.([]interface{})
		if bindStr == "" {
			bindStr = genBindVars(len(v))
		}
		sql.WriteString(fmt.Sprintf("(%v)", bindStr))
		if i+1 != len(values) {
			sql.WriteString(", ")
		}
		vars = append(vars, v...)
	}
	return sql.String(), vars
}

//拼接select field1,field2 from tableName 这样的语句
func _select(values ...interface{}) (string, []interface{}) {
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("select %v from %s", fields, tableName), []interface{}{}

}

//返回 where field1=
func _where(values ...interface{}) (string, []interface{}) {
	desc, vars := values[0], values[1:]
	return fmt.Sprintf("WHERE %s", desc), vars
}

//拼接limit 语句
//返回 limit 10 或者limit 1,10
func _limit(values ...interface{}) (string, []interface{}) {
	length := len(values)
	if length == 2 {
		return fmt.Sprintf("limit %d,%d", values[0], values[1]), []interface{}{}
	} else {
		return fmt.Sprintf("limit %d", values[0]), []interface{}{}
	}
}

//拼接order by 语句
//返回 order by id asc
func _orderBy(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("ORDER BY %s", values[0]), []interface{}{}
}

//拼接 update语句
func _update(values ...interface{}) (string, []interface{}) {
	tableName := values[0]
	m := values[1].(map[string]interface{})
	var keys []string
	var value []interface{}
	for k, v := range m {
		keys = append(keys, k+"= ?")
		value = append(value, v)

	}
	return fmt.Sprintf("UPDATE %s SET %s", tableName, strings.Join(keys, ", ")), value
}

//拼接delete语句
func _delete(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("DELETE FROM %s", values[0]), []interface{}{}
}

//count 语句
func _count(values ...interface{}) (string, []interface{}) {
	return _select(values[0], []string{"count(*)"})
}
