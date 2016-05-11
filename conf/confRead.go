package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Conf struct {
	FuncationName string
	FileName      string
	Params        []string
	Description 	string
}

func ReadConf() (conf *[]Conf){
	dat, err := ioutil.ReadFile("./lambdaSetup.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(dat, &conf)
	return conf
}
