package core

import (
	"fmt"
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/initialize"
	//"github.com/firwoodlin/letstrans/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	//if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
	//	// 初始化redis服务
	//	initialize.Redis()
	//}
	//if global.GVA_CONFIG.System.UseMongo {
	//	err := initialize.Mongo.Initialization()
	//	if err != nil {
	//		zap.L().Error(fmt.Sprintf("%+v", err))
	//	}
	//}
	//// 从db加载jwt数据
	//if global.GVA_DB != nil {
	//	system.LoadAll()
	//}

	Router := initialize.Routers()
	//Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	//	fmt.Printf(`
	//	欢迎使用 gin-vue-admin
	//	当前版本:v2.6.2
	//    加群方式:微信号：shouzi_1994 QQ群：470239250
	//	插件市场:https://plugin.gin-vue-admin.com
	//	GVA讨论社区:https://support.qq.com/products/371961
	//	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	//	默认前端文件运行地址:http://127.0.0.1:8080
	//	如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/coffee/index.html
	//`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
