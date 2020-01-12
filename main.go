package main

import (
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	eventRepo "github.com/minas528/Online-voting-System/Event/repository"
	eventServ "github.com/minas528/Online-voting-System/Event/service"
	"github.com/minas528/Online-voting-System/delivery/http/handler"
	partyRepo "github.com/minas528/Online-voting-System/parties/repository"
	partyServ "github.com/minas528/Online-voting-System/parties/service"
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
func parties(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "parties", nil)
}
func RoutesForAdmin() {
}
func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:minpass@localhost:9090/votes?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

<<<<<<< HEAD
	/*errs := dbconn.CreateTable(&entities.Parties{}).GetErrors()
=======
<<<<<<< HEAD
	//errs := dbconn.CreateTable(&entities.Events{}).GetErrors()
	//if 0 < len(errs) {
	//	panic(errs)
	//}
=======
	/*errs := dbconn.CreateTable(&entities.Events{}).GetErrors()
>>>>>>> 16e7adbc68177c043a8fc6c3f98223984f6335a7
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

	partyRepo := partyRepo.NewPartiesGormRepo(dbconn)
	partyserv := partyServ.NewPartiesService(partyRepo)
	partyHandler := handler.NewAdminPartiesHandler(temp, partyserv)

	eventRep := eventRepo.NewEventRepository(dbconn)
	eventserv := eventServ.NewEventService(eventRep)
	eventHandle := handler.NewEventHandler(temp, eventserv)

	fs := http.FileServer(http.Dir("ui/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	//http.HandleFunc("/upost", postHandler.PostNew)
	http.HandleFunc("/posts", postHandler.Posts)

	http.HandleFunc("/parties", partyHandler.PartiesNew)
	http.HandleFunc("/party", partyHandler.Parties)
	http.HandleFunc("/", index)
	//http.HandleFunc("/newevent",newEvnet)

	http.HandleFunc("/events", eventHandle.Events)
	http.HandleFunc("/newevent", eventHandle.EventNew)

	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8181", nil)
}
