package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

var Conf Config = Config{}

type Config struct {
	Ftp struct {
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		LiveTime int    `yaml:"liveTime"`
	} `yaml:"ftp"`

	Ssh struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Port     int    `yaml:"port"`
		LiveTime int    `yaml:"liveTime"`
	} `yaml:"ssh"`

	App struct {
		VideoCount  int    `yaml:"videoCount"`
		PlayUrlPre  string `yaml:"playUrlPre"`
		CoverUrlPre string `yaml:"coverUrlPre"`
	} `yaml:"app"`

	Database struct {
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
		Addr   string `yaml:"addr"`
		Port   string `yaml:"port"`
		Dbname string `yaml:"dbname"`
	} `yaml:"database"`

	Security struct {
		KeyString string `yaml:"keyString"`
		JwtKey    string `yaml:"jwtKey"`
	} `yaml:"security"`

	Redis struct {
		Addr string `yaml:"addr"`
		Pass string `yaml:"pass"`
	} `yaml:"redis"`
}

func InitConf() {
	remoteConfigURL := "http://39.101.72.240:9002/config/configration.yaml"

	resp, err := http.Get(remoteConfigURL)

	if err != nil {
		fmt.Println("获取配置文件失败")
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败")
	}
	// 解析yaml到结构体

	Conf = Config{}

	err = yaml.Unmarshal(body, &Conf)
	if err != nil {
		fmt.Println("Error unmarshaling YAML", err)
	}
	fmt.Println(Conf.Database.Port)

}
