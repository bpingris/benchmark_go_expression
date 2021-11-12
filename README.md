# Benchmarks


```
goos: darwin
goarch: arm64
pkg: github.com/benchmark_go_expression
Benchmark/govaluate/parsing_single-8         	 2443224	       457.0 ns/op	     392 B/op	      10 allocs/op
Benchmark/govaluate/evaluating_single-8      	51203276	        22.32 ns/op	       0 B/op	       0 allocs/op
Benchmark/gval/parsing_single-8              	 1994931	       551.7 ns/op	    1616 B/op	       8 allocs/op
Benchmark/gval/evaluating_single-8           	580653591	         2.077 ns/op	       0 B/op	       0 allocs/op
Benchmark/govaluate#01/parsing_simple-8      	  921732	      1322 ns/op	    1056 B/op	      27 allocs/op
Benchmark/govaluate#01/evaluating_simple-8   	49599159	        22.34 ns/op	       0 B/op	       0 allocs/op
Benchmark/gval#01/parsing_simple-8           	 1365794	       880.9 ns/op	    1768 B/op	      16 allocs/op
Benchmark/gval#01/evaluating_simple-8        	557657311	         2.030 ns/op	       0 B/op	       0 allocs/op
Benchmark/govaluate#02/parsing_var-8         	  908763	      1237 ns/op	     952 B/op	      25 allocs/op
Benchmark/govaluate#02/evaluating_var-8      	12611107	        95.08 ns/op	      24 B/op	       2 allocs/op
Benchmark/gval#02/parsing_var-8              	 1000000	      1021 ns/op	    1912 B/op	      21 allocs/op
Benchmark/gval#02/evaluating_var-8           	10098129	       122.5 ns/op	      19 B/op	       2 allocs/op
Benchmark/govaluate#03/parsing_complex-8     	   68862	     15626 ns/op	   13288 B/op	     226 allocs/op
Benchmark/govaluate#03/evaluating_complex-8  	 2033379	       576.5 ns/op	      48 B/op	       5 allocs/op
Benchmark/gval#03/parsing_complex-8          	  112210	     10356 ns/op	    5376 B/op	     175 allocs/op
Benchmark/gval#03/evaluating_complex-8       	 1000000	      1101 ns/op	     336 B/op	      17 allocs/op
PASS
ok  	github.com/benchmark_go_expression	22.062s
```
