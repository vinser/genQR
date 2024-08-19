package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	qr "github.com/skip2/go-qrcode"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	content := r.FormValue("content")
	size := r.FormValue("size")

	if content == "" {
		responseError(w, "Could not determine the desired QR code content.")
		return
	}

	qrSize, err := strconv.Atoi(size)
	if err != nil || size == "" {
		responseError(w, "Could not determine the desired QR code size.")
		return
	}

	codeData, err := qr.Encode(content, qr.Medium, qrSize)

	if err != nil {
		responseError(w, fmt.Sprintf("Could not generate QR code. %v", err))
		return
	}
	w.Header().Add("Content-Type", "image/png")
	w.Write(codeData)
}

func responseError(w http.ResponseWriter, errMsg string) {
	w.WriteHeader(http.StatusBadRequest)
	responseData := make(map[string]string)
	responseData["error"] = errMsg

	response, err := json.Marshal(responseData)
	if err == nil {
		w.Write(response)
	}
}

func main() {
	port := flag.Int("port", 443, "HTTPS port to listen on - optional")
	cert := flag.String("cert", "", "path to public certificate file - mandatory")
	key := flag.String("key", "", "path to private key file - mandatory")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: genQR [flags]\n")
		fmt.Fprintf(os.Stderr, "REST service for QR Code generation\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()
	if *cert == "" || *key == "" {
		flag.Usage()
	}

	addr := fmt.Sprintf(":%d", *port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	log.Printf("Starting server on localhost:%d", *port)
	err := http.ListenAndServeTLS(addr, *cert, *key, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
