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

var jdata string = `[{"dstype": "GAUGE", "step": 60, "endpoint": "test-zj-223", "Values": [{"timestamp": 1462347600, "value": null}, {"timestamp": 1462348800, "value": 99.649811}, {"timestamp": 1462350000, "value": 99.399342}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "test-zj-224", "Values": [{"timestamp": 1462347600, "value": 97.610744}, {"timestamp": 1462348800, "value": 98.423146}, {"timestamp": 1462350000, "value": 99.699843}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "test-zj-225", "Values": [{"timestamp": 1462347600, "value": 98.977213}, {"timestamp": 1462348800, "value": 99.474465}, {"timestamp": 1462350000, "value": 99.399874}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "test-sd-223", "Values": [{"timestamp": 1462347600, "value": 88.61222}, {"timestamp": 1462348800, "value": 90.017757}, {"timestamp": 1462350000, "value": 93.540823}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "test-jx", "Values": [{"timestamp": 1462347600, "value": 89.77683}, {"timestamp": 1462348800, "value": 93.723522}, {"timestamp": 1462350000, "value": 90.842514}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "test-ah-11", "Values": [{"timestamp": 1462347600, "value": 85.020967}, {"timestamp": 1462348800, "value": 92.762894}, {"timestamp": 1462350000, "value": 89.661321}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "test-js-183", "Values": [{"timestamp": 1462347600, "value": 61.734885}, {"timestamp": 1462348800, "value": 70.364942}, {"timestamp": 1462350000, "value": 66.678902}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "test-gd-184", "Values": [{"timestamp": 1462347600, "value": null}, {"timestamp": 1462348800, "value": null}, {"timestamp": 1462350000, "value": null}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "test-gx-111", "Values": [{"timestamp": 1462347600, "value": 90.75715}, {"timestamp": 1462348800, "value": 91.313158}, {"timestamp": 1462350000, "value": 93.267493}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "test-gd-189", "Values": [{"timestamp": 1462347600, "value": null}, {"timestamp": 1462348800, "value": null}, {"timestamp": 1462350000, "value": null}], "counter": "cpu.idle"}]`

func main() {
	var t []*Result
	json.Unmarshal([]byte(jdata), &t)
	dat, err := ioutil.ReadFile("./js/sum.js")
	if err != nil {
		fmt.Println(err)
	}
	codes := string(dat)
	//fmt.Println(codes)
	vm := otto.New()
	vm.Set("input", t)
	vm.Run(codes)
	x, _ := vm.Get("output")
	fmt.Printf("%v", x)
}
