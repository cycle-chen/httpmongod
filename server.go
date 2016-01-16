package main

import (
    "log"
    "net/http"
    "os"

    "github.com/golangframework/httpmongo"
    "github.com/golangframework/moeregexp"
)

var (
    root    = ""
    mongodb = "127.0.0.1:27017"
)

func main() {
    //检查根目录
    root, _ = os.Getwd()
    var mux = http.NewServeMux()
    mux.HandleFunc("/", router)
    err := http.ListenAndServe(":8090", mux)
    log.Println("http.ListenAndServe(:8090)")
    if err != nil {
        log.Fatal("http.ListenAndServe:", err.Error())
    }
}
func router(w http.ResponseWriter, r *http.Request) {
    urlpath := r.URL.Path
  //路由匹配正则 "^/mongo.+"
    if moeregexp.IsMatch(httpmongo.Mongo_path, urlpath) {
      //调用handler_mongo，处理 /mongo路由下的所有请求
        httpmongo.Httphandler_mongo(mongodb, w, r)
    }
}