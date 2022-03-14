package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var port int = 8000

func getIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header

	log.Println(r.Header.Get("User-Agent"))
	ip := r.Header.Get("X-REAL-IP")

	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	log.Println(ips)
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {

		return ip, nil
	}
	return "", fmt.Errorf("No valid ip found")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/db", getDbPageHandler)
	router.HandleFunc("/create", createQuestionPageHandler)
	router.HandleFunc("/addQuestion", createQuestionHandler).Methods("POST")
	router.HandleFunc("/register", registerPageHandler)
	router.HandleFunc("/addUser", registerHandler).Methods("POST")
	router.HandleFunc("/login", loginHandler).Methods("POST")

	http.Handle("/", router)

	log.Println("main: running simple server on port", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("Could not start server", err)
	}
}
