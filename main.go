package main

import (
	"encoding/json"
	"golang.org/x/net/webdav"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

type DirMaps struct {

	Prefix string `json:"prefix"`
	Dir string `json:"dir"`
}

type Dirs struct {
	Port	string		`json:"port"`
	Mapping []DirMaps `json:"dirs"`
}

func main(){
	contentBytes,err := ioutil.ReadFile("config.json")
	if err != nil{
		log.Fatalf(err.Error())
	}
	var config Dirs
	err = json.Unmarshal(contentBytes,&config)
	if err != nil{
		log.Fatalf(err.Error())
	}
	handlersMapping := make(map[string]*webdav.Handler)

	for _,cfg := range config.Mapping{
		prefix := cfg.Prefix
		dir := cfg.Dir
		h := &webdav.Handler{
			Prefix:     prefix,
			FileSystem: webdav.Dir(dir),
			LockSystem: webdav.NewMemLS(),
		}
		handlersMapping[prefix] = h
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		uri := request.RequestURI
		uris := strings.FieldsFunc(uri, func(r rune) bool {
			r1,_ :=  utf8.DecodeRuneInString("/")
			return r1 == r
		})
		//fmt.Println("req: " + uri)
		//fmt.Println("uris:",uris)
		h := handlersMapping["/"+uris[0]]
		if h != nil{
			h.ServeHTTP(writer,request)
		}else{
			writer.WriteHeader(404)
		}

	})

	if config.Port == ""{
		_ = http.ListenAndServe(":7212", mux)
	}
	_ = http.ListenAndServe(":"+config.Port, mux)

}
