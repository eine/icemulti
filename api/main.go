// Copyright © 2018 eine <@github>
//
// Unless required by applicable law or agreed to in writing, this
// software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	flag "github.com/ogier/pflag"

	"github.com/eine/icemulti/api/handler"
)

//> Args
var port = flag.IntP("port", "p", 8080, "port to serve at")
var dir = flag.StringP("dir", "d", "./public", "dir to serve from")
var verb = flag.BoolP("verbose", "v", false, "")
var nofs = flag.BoolP("nofs", "n", false, "whether to disable serving files")

//<

//> Global vars
var httpReqDump func(r *http.Request) = func(r *http.Request) {}

func init() {
	flag.Parse()
	if *verb {
		httpReqDump = httputilReqDump
	}
}

//<

//> Print requests to stdout
func httputilReqDump(r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(requestDump))
}

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		httpReqDump(r)
		h.ServeHTTP(rw, r)
	})
}

//<

func main() {
	log.Println("[DTD service] Start")

	str := ""
	if *dir != "none" {
		istr, err := filepath.Abs(*dir)
		if err != nil {
			os.Exit(1)
		}
		str = istr
	} else {
		*nofs = true
	}

	r := InitHandlers(*dir)

	log.Printf("Serving at port %d\n", *port)
	if *nofs == false {
		log.Printf("Serving from dir %s\n", str)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", *port), r)
}

func InitHandlers(d string) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ajax{rest:.*}", handler.AJAX)
	if *nofs == false {
		r.PathPrefix("/").Handler(loggingHandler(http.FileServer(handler.FilesOnlyFs{http.Dir(d)})))
	}
	return r
}
