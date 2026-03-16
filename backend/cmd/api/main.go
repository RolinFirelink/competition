package main

import (
	"log"
	"net/http"

	"sf6-academy/internal"
)

func main() {
	addr := ":" + envDefault("PORT", "8080")
	adminPass := envDefault("ADMIN_PASSWORD", "feigeA.5200....")

	store := internal.NewMemoryStore()
	svc := internal.NewService(store, adminPass)
	router := internal.NewRouter(svc)

	log.Printf("backend listening on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}

func envDefault(key, def string) string {
	if v := env(key); v != "" {
		return v
	}
	return def
}

func env(key string) string {
	return internal.GetEnv(key)
}
