package global

import (
	"github.com/firwoodlin/letstrans/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB                  *gorm.DB
	GVA_REDIS               redis.UniversalClient
	GVA_CONFIG              config.Server
	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	lock                    sync.RWMutex
	GVA_Concurrency_Control = &singleflight.Group{}
	//GVA_Timer  timer.Timer = timer.NewTimerTask()
)
