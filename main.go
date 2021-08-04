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

func main() {
	data := make(map[string]interface{})

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
