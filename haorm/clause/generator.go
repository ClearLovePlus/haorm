package clause

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
}

func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {

	}
}
