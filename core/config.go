package core

import (
	"fast_gin/config"
	"fast_gin/flags"
	"fast_gin/global"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// 读取配置文件
func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile(flags.Options.File)
	if err != nil {
		logrus.Fatalf("读取配置文件失败 %s", err)
		return
	}
	err = yaml.Unmarshal(byteData, cfg) //将字节流转化为结构化对象
	if err != nil {
		logrus.Fatalf("配置文件格式错误%s", err)
		return
	}
	logrus.Infof("%s配置文件读取成功", flags.Options.File)
	return
}

// 写入配置文件
func DumpConfig() {
	byteData, err := yaml.Marshal(global.Config) //将结构化对象转化为字节流写入文件
	if err != nil {
		logrus.Errorf("配置文件转换错误%s", err)
		return
	}
	err = os.WriteFile(flags.Options.File, byteData, 0666)
	if err != nil {
		logrus.Errorf("写入配置文件失败%s", err)
		return
	}
	logrus.Infof("配置文件写入成功")
}
