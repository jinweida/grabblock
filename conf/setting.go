package conf

import (
	"encoding/json"
	"os"
	"time"
)

var Context *Config

// config
type Config struct {
	Mode     string   `json:"mode"`
	Name     string   `json:"name"`
	Db       Database `json:"database"`
	Server   Server   `json:"server"`
	Node     Node     `json:"node"`
	Influxdb Influxdb `json:"influxdb"`
	Redis    Redis    `json:"redis"`
}

type Redis struct {
	Host      string
	Port      int
	Maxidle   int
	Maxactive int
	Password  string
}

// db model
type Influxdb struct {
	Address       string `json:"address" from:"address"`
	Dbname        string `json:"dbname" from:"dbname"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	LogMode       bool   `json:"logmode"`
	BlockStart    string `json:"blockstart"`
	BlockDuration int64  `json:"blockduration"`
}

// db model
type Database struct {
	Address      string        `json:"address" from:"address"`
	Port         string        `json:"port" from:"port"`
	Dbname       string        `json:"dbname" from:"dbname"`
	Username     string        `json:"username"`
	Password     string        `json:"password"`
	Maxopenconns int           `json:"maxopenconns"`
	Maxidleconns int           `json:"maxidleconns"`
	Maxlifetime  time.Duration `json:"maxlifetime"`
	LogMode      bool          `json:logmode`
}

// server port
type Server struct {
	Port         int           `json:"port"`
	Readtimeout  time.Duration `json:"readtimeout"`
	Writetimeout time.Duration `json:"writetimeout"`
}

// node info
type Node struct {
	Url             string        `json:"url"`
	Count           int64         `json:"count"`
	Interval        time.Duration `json:"interval"`
	Proxy           string        `json:proxy`
	FailCount       int64         `json:failcount`
	SafeDistance    int64         `json:safedistance`
	Contractparsing bool          `json:contractparsing`
	Tokenparsing    bool          `json:tokenparsing`
	Contractsize    int64         `json:contractsize`
	IsEvfs          bool          `json:isevfs`
}

// parseconf
func ParseConf(config string) error {
	var c Config
	conf, err := os.Open(config)
	if err != nil {
		return err
	}
	err = json.NewDecoder(conf).Decode(&c)
	Context = &c
	return err
}
