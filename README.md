# Benchmarks

goos: darwin
goarch: arm64
pkg: github.com/benchmark_go_expression
Benchmark/govaluate/parsing_single-8             2514556               456.5 ns/op
Benchmark/govaluate/evaluating_single-8         52046814                22.44 ns/op
Benchmark/gval/parsing_single-8                  2221927               539.3 ns/op
Benchmark/gval/evaluating_single-8              584697728                2.040 ns/op
Benchmark/govaluate#01/parsing_simple-8           909288              1272 ns/op
Benchmark/govaluate#01/evaluating_simple-8      50682453                22.34 ns/op
Benchmark/gval#01/parsing_simple-8               1379097               867.9 ns/op
Benchmark/gval#01/evaluating_simple-8           584877502                2.044 ns/op
Benchmark/govaluate#02/parsing_var-8              960124              1220 ns/op
Benchmark/govaluate#02/evaluating_var-8         12614542                93.51 ns/op
Benchmark/gval#02/parsing_var-8                  1222411               984.8 ns/op
Benchmark/gval#02/evaluating_var-8               9715621               120.7 ns/op
Benchmark/govaluate#03/parsing_complex-8           71268             16654 ns/op
Benchmark/govaluate#03/evaluating_complex-8     28613354                41.73 ns/op
Benchmark/gval#03/parsing_complex-8               155883              7652 ns/op
Benchmark/gval#03/evaluating_complex-8          203008358                5.946 ns/op
PASS
ok      github.com/benchmark_go_expression      23.558s
