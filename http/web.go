package http

import (
	"encoding/json"
	"io/ioutil"

	"log"
	"strings"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/masato25/lambdaQuery/conf"
	"github.com/masato25/lambdaQuery/model"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

func getFakeData() (t []*model.Result) {
	var fakedataf []byte
	fakedataf, _ = ioutil.ReadFile("./test/testdata")
	var jdata string = string(fakedataf)
	json.Unmarshal([]byte(jdata), &t)
	return
}

func StartWeb() {
	conf.GetAvaibleFun()
	r := gin.Default()
	r.GET("/funcations", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"funcations": conf.GetAvaibleFun(),
		})
	})

	r.GET("/data", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": getFakeData(),
		})
	})

	r.GET("/compute", func(c *gin.Context) {
		funkey := c.DefaultQuery("func", "")
		if funkey == "" {
			log.Println("Get params fun error")
		}

		funobj := conf.GetFunc(funkey)
		//fmt.Println(codes)
		vm := otto.New()
		vm.Set("input", getFakeData())
		var tmpparams []interface{}
		for _, params := range funobj.Params {
			ss := strings.Split(params, ":")
			tmpparams = append(tmpparams, ss[0])
			params_key := ss[0]
			params_type := ss[1]
			paramset := c.DefaultQuery(params_key, "")
			if paramset != "" {
				switch params_type {
				case "string":
					vm.Set(params_key, paramset)
					tmpparams = append(tmpparams, paramset)
				case "int":
					i, err := strconv.Atoi(paramset)
					if err != nil {
						log.Println(err.Error())
					} else {
						vm.Set(params_key, i)
						tmpparams = append(tmpparams, i)
					}
				}
			}
		}
		vm.Run(funobj.Codes)
		output, _ := vm.Get("output")
		c.JSON(200, gin.H{
			"compted_data": output.String(),
			"func":         funkey,
			"params":       tmpparams,
		})
	})
	r.Run(":8888")
}
