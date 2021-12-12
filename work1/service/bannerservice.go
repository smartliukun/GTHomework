package service

import (
	"database/sql"
	"example/gthomework/work1/api"
	"example/gthomework/work1/dao"
	"example/gthomework/work1/model"
)

type BannerService struct {
	bannerDao *dao.BannerDao
}

var BannerServiceImpl *BannerService

func init() {
	BannerServiceImpl = new(BannerService)
	BannerServiceImpl.bannerDao = dao.BannerDaoImpl
}

func (b *BannerService) QueryExactlyOne(bannerId *int64) (banner *model.Banner, err error) {
	banner, err = b.bannerDao.QueryExactlyOne(bannerId)
	if err == sql.ErrNoRows {
		err = api.BizError{Ret: api.RECORD_NOT_FOUND, Err: err}
		return
	}
	// fmt.Println(fmt.Sprintf("%v", banner))
	// 处理业务逻辑。。。
	return
}

func (b *BannerService) QueryOneOrNot(bannerId *int64) (banner *model.Banner, err error) {
	banner, err = b.bannerDao.QueryOne(bannerId)
	if err != nil {
		err = api.BizError{Ret: api.UNKNOW_ERROR, Err: err}
		return
	}
	// fmt.Println(fmt.Sprintf("%v", banner))
	// 处理业务逻辑。。。
	return
}
