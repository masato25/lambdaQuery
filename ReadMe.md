# LambdaQuery. 可擴充運算聚合框架.
![](./lambdaQuery.jpg)
* 支援RRDTOOL data / open-falcon graph

### 如何擴充
1. 定義自己的js聚合計算.
2. 設定 `conf/lambdaSetup.json`
  * funcation_name 名字
  * file_path 位址
  * params 可定義傳遞參數
    * "params_name:type"
  * description 描述

### 如何觀看demo
  `go get -u github.com/masato25/lambdaQuery`
```
cd $GOPATH/src/github.com/masato25/lambdaQuery
cp cfg.example.json cfg.json
開啟bowser: http://localhost:8888/func/compute?funcName=top&orderby=desc&limit=10&source=fake
```
