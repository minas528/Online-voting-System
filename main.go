package main

import (
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	eventRepo "github.com/minas528/Online-voting-System/Event/repository"
	eventServ "github.com/minas528/Online-voting-System/Event/service"
	"github.com/minas528/Online-voting-System/delivery/http/handler"
	postRepo "github.com/minas528/Online-voting-System/post/repository"
	postServ "github.com/minas528/Online-voting-System/post/service"
)

var temp = template.Must(template.ParseGlob("ui/templates/*"))

func login(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "login", nil)
}
func signup(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "signup", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "", nil)
}
func newEvnet(w http.ResponseWriter, req *http.Request) {
	temp.ExecuteTemplate(w, "new.event", nil)
}
<<<<<<< HEAD
func parties(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "parties",nil)
}
func RoutesForAdmin()  {
=======

func RoutesForAdmin() {
>>>>>>> aa1189b6461a32fdafb119ec0aa96fb2336f55e2

}
func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:minpass@localhost:9090/votes?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

<<<<<<< HEAD
	//errs := dbconn.CreateTable(&entities.Events{}).GetErrors()
	//if 0 < len(errs) {
	//	panic(errs)
	//}
=======
	/*errs := dbconn.CreateTable(&entities.Events{}).GetErrors()
	if 0 < len(errs) {
		panic(errs)
	}*/
	// errs := dbconn.CreateTable(&entities.RegParties{}, &entities.RegVoters{}).GetErrors()
	// if 0 < len(errs) {
	//	panic(errs)
	// }
>>>>>>> 2bc9a017d1f15b2af48c0a8c30a8390b93e967e3

	postRepo := postRepo.NewPostGormRepo(dbconn)
	postserv := postServ.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(temp, postserv)

	eventRep := eventRepo.NewEventRepository(dbconn)
	eventserv := eventServ.NewEventService(eventRep)
	eventHandle := handler.NewEventHandler(temp, eventserv)

	fs := http.FileServer(http.Dir("ui/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/upost", postHandler.PostNew)
	http.HandleFunc("/posts", postHandler.Posts)
	http.HandleFunc("/", index)
	//http.HandleFunc("/newevent",newEvnet)

	http.HandleFunc("/events", eventHandle.Events)
	http.HandleFunc("/newevent", eventHandle.EventNew)

	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8181", nil)
}
