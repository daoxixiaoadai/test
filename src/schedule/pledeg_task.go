package main

import "github.com/daoxixiaoadai/test/src/db"

func main() {
	db.InitMysql()
	db.InitRedis()

	println("Hello, world!")
}
