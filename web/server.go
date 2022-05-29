package web 

import (
	"log"
	"net/http"
	"os"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, there!\n"))
}

func ServerLaunch() {
	addr := os.Getenv("ADDR")

	if len(addr) == 0 {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", HelloHandler)

	log.Printf("server booted...")
	log.Printf("Visit http://localhost%s ...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
