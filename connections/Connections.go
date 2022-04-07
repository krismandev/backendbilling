package connections

import (
	"billingdashboard/config"
	"billingdashboard/lib"
	"context"

	redis "github.com/go-redis/redis/v7"
)

//Connections Holds all passing value to functions
type Connections struct {
	DBAppConn    *lib.DBConnection
	DBRedis      *redis.Client
	Context      context.Context
	JWTSecretKey string
}

//InitiateConnections is for Initiate Connection
func InitiateConnections(param config.Configuration) *Connections {
	var conn Connections
	conn.JWTSecretKey = param.JWTSecretKey

	// add redis connection
	conn.DBRedis = InitRedisConnection(param)

	// add mysql connection
	dbAppconn := lib.InitDB(param.DBList["app"].DBType, param.DBList["app"].DBUrl)
	conn.DBAppConn = &dbAppconn

	return &conn

}
