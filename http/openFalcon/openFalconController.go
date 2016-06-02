package openFalcon

import (
	"time"

	"strconv"

	cmodel "github.com/Cepave/common/model"
	"github.com/gin-gonic/gin"
	"github.com/masato25/lambdaQuery/graph"
	"github.com/masato25/lambdaQuery/model"
)

func GetEndpoints(c *gin.Context) {
	enps := model.EndpointQuery()
	c.JSON(200, gin.H{
		"status": "ok",
		"data": map[string][]string{
			"endpoints": enps,
		},
	})
}

func getEndpoints(endpoint string) {

}

func QDataGet(c *gin.Context) []*cmodel.GraphQueryResponse {
	startTmp := c.DefaultQuery("startTs", string(time.Now().Unix()-(86400)))
	startTmp2, _ := strconv.Atoi(startTmp)
	startTs := int64(startTmp2)
	endTmp := c.DefaultQuery("endTs", string(time.Now().Unix()))
	endTmp2, _ := strconv.Atoi(endTmp)
	endTs := int64(endTmp2)
	consolFun := c.DefaultQuery("consolFun", "AVERAGE")
	counter := c.DefaultQuery("counter", "cpu.idle")
	endpoints := model.EndpointQuery()
	var result []*cmodel.GraphQueryResponse
	for _, enp := range endpoints {
		q := cmodel.GraphQueryParam{
			Start:     startTs,
			End:       endTs,
			ConsolFun: consolFun,
			Endpoint:  enp,
			Counter:   counter,
		}
		res, _ := graph.QueryOne(q)
		result = append(result, res)
	}
	return result
}

func QueryData(c *gin.Context) {
	result := QDataGet(c)
	c.JSON(200, gin.H{
		"status": "ok",
		"data":   result,
	})
}
