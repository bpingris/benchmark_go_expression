# Benchmarks


```
goos: darwin
goarch: arm64
pkg: github.com/benchmark_go_expression
Benchmark/govaluate/parsing_single-8             2481657               459.7 ns/op           392 B/op         10 allocs/op
Benchmark/govaluate/evaluating_single-8         50350353                22.31 ns/op            0 B/op          0 allocs/op
Benchmark/gval/parsing_single-8                  2248018               534.2 ns/op          1616 B/op          8 allocs/op
Benchmark/gval/evaluating_single-8              583142328                2.046 ns/op           0 B/op          0 allocs/op
Benchmark/expr/parsing_single-8                  1337137               896.1 ns/op          1240 B/op         24 allocs/op
Benchmark/expr/evaluating_single-8              32718163                35.73 ns/op           32 B/op          1 allocs/op
Benchmark/govaluate#01/parsing_simple-8           870778              1279 ns/op            1056 B/op         27 allocs/op
Benchmark/govaluate#01/evaluating_simple-8      50152022                22.24 ns/op            0 B/op          0 allocs/op
Benchmark/gval#01/parsing_simple-8               1375668               869.6 ns/op          1768 B/op         16 allocs/op
Benchmark/gval#01/evaluating_simple-8           583062400                2.045 ns/op           0 B/op          0 allocs/op
Benchmark/expr#01/parsing_simple-8                651802              1667 ns/op            1672 B/op         30 allocs/op
Benchmark/expr#01/evaluating_simple-8           25006794                47.00 ns/op           32 B/op          1 allocs/op
Benchmark/govaluate#02/parsing_var-8              918428              1223 ns/op             952 B/op         25 allocs/op
Benchmark/govaluate#02/evaluating_var-8         12753966                93.62 ns/op           24 B/op          2 allocs/op
Benchmark/gval#02/parsing_var-8                  1215367               986.6 ns/op          1912 B/op         21 allocs/op
Benchmark/gval#02/evaluating_var-8               9837984               119.9 ns/op            19 B/op          2 allocs/op
Benchmark/expr#02/parsing_var-8                   660237              1681 ns/op            1704 B/op         31 allocs/op
Benchmark/expr#02/evaluating_var-8              10712137               111.4 ns/op            48 B/op          2 allocs/op
Benchmark/govaluate#03/parsing_complex-8           77148             15437 ns/op           13288 B/op        226 allocs/op
Benchmark/govaluate#03/evaluating_complex-8      2090036               572.2 ns/op            48 B/op          5 allocs/op
Benchmark/gval#03/parsing_complex-8               115651             10250 ns/op            5376 B/op        175 allocs/op
Benchmark/gval#03/evaluating_complex-8           1000000              1083 ns/op             336 B/op         17 allocs/op
Benchmark/expr#03/parsing_complex-8                66877             17812 ns/op           15730 B/op        113 allocs/op
Benchmark/expr#03/evaluating_complex-8           2093114               571.3 ns/op           136 B/op          8 allocs/op
PASS
ok      github.com/benchmark_go_expression      34.512s
```
