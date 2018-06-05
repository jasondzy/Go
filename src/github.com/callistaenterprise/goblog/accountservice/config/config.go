package config

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/spf13/viper"
)

//这是一个加载config文件的主函数
func LoadConfigurationFromBranch(configServerUrl string, appName string, profile string, branch string) {
	
	url := fmt.Sprintf("%s/%s/%s/%s", configServerUrl, appName, profile, branch)
	fmt.Printf("Loading config from %s\n", url)

	body, err := fetchConfiguration(url) //调用函数获取config信息
	if err != nil {
		panic("Could not load configuration, cannot start. Terminating.Error: " + err.Error())
	}

	parserConfiguration(body) //对获取的信息进行json解析
}


func fetchConfiguration(url string) ([]byte, error) {
	resp, err := http.Get(url)  //通过 http的get函数来获取远程的config信息
	if err != nil {
		panic("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func parserConfiguration(body []byte) {
	var cloudConfig springCloudConfig

	err := json.Unmarshal(body, &cloudConfig) //json解析的关键就是将json数据用struct进行描述即可
	if err != nil {
		 panic("Cannot parse configuration, message: " + err.Error())
	}

	for key, value := range cloudConfig.PropertySources[0].Source {
		viper.Set(key, value)
		fmt.Printf("Loading config property %v => %v\n", key, value)
	}

	if viper.IsSet("server_name") {
		 fmt.Printf("Successfully loaded configuration for service %s\n", viper.GetString("server_name"))
	}

}

//如下的两个结构体是对json中的数据结构进行描述，这样就能够直接解析出结构体数据
type springCloudConfig struct {
	Name			string				`jason:"name"`
	Profiles		[]string			`jason:"profiles"`
	Label			string				`jason:"label"`
	Version			string				`jason:"version"`
	PropertySources	[]propertySource	`jason:"PropertySources"`
}

type propertySource struct {
	Name	string					`jason:"name"`
	Source	map[string]interface{}	`jason:"source"`
}

