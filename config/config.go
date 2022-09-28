package config

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type TDB struct {
	DBUrl  string `yaml:"dbUrl"`
	DBType string `yaml:"dbType"`
}

// Configuration stores global configuration
type Configuration struct {
	Env                      string         `yaml:"env"`
	ListenPort               string         `yaml:"listenPort"`
	DBList                   map[string]TDB `yaml:"dblist"`
	ManagementUrl            string         `yaml:"managementUrl"`
	AppName                  string         `yaml:"appName"`
	UseJWT                   bool           `yaml:"useJWT"`
	JWTSecretKey             string         `yaml:"JWTSecretKey"`
	RequestTimeout           int            `yaml:"requestTimeout"`
	MaxConcurrentProcessData int            `yaml:"maxConcurrentProcessData"`
	ConcurrentWaitLimit      int            `yaml:"concurrentWaitLimit"`
	MaxBodyLogLength         int            `yaml:"maxBodyLogLength"`

	UseRedis         bool `yaml:"useRedis"`
	UseRedisSentinel bool `yaml:"useRedisSentinel"`
	RedisSentinel    struct {
		MasterName       string   `yaml:"masterName"`
		SentinelPassword string   `yaml:"sentinelPassword"`
		SentinelURL      []string `yaml:"sentinelUrl"`
	} `yaml:"redisSentinel"`
	Redis struct {
		RedisURL      string `yaml:"redisUrl"`
		RedisPassword string `yaml:"redisPassword"`
		DB            int    `yaml:"db"`
	} `yaml:"redis"`
	Log struct {
		FileNamePrefix string `yaml:"filenamePrefix"`
		FileName       string
		Level          string `yaml:"level"`
	} `yaml:"log"`
}

//Param is use for
var Param Configuration

//LoadConfig is use for
func LoadConfig(fn *string) {

	if err := LoadYAML(fn, &Param); err != nil {
		log.Errorf("LoadConfigFromFile() - Failed opening config file %s\n%s", &fn, err)
		os.Exit(1)
	}
	//log.Logf("Loaded configs: %v", Param)
	t := time.Now()
	sDate := fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
	if Param.Env == "local" {
		Param.Log.FileName = Param.Log.FileNamePrefix + sDate + ".log"
	} else {
		Param.Log.FileName = Param.Log.FileNamePrefix + ".log"
	}
}
