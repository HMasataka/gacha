package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"time"
)

func randomFloat(max, min float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}

type Data map[string]interface{}

func (d Data) Keys() []string {
	keys := make([]string, 0, len(d))
	for key := range d {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	data := make(Data)

	j, err := os.ReadFile("example.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(j, &data)
	for _, value := range data {
		fmt.Println(value)
		fmt.Println(reflect.TypeOf(value))
		fmt.Println(reflect.TypeOf(value).Kind() == reflect.Slice)
	}
}
