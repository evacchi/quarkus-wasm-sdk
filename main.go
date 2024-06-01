package main

import (
	"github.com/extism/go-pdk"
)

type Request struct {
	Headers map[string][]string `json:"headers"`
}

//go:export request_headers
func RequestHeaders() int32 {
	req := &Request{}
	err := pdk.InputJSON(req)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	req.appendHeaders("X-Wasm-Plugin", "Hello World!")

	err = pdk.OutputJSON(req)
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	return 0
}

func (r *Request) appendHeaders(k, v string) {
	hs := r.Headers[k]
	hs = append(hs, v)
	r.Headers[k] = hs

}

func main() {}
