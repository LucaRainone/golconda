package golconda

import (
	"fmt"
	"strings"
)

type Condition struct {
	mode string
	ops  []string
	vals []interface{}
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

func (c *Condition) Append(ops ...Operator) *Condition {
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		if op.Expression != "" {
			c.ops = append(c.ops, op.Expression)
			c.vals = append(c.vals, op.Vals...)
		}
	}
	return c
}

func (c Condition) Build() string {
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
