package define

// 状态码
const (
	CodeSuccess      = 0
	CodeDBError      = 1
	CodeCacheError   = 2
	CodeParamError   = 3
	CodeUnknownError = 4
)

const (
	NumberZero = 0
	ServerName = "userapisrv"
)

// 错误MSG
const (
	ErrParam = "參數錯誤"
	ErrDB    = "數據庫操作出錯"
	ErrCache = "緩存操作出錯"
	StatusOk = "successful"
)

// 数据库参数
const (
	// DB NAME
	DBExample = "example"

	// Table NAME
	TableUserUser       = "user_user"
	TableRole           = "role"
	TablePermission     = "permission"
	TableUserRole       = "user_role"
	TableRolePermission = "role_permission"
)

// 年月日常量
const (
	YMDLayout         = "2006-01-02"
	YMDHMSLayout      = "2006-01-02 15:04:05"
	LayoutDayForCache = "20060102"
)
