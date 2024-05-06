package system

type ServiceGroup struct {
	JwtService
	UserService
	InitDBService
	SystemConfigService
	OperationRecordService
}
