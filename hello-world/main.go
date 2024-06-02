package main

import (
	"github.com/evacchi/quarkus-wasm-sdk/sdk"
)

type HelloWorldPlugin struct{}

func (h *HelloWorldPlugin) OnRequestHeaders(req *sdk.Request) error {
	req.AppendHeaders("X-Wasm-Plugin", "Hello World!")
	return nil
}

func (h *HelloWorldPlugin) OnResponseHeaders(resp *sdk.Response) error { return nil }

func init() {
	sdk.SetPlugin(&HelloWorldPlugin{})
}

func main() {}
