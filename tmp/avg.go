package main

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
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

type User struct {
	Name  string
	Phone string
}

func main() {
	jdata := `[{"dstype": "GAUGE", "step": 60, "endpoint": "cmb-zj-223-094-095-185", "Values": [{"timestamp": 1462347600, "value": null}, {"timestamp": 1462348800, "value": 99.649811}, {"timestamp": 1462350000, "value": 99.399342}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "cmb-zj-223-094-095-187", "Values": [{"timestamp": 1462347600, "value": 97.610744}, {"timestamp": 1462348800, "value": 98.423146}, {"timestamp": 1462350000, "value": 99.699843}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "cmb-zj-223-094-095-186", "Values": [{"timestamp": 1462347600, "value": 98.977213}, {"timestamp": 1462348800, "value": 99.474465}, {"timestamp": 1462350000, "value": 99.399874}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "cmb-sd-223-099-243-165", "Values": [{"timestamp": 1462347600, "value": 88.61222}, {"timestamp": 1462348800, "value": 90.017757}, {"timestamp": 1462350000, "value": 93.540823}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "cmb-jx-117-169-016-150", "Values": [{"timestamp": 1462347600, "value": 89.77683}, {"timestamp": 1462348800, "value": 93.723522}, {"timestamp": 1462350000, "value": 90.842514}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "cmb-ah-112-029-150-012", "Values": [{"timestamp": 1462347600, "value": 85.020967}, {"timestamp": 1462348800, "value": 92.762894}, {"timestamp": 1462350000, "value": 89.661321}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "cmb-js-183-213-022-038", "Values": [{"timestamp": 1462347600, "value": 61.734885}, {"timestamp": 1462348800, "value": 70.364942}, {"timestamp": 1462350000, "value": 66.678902}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "cmb-gd-183-232-066-078", "Values": [{"timestamp": 1462347600, "value": null}, {"timestamp": 1462348800, "value": null}, {"timestamp": 1462350000, "value": null}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "cmb-gx-111-012-014-147", "Values": [{"timestamp": 1462347600, "value": 90.75715}, {"timestamp": 1462348800, "value": 91.313158}, {"timestamp": 1462350000, "value": 93.267493}], "counter": "cpu.idle"}, {"dstype": "GAUGE", "step": 60, "endpoint": "cmb-gd-183-232-066-051", "Values": [{"timestamp": 1462347600, "value": null}, {"timestamp": 1462348800, "value": null}, {"timestamp": 1462350000, "value": null}], "counter": "cpu.idle"}]`
	var t []*Result
	json.Unmarshal([]byte(jdata), &t)
	fmt.Sprintf("%v", t)
	vm := otto.New()
	vm.Set("t", t)
	vm.Run(`
	    abc = 2 + 2;
			console.log("The value of abc is " + abc); // 4
			t2 = _.map(t, function(res){
				res.avg = _.reduce(res.Values, function(sum,v){
					return (sum+v.Value)
				},0) / (res.Values.length === 0 ? 1 : res.Values.length)
				return res;
			}) 
			
			_.each(t2, function(x){
				console.log(x.Endpoint, x.avg)
			})
	`)
}
