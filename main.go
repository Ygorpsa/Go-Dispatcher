package main

import (
    //"encoding/base64"
    "encoding/json"
    //"fmt"
    //"log"
    "net/http"
    "os"

    //"github.com/GrooveCommunity/dispatcher-jira/internal"
    //"github.com/GrooveCommunity/glib-noc-event-structs/entity"
    "github.com/gorilla/mux"
    "google.golang.org/api/pubsub/v1"
)


func main() {
    router := mux.NewRouter()
    //router.HandleFunc("/healthy", handleValidateHealthy).Methods("GET")
    //router.HandleFunc("/queue-dispatcher-jira", handleQueueDispatcher).Methods("POST")
    //router.HandleFunc("/put-rule", handlePutRule).Methods("POST")
    router.HandleFunc("/rules", handleRules).Methods("GET")

    //username = os.Getenv("JIRA_USERNAME")
    //token = os.Getenv("JIRA_TOKENAPI")
    //endpoint = os.Getenv("JIRA_ENDPOINT")
    appPort := os.Getenv("APPPORT")

    /if username == ""  token == ""  endpoint == "" || appPort == "" {
        log.Fatal("Nem todas as variáveis de ambiente requeridas foram fornecidas. ")
    }/

    panic(http.ListenAndServe(":"+appPort, router))
}



type Field struct {
    ID    string
    Name  string
    Value string
}


func handleRules(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(GetRules{})
}


func GetRules() []entity.Rule {
    var rules []entity.Rule

    dataObjects := gcp.GetObjects("rules-dispatcher")

    for , b := range dataObjects {
        var rule entity.Rule
        errUnmarsh := json.Unmarshal(b, &rule)

        if errUnmarsh != nil {
            log.Fatal("Erro no unmarshal\n", errUnmarsh.Error())
        }

        rules = append(rules, rule)
}