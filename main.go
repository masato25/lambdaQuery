package main

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
	"io/ioutil"
)

type Result struct {
	Dstype   string
	Step     int
	Endpoint string
	Counter  string
	Values   []*TimeSeriesData
}

type TimeSeriesData struct {
	Timestamp int
	Value     float32
}

func main() {
	var fakedataf []byte
	fakedataf, _ = ioutil.ReadFile("./test/testdata")
	var jdata string = string(fakedataf)
	var t []*Result
	json.Unmarshal([]byte(jdata), &t)
	dat, err := ioutil.ReadFile("./js/sumAll.js")
	if err != nil {
		fmt.Println(err)
	}
	codes := string(dat)
	//fmt.Println(codes)
	vm := otto.New()
	vm.Set("input", t)
	vm.Set("limit", 1)
	vm.Run(codes)
	x, _ := vm.Get("output")
	fmt.Printf("%v", x)
}
