package utils

import "encoding/json"

func parseJSON(data interface{}, c *Context) error {
	dec := json.NewDecoder(c.Request.Body)
	return dec.Decode(data)
}

func writeJSON(data interface{}, c *Context) error {
	enc := json.NewEncoder(c.response)
	return enc.Encode(data)
}
