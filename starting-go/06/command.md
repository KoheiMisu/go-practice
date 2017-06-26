## Command

```bash
$ go version

$ go env

$ go fmt # 内部ではgofmtが使われている

$ go doc # パッケージのドキュメントを参照する
ex.) $ go doc math/rand  

$ go build # -x で冗長な出力. -o [name] でビルド結果をnameにする

$ go get [package] # phpでいうcomposer
ex.) $ go get golang.org/x/net/http2 # http2パッケージをインストール
$GOPATHのpkg配下に入る

```

## Vendor

$GOPATH/src/vendor  
vendor配下に置かれているパッケージは最後に読み込まれる  

そのため、packageのバージョン管理が可能となる