
# go-echo-sample

goとechoを使った簡単なapiサーバーのサンプルコード。

動作確認バージョンはgo1.8+

### 操作

起動

```
$ start_server --port 8080 --pid-file app.pid -- ./bin/server
```

再起動　

```
$ kill -HUP `cat app.pid`
```

停止  

```
$ kill -TERM `cat app.pid`
```

### start_server

インストール
```
go get github.com/lestrrat/go-server-starter/cmd/start_server
```
