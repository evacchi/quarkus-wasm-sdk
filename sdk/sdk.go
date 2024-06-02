package sdk

import (
	"github.com/extism/go-pdk"
)

type HttpPlugin interface {
	OnRequestHeaders(req *Request) error
}

var plugin HttpPlugin

type Request struct {
	Headers map[string][]string `json:"headers"`
}

func SetPlugin(p HttpPlugin) {
	plugin = p
}

//go:export request_headers
func RequestHeaders() int32 {
	req := &Request{}
	err := pdk.InputJSON(req)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	err = plugin.OnRequestHeaders(req)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	err = pdk.OutputJSON(req)
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	return 0
}

func (r *Request) AppendHeaders(k, v string) {
	hs := r.Headers[k]
	hs = append(hs, v)
	r.Headers[k] = hs
}
