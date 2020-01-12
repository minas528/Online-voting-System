package main

import (
	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/rtoken"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	eventRepo "github.com/minas528/Online-voting-System/Event/repository"
	eventServ "github.com/minas528/Online-voting-System/Event/service"
	"github.com/minas528/Online-voting-System/delivery/http/handler"
	partyRepo "github.com/minas528/Online-voting-System/parties/repository"
	partyServ "github.com/minas528/Online-voting-System/parties/service"
	postRepo "github.com/minas528/Online-voting-System/post/repository"
	postServ "github.com/minas528/Online-voting-System/post/service"
	authRepo "github.com/minas528/Online-voting-System/voters/repository"
	authServ "github.com/minas528/Online-voting-System/voters/service"
)

var temp = template.Must(template.ParseGlob("ui/templates/*"))

func login(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "login", nil)
}

func signup(w http.ResponseWriter, r *http.Request)  {
	log.Println("just here")
	temp.ExecuteTemplate(w,"home.l.layout",nil)
}

func index(w http.ResponseWriter, r *http.Request)  {
	temp.ExecuteTemplate(w,"index.html",nil)
}
func newEvnet(w http.ResponseWriter, req *http.Request) {
	temp.ExecuteTemplate(w, "new.event", nil)
}
<<<<<<< HEAD

func parties(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "parties",nil)
}

func createTables(dbconn *gorm.DB) []error  {
	errs := dbconn.CreateTable(&entities.Voters{},&entities.Session{},&entities.Role{},&entities.Parties{},&entities.Events{},&entities.RegParties{},&entities.RegVoters{}).GetErrors()
	if errs != nil{
		return errs
	}
	return nil
=======
func parties(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "parties", nil)
}
func RoutesForAdmin() {
>>>>>>> 345e47e2cd443abfe01044e928281286fec9a418
}

func main() {
	scrfSingKey := []byte(rtoken.GenerateRandomID(32))

<<<<<<< HEAD
	dbconn, err := gorm.Open("postgres", "postgres://postgres:minpass@localhost:9090/electe?sslmode=disable")
=======
	dbconn, err := gorm.Open("postgres", "postgres://postgres:minpass@localhost:9090/votes?sslmode=disable")
>>>>>>> 345e47e2cd443abfe01044e928281286fec9a418
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

<<<<<<< HEAD
	//createTables(dbconn)

	sessionRepo := authRepo.NewSessionGormRepo(dbconn)
	sessionServ := authServ.NewSessionService(sessionRepo)

=======
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
>>>>>>> 345e47e2cd443abfe01044e928281286fec9a418

	postRepo := postRepo.NewPostGormRepo(dbconn)
	postserv := postServ.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(temp, postserv)

	partyRepo := partyRepo.NewPartiesGormRepo(dbconn)
	partyserv := partyServ.NewPartiesService(partyRepo)
	partyHandler := handler.NewAdminPartiesHandler(temp, partyserv)

	eventRep := eventRepo.NewEventRepository(dbconn)
	eventserv := eventServ.NewEventService(eventRep)
	eventHandle := handler.NewEventHandler(temp, eventserv)

	voterrole := authRepo.NewRoleGormRepo(dbconn)
	voterroleserv := authServ.NewRoleService(voterrole)
	sess := configSess()

	voterRepo := authRepo.NewVoterGormRepo(dbconn)
	voterserv := authServ.NewAuthService(voterRepo)
	aph := handler.NewAdminPostHandler(temp,postserv,scrfSingKey)
	vth := handler.NewVoterHandler(temp,voterserv,sessionServ,voterroleserv,sess,scrfSingKey)

	fs := http.FileServer(http.Dir("ui/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	//http.HandleFunc("/upost", postHandler.PostNew)
	http.HandleFunc("/posts", postHandler.Posts)
<<<<<<< HEAD
	http.HandleFunc("/", signup)
=======

	http.HandleFunc("/parties", partyHandler.PartiesNew)
	http.HandleFunc("/party", partyHandler.Parties)
	http.HandleFunc("/", index)
>>>>>>> 345e47e2cd443abfe01044e928281286fec9a418
	//http.HandleFunc("/newevent",newEvnet)

	http.HandleFunc("/events", eventHandle.Events)
	http.HandleFunc("/newevent", eventHandle.EventNew)

<<<<<<< HEAD
	http.HandleFunc("/signup", vth.Signup)
	http.HandleFunc("/voters", login)
	http.HandleFunc("/login", vth.Login)
	http.Handle("/logout",vth.Authenticated(http.HandlerFunc(vth.Logout)))
	http.Handle("/admin/addposts",vth.Authenticated(vth.Authorized(http.HandlerFunc(aph.PostNew))))
=======
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
>>>>>>> 345e47e2cd443abfe01044e928281286fec9a418
	http.ListenAndServe(":8181", nil)
	
}
func configSess() *entities.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := rtoken.GenerateRandomID(32)
	signingString, err := rtoken.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &entities.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}