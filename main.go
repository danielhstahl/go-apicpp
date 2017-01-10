package main
import (
    "github.com/abbot/go-http-auth"
    "log"
    "net/http"
    "os/exec"
    "io/ioutil"
    "encoding/json"
)

const isDev bool = false
type marketRiskInput struct{
    Port string 
    Date string
    Amount int
    Material string
    Comment string
}
type creditRiskInput struct{
    Port string 
    Date string
    Amount int
    Material string
    Comment string
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
}
func writePortType(w http.ResponseWriter, r *auth.AuthenticatedRequest){
    if(isDev){
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    }
    decoder := json.NewDecoder(r.Body)
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
    json.NewEncoder(w).Encode(results)
}


func main(){
    //defer db.Close()
    
    //ath := auth.NewBasicAuthenticator("example.com", Secret)
    http.HandleFunc("/", ath.Wrap(func(res http.ResponseWriter, req *auth.AuthenticatedRequest) {
        http.FileServer(http.Dir(currEnv.Statichtml)).ServeHTTP(res, &req.Request)
    }))
    http.HandleFunc("/writePort", ath.Wrap(writePortType))
    http.HandleFunc("/writeMaterial", ath.Wrap(writeMaterialType))
    http.HandleFunc("/writeTransaction", ath.Wrap(writeTransaction))
    http.HandleFunc("/getResults", ath.Wrap(getAsOfMaterials))
    http.HandleFunc("/getPorts", ath.Wrap(getPort))
    http.HandleFunc("/getMaterials", ath.Wrap(getMaterial))
    http.HandleFunc("/getAllResults", ath.Wrap(getAllResults))
    log.Println("Listening...")
    http.ListenAndServe(":3000", nil)
}