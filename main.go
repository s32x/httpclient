package main

import "github.com/starboy/httpclient"

func main() {
	bytes, err := httpclient.GetString("https://api.github.com/users/starboy/repos")
	if err != nil {
		panic(err)
	}
	println(bytes)
}
