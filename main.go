package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


//var temp = template.Must(template.ParseGlob("ui/*"))
func main() {

	dbconn ,err := gorm.Open("postgres","postgres://postgres:minpass@localhost:9090/votingdb?sslmode=disable")

	if err != nil{
		panic(err)
	}
	defer dbconn.Close()
	//errs := dbconn.CreateTable(&entities.Events{},&entities.Parties{}).GetErrors()
	//
	//if len(errs)>0{
	//	panic(errs)
	//}
	//fs := http.FileServer(http.Dir("ui"))
	//
	//mux := http.NewServeMux()
	//mux.Handle("/ui/",http.StripPrefix("/ui",fs))
	//mux.HandleFunc("/",index)
	//http.ListenAndServe(":8989",mux)
}

//func index(w http.ResponseWriter,r *http.Request)  {
//	temp.ExecuteTemplate(w, "index.html",nil)
//}


