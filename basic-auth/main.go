package main

import (
	"encoding/base64"
	"fmt"
	"github.com/evacchi/quarkus-wasm-sdk/sdk"
	"strings"
)

type BasicAuthPlugin struct{}

func (h *BasicAuthPlugin) OnRequestHeaders(req *sdk.Request) error {
	if hh, ok := req.Headers["Authorization"]; ok && hh != nil {
		authHeader := strings.Split(hh[0], " ")
		if authHeader[0] == "Basic" {
			data, err := base64.StdEncoding.DecodeString(authHeader[1])
			if err != nil {
				return err
			}
			pair := string(data)
			userPwd := strings.Split(pair, ":")
			if userPwd[0] == "admin" && userPwd[1] == "admin" {
				// Access granted.
				delete(req.Headers, "Authorization")
				req.Headers["X-Authorized"] = []string{userPwd[1]}
			} else {
				return fmt.Errorf("invalid authorization header %s", pair)
			}
		}
	}
	return nil
}

func init() {
	sdk.SetPlugin(&BasicAuthPlugin{})
}

func main() {}
