## 実行方法
popcount配下で下記を実行する
```
$ go test -bench=.
```

## ベンチマークテストの実行結果
練習問題2.4と比較すると1.5倍程度早い
```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/sh-tatsuno/go-training/ch02/2.5/popcount
BenchmarkPopCount-4     2000000000               0.27 ns/op
BenchmarkPopCount3-4    100000000               16.4 ns/op
BenchmarkPopCount4-4    100000000               11.2 ns/op
PASS
ok      github.com/sh-tatsuno/go-training/ch02/2.5/popcount     3.359s
```