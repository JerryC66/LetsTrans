package system

type RouterGroup struct {
	JwtRouter
	SysRouter
	InitRouter
	UserRouter
	OperationRecordRouter
}
