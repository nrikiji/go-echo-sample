package config

import(
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"
)

var Database map[interface{}]interface{}

func SetDB(env string) {
	
	yml, err := ioutil.ReadFile("src/app/config/database.yml")
	if err != nil {
		panic(err)
	}

	t := make(map[interface{}]interface{})
	_ = yaml.Unmarshal([]byte(yml), &t)
	conn := t[env].(map[interface {}]interface {})

	Database = conn
}
