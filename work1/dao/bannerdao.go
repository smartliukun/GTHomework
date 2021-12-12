package dao

import (
	"database/sql"
	"example/gthomework/work1/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	DBHost     = "127.0.0.1"
	DBUser     = "root"
	DBPassword = "123456"
	DBName     = "test"
)

type BannerDao struct {
	db *sql.DB
}

var BannerDaoImpl *BannerDao

func init() {
	var err error
	BannerDaoImpl = new(BannerDao)
	BannerDaoImpl.db, err = sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", DBUser, DBPassword, DBHost, DBName))
	if err != nil {
		log.Panic(err)
	}
}

// QueryExactlyOne 应该查到一条数据，否则返回err
func (e *BannerDao) QueryExactlyOne(bannerId *int64) (banner *model.Banner, err error) {
	banner = new(model.Banner)
	err = e.db.QueryRow("select banner_id,banner_name from banner where banner_id = ?", bannerId).Scan(&banner.BannerId, &banner.BannerName)
	return
}

// QueryOne 尝试查询一条数据，如果没有，则返回空指针
func (e *BannerDao) QueryOne(bannerId *int64) (banner *model.Banner, err error) {
	banner = new(model.Banner)
	err = e.db.QueryRow("select banner_id,banner_name from banner where banner_id = ?", bannerId).Scan(&banner.BannerId, &banner.BannerName)

	if err == sql.ErrNoRows {
		log.Println(fmt.Sprintf("there is no row where banner_id=%d", bannerId))
		banner = nil
		err = nil
	}

	return
}
