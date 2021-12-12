package api

var (
	// SUCCESS 成功
	SUCCESS = RetEnm{0, "success"}
	// RECORD_NOT_FOUND 记录缺失
	RECORD_NOT_FOUND = RetEnm{201, "record not found"}
	// UNKNOW_ERROR 未知异常
	UNKNOW_ERROR = RetEnm{999, "unknow error"}
)

type RetEnm struct {
	Code int
	Msg  string
}
