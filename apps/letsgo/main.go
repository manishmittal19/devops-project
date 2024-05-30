package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, string("let's go app!"))
}

func readiness(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string("i'm ready!!"))
}

func liveness(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string("i'm alive!!"))
}

func appinfo(w http.ResponseWriter, r *http.Request) {

	deployType := os.Getenv("DEPLOYMENT_TYPE")
	if deployType == "" {
		deployType = "UNKNOWN"
	}

	appInfo := map[string]string{
		"appname":         "letsgo",
		"deployment_type": deployType,
	}

	if appInfoJSON, err := json.Marshal(appInfo); err != nil {
		http.Error(w, "Internal Server Error", 500)
	} else {
		fmt.Fprint(w, string(appInfoJSON))
	}

	return
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/ready", readiness)
	mux.HandleFunc("/live", liveness)
	mux.HandleFunc("/appinfo", appinfo)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
