package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Hello")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("Port is not set")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.ListenAndServe(":"+port, nil)
}
