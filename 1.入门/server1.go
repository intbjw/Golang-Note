package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 把发送到/的请求和handler函数关联起来
	http.HandleFunc("/", handler)
	// 服务
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
