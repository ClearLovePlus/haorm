package haosession

import haolog "github.com/ClearLovePlus/haorm/log"

//开启事务方法
func (s *Session) Begin() (err error) {
	haolog.Info("transaction begin")
	if s.tx, err = s.db.Begin(); err != nil {
		haolog.Error(err)
	}
	return

}

//提交事务方法
func (s *Session) Commit() (err error) {
	haolog.Info("transaction commit")
	if err = s.tx.Commit(); err != nil {
		haolog.Error(err)
	}
	return
}

//回滚事务方法
func (s *Session) Rollback() (err error) {
	haolog.Info("transaction rollback")
	if err = s.tx.Rollback(); err != nil {
		haolog.Error(err)
	}
	return
}
