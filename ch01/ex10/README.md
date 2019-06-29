# fetchの2回実行について

```
go run fetchall.go https://go-tour-jp.appspot.com/welcome/1 https://go-tour-jp.appspot.com/welcome/1             18:50:37  ☁  master ☀
0.38s    1617 https://go-tour-jp.appspot.com/welcome/1
0.41s    1617 https://go-tour-jp.appspot.com/welcome/1
0.41s elapsed
```

```
go run fetchall.go https://play.golang.org/ https://play.golang.org/                                           18:52:08  ☁  master ☂ ✭
0.42s    6013 https://play.golang.org/
0.65s    6013 https://play.golang.org/
0.65s elapsed
```

```
go run fetchall.go https://play.golang.org/ https://play.golang.org/                                           18:52:18  ☁  master ☂ ✭
1.44s    6013 https://play.golang.org/
1.44s    6013 https://play.golang.org/
1.44s elapsed
```

```
go run fetchall.go https://play.golang.org/ https://play.golang.org/                                           18:52:23  ☁  master ☂ ✭
0.68s    6013 https://play.golang.org/
0.68s    6013 https://play.golang.org/
0.68s elapsed
```