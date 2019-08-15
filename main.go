package main // import "github.com/pinghe/caddycors"

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
	_ "github.com/captncraig/cors/caddy"
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false	
	caddymain.Run()
}