package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/jessevdk/go-flags"
)

func randomFloat(max, min float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}

func toStrings(v []interface{}) []string {
	s := make([]string, len(v))
	for i, vv := range v {
		s[i] = fmt.Sprint(vv)
	}
	return s
}

func toData(target interface{}) (Data, error) {
	v, ok := target.(Data)
	if ok {
		return v, nil
	}

	vv, ok := target.(map[string]interface{})
	if ok {
		return Data(vv), nil
	}
	return nil, errors.New("cast error")
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

func run(data interface{}) (interface{}, error) {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Map:
		target, err := toData(data)
		if err != nil {
			return nil, err
		}
		isValidPercent, err := target.IsValidPercent()
		if !isValidPercent || err != nil {
			return nil, errors.New("invalid percentage")
		}

		var percent float64
		r := randomFloat(100, 0)
		for key, value := range target {
			f, err := strconv.ParseFloat(key, 64)
			if err != nil {
				return nil, err
			}
			percent += f
			if r < percent {
				return run(value)
			}
		}
	case reflect.Slice:
		target := toStrings(data.([]interface{}))
		r := rand.Intn(len(target))
		return run(target[r])
	case reflect.String:
		return data, nil
	}

	return nil, errors.New("error")
}

type Opts struct {
	Input string `long:"input" short:"i" description:"Input file path" required:"true"`
}

func main() {
	var opts Opts
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	data := make(Data)

	j, err := os.ReadFile(opts.Input)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(j, &data)

	isValidPercent, err := data.IsValidPercent()
	if !isValidPercent {
		panic(errors.New("Invalid percent"))
	}

	fmt.Println(run(data))
}
