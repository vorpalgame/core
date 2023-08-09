package lib

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Marshal(file string, data interface{}) {
	marshaled, _ := yaml.Marshal(data)

	e := os.WriteFile(file, marshaled, 0644)
	if e != nil {
		panic(e)
	}
}

func UnmarshalFile(fileName string, toType interface{}) {
	UnmarshalBytes(ReadFile(fileName), toType)
}

func UnmarshalBytes(contents []byte, toType interface{}) {
	e := yaml.Unmarshal(contents, toType)
	if e != nil {
		panic(e)
	}
}

func ReadFile(fileName string) []byte {
	f, e := os.ReadFile(fileName)
	if e != nil {
		log.Default().Println(e)
		os.Exit(1)
	}
	return f
}
