module github.com/zw8677174/game_site

go 1.13

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.48.0 // indirect
	github.com/jinzhu/gorm v1.9.11 // indirect
	github.com/unknwon/com v1.0.1 // indirect
	google.golang.org/appengine v1.6.5 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => ./pkg/conf
	github.com/EDDYCJY/go-gin-example/middleware => ./middleware
	github.com/EDDYCJY/go-gin-example/models => ./models
	github.com/EDDYCJY/go-gin-example/pkg/setting => ./pkg/setting
	github.com/EDDYCJY/go-gin-example/routers => ./routers
)
