package caches

import (
	"fmt"
	"github.com/astaxie/beego/utils"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis"
	"sync"
)

type RedisConnConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

var Cache *redis.Client
var mu sync.Mutex

func NewCache() *redis.Client {
	if Cache != nil {
		return Cache
	}
	mu.Lock()
	defer mu.Unlock()
	if Cache != nil {
		return Cache
	}

	return RegisterRedis()
}

func loadRedisConfig() RedisConnConfig {
	redisConf := getConfPath()
	if !utils.FileExists(redisConf) {
		panic("Redis config not exist")
	}

	cfg, _ := ini.Load(redisConf)
	section := cfg.Section("default")
	host, _ := section.GetKey("host")
	port, _ := section.GetKey("port")
	password, _ := section.GetKey("password")
	db, _ := section.GetKey("db")
	dbIndex, _ := db.Int()

	return RedisConnConfig{
		Host:     host.String(),
		Port:     port.String(),
		Password: password.String(),
		DB:       dbIndex,
	}
}

func RegisterRedis() *redis.Client {
	config := loadRedisConfig()
	host := fmt.Sprintf("%s:%s", config.Host, config.Port)
	Cache = redis.NewClient(&redis.Options{
		Network:    "tcp",
		Addr:       host,
		Password:   config.Password,
		DB:         config.DB,
		MaxRetries: 3,
		PoolSize:   32,
	})

	return Cache
}

func getConfPath() string {
	return "conf/databases/redis.ini"
}
