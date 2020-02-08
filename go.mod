module go-gin-cli

go 1.13

require (
	github.com/denisenkom/go-mssqldb v0.0.0-20191124224453-732737034ffd // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.52.0
	github.com/lib/pq v1.1.1
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	go.mongodb.org/mongo-driver v1.3.0
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd // indirect
	gopkg.in/ini.v1 v1.52.0 // indirect
	xorm.io/xorm v0.8.1
)

replace (
	go-gin-cli/database => /Users/jensonsu/MyProject/go-gin-cli/database
	go-gin-cli/middleware => /Users/jensonsu/MyProject/go-gin-cli/middleware
	go-gin-cli/models => /Users/jensonsu/MyProject/go-gin-cli/models
	go-gin-cli/pkg => /Users/jensonsu/MyProject/go-gin-cli/pkg
	go-gin-cli/router => /Users/jensonsu/MyProject/go-gin-cli/router
)
