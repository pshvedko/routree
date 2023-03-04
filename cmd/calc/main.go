package main

import "github.com/pshvedko/routree/router"

func main() {
	router.CalcParse(&router.CalcLex{E: "1+2\n"})
}
