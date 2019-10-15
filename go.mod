module github.com/zw8677174/game_site

go 1.13

require github.com/gin-gonic/gin v1.4.0


replace (
        github.com/EDDYCJY/go-gin-example/pkg/setting => ~/go-application/go-gin-example/pkg/setting
        github.com/EDDYCJY/go-gin-example/conf          => ~/go-application/go-gin-example/pkg/conf
        github.com/EDDYCJY/go-gin-example/middleware  => ~/go-application/go-gin-example/middleware
        github.com/EDDYCJY/go-gin-example/models       => ~/go-application/go-gin-example/models
        github.com/EDDYCJY/go-gin-example/routers       => ~/go-application/go-gin-example/routers
)
