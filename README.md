# Golconda

## Why

This is my first attempt with Golang, so I trained myself with a porting of my [Condition-builder](https://github.com/LucaRainone/condition-builder) in PHP.

Due the fact this is a try, feel free to help me to understand the basics of Go ~~if~~ when you see mistakes or best practice violations. Thanks.

## What

Golconda stands for **GOL**ang **COND**ition **A**id: it gives you a simple help to build conditions for SQL queries. It supports params placeholder

## How

this is a full example

```go
package main

import (
    "LucaRainone/golconda"
    "fmt"
)

type filters struct {
    byId        interface{}
    byDateStart interface{}
    byDateEnd   interface{}
    byEmail     interface{}
}

func buildCondition(filters filters) (string, []interface{}) {
    c := golconda.NewAnd()
    c.Append(
        golconda.IsEqual("id", filters.byId),
        golconda.IsEqual("email", filters.byEmail),
        golconda.IsBetween("date", filters.byDateStart, filters.byDateEnd),
    )
    return c.Build()
}

func main() {

    filters := filters{}
    filters.byId = []int{1, 2, 3}
    filters.byDateStart = "2021-08-01"
    // note that we're missing filters.byEmail and byDateEnd.
    // when filters are nil, they will be ignored by golconda

    whereString, whereValues := buildCondition(filters)
    fmt.Println(whereString) // -> (id IN (?,?,?) AND date >= ?)
    fmt.Printf("%#v\n", whereValues) // -> []interface {}{1, 2, 3, "2021-08-01"}

    // Do you love Squirrel?
    users := sq.Select("*").From("users")
    users.Where(whereString, whereValues)
    // ... and so on
}


```

What if we want to filter by expressions?

```go

filters.byDateStart = golconda.SqlExpression("NOW()")

```

What if we have multiple conditions?

```go
condition := golconda.NewAnd()
subCondition := golconda.NewOr()
// [... fill your conditions]

c.Append(subCondition.AsOperator())

```

What if we want to use Postgres?

```go

golcondaP := golconda.GolcondaBuilder.PlaceholderFormat(golconda.PlaceholderDollar)

c := golcondaP.NewAnd()
c.Append(golconda.IsEqual("id", 23))
query, values := c.Build()

fmt.Printf("%s -> %v", query, values)
// (id = $1) -> [23]

```

That's all.
