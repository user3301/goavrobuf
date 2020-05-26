package main

import (
	"fmt"
	"github.com/user3301/goavrobuf"
	"io/ioutil"
)

func main() {
	schema, err := ioutil.ReadFile("fixture/avro.avsc")
	if err != nil {
		panic(err)
	}
	root, err := goavrobuf.NewSchema(string(schema))
	fmt.Printf("%#v", root)
}
