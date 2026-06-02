// Command hka-gateway runs the HKA SDK as a language-neutral JSON/HTTP service.
//
// Any language can call it; it hides SOAP and computes all document totals. Point
// it at the demo or production HKA endpoint and pass per-tenant credentials on each
// request via the X-HKA-Token-Empresa / X-HKA-Token-Password headers.
//
// Usage:
//
//	HKA_ENDPOINT=https://demoemision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc \
//	ADDR=:8080 go run ./cmd/hka-gateway
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	hka "github.com/angelnereira/hka-sdk"
	"github.com/angelnereira/hka-sdk/gateway"
)

func main() {
	addr := envOr("ADDR", ":8080")
	endpoint := os.Getenv("HKA_ENDPOINT") // empty -> demo endpoint

	var client *hka.Client
	if endpoint == "" {
		client = hka.NewDemo()
		log.Printf("hka-gateway: using HKA demo endpoint")
	} else {
		client = hka.New(&hka.Config{Endpoint: endpoint, Timeout: 60 * time.Second})
		log.Printf("hka-gateway: using HKA endpoint %s", endpoint)
	}

	h := gateway.NewHandler(client)
	srv := &http.Server{
		Addr:              addr,
		Handler:           h.Routes(),
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Printf("hka-gateway: listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
