package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Conf struct {
	FuncationName string
	FileName      string
	Params        map[string]string
}

func ReadConf() (conf *[]Conf){
	dat, err := ioutil.ReadFile("./funjs.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(dat, &conf)
	return conf
}
