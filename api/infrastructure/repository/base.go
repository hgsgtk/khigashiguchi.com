package repository

import "database/sql"

type (
	Execer interface {
		Exec(string, ...interface{}) (sql.Result, error)
	}
	Queryer interface {
		Query(string, ...interface{}) (*sql.Rows, error)
		QueryRow(string, ...interface{}) *sql.Row
	}
	Preparer interface {
		Prepare(string) (*sql.Stmt, error)
	}
	Beginner interface {
		Begin() (Tx, error)
	}
)

type Executor interface {
	Execer
	Queryer
	Preparer
	//Beginner
}

type Tx interface {
	Rollback() error
	Commit() error
	Executor
}