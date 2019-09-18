package db

import (
	"fmt"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var logger = util.GetLogger("DB")
var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		logger.Warning(".env file not existing or bad format", e)
	}

	logger.Info("Initializing Database Connection...")

	username := os.Getenv("dbuser")
	password := os.Getenv("dbpass")
	dbName := os.Getenv("dbname")
	dbhost := os.Getenv("dbhost")
	if dbhost == "" {
		dbhost = "localhost"
	}
	dbport := os.Getenv("dbport")
	if dbport == "" {
		dbport = "32700"
	}

	dbUri := fmt.Sprintf("user = %s password = %s host = %s port = %s dbname = %s sslmode=disable",
		username,
		password,
		dbhost,
		dbport,
		dbName,
	)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		logger.Errorf("error while opening database connection: %v", err)
		os.Exit(-1)
	} else {
		logger.Info("connected to database")
	}
	db = conn.LogMode(true)

	// Redis
	//redisConn := GetRedis()
	//if redisConn == nil {
	//	os.Exit(-1)
	//}
	//if redis != nil {
	//	logger.Notice("flushing redis")
	//	_, err := redis.FlushDB().Result()
	//	if err != nil {
	//		logger.Warning("error while flushing redis", err)
	//	} else {
	//		logger.Notice("redis server is flushed")
	//	}
	//}
}

func GetDB() *gorm.DB {
	return db.Set("gorm:auto_preload", true)
}

var sharedRedis *redis.Client = nil

func GetRedis() *redis.Client {
	if sharedRedis == nil {
		client := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("redisAddr"),
			Password: os.Getenv("redisPasswd"),
			DB:       0,
		})

		_, err := client.Ping().Result()
		if err != nil {
			logger.Errorf("failed to connect redis: %v", err)
		} else {
			logger.Info("connected to redis server")
			sharedRedis = client
		}
	}

	return sharedRedis
}

func RedisGetKey(module string, key string) string {
	return fmt.Sprintf("code_sso:%s:%s", module, key)
}
