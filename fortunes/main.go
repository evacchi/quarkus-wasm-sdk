package main

import (
	_ "embed"
	"github.com/evacchi/quarkus-wasm-sdk/sdk"
	"math/rand"
	"strings"
	"time"
)

//go:embed fortunes.txt
var fortunes string

type FortunesPlugin struct{}

func (h *FortunesPlugin) OnRequestHeaders(req *sdk.Request) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	lines := strings.Split(fortunes, "\n")
	l := r.Intn(len(lines))
	req.AppendHeaders("X-Fortune-Plugin", lines[l])
	return nil
}

func init() {
	p := &FortunesPlugin{}
	sdk.SetPlugin(p)
}

func main() {}
