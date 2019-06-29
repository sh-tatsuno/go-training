## 実行方法
popcount配下で下記を実行する
```
$ go test -bench=.
```

## ベンチマークテストの実行結果

およそ28倍ほどの差が生じる
```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/sh-tatsuno/go-training/ch02/2.3/popcount
BenchmarkPopCount-4     2000000000               0.54 ns/op
BenchmarkPopCount2-4    100000000               15.3 ns/op
PASS
ok      github.com/sh-tatsuno/go-training/ch02/2.3/popcount     2.685s
```