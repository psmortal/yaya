package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)
var number int
func init() {
	rand.Seed(time.Now().UnixNano())
	number = rand.Int()
}

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(x + strconv.Itoa(number)))
	})
	_ = http.ListenAndServe(":8080",nil)
	//engine := gin.Default()
	//engine.Handle("GET","/", func(context *gin.Context) {
	//
	//})

}
