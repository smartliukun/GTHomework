package main

import (
	"database/sql"
	"example/gthomework/work1/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type ExampleApp struct {
	db *sql.DB
}

var exampleApp ExampleApp

func main() {
	fmt.Println("hello world!")

	var bannerId int64 = 1
	//eg.1 必须返回一条记录,否则返回err != nil
	log.Println("case 1 =====================================>")

	banner1, err1 := service.BannerServiceImpl.QueryExactlyOne(&bannerId)

	log.Println(fmt.Sprintf("banner1=%v,err1=%v", banner1, err1))

	//eg.2 如果没有数据，则返回空指针，且err == nil
	log.Println("case 2 =====================================>")

	banner2, err2 := service.BannerServiceImpl.QueryOneOrNot(&bannerId)

	log.Println(fmt.Sprintf("banner2=%v,err2=%v", banner2, err2))
}
