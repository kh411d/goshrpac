package sqx

import (
	"fmt"
	"testing"
)

func TestToSql(t *testing.T) {
	sq := `SELECT ta.* , tb.*
	   FROM tableA ta INNER JOIN tableB tb`

	eq := Eq{"ta.column_a": []int64{1, 3, 5}, "tb.column_b": 2345}
	noteq := NotEq{"ta.column_a": []int64{1, 3, 5}, "tb.column_b": 2345}
	like := Like{"ta.column_c": "%same"}
	likes := Like{"ta.column_c": "%same"}

	x := And{
		Eq{"ta.and": []int64{1, 3, 5}},
		Eq{"ta.and": []int64{1, 3, 5}, "tb.and": 2345},
	}

	query, args, err := ToSql(Expr(sq), Expr("WHERE"), eq, Expr("AND"), like, Expr("OR"), likes, Expr("AND"), noteq, Expr("AND"), x)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(query)
	fmt.Printf("Args: %#v\n", args)

}

func TestToSqlBy(t *testing.T) {

	sq := `INSERT INTO stats(totalProduct, totalCustomer, totalOrder)
	VALUES `

	expr := Expr("(?, ?, ?)", 123, "12334", 23)

	query, args, err := ToSqlWith(",")(expr, expr)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(sq + query)
	fmt.Printf("Args: %#v\n", args)

}
