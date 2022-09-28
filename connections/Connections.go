package connections

import (
	"backendbillingdashboard/config"
	"backendbillingdashboard/lib"
	"context"

	redis "github.com/go-redis/redis/v7"
)

//Connections Holds all passing value to functions
type Connections struct {
	DBAppConn     *lib.DBConnection
	DBDashbConn   *lib.DBConnection
	DBOcsConn     *lib.DBConnection
	DBRedis       *redis.Client
	Context       context.Context
	ManagementUrl string
	JWTSecretKey  string
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

	dbDashbconn := lib.InitDB(param.DBList["dashb"].DBType, param.DBList["dashb"].DBUrl)
	conn.DBDashbConn = &dbDashbconn

	dbOcsbconn := lib.InitDB(param.DBList["ocs"].DBType, param.DBList["ocs"].DBUrl)
	conn.DBOcsConn = &dbOcsbconn

	conn.ManagementUrl = param.ManagementUrl

	return &conn

}
