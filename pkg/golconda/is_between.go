package golconda

func IsBetween(field string, from interface{}, to interface{}) func(paramPlaceholder func() string) Operator {
	return func(paramPlaceholder func() string) Operator {
		operator := Operator{}

		if from != nil || to != nil {
			if from == nil {
				return IsLessOrEqual(field, to)(func() string { return "?" })
			} else if to == nil {
				return IsGreaterOrEqual(field, from)(func() string { return "?" })
			} else {
				operator.Expression = field + " BETWEEN " + paramPlaceholder() + " AND " + paramPlaceholder()
				operator.Vals = append(operator.Vals, from)
				operator.Vals = append(operator.Vals, to)
			}
		}

		return operator
	}
}
