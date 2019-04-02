package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/v2af/file"
)

type CookieConfig struct {
	Secure bool   `json:"secure"`
	MaxAge int    `json:"max_age"`
	Damain string `json:"damain"`
}

type HttpConfig struct {
	Listen string `json:"listen"`
	Secret string `json:"secret"`
}

type MysqlConfig struct {
	Addr    string `json:"addr"`
	Idle    int    `json:"idle"`
	Max     int    `json:"max"`
	ShowSql bool   `json:"show_sql"`
}

type RedisTimeoutConfig struct {
	Conn  int64 `json:"conn"`
	Read  int64 `json:"read"`
	Write int64 `json:"write"`
}

type RedisConfig struct {
	Addr         string              `json:"addr"`
	Idle         int                 `json:"idle"`
	Max          int                 `json:"max"`
	RedisTimeout *RedisTimeoutConfig `json:"timeout"`
}

type CacheConfig struct {
	Provider    string        `json:"provider"`
	ExpireInt64 int64         `json:"expire"`
	Expire      time.Duration `json:"-"`
}

type QiniuConfig struct {
	Enabled   bool   `json:"enabled"`
	Bucket    string `json:"bucket"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Domain    string `json:"domain"`
}

type GlobalConfig struct {
	Debug    bool          `json:"debug"`
	Portrait string        `json:"portrait"`
	Motto    string        `json:"motto"`
	Cookie   *CookieConfig `json:"cookie"`
	Http     *HttpConfig   `json:"http"`
	Mysql    *MysqlConfig  `json:"mysql"`
	Cache    *CacheConfig  `json:"cache"`
	Redis    *RedisConfig  `json:"redis"`
	Upload   string        `json:"upload"`
	Qiniu    *QiniuConfig  `json:"qiniu"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func Parse(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("configuration file %s is nonexistent", cfg)
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		return fmt.Errorf("read configuration file %s fail %s", cfg, err.Error())
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		return fmt.Errorf("parse configuration file %s fail %s", cfg, err.Error())
	}

	c.Cache.Expire = time.Duration(c.Cache.ExpireInt64) * time.Second

	// if c.Qiniu != nil && c.Qiniu.Enabled {
	// 	conf.ACCESS_KEY = c.Qiniu.AccessKey
	// 	conf.SECRET_KEY = c.Qiniu.SecretKey
	// }

	config = &c

	log.Println("load configuration file", cfg, "successfully")
	return nil
}
