all-pairs shortest paths (APSP)
===============================

Go implementation of floyd-warshall, benchmarked (each "op" is solving
an instance of APSP):

```
[bpowers@vyse floyd-warshal]$ go test -bench .
BenchmarkFW10-4     	  500000	      3524 ns/op
BenchmarkFW20-4     	  100000	     22802 ns/op
BenchmarkFW50-4     	    5000	    300667 ns/op
BenchmarkFW100-4    	    1000	   2278273 ns/op
BenchmarkFW500-4    	       5	 306837154 ns/op
BenchmarkFW1000-4   	       1	2585802126 ns/op
BenchmarkFW5000-4   	       1	372099773461 ns/op
PASS
ok  	floyd-warshal	13.743s

```

So, for a 1000 vertex graph, the algorithm runs in about 2.6 seconds,
and for a 5000 vertex graph it takes over 6 minutes.

[This
link](https://goparallel.sourceforge.net/speed-up-floyd-warshall-algorithm-with-parallelization/)
suggests you can get a linear speedup by building with C++ and using
an OpenMP pragma (so on a 4 core machine, 5000 vertexes in ~ 1m15s).