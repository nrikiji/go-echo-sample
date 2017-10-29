package redis

import(
	"fmt"
	"flag"
	"app/config"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

type Campaigns struct {
	Campaigns []Campaign `json:"campaigns"`
}

type Campaign struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

var conn redis.Conn
var err error

func Init() {

	c := config.Config.Redis

	var host = flag.String("hostname", c.Host, "Set Hostname")
	var port = flag.String("port", c.Port, "Set Port")

	flag.Parse()

	conn, err = redis.Dial("tcp", *host + ":" + *port)
	if err != nil {
		panic(err)
	}

	jsonStr, err := redis.String(conn.Do("GET", "campaigns"))
	if err != nil {
		fmt.Println("redis getでエラーが発生")
	}

	fmt.Println("%s", jsonStr)

	jsonBytes := ([]byte)(jsonStr)
	var campaigns Campaigns
	
	err = json.Unmarshal(jsonBytes, &campaigns)
	if err != nil {
		panic(err)
	}

	fmt.Println("%v", campaigns)
	
}

func GetConnection() redis.Conn {
	return conn
}

