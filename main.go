package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const isDev bool = false

type env struct { //environmental variables
	Statichtml string
}

var currEnv env

func handleExecute(w http.ResponseWriter, r *http.Request /**auth.AuthenticatedRequest*/, command string) {
	if isDev {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	}
	var anyJson map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&anyJson)
	if err != nil {
		log.Println(err)
	}
	jsonString, err := json.Marshal(anyJson)
	if err != nil {
		log.Println(err)
	}
	log.Println(jsonString)
	cppCmd, err := exec.Command(command, string(jsonString)).Output()
	if err != nil {
		log.Println(err)
	}
	result := string(cppCmd)
	json.NewEncoder(w).Encode(result)
}

func init() {
	f, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Couldn't open log")
	}
	log.SetOutput(f)
	pwd, _ := os.Getwd()
	file, err1 := ioutil.ReadFile(pwd + "/environment.json") // For read access
	if err1 != nil {
		log.Fatal(err1)
	}
	err2 := json.Unmarshal(file, &currEnv)
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Println(currEnv)

}
func main() {
	log.Println(currEnv.Statichtml)
	http.Handle("/", http.FileServer(http.Dir(currEnv.Statichtml)))
	http.HandleFunc("/RunModel/creditRisk", func(w http.ResponseWriter, r *http.Request) {
		handleExecute(w, r, "./creditRisk")
	})
	http.HandleFunc("/RunModel/marketRisk", func(w http.ResponseWriter, r *http.Request) {
		handleExecute(w, r, "./marketRisk")
	})
	http.HandleFunc("/RunModel/opsRisk", func(w http.ResponseWriter, r *http.Request) {
		handleExecute(w, r, "./opsRisk")
	})
	log.Println("Listening...")
	http.ListenAndServe(":80", nil)

}
