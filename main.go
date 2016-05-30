package main

import (
	"github.com/masato25/lambdaQuery/conf"
	"github.com/masato25/lambdaQuery/http"
)

func main() {
	conf.ReadConf("./conf/lambdaSetup.json")
	http.StartWeb()
}
