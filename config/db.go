package config

import(
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"
)

var Database map[interface{}]interface{}

func SetDB(env string) {
	yml, err := ioutil.ReadFile("src/app/config/db.yml")
	if err != nil {
		panic(err)
	}

	t := make(map[interface{}]interface{})
	yaml.Unmarshal([]byte(yml), &t)
	conf := t[env].(map[interface {}]interface {})
	Database = conf
}
