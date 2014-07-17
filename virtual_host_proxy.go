// Copyright 2013, Amahi.  All rights reserved.
// Use of this source code is governed by the
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
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
		remote, err := url.Parse("http://" + Host)
		if err != nil {
			fmt.Println("Error:", err)
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", RequestHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
