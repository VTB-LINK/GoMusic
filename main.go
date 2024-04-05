package main

import (
	"GoMusic/handler"
	"GoMusic/misc/log"
	"GoMusic/repo/cache"
	"GoMusic/repo/db"
	_ "GoMusic/repo/db"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func main() {
	initDSNFormEnv()
	initRdbCfgFormEnv()

	db.InitDB()
	cache.InitRedis()

	r := handler.NewRouter()
	if err := r.Run(getPortFromEnv()); err != nil {
		log.Errorf("fail to run server: %v", err)
		panic(err)
	}
}

func getPortFromEnv() string {
	if port := os.Getenv("HTTP_PORT"); port != "" {
		return ":" + port
	}
	return ":8081"
}

func initRdbCfgFormEnv() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		return
	}

	op := &redis.Options{
		Addr:     addr,
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	dbStr := os.Getenv("REDIS_DB")
	if dbStr != "" {
		dbNum, err := strconv.Atoi(dbStr)
		if err != nil {
			panic(err)
		}
		op.DB = dbNum
	}

	cache.RdbOptions = op
}

func initDSNFormEnv() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn != "" {
		db.DSN = dsn
	}
}
