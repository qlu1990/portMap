package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"manager/node/controlers"
	"manager/server/configure"
	"net/http"
	"path/filepath"
)

var (
	Tasks    chan map[string]string
	HtmlPath string
)

func init() {
	controlers.ConfControl(configure.Conf)
	HtmlPath = configure.GetHtmlPath()

}

func index(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.Method)
	if r.Method == "GET" {
		//t,_ :=template.ParseFiles("index.html")

		indexdata, _ := ioutil.ReadFile(filepath.Join(HtmlPath, "index.html"))
		w.Write(indexdata)
	}
}
func control(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)

	if r.Method == "POST" {
		r.ParseForm()
		task := make(map[string]string)
		task["servicename"] = r.FormValue("softname")
		task["operate"] = r.FormValue("operate")
		Tasks <- task
		ret := make(map[string]string)
		for k, v := range task {
			ret[k] = v
		}
		ret["status"] = "success"
		header := w.Header()
		header.Set("Content-Type", "application/json")
		str, err := json.Marshal(ret)
		if err != nil {
			io.WriteString(w, "{\"status\":\"error\"}")
			return
		}
		io.WriteString(w, string(str))
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.Method)

	if r.Method == "GET" {

		ret := controlers.GetServicesInfo()
		header := w.Header()
		header.Set("Content-Type", "application/json")
		w.Write(ret)
	}

}

func main() {
	
}
