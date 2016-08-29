package executors

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/SierraSoftworks/chieftan-server/utils"

	"net/http"
)

type HTTP struct {
}

func (e *HTTP) Name() string {
	return "HTTP"
}

func (r *HTTP) Run(ctx *Execution) error {
	var data io.Reader
	i := utils.NewInterpolator(ctx.Variables)

	ctx.WriteLn("%s %s", ctx.Action.HTTP.Method, ctx.Action.HTTP.URL)

	headers := map[string]string{}

	if ctx.Action.HTTP.Headers != nil {
		for key, val := range ctx.Action.HTTP.Headers {
			iVal, err := i.Run(val)
			if err != nil {
				return err
			}

			headers[key] = iVal.(string)
			ctx.WriteLn("%s: %s", key, iVal.(string))
		}
	}

	ctx.WriteLn("")

	if ctx.Action.HTTP.Data != nil {
		dataBuffer := bytes.NewBuffer([]byte{})

		encodedData, err := i.Run(ctx.Action.HTTP.Data)
		if err != nil {
			return err
		}

		switch encodedData.(type) {
		case string:
			dataBuffer.WriteString(encodedData.(string))
		default:
			j := json.NewEncoder(dataBuffer)
			err = j.Encode(encodedData)
			if err != nil {
				return err
			}
		}

		data = bufio.NewReader(dataBuffer)
		ctx.WriteLn(dataBuffer.String())
		ctx.WriteLn("")
	}

	req, err := http.NewRequest(ctx.Action.HTTP.Method, ctx.Action.HTTP.URL, data)
	if err != nil {
		return err
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	ctx.WriteLn("%s %d %s", res.Proto, res.StatusCode, res.Status)
	for key, val := range res.Header {
		ctx.WriteLn("%s: %s", key, val)
	}
	ctx.WriteLn("")

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	ctx.WriteLn(buf.String())

	if res.StatusCode >= 400 {
		return fmt.Errorf("Request failed with status %d %s", res.StatusCode, res.Status)
	}

	return nil
}
