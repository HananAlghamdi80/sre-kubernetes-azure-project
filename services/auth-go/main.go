package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var loginRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "auth_login_requests_total",
		Help: "Total login requests",
	},
)

func init() {
	prometheus.MustRegister(loginRequests)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	loginRequests.Inc()

	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var user LoginRequest
	json.NewDecoder(r.Body).Decode(&user)

	// Dummy authentication logic
	if user.Username == "admin" && user.Password == "123" {
		w.Write([]byte("login successful ✅"))
	} else {
		w.Write([]byte("invalid credentials ❌"))
	}
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ready"))
	})

	http.Handle("/metrics", promhttp.Handler())

	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	http.ListenAndServe(":"+port, nil)
}
