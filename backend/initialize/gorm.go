package initialize

import (
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/model/letstrans"
	"github.com/firwoodlin/letstrans/model/system"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "pgsql":
		return GormPgSql()
	default:
		return GormPgSql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		//system.SysApi{},
		system.SysUser{},
		//system.SysBaseMenu{},
		//system.JwtBlacklist{},
		//system.SysAuthority{},
		//system.SysDictionary{},
		//system.SysOperationRecord{},
		//system.SysAutoCodeHistory{},
		//system.SysDictionaryDetail{},
		//system.SysBaseMenuParameter{},
		//system.SysBaseMenuBtn{},
		//system.SysAuthorityBtn{},
		//system.SysAutoCode{},
		//system.SysExportTemplate{},
		//system.Condition{},
		//system.JoinTemplate{},

		//example.ExaFile{},
		//example.ExaCustomer{},
		//example.ExaFileChunk{},
		//example.ExaFileUploadAndDownload{},
		letstrans.FileRecord{},
		letstrans.Document{},
		letstrans.Project{},
		letstrans.Segment{},
		letstrans.Glossary{},
		letstrans.Term{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
