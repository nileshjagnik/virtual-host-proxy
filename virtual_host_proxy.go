// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

// Frame related functions for generic for control/data frames
// as well as specific frame related functions

package main

import (
	"fmt"
	"net/http"
	"strings"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	Host := r.Host
	if Host == "" {
		http.Error(w, "Host header required", 401)
	} else {
		//strip the port from the header
		parts := strings.Split(Host, ":")
		if len(parts) > 1 {
			Host = parts[0]
		}
		fmt.Println("Forwarding Request to:", Host)

		// Make a new client
		client := &http.Client{}

		//	RequestURI may not be sent to client
		//	URL.Scheme must be lower-case
		r.RequestURI = ""
		r.URL.Scheme = "http"
		r.URL.Host = Host
		fmt.Println(r.URL)

		// proxy the requests
		resp, err := client.Do(r)
		if err != nil {
			http.Error(w, "Proxy didnt receive response from upstream server", 502)
		} else {
			resp.Write(w)
		}
	}
}

func main() {
	http.HandleFunc("/", RequestHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
