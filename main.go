package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
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

func (d Data) IsValidPercent() (bool, error) {
	var percent float64
	for key := range d {
		f, err := strconv.ParseFloat(key, 64)
		if err != nil {
			return false, err
		}
		percent += f
	}
	return percent == 100.0, nil
}

func main() {
	data := make(Data)

	j, err := os.ReadFile("example.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(j, &data)

	fmt.Println(data.Keys())
	fmt.Println(data.IsValidPercent())

	for _, value := range data {
		fmt.Println(value)
		fmt.Println(reflect.TypeOf(value))
		fmt.Println(reflect.TypeOf(value).Kind() == reflect.Slice)
	}
}
