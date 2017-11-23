package config

import (
	"path"
	"strings"
	"io/ioutil"
	"errors"
	"github.com/dingdayu/golangtools/config/ini"
	"github.com/dingdayu/golangtools/config/yaml"
	"github.com/dingdayu/golangtools/config/json"
	"github.com/dingdayu/golangtools/config/xml"
	"github.com/dingdayu/golangtools/config/toml"
)

func New(f string, c interface{}) error  {

	buf, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	//fmt.Println(string(buf))

	fileSuffix := getFileSuffix(f)
	
	switch fileSuffix {
	case "json":
		err = json.Unmarshal(buf, c)
	case "yaml":
		err = yaml.Unmarshal(buf, c)
	case "xml":
		err = xml.Unmarshal(buf, c)
	case "ini":
		err = ini.Unmarshal(buf, c)
	case "toml":
		err = toml.Unmarshal(buf, c)
	default:
		err = errors.New("格式不支持")
	}
	return  err
}

// 自动获取文件后缀
func getFileSuffix (f string) string  {
	filePath := path.Base(f)
	fileSuffix := path.Ext(filePath)
	return strings.Trim(fileSuffix, ".")
}