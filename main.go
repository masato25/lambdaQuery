package main

import (
	"github.com/masato25/lambdaQuery/conf"
	"github.com/masato25/lambdaQuery/database"
	"github.com/masato25/lambdaQuery/g"
	"github.com/masato25/lambdaQuery/graph"
	"github.com/masato25/lambdaQuery/http"
)

func main() {
	conf.ReadConf("./conf/lambdaSetup.json")
	g.ParseConfig("./cfg.json")
	database.Init()
	graph.Start()
	http.StartWeb()
}
