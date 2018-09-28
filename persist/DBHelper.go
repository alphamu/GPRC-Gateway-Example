package persist

import (
	"database/sql"
	"strings"
	"fmt"
	"github.com/alphamu/goecho/db"
	)

func WriteMessageToDb(messages []string, messageFrom, messageTo string) (sql.Result, error) {
	conn := db.GetConnection()
	return conn.ExecuteInsert(func(db *sql.DB) (sql.Result, error) {
		stmtIn, err := PrepareInsert(db, "messages", []string{"message", "message_from", "message_to"})
		defer stmtIn.Close()
		if err != nil {
			panic(err.Error())
		}


		for i := 0; i < len(messages); i++ {
			stmtIn.Exec(messages, messageFrom, messageTo)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
}

func ReadMessagesForUser(user string) ([]string, error) {
	conn := db.GetConnection()
	rows, err := conn.ExecuteSelect(func(db *sql.DB) (*sql.Rows, error) {
		where := &WhereCondition{
			Column:  "message_to",
			Operand: Equals,
		}
		stmtOut, err := PrepareSelect(db, "messages", []string{"message"}, where)

		defer stmtOut.Close()
		if err != nil {
			return nil, err
		}
		return stmtOut.Query(user)
	})

	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	count := 0
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		var message string
		err = rows.Scan(&message)
		if err != nil {
			return nil, err
		}
		results = append(results, message)
		count++
	}

	return results, nil
}

func PrepareInsert(db *sql.DB, table string, columns []string) (*sql.Stmt, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, arrayToCsv(columns), getCsvQuestionMarks(len(columns)))
	return db.Prepare(query)
}

func PrepareSelect(db *sql.DB, table string, columns []string, condition *WhereCondition) (*sql.Stmt, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", arrayToCsv(columns), table)
	if condition != nil {
		query = fmt.Sprintf("%s %s", query, condition.ToString())
	}
	return db.Prepare(query)
}

func arrayToCsv(array []string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ","), "[]")
}

func getCsvQuestionMarks(count int) string {
	var q = make([]string, count)
	for i := 0; i < count; i++ {
		q[i] = "?"
	}
	return arrayToCsv(q)
}
