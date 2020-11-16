package cqa

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
    "github.com/brugger/kbr-go-tools/http"
	"github.com/gorilla/mux"
    //	"internal/dbFacade"
)


var version string = "0.0.0"


func checkErr(err error) {
    if err != nil {
        log.Fatalln("Error")
        log.Fatal(err)
    }
}

func dbGetProbes(filter map[string]string) ([]map[string]interface{}) {
    stmt := "SELECT * FROM probes"

    var conds []string

    for key, value := range filter {
        switch key {
            case "from":
                conds = append(conds, fmt.Sprintf(" pos_%s_vcf >= '%s'", filter["coords"], value))
            case "to":
                conds = append(conds, fmt.Sprintf(" pos_%s_vcf <= '%s'", filter["coords"], value))
            case "coords":
            default:
                conds = append( conds, fmt.Sprintf(" %s = '%s'", key, value))
        }
    }

    if len( conds) > 0 {
        stmt = fmt.Sprintf("%s WHERE %s ", stmt, strings.Join(conds[:], " AND "))
    }

    dbUtils.Connect( "sqlite3", sqlite_file )
    rows := dbUtils.AsList( stmt )
    dbUtils.Close()
    return rows
}


func infoPage(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: infoPage")
    json.NewEncoder(w).Encode(map[string]string{"name":"array-api", "version":version})
}

func getInstruments(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: instruments")
    query := r.URL.Query()


    params = httpUtils.ArrayToMap( query.Get(key) )
    httpUtils.ValidArguments( params, ["name", "id"])

    json.NewEncoder(w).Encode( dbGetInstruments( params ) )
}






func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", infoPage)
    myRouter.HandleFunc("/instruments/", getInstruments)
//    myRouter.HandleFunc("/probe/{id}", createProbe).Methods("POST")
//    myRouter.HandleFunc("/probe/{id}/", readProbes)
//    myRouter.HandleFunc("/probe/{id}", updateProbe).Methods("PATCH")
//    myRouter.HandleFunc("/probe/{id}", deleteProbe).Methods("DELETE")
    var port = 10000
    var portString = fmt.Sprintf(":%d" , port)
    fmt.Println("Listening on port" , portString)
    log.Fatal(http.ListenAndServe( portString , myRouter))
    
}
func main() {
    dbUtils.Connect( "sqlite3", "host=localhost port=5432 user=cqa dbname=cqa sslmode=disable password=cqa" )


    handleRequests()
}
