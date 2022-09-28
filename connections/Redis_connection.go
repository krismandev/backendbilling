package connections

import (
	"backendbillingdashboard/config"
	"os"
	"time"

	redis "github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
)

func InitRedisConnection(param config.Configuration) *redis.Client {
	var dbredis *redis.Client

	redisType := ""
	if param.UseRedisSentinel {
		log.Infof("Connecting Redis Sentinel Master "+param.RedisSentinel.MasterName+" : %v", param.RedisSentinel.SentinelURL)
		redisType = "redis sentinel"
		dbredis = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    param.RedisSentinel.MasterName,
			SentinelAddrs: param.RedisSentinel.SentinelURL,
			DialTimeout:   time.Duration(param.RequestTimeout) * time.Second,
		})

	} else if param.UseRedis {
		log.Infof("Connecting Redis : %v", param.Redis.RedisURL)
		redisType = "redis"
		dbredis = redis.NewClient(&redis.Options{
			Addr:     param.Redis.RedisURL,
			Password: param.Redis.RedisPassword,
			DB:       param.Redis.DB,
		})
	}

	if param.UseRedis || param.UseRedisSentinel {
		dbStatus := dbredis.Ping()
		if dbStatus.Err() != nil {
			log.Errorf("Error connecting to redis : %v", dbStatus.Err().Error())
			log.Errorf("Unable to connect Redis server %v", dbStatus.Err())
			os.Exit(1)
		}
		log.Infof("Connected to " + redisType)
	} else {
		log.Infof("App doesnt use redis")
	}

	return dbredis
}
