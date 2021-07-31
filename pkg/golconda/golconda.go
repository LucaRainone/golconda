package golconda

import (
	"fmt"
	"strings"
)

type Condition struct {
	built      bool
	mode       string
	ops        []string
	vals       []interface{}
	operations []func(paramPlaceholder func() string) Operator
}

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

func (c *Condition) Append(ops ...func(paramPlaceholder func() string) Operator) *Condition {

	c.operations = append(c.operations, ops...)

	return c
}

func preBuild(c *Condition) {
	if !c.built {
		for i := 0; i < len(c.operations); i++ {
			operationBuilder := c.operations[i]
			op := operationBuilder(func() string { return "?" })
			if op.Expression != "" {
				c.ops = append(c.ops, op.Expression)
				c.vals = append(c.vals, op.Vals...)
			}
		}
		c.built = true
	}

}

func (c *Condition) Build() string {
	preBuild(c)
	if len(c.ops) == 0 {
		if c.mode == "OR" {
			return "(FALSE)"
		}
		return "(TRUE)"
	}
	return fmt.Sprintf("(%s)", strings.Join(c.ops, fmt.Sprintf(" %s ", c.mode)))
}

func (c Condition) String() string {
	return c.Build()
}

func (c *Condition) Values() []interface{} {
	return c.vals
}
