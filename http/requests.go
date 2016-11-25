package http

import (
	"encoding/json"
	"net/http"
)

func ParseResponse(resp *http.Response, err error, target interface{}) error {
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
