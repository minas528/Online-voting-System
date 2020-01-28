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



func signup(w http.ResponseWriter, r *http.Request) {
	log.Println("just here")
	temp.ExecuteTemplate(w, "home.l.layout", nil)
}
func test(w http.ResponseWriter,r *http.Request)  {
	temp.ExecuteTemplate(w,"update.party",nil)
}






func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable( &entities.Post{},&entities.Voters{},
		&entities.Parties{},&entities.Votes{},
		&entities.RegParties{},&entities.RegVoters{},
		&entities.Events{},&entities.Session{},
		&entities.Role{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}


func main() {
	scrfSingKey := []byte(rtoken.GenerateRandomID(32))


	dbconn, err := gorm.Open("postgres", "postgres://postgres:minpass@localhost:9090/electe?sslmode=disable")


	if err != nil {
		panic(err)
	}

	defer dbconn.Close()


	//createTables(dbconn)

	sessionRepo := authRepo.NewSessionGormRepo(dbconn)
	sessionServ := authServ.NewSessionService(sessionRepo)



	postRepo := postRepo.NewPostGormRepo(dbconn)
	postserv := postServ.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(temp, postserv)

	//adpsthand := handler.NewAdminPostHandler(temp,postserv,scrfSingKey)

	partyRepo := partyRepo.NewPartiesGormRepo(dbconn)
	partyserv := partyServ.NewPartiesService(partyRepo)
	partyHandler := handler.NewPartiesHandler(temp, partyserv)
	adminPartyhandler := handler.NewAdminPartiesHandler(temp,partyserv)

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



	http.Handle("/admin/users", vth.Authenticated(vth.Authorized(http.HandlerFunc(vth.AdminUsers))))
	http.Handle("/admin/users/new", vth.Authenticated(vth.Authorized(http.HandlerFunc(vth.AdminUsersNew))))
	http.Handle("/admin/users/update", vth.Authenticated(vth.Authorized(http.HandlerFunc(vth.AdminUsersUpdate))))
	http.Handle("/admin/users/delete", vth.Authenticated(vth.Authorized(http.HandlerFunc(vth.AdminUsersDelete))))
	http.Handle("/admin/posts",vth.Authenticated(vth.Authorized(http.HandlerFunc(aph.Posts))))
	http.Handle("/admin/posts/new",vth.Authenticated(vth.Authorized(http.HandlerFunc(aph.PostNew))))
	http.Handle("/admin/posts/update",vth.Authenticated(vth.Authorized(http.HandlerFunc(aph.AdminPostsUpdate))))
	http.Handle("/admin/posts/delete",vth.Authenticated(vth.Authorized(http.HandlerFunc(aph.AdminPostDelete))))
	http.Handle("/admin/events",vth.Authenticated(vth.Authorized(http.HandlerFunc(eventHandle.Events))))
	http.Handle("/admin/events/new",vth.Authenticated(vth.Authorized(http.HandlerFunc(eventHandle.EventNew))))
	http.Handle("/admin/events/update",vth.Authenticated(vth.Authorized(http.HandlerFunc(eventHandle.UpdateEvents))))
	http.Handle("/admin/events/delete",vth.Authenticated(vth.Authorized(http.HandlerFunc(eventHandle.DeleteEvents))))
	http.Handle("/admin/parties",vth.Authenticated(vth.Authorized(http.HandlerFunc(adminPartyhandler.AdminParties))))
	http.Handle("/admin/parties/new",vth.Authenticated(vth.Authorized(http.HandlerFunc(adminPartyhandler.PartiesNew))))
	http.Handle("/admin/parties/update",vth.Authenticated(vth.Authorized(http.HandlerFunc(adminPartyhandler.AdminPartiesUpdate))))
	http.Handle("/admin/parties/delete",vth.Authenticated(vth.Authorized(http.HandlerFunc(adminPartyhandler.AdminPartiesDelete))))


	http.Handle("/home",vth.Authenticated(vth.Authorized(http.HandlerFunc(signup))))
	http.Handle("/events",vth.Authenticated(vth.Authorized(http.HandlerFunc(eventHandle.Events))))
	http.Handle("/vote",vth.Authenticated(vth.Authorized(http.HandlerFunc(vth.Vote))))
	http.Handle("/posts",vth.Authenticated(vth.Authorized(http.HandlerFunc(postHandler.Posts))))
	http.Handle("/register",vth.Authenticated(vth.Authorized(http.HandlerFunc(vth.Register4Event))))
	http.Handle("/parties",vth.Authenticated(vth.Authorized(http.HandlerFunc(partyHandler.Parties))))
	http.Handle("/logout",vth.Authenticated(http.HandlerFunc(vth.Logout)))


	http.HandleFunc("/login", vth.Login)
	http.HandleFunc("/signup", vth.Signup)
	http.HandleFunc("/",test)



	//http.HandleFunc("/login", login)
	//http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8182", nil)
	
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