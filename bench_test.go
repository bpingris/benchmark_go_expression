package main

import (
	"context"
	"testing"

	"github.com/Knetic/govaluate"
	"github.com/PaesslerAG/gval"
	"github.com/antonmedv/expr"
)

func Benchmark(b *testing.B) {
	benchs := []struct {
		name   string
		expr   string
		params map[string]interface{}
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
			expr: "foo > 20",
			params: map[string]interface{}{
				"foo": 10,
			},
		},
		{
			name: "complex",
			expr: `foo > 20 && 
            2 > 1 &&
			jon != "nothing" ||
            smith == jon ||
			bar <= 50 &&
            baz == true &&
			modifierTest + 1000 / 2 > (80 * 100 % 2)`,
			params: map[string]interface{}{
				"foo":          10,
				"modifierTest": 5,
				"bar":          10,
				"baz":          true,
				"jon":          "hello",
				"smith":        "world",
			},
		},
	}
	for _, bb := range benchs {
		exp, err := govaluate.NewEvaluableExpression(bb.expr)
		if err != nil {
			b.Fatal(err)
		}
		b.Run("govaluate", func(b *testing.B) {
			b.Run("parsing_"+bb.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					govaluate.NewEvaluableExpression(bb.expr)
				}
			})
			b.Run("evaluating_"+bb.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, err := exp.Evaluate(bb.params)
					if err != nil {
						b.Fatal(err)
					}
				}

			})
		})
		eval, err := gval.Full().NewEvaluable(bb.expr)
		if err != nil {
			b.Fatal(err)
		}
		b.Run("gval", func(b *testing.B) {
			b.Run("parsing_"+bb.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					gval.Full().NewEvaluable(bb.expr)
				}
			})
			b.Run("evaluating_"+bb.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, err := eval(context.Background(), bb.params)
					if err != nil {
						b.Fatal(err)
					}
				}

			})
		})

		out, err := expr.Compile(bb.expr)
		if err != nil {
			b.Fatal(err)
		}
		b.Run("expr", func(b *testing.B) {
			b.Run("parsing_"+bb.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					expr.Compile(bb.expr)
				}
			})
			b.Run("evaluating_"+bb.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, err := expr.Run(out, bb.params)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		})
	}
}
