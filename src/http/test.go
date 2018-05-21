package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {
	resp, err := http.Get("https://m.imooc.com")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	response, err := httputil.DumpResponse(resp, true)
	if err != nil{
		panic(err)
	}
	fmt.Print("%response\n",response)
}
