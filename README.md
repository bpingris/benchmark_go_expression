# Benchmarks

```
goos: darwin
goarch: arm64
pkg: github.com/benchmark_go_expression
Benchmark/govaluate/parsing_single-8         	 2429138	       486.3 ns/op	     392 B/op	      10 allocs/op
Benchmark/govaluate/evaluating_single-8      	51002877	        22.30 ns/op	       0 B/op	       0 allocs/op
Benchmark/gval/parsing_single-8              	 2244898	       570.7 ns/op	    1616 B/op	       8 allocs/op
Benchmark/gval/evaluating_single-8           	575525192	         2.147 ns/op	       0 B/op	       0 allocs/op
Benchmark/govaluate#01/parsing_simple-8      	  945697	      1294 ns/op	    1056 B/op	      27 allocs/op
Benchmark/govaluate#01/evaluating_simple-8   	50729768	        22.76 ns/op	       0 B/op	       0 allocs/op
Benchmark/gval#01/parsing_simple-8           	 1378144	       869.6 ns/op	    1768 B/op	      16 allocs/op
Benchmark/gval#01/evaluating_simple-8        	589021857	         2.029 ns/op	       0 B/op	       0 allocs/op
Benchmark/govaluate#02/parsing_var-8         	  970818	      1214 ns/op	     952 B/op	      25 allocs/op
Benchmark/govaluate#02/evaluating_var-8      	12689644	        93.90 ns/op	      24 B/op	       2 allocs/op
Benchmark/gval#02/parsing_var-8              	 1220355	      1003 ns/op	    1912 B/op	      21 allocs/op
Benchmark/gval#02/evaluating_var-8           	 9717634	       123.0 ns/op	      19 B/op	       2 allocs/op
Benchmark/govaluate#03/parsing_complex-8     	   86655	     13662 ns/op	   10480 B/op	     196 allocs/op
Benchmark/govaluate#03/evaluating_complex-8  	 3924668	       335.8 ns/op	      48 B/op	       5 allocs/op
Benchmark/gval#03/parsing_complex-8          	  149298	      7835 ns/op	    4320 B/op	     123 allocs/op
Benchmark/gval#03/evaluating_complex-8       	 2462046	       490.0 ns/op	      80 B/op	       7 allocs/op
PASS
ok  	github.com/benchmark_go_expression	24.024s
```

