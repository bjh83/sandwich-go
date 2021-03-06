package main

import "net/http"
import "log"

// this will eventually resize, but right now you can't have more than 500 peers.
var defaultPeerListSize = 500

func pingHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("pong\n"))
}

func indexForHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func peerListHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/json")
	addressList := AddressList.Contents() //Gets a copy of the underlying IPSlice
	writer.Write([]byte(addressList.String())) //creates a json representation of addressList
}

func main() {

	InitializeAddressList()

	http.HandleFunc("/", indexForHandler)
	http.HandleFunc("/peerlist", peerListHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/indexfor", indexForHandler)

	log.Printf("About to listen on 8000. Go to http://127.0.0.1:8000/")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

