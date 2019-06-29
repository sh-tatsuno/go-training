## 実行方法
popcount配下で下記を実行する
```
$ go test -bench=.
```

## ベンチマークテストの実行結果

およそ31倍ほどの差が生じる
```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/sh-tatsuno/go-training/ch02/2.4/popcount
BenchmarkPopCount-4     2000000000               0.54 ns/op
BenchmarkPopCount3-4    100000000               16.6 ns/op
PASS
ok      github.com/sh-tatsuno/go-training/ch02/2.4/popcount     2.811s
```