package model

import (
	"time"

	"github.com/masato25/lambdaQuery/database"
)

type Endpoint struct {
	Id       int64     `json:"id"`
	Endpoint string    `json:"endpoint"`
	Ts       int64     `json:"ts"`
	TCreate  time.Time `json:"-"`
	TModify  time.Time `json:"-"`
	Ipv4     string    `json:"-"`
}

func EndpointQuery() (endpointList []string) {
	database.Init()
	db := database.DBConn()
	var enps []Endpoint
	db.Raw("SELECT * from graph.endpoint").Scan(&enps)
	if len(enps) != 0 {
		for _, host := range enps {
			endpointList = append(endpointList, host.Endpoint)
		}
	}
	return
}
