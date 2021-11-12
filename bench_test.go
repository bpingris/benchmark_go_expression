package main

import (
	"testing"

	"github.com/Knetic/govaluate"
	"github.com/PaesslerAG/gval"
)

func Benchmark_Parsing(b *testing.B) {
	benchs := []struct {
		name string
		expr string
	}{
		{
			name: "single",
			expr: "1",
		},
		{
			name: "simple",
			expr: "10 < 20",
		},
		{
			name: "var",
			expr: "foo < 20",
		},
		{
			name: "complex",
			expr: `2 > 1 &&
			"something" != "nothing" ||
			date("2014-01-20") < date("Wed Jul  8 23:07:35 MDT 2015") && 
			object["Variable name with spaces"] <= array[0] &&
			modifierTest + 1000 / 2 > (80 * 100 % 2)`,
		},
	}

	for _, bb := range benchs {
		b.Run("govaluate_"+bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				govaluate.NewEvaluableExpression(bb.expr)
			}
		})
		b.Run("gval_"+bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gval.Full().NewEvaluable(bb.expr)
			}
		})
	}
}
