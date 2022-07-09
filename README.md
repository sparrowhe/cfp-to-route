# cfp-to-route

提供 CFP 航路转为航路点的工具。

---

# Benchmark

```
goos: linux
goarch: amd64
pkg: cfptoroute/internal/service
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
Benchmark_ParseCFPRoute-8                1022570              1109 ns/op            1088 B/op          5 allocs/op
Benchmark_SegmentToPointsList-8               12          94625083 ns/op          127270 B/op       3987 allocs/op
Benchmark_MemoryConvert-8               252090301                4.420 ns/op           0 B/op          0 allocs/op
PASS
ok      cfptoroute/internal/service     5.440s
```