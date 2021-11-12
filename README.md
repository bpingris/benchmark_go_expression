# Benchmarks

goos: darwin
goarch: arm64
pkg: github.com/benchmark_go_expression
Benchmark_Parsing/govaluate_single-8         	 2458281	       457.5 ns/op
Benchmark_Parsing/gval_single-8              	 2263848	       530.0 ns/op
Benchmark_Parsing/govaluate_simple-8         	  921517	      1268 ns/op
Benchmark_Parsing/gval_simple-8              	 1378863	       871.3 ns/op
Benchmark_Parsing/govaluate_var-8            	  972102	      1216 ns/op
Benchmark_Parsing/gval_var-8                 	 1217625	       981.2 ns/op
Benchmark_Parsing/govaluate_complex-8        	   95074	     12496 ns/op
Benchmark_Parsing/gval_complex-8             	  128590	      9216 ns/op
PASS
ok  	github.com/benchmark_go_expression	12.699s
