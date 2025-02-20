package global

import (
	"gorm.io/gorm"
	"github.com/go-redis/redis"

)
var (
	Db *gorm.DB
	RedisDB *redis.Client
)