package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RequestAndUpdate(url string, data *StackOverflowItemList) error {
	// Send GET request
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	// Close body when all of the code in this block is finished.
	defer resp.Body.Close()

	// Read bytes
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	// Decode JSON
	err = json.Unmarshal(bodyBytes, data)

	if err != nil {
		return err
	}

	return nil
}
