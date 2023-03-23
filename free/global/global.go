package global

import (
	"gorm-demo/config"
	"sync"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 通用表数据
type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

var (
	GVA_CONFIG config.Server //配置信息
	GVA_VP     *viper.Viper  //配置文件读取库
	GVA_LOG    *zap.Logger   //日志库
	MY_SQL     *gorm.DB      //数据库
	RW_LOCK    sync.RWMutex  //读写锁
	LOCK       sync.Mutex    //互斥锁
)

const (
	ConfigEnv         = "GVA_CONFIG"
	ConfigDefaultFile = "config.yaml"
)
