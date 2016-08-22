package executors

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/SierraSoftworks/chieftan-server/utils"

	"net/http"
	"strings"
)

type Request struct {
	*Executor

	Client http.Client
}

func (r *Request) run(ctx *executorContext) error {
	out := []string{}

	data := bytes.NewBuffer([]byte{})

	if ctx.Action.HTTP.Data != nil {
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

	for key, val := range ctx.Action.HTTP.Headers {
		req.Header.Set(key, val)
	}

	res, err := r.Client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	out = append(out, buf.String())

	ctx.Task.Output = fmt.Sprintf("%s\n%s", ctx.Task.Output, strings.Join(out, "\n"))

	if res.StatusCode >= 400 {
		return fmt.Errorf("::[error] Request failed with status %d %s:: ", res.StatusCode, res.Status)
	}

	return nil
}
