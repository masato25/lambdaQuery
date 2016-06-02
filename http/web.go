package http

import (
	"time"

	cmodel "github.com/Cepave/common/model"
	// "github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/masato25/lambdaQuery/g"
	"github.com/masato25/lambdaQuery/http/computeFunc"
	"github.com/masato25/lambdaQuery/http/openFalcon"
)

type QueryInput struct {
	StartTs       time.Time
	EndTs         time.Time
	ComputeMethod string
	Endpoint      string
	Counter       string
}

//this function will generate query string obj for QueryRRDtool
func getq(q QueryInput) cmodel.GraphQueryParam {
	request := cmodel.GraphQueryParam{
		Start:     q.StartTs.Unix(),
		End:       q.EndTs.Unix(),
		ConsolFun: q.ComputeMethod,
		Endpoint:  q.Endpoint,
		Counter:   q.Counter,
	}
	return request
}

func StartWeb() {
	handler := gin.Default()
	compute := handler.Group("/func")
	conf := g.Config()
	compute.GET("/compute", computeFunc.Compute)
	compute.GET("/funcations", computeFunc.GetAvaibleFun)
	compute.GET("/smapledata", computeFunc.GetTestData)
	openfalcon := handler.Group("owl")
	openfalcon.GET("/endpoints", openFalcon.GetEndpoints)
	openfalcon.GET("/queryrrd", openFalcon.QueryData)
	// endless.ListenAndServe(":8888", handler)
	handler.Run(conf.Http.Listen)
}
