package golconda

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PlaceholderDollar       string = "$"
	PlaceholderQuestionMark string = "?"
)

type Condition struct {
	built       bool
	mode        string
	placeholder string
	ops         []string
	vals        []interface{}
	operations  []operatorBuilder
}

type operatorParamBuilder func() string

type operatorBuilder func(operatorParamBuilder) Operator

type Operator struct {
	Expression string
	Vals       []interface{}
}

type GolcondaBuilderType struct {
}

type golcondaBuilderApi struct {
	NewAnd func() *Condition
	NewOr  func() *Condition
}

func (builder *GolcondaBuilderType) PlaceholderFormat(placeholderFormat string) *golcondaBuilderApi {
	api := golcondaBuilderApi{}

	api.NewAnd = func() *Condition {
		c := Condition{}
		c.mode = "AND"
		c.placeholder = placeholderFormat
		return &c
	}

	api.NewOr = func() *Condition {
		c := Condition{}
		c.mode = "OR"
		c.placeholder = placeholderFormat
		return &c
	}

	return &api
}

func (condition *Condition) AsOperator() operatorBuilder {
	return func(opb operatorParamBuilder) Operator {
		operator := Operator{}
		expression, values := condition.Build()
		operator.Expression = expression
		operator.Vals = values
		return operator
	}
}

var GolcondaBuilder GolcondaBuilderType = GolcondaBuilderType{}

// defaults
func NewAnd() *Condition {
	return GolcondaBuilder.PlaceholderFormat(PlaceholderQuestionMark).NewAnd()
}
func NewOr() *Condition {
	return GolcondaBuilder.PlaceholderFormat(PlaceholderQuestionMark).NewOr()
}

func (c *Condition) Append(ops ...operatorBuilder) *Condition {

	c.operations = append(c.operations, ops...)

	return c
}

func preBuild(c *Condition) {
	if !c.built {
		index := 0
		var placeholder operatorParamBuilder = func() string {
			return c.placeholder
		}
		if c.placeholder == PlaceholderDollar {
			placeholder = func() string {
				index++
				return PlaceholderDollar + strconv.Itoa(index)
			}
		}
		for i := 0; i < len(c.operations); i++ {
			operationBuilder := c.operations[i]
			op := operationBuilder(placeholder)
			if op.Expression != "" {
				c.ops = append(c.ops, op.Expression)
				c.vals = append(c.vals, op.Vals...)
			}
		}
		c.built = true
	}

}

func (c *Condition) Build() (string, []interface{}) {
	preBuild(c)
	if len(c.ops) == 0 {
		if c.mode == "OR" {
			return "(FALSE)", make([]interface{}, 0)
		}
		return "(TRUE)", make([]interface{}, 0)
	}
	return fmt.Sprintf("(%s)", strings.Join(c.ops, fmt.Sprintf(" %s ", c.mode))), c.values()
}

func (c *Condition) values() []interface{} {
	return c.vals
}
