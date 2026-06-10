package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8501"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		code := http.StatusOK
		if c := q.Get("code"); c != "" {
			if n, err := fmt.Sscanf(c, "%d", &code); err != nil || n != 1 {
				code = http.StatusOK
			}
		}

		if d := q.Get("delay"); d != "" {
			var secs int
			if n, err := fmt.Sscanf(d, "%d", &secs); err == nil && n == 1 {
				time.Sleep(time.Duration(secs) * time.Second)
			}
		}

		log.Printf("[%s] %d %s %s %s", time.Now().Format(time.RFC3339), code, r.Method, r.URL.String(), r.UserAgent())
		w.WriteHeader(code)
		fmt.Fprintln(w, http.StatusText(code))
	})

	addr := ":" + port
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
