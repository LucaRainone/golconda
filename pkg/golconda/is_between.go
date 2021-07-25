package golconda

func IsBetween(field string, from interface{}, to interface{}) Operator {
	operator := Operator{}

	if from != nil || to != nil {
		if from == nil {
			operator.Expression = field + " <= ?"
			operator.Vals = append(operator.Vals, to)
		} else if to == nil {
			operator.Expression = field + " >= ?"
			operator.Vals = append(operator.Vals, from)
		} else {
			operator.Expression = field + " BETWEEN ? AND ?"
			operator.Vals = append(operator.Vals, from)
			operator.Vals = append(operator.Vals, to)
		}
	}

	return operator
}
