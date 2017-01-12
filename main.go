package main
//"github.com/abbot/go-http-auth"
//"log"\
// "encoding/json"
//"os/exec"
import (
    
    "log"
    "net/http"
    "os"
    "io/ioutil"
   "encoding/json"
)

const isDev bool = true

type env struct{ //environmental variables
    Statichtml string
    
}
var currEnv env
/*type marketRiskInput struct{
    n int 
    sigma float64
    a float64
    t float64
    T float64
    delta float64
    Tm float64
}
type creditRiskInput struct{
    n int 
    t float64
    x0 float64
    alpha float64
    sigma float64
    q float64
    lambda float64
    xNum int
    uNum int
}
type opsRiskInput struct{
    Port string 
    Date string
    Amount int
    Material string
    Comment string
}
type marketRiskOutput struct{
    Port string 
    Date string
    Amount int
    Material string
    Comment string
}
type creditRiskOutput struct{
    Port string 
    Date string
    Amount int
    Material string
    Comment string
}
type opsRiskOutput struct{
    Port string 
    Date string
    Amount int
    Material string
    Comment string
}*/
func handleExecute(w http.ResponseWriter, r *http.Request/**auth.AuthenticatedRequest*/, command string){
    if(isDev){
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    }
    log.Println(r.Body)
    //grepCmd := exec.Command(command, "hello")

    /*decoder := json.NewDecoder(r.Body)
    var t portType   
    err := decoder.Decode(&t)
    if err != nil {
        log.Println(err) 
    }
    _, err1:=db.Exec(`INSERT INTO main.possibleports VALUES($1)`, t.Port)
    if err1!=nil{
        log.Println(err1)
    }
    results:=new(success)
    results.Success=true
    json.NewEncoder(w).Encode(results)*/
}

func init(){
    f, err := os.OpenFile("log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("Couldn't open log")
    }
    log.SetOutput(f)
    pwd, _ := os.Getwd()
    file, err1 := ioutil.ReadFile(pwd+"/environment.json") // For read access
    if err1!=nil{
        log.Fatal(err1)
    }
    err2:=json.Unmarshal(file,&currEnv)
    if err2!=nil{
        log.Fatal(err2)
    }
    log.Println(currEnv)
    
}
func main(){
    //defer db.Close()
    
    //ath := auth.NewBasicAuthenticator("example.com", Secret)
    //http.HandleFunc("/", /*ath.Wrap(func(res http.ResponseWriter, req *auth.AuthenticatedRequest*/func(res http.ResponseWriter, req *http.Request) {
        //http.FileServer(http.Dir(currEnv.Statichtml)).ServeHTTP(res, &req.Request)
    //})
    log.Println(currEnv.Statichtml)
    http.Handle("/", http.FileServer(http.Dir(currEnv.Statichtml)))
    http.HandleFunc("/RunModel/creditRisk", func(w http.ResponseWriter, r *http.Request){
        handleExecute(w, r, "./creditRisk")
    })
    http.HandleFunc("/RunModel/marketRisk", func(w http.ResponseWriter, r *http.Request){
        handleExecute(w, r, "./marketRisk")
    })
    http.HandleFunc("/RunModel/opsRisk", func(w http.ResponseWriter, r *http.Request){
        handleExecute(w, r, "./opsRisk")
    })
    log.Println("Listening...")
    http.ListenAndServe(":3001", nil)
     
}