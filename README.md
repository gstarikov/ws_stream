# ws_stream

```
BenchmarkStdJson-2                300000              3913 ns/op            1088 B/op          8 allocs/op
BenchmarkFastJson-2               300000              4279 ns/op            1760 B/op         13 allocs/op
BenchmarkGobJson-2                 50000             37047 ns/op            7904 B/op        195 allocs/op
BenchmarkJsoniter-2               300000              3492 ns/op            1624 B/op         17 allocs/op
BenchmarkSimpleCopy-2           50000000                33.7 ns/op             0 B/op          0 allocs/op
BenchmarkJsonparser-2             500000              2864 ns/op             992 B/op         13 allocs/op
BenchmarkEasyjson-2              2000000               645 ns/op             136 B/op          2 allocs/op
BenchmarkCapnproto-2             1000000              1853 ns/op             824 B/op         13 allocs/op
```