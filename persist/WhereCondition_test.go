package persist

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestWhereAndCondition(t *testing.T) {

	var where = &WhereCondition{
		Column: "user",
		Operand:Equals,
		Value: "Ali",
		And: &WhereCondition{
			Column: "password",
			Operand:Equals,
			Value: "password",
		},
	}

	assert.Equal(t, "WHERE user = ? AND password = ?", where.ToString())

}

func TestWhereOrCondition(t *testing.T) {

	var where = &WhereCondition{
		Column: "user",
		Operand:Equals,
		Value: "Ali",
		Or: &WhereCondition{
			Column: "password",
			Operand:Equals,
			Value: "password",
		},
	}

	assert.Equal(t, "WHERE user = ? OR password = ?", where.ToString())

}

func TestWhereAndOrCondition(t *testing.T) {

	var where = &WhereCondition{
		Column:  "user",
		Operand: Equals,
		Value:   "Ali",
		And: &WhereCondition{
			Column:  "password",
			Operand: Equals,
			Value:   "password",
			Or: &WhereCondition{
				Column:  "skip_password",
				Operand: Equals,
				Value:   "true",
			},
		},
	}

	assert.Equal(t, "WHERE user = ? AND password = ? OR skip_password = ?", where.ToString())

}