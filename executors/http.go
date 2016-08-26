package executors

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/SierraSoftworks/chieftan-server/utils"

	"net/http"
)

type HTTP struct {
	*ExecutorBase

	Client http.Client
}

func (e *HTTP) Name() string {
	return "HTTP"
}

func (r *HTTP) Run(ctx *Execution) error {
	var data *bytes.Buffer

	if ctx.Action.HTTP.Data != nil {
		data = bytes.NewBuffer([]byte{})
		i := utils.NewInterpolator(ctx.Variables)

		encodedData, err := i.Run(ctx.Action.HTTP.Data)
		if err != nil {
			return err
		}

		switch encodedData.(type) {
		case string:
			data.WriteString(encodedData.(string))
		default:
			j := json.NewEncoder(data)
			err = j.Encode(encodedData)
			if err != nil {
				return err
			}
		}
	}

	req, err := http.NewRequest(ctx.Action.HTTP.Method, ctx.Action.HTTP.URL, data)
	if err != nil {
		return err
	}

	ctx.WriteLn("%s %s", req.Method, req.URL)

	for key, val := range ctx.Action.HTTP.Headers {
		req.Header.Set(key, val)
		ctx.WriteLn("%s: %s", key, val)
	}

	ctx.WriteLn("")

	if data != nil {
		ctx.WriteLn(data.String())
		ctx.WriteLn("")
	}

	res, err := r.Client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	ctx.WriteLn(buf.String())

	if res.StatusCode >= 400 {
		return fmt.Errorf("Request failed with status %d %s", res.StatusCode, res.Status)
	}

	return nil
}
