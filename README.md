[![License](https://img.shields.io/badge/license-MIT-green)](https://github.com/longhaoteng/wineglass/blob/master/LICENSE)

## Wineglass

🍸🍹 Wineglass is minimalist scaffolding based on [gin](https://github.com/gin-gonic/gin) .

## Install
```shell
go get github.com/longhaoteng/wineglass
```

## Getting Started
```shell
export ENV=dev # dev,test,prod/release
export LOG_LEVEL=info # trace,debug,info,warn,error,fatal,panic
export VERSION=latest
export HTTP_ADDR=":8080"
export GRPC_ADDR=":50051"
export ALLOW_ORIGINS="*"

export ENABLE_PPROF=false # false,true
export DISABLE_DB=false # false,true
export DISABLE_AUTH=false # false,true
export DISABLE_REDIS=false # false,true

export LIMITER_STORE=memory # memory,cookie,redis
export LIMITER_LIMIT="10-S" # format:<limit>-<period>
# 5 reqs/second: "5-S"
# 10 reqs/minute: "10-M"
# 1000 reqs/hour: "1000-H"
# 2000 reqs/day: "2000-D"

export SESSION_STORE=memory # memory,cookie,redis
export SESSION_MAX_AGE=604800
export SESSION_SECRET="wineglass"
export SESSION_DB=0 # SESSION_STORE=redis
export SESSION_HTTP_ONLY=false # false,true

export REDIS_DB=0
export REDIS_ADDRS="localhost:6379"
export REDIS_PASSWORD=""
export REDIS_PREFIX=${serviceName}

export DB_HOST="localhost"
export DB_PORT=3306
export DB_USER="root"
export DB_PASSWORD=""
export DB_NAME=${serviceName}
export DB_LOW_THRESHOLD=0
export DB_MAX_OPEN_CONNS=100
export DB_MAX_IDLE_CONNS=25
```

```go
import (
    _ "github.com/longhaoteng/wineglass/_examples/api"
    "github.com/longhaoteng/wineglass/server"
)

func main() {
    // Init server
    server.Init(
        server.Name("helloworld"),
    )
    
    // Run server
    server.Run()
}
```

## [More examples](https://github.com/longhaoteng/wineglass/blob/master/_examples)

## License
[MIT License](https://github.com/longhaoteng/wineglass/blob/master/LICENSE)