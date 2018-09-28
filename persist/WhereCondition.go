package persist

import (
	"fmt"
	"strings"
)

type WhereOperand string
type WhereAndOr string

const (
	Equals 				WhereOperand = "="
	GreaterThan 		WhereOperand = ">"
	GreaterThanEquals 	WhereOperand = ">="
	LessThan 			WhereOperand = "<"
	LessThanEquals 		WhereOperand = "<="

	And WhereAndOr = "AND"
	Or WhereAndOr = "OR"
)



type WhereCondition struct {
	Column 			string
	Operand 		WhereOperand
	Value 			interface{}
	And				*WhereCondition
	Or 				*WhereCondition
}

func (w *WhereCondition) ToString() string {
	conditions := flatten(w, "WHERE %s")

	return strings.Trim(conditions, " ")
}

func flatten(w *WhereCondition, template string) string {
	conditions := fmt.Sprintf(template, w.getString())
	if w.And != nil {
		conditions = fmt.Sprintf("%s %s", conditions, flatten(w.And, "%s"))
	} else if w.Or != nil {
		conditions = fmt.Sprintf("%s %s", conditions, flatten(w.Or, "%s"))
	}
	return conditions
}

func (w *WhereCondition) getString() string {
	return fmt.Sprintf(
		"%s %s %s %s",
		w.Column,
		w.Operand,
		"?",
		w.getNextConditionString())
}

func (w *WhereCondition) getNextConditionString() string {
	if w.And != nil {
		return "AND"
	}  else if w.Or != nil {
		return "OR"
	} else {
		return ""
	}
}


