package main // import "github.com/pinghe/caddycors"

import (
	"github.com/mholt/caddy/caddy/caddymain"
	_ "github.com/captncraig/cors/caddy"
)

func main() {
	caddymain.Run()
}