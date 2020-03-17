// Copyright © 2018 1138-4EB <@github>
//
// Unless required by applicable law or agreed to in writing, this
// software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.

package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/1138-4EB/icemulti/lib"
)

func AJAX(w http.ResponseWriter, r *http.Request) {
	rest := mux.Vars(r)["rest"]
	switch {
	case strings.HasPrefix(rest, "/setDTR/"):
		data, err := ioutil.ReadAll(r.Body)
		if err == nil {
			b, err := strconv.ParseBool(string(data))
			log.Println("setDTR request:", b)
			if err == nil {
				w.Write([]byte(data))
				return
			}
			w.Write([]byte("ERROR"))
		}
	case strings.HasPrefix(rest, "/list/"):
		bins, err := lib.List([]string{"/src/test"}, true)
		if err != nil {
			log.Printf("list error [%v]\n", err)
			return
		}
		w.Write(lib.List2JSON(bins))
		return
	default:
		http.Error(w, "Unknown Status", http.StatusInternalServerError)
	}
}
