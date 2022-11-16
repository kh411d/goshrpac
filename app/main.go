package main

import (
	"fmt"

	"github.com/kh411d/goshrpac/app/internal/cmd"
	"github.com/kh411d/goshrpac/pkg/sqx"
)

func main() {
	cmd.New().Execute()
	x := sqx.Eq{"hello": "world"}
	a, b, c := x.ToSql()
	fmt.Printf("%v %v %v\n", a, b, c)
}
