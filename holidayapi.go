package holidayapi

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/simplereach/timeutils"
	"io/ioutil"
	"net/http"
	"net/url"
)

type V1 struct {
	Url string
	Key string
}

func NewV1(key string) *V1 {
	v1 := &V1{
		Key: key,
		Url: "https://holidayapi.com/v1/holidays?",
	}

	return v1
}

type Respone struct {
	Status   int32     `json:"status"`
	Requests Request   `json:"requests"`
	Holidays []Holiday `json:"holidays"`
}

type Request struct {
	Used      int32          `json:"used"`
	Available int32          `json:"available"`
	Resets    timeutils.Time `json:"resets"`
}

type Holiday struct {
	Name     string         `json:"name"`
	Date     timeutils.Time `json:"date,omniempty"`
	Observed timeutils.Time `json:"observed"`
	Public   bool           `json:"public"`
}

func (v1 *V1) Holidays(args map[string]string) (Respone, error) {
	var data Respone

	if _, ok := args["key"]; !ok {
		args["key"] = v1.Key
	}

	params := url.Values{}

	for k, v := range args {
		params.Add(k, v)
	}

	resp, err := http.Get(v1.Url + params.Encode())

	if err != nil {
		return Respone{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Respone{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return Respone{}, err
	}

	if resp.StatusCode != 200 {
		return Respone{}, errors.New("Unknown error!")
	}

	return data, nil
}
