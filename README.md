# go_eagle_service simple

>
> 简单体验来go-zero框架生成项目
>

# build

### Go 1.15 及之前版本

```shell
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/zeromicro/go-zero/tools/goctl@latest
```

### Go 1.16 及以后版本

```shell
GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest
```

### For Mac

```shell
brew install goctl

goctl api new go_eagle_service
cd go_eagle_service
go mod init
go mod tidy
go run main.go -f etc/dev-api.yaml
```

```shell
curl -i http://localhost:8888/from/you


HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Traceparent: 00-281f59e2749fc89ab521d549908245e3-0f2c3d02f0eb0dfa-00
Date: Mon, 18 Jul 2022 03:55:52 GMT
Content-Length: 24

{"message":"Hello, you"}
```

# go evn

```shell
GO111MODULE=""
GOARCH="arm64"
GOBIN="/Users/edwin/gopath/bin"
GOCACHE="/Users/edwin/Library/Caches/go-build"
GOENV="/Users/edwin/Library/Application Support/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="arm64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/edwin/gopath/pkg/mod"
GONOPROXY="*.inke.cn"
GONOSUMDB="*.inke.cn"
GOOS="darwin"
GOPATH="/Users/edwin/gopath"
GOPRIVATE="*.inke.cn"
GOPROXY="direct"
GOROOT="/opt/homebrew/Cellar/go@1.16/1.16.10/libexec"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/opt/homebrew/Cellar/go@1.16/1.16.10/libexec/pkg/tool/darwin_arm64"
GOVCS=""
GOVERSION="go1.16.10"
GCCGO="gccgo"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/dev/null"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -arch arm64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/k1/y7_lmrn152j4d1t5q4g_vmqr0000gn/T/go-build592013982=/tmp/go-build -gno-record-gcc-switches -fno-common"
```