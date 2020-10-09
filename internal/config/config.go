package config

import (
	"github.com/tinyhubs/properties"
	"go-web-template/pkg/logger"
	"go-web-template/pkg/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Conf struct {
	Env                 string
	Port                string `yaml:"port"`

}

var (
	conf     Conf
	env      string
	hostname string
)

func init() {
	//校验是否已经加载过配置
	if conf.Env != "" {
		return
	}
	//获取环境变量
	var env = getEnv()
	conf.Env = env
	//获取项目路径
	dir := utils.GetSourcePath()
	filePath := dir + "/../../config/config-" + env + ".yml"
	logger.Infow("load config file","filePath", filePath)
	//load config file
	ymlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Infow("yamlFile.Get err", "err", err)
		panic(err)
	}
	err = yaml.Unmarshal(ymlFile, &conf)
	if err != nil {
		logger.Infow("yamlFile Unmarshal", "err", err)
		panic(err)
	}
}

func GetConf() *Conf {
	return &conf
}

//Get env
func getEnv() string {
	if env != "" {
		return env
	}
	//获取环境变量
	env = "fat"
	envFilePath := os.Getenv("ENV_FILE_PATH")
	file, err := os.OpenFile(envFilePath, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		env = "fat"
	} else {
		doc, err := properties.Load(file)
		if nil != err {
			logger.Error("加载失败")
		} else {
			env = doc.String("env")
			hostname = doc.String("hostname")
		}
	}
	return env
}
