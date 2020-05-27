package main

import (
	"github.com/user3301/goavrobuf"
	"io/ioutil"
)

func main() {
	schema, err := ioutil.ReadFile("fixture/avro.avsc")
	if err != nil {
		panic(err)
	}
	root, err := goavrobuf.NewSchema(string(schema))
	proto3 := goavrobuf.GenerateProto3(root)
	err = ioutil.WriteFile("avro.proto", []byte(proto3), 0644)
	if err != nil {
		panic("damn!")
	}
}
