package session

import (
	"database/sql"
	"gee/log"
)

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars); err != nil {
		log.Error(err)
	}
	return result, err
}
