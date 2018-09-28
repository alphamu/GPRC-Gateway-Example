package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/alphamu/goecho/config"
)

type dbConn struct {
	driver *sql.DB
}

var conn *dbConn

func open(dbDetails *config.DbDetails) *sql.DB {
	db, err := sql.Open("mysql", dbDetails.ConnectionString())
	if err != nil {
		panic(err.Error())
	}
	return db
}

func connect() *sql.DB {
	return open(config.GetConfig().MySql)
}

func GetConnection() *dbConn {
	if conn == nil {
		conn = &dbConn{
			driver: connect(),
		}
	}
	return conn
}

func (db *dbConn) Execute(run func(db *sql.DB) (*sql.Stmt, error)) (*sql.Stmt, error) {
	return run(db.driver)
}

func (db *dbConn) ExecuteInsert(run func(db *sql.DB) (sql.Result, error)) (sql.Result, error) {
	return run(db.driver)
}
func (db *dbConn) ExecuteSelect(run func(db *sql.DB) (*sql.Rows, error)) (*sql.Rows, error) {
	return run(db.driver)
}


