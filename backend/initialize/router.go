package initialize

import (
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/middleware"
	"github.com/firwoodlin/letstrans/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	letsTransRouter := router.RouterGroupApp.LetsTrans

	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.GVA_LOG.Info("use middleware cors")
	//docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	//Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		//systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//systemRouter.InitInitRouter(PublicGroup)    // 自动初始化相关
		letsTransRouter.InitFileRouter(PublicGroup) // 文件上传下载功能
		letsTransRouter.InitProjectRouter(PublicGroup)
		letsTransRouter.InitSegmentRouter(PublicGroup)
		letsTransRouter.InitGlossaryRouter(PublicGroup)
		letsTransRouter.InitThirdPartyRouter(PublicGroup)
		letsTransRouter.InitTranslationMemoryRouter(PublicGroup)
	}
	//PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//{
	//
	//}

	global.GVA_LOG.Info("router register success")
	return Router
}
