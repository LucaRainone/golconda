package golconda

import (
	"fmt"
	"strings"
)

type Condition struct {
	mode string
	ops  []string
	vals []string
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

func (c *Condition) IsEqual(field string, value string) *Condition {
	c.ops = append(c.ops, fmt.Sprintf("%s = ?", field))
	c.vals = append(c.vals, value)
	return c
}

func (c Condition) Build() string {
	return fmt.Sprintf("(%s)", strings.Join(c.ops, fmt.Sprintf(" %s ", c.mode)))
}

func (c Condition) String() string {
	return c.Build()
}

func (c *Condition) Values() []string {
	return c.vals
}
