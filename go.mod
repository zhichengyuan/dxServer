module dx.com

go 1.16

replace base.com => /Users/nn/Desktop/newversion/server/dx-baseserver //工程外的包

require (
	base.com v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/kataras/iris/v12 v12.1.8
)
