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
	temp.ExecuteTemplate(w, "admin.voters", nil)
}
<<<<<<< HEAD
func signup(w http.ResponseWriter, r *http.Request)  {
	temp.ExecuteTemplate(w,"home.l.layout",nil)
}

func index(w http.ResponseWriter, r *http.Request)  {
	temp.ExecuteTemplate(w,"index.html",nil)
=======
func signup(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "signup", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "", nil)
>>>>>>> aa1189b6461a32fdafb119ec0aa96fb2336f55e2
}
func newEvnet(w http.ResponseWriter, req *http.Request) {
	temp.ExecuteTemplate(w, "new.event", nil)
}

func RoutesForAdmin() {

}
func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:berekettussa@localhost:8080/votes?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	/*errs := dbconn.CreateTable(&entities.Events{}).GetErrors()
	if 0 < len(errs) {
		panic(errs)
	}*/

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

	http.HandleFunc("/voters", login)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8181", nil)
}
