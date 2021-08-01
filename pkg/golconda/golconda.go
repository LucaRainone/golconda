package golconda

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PlaceholderDollar   string = "$"
	PlaceholderQuestion string = "?"
)

type Condition struct {
	built      bool
	mode       string
	ops        []string
	vals       []interface{}
	operations []operatorBuilder
}

type operatorParamBuilder func() string

type operatorBuilder func(operatorParamBuilder) Operator

var currentPlaceholderMethod = PlaceholderQuestion

type Operator struct {
	Expression string
	Vals       []interface{}
}

func NewAnd() *Condition {
	c := Condition{}
	c.mode = "AND"
	return &c
}

func NewOr() *Condition {
	c := Condition{}
	c.mode = "OR"
	return &c
}

func SetPlaceholder(pl string) {
	currentPlaceholderMethod = pl
}

func (c *Condition) Append(ops ...operatorBuilder) *Condition {

	c.operations = append(c.operations, ops...)

	return c
}

func preBuild(c *Condition) {
	if !c.built {
		index := 0
		var placeholder operatorParamBuilder = func() string {
			return currentPlaceholderMethod
		}
		if currentPlaceholderMethod == PlaceholderDollar {
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
