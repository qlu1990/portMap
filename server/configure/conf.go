package configure

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Service struct {
	dockername string   `json:"dockername`
	requires   []string `json:"requires"`
}

var (
	Conf map[string]interface{}
)

func init() {

	// file, err := os.OpenFile("conf/manager.json", os.O_RDONLY, 0666)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	dir := getdir()

	byte_Conf, err := ioutil.ReadFile(filepath.Join(dir, "../conf/manager.json"))
	if err != nil {
		byte_Conf, err = ioutil.ReadFile(filepath.Join(dir, "conf/manager.json"))
		if err != nil {
			log.Fatal(err)
		}
	}

	json.Unmarshal(byte_Conf, &Conf)
	// log.Printf("%T is : %v", serviceconf, serviceconf)
	// for dec.More() {

	// 	log.Printf("%T is : %v", serviceconf, serviceconf)
	// }
	setlog() //设置log

}
func setlog() {

	logpath, exit := Conf["logpath"].(string)
	if !exit {
		logpath = getdir()
	}

	logfile := filepath.Join(logpath, "control.log")

	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
		log.Println("open control.log failed")
		os.Exit(1)
	}
	log.SetOutput(f)
	log.Printf("conf log file is:%s", logfile)
}
func getdir() string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
		return ""

	} else {
		return dir
	}

}

func GetHtmlPath() string {
	htmlPath, exit := Conf["htmlpath"]
	if !exit {
		htmlPath := getdir()
		return htmlPath
	}
	return htmlPath.(string)
}
