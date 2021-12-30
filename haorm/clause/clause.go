package clause

type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]interface{}
}

type Type int

const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
)

func (c *Clause) set(name Type, vars ...interface{}) {
	if c.sql == nil {
	}
}
