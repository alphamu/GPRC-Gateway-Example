package config

import "fmt"

type DbDetails struct {
	Username string
	Password string
	Protocol string
	Address  string
	DbName   string
}

func (db *DbDetails) ConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@%s%s/%s",
		db.Username,
		db.Password,
		db.Protocol,
		db.Address,
		db.DbName)
}
