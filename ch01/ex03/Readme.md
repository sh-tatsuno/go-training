## 実行方法
下記を実行する
```
$ go test -bench=.
```

## ベンチマークテストの実行結果

およそ1.87倍ほどの差が生じる
```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/sh-tatsuno/go-training/ch01/ex03
BenchmarkEchoSlow-4     10000000               133 ns/op
BenchmarkEchoFast-4     20000000                73.1 ns/op
PASS
ok      github.com/sh-tatsuno/go-training/ch01/ex03     3.025s
```