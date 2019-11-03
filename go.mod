module github.com/zw8677174/game_site

go 1.13

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.49.0
	github.com/jinzhu/gorm v1.9.11 // indirect
	github.com/unknwon/com v1.0.1 // indirect
	google.golang.org/appengine v1.6.5 // indirect
)

replace (
	github.com/zw8677174/game_site/conf => ./pkg/conf
	github.com/zw8677174/game_site/middleware => ./middleware
	github.com/zw8677174/game_site/models => ./models
	github.com/zw8677174/game_site/pkg/setting => ./pkg/setting
	github.com/zw8677174/game_site/routers => ./routers
)
