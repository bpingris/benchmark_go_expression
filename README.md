# Benchmarks

```
goos: darwin
goarch: arm64
pkg: github.com/benchmark_go_expression
Benchmark/govaluate/parsing_single-8             2462026               466.3 ns/op           392 B/op         10 allocs/op
Benchmark/govaluate/evaluating_single-8         50183656                22.44 ns/op            0 B/op          0 allocs/op
Benchmark/gval/parsing_single-8                  2240035               536.4 ns/op          1616 B/op          8 allocs/op
Benchmark/gval/evaluating_single-8              580352409                2.062 ns/op           0 B/op          0 allocs/op
Benchmark/govaluate#01/parsing_simple-8           901184              1286 ns/op            1056 B/op         27 allocs/op
Benchmark/govaluate#01/evaluating_simple-8      50348416                22.58 ns/op            0 B/op          0 allocs/op
Benchmark/gval#01/parsing_simple-8               1359560               879.5 ns/op          1768 B/op         16 allocs/op
Benchmark/gval#01/evaluating_simple-8           562839902                2.073 ns/op           0 B/op          0 allocs/op
Benchmark/govaluate#02/parsing_var-8              959343              1226 ns/op             952 B/op         25 allocs/op
Benchmark/govaluate#02/evaluating_var-8         12437332                94.20 ns/op           24 B/op          2 allocs/op
Benchmark/gval#02/parsing_var-8                  1207848              1005 ns/op            1912 B/op         21 allocs/op
Benchmark/gval#02/evaluating_var-8               9251005               126.1 ns/op            19 B/op          2 allocs/op
Benchmark/govaluate#03/parsing_complex-8           70939             16945 ns/op           12584 B/op        216 allocs/op
Benchmark/govaluate#03/evaluating_complex-8     26078617                43.53 ns/op           16 B/op          1 allocs/op
Benchmark/gval#03/parsing_complex-8               136020              7854 ns/op            4408 B/op        122 allocs/op
Benchmark/gval#03/evaluating_complex-8          199876783                6.100 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/benchmark_go_expression      23.423s
```
