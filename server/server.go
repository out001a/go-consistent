package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/out001a/go-consistent"
)

const (
	apiVersion  = "v1"
	apiBasePath = "/api/" + apiVersion + "/"

	consistentApiPath = apiBasePath + "consistent/"
)

var (
	port int

	consist *consistent.Consistent
)

func init() {
	flag.IntVar(&port, "port", 9871, "The port to listen on.")
}

func main() {
	flag.Parse()
	consist = consistent.NewConsistent()
	http.Handle(consistentApiPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			httpGetHandler(w, r)
		case http.MethodPost:
			httpPostHandler(w, r)
		case http.MethodDelete:
			httpDeleteHandler(w, r)
		}
	}))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
	// use `lsof -i :{port}` to see the info of process which binding this port.
}

func httpGetHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Path[len(consistentApiPath):]
	if target == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't lookup if there is no target."))
		log.Print("empty request.")
		return
	}
	if result, ok := consist.Lookup(target); ok {
		w.Write([]byte(result))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func httpPostHandler(w http.ResponseWriter, r *http.Request) {
	target, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(target) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't add if there is no target."))
		log.Print("empty request.")
		return
	}
	consist.Add(string(target))
	w.WriteHeader(http.StatusCreated)
}

func httpDeleteHandler(w http.ResponseWriter, r *http.Request) {
	target, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(target) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can't remove if there is no target."))
		log.Print("empty request.")
		return
	}
	consist.Remove(string(target))
	w.WriteHeader(http.StatusOK)
}
