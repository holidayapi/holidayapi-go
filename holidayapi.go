//package holidayapi
package main

import "fmt"

// TODO: import ( encoding/json net/http net/url )

type V1 struct {
	Key string
}

func NewV1(key string) *V1 {
	v1 := &V1{
		Key: key,
	}

	return v1
}

func (v1 *V1) Holidays(parameters map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"test": "ing",
	}
}

func main() {
	hapi := NewV1("_MY_API_KEY_")

	fmt.Println(hapi.Key)

	holidays := hapi.Holidays(map[string]interface{}{
		"country": "US",
	})

	fmt.Println("%#v\n", holidays)
}
