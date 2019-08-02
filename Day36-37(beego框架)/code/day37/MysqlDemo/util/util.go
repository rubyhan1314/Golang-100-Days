package util

import (
	"encoding/json"
	"github.com/astaxie/beego"
	//redis缓存引擎
	_ "github.com/astaxie/beego/cache/redis"
	//引入缓存模块
	"github.com/astaxie/beego/cache"
	"os"
)

const PAGELIMIT = 20

/**
 * 获取redis连接实例
 */
func GetRedis() (adapter cache.Cache, err error) {

	redisKey := beego.AppConfig.String("rediskey")
	redisAddr := beego.AppConfig.String("redisaddr")
	redisPort := beego.AppConfig.String("redisport")
	redisdbNum := beego.AppConfig.String("redisdbnum")

	redis_config_map := map[string]string{
		"key":   redisKey,
		"conn":  redisAddr + ":" + redisPort,
		"dbNum": redisdbNum,
	}
	redis_config, _ := json.Marshal(redis_config_map)

	cache_conn, err := cache.NewCache("redis", string(redis_config))
	if err != nil {
		return nil, err
	}
	return cache_conn, nil
}

/**
 * 判断当前path是否存在的工具方法
 */
func IsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
 *json格式数据转换为实体对象
 */
func JsonToEntity(data []byte, object interface{}) error {
	if len(data) <= 0 {
		return nil
	}
	return json.Unmarshal(data, object)
}

//向Map中存放数据
func PutParamToMap(mapp map[string]interface{}, key string, value interface{}) (map[string]interface{}) {
	mapp[key] = value
	return mapp
}

/**
 * 根据开发模式进行判断是否输出日志
 */
func LogInfo(v ...interface{}) {

	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		beego.Info(v)
	}
}

/**
 * 错误
 */
func LogError(v ...interface{}) {
	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		beego.Error(v)
	}
}

func LogWarn(v ...interface{}) {
	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		beego.Warn(v)
	}
}

func LogDebug(v ...interface{}) {
	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		beego.Debug(v)
	}
}

func LogNotice(v ...interface{}) {
	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		beego.Notice(v)
	}
}
