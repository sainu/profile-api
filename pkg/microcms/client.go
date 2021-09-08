package microcms

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type client struct{}

func NewClient() *client {
	return &client{}
}

func (c *client) Do(req *http.Request, v interface{}) error {
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, v); err != nil {
		return err
	}
	return nil
}
