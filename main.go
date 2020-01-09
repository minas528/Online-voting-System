package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
<<<<<<< HEAD
	"../../../github.com/outThabox/Online-voting-System/delivery/http/handler"
	"../../../github.com/outThabox/Online-voting-System/post/repository"
	"../../../github.com/outThabox/Online-voting-System/post/service"
=======
	eventRepo "github.com/minas528/Online-voting-System/Event/repository"
	eventServ "github.com/minas528/Online-voting-System/Event/service"
	"github.com/minas528/Online-voting-System/delivery/http/handler"
	postRepo "github.com/minas528/Online-voting-System/post/repository"
	postServ "github.com/minas528/Online-voting-System/post/service"
>>>>>>> 22a1904d57f4055e35f6cb753c2113699a2fb359
	"html/template"
	"net/http"
)


var temp = template.Must(template.ParseGlob("ui/templates/*"))

func index(w http.ResponseWriter, r *http.Request)  {
	temp.ExecuteTemplate(w,"",nil)
}
func newEvnet(w http.ResponseWriter,req *http.Request)  {
	temp.ExecuteTemplate(w,"new.event",nil)
}

func RoutesForAdmin()  {

}
func main()  {

	dbconn,err := gorm.Open("postgres","postgres://postgres:default@localhost:5432/votes?sslmode=disable")
	if err != nil{
		panic(err)
	}

	defer dbconn.Close()

	//errs := dbconn.CreateTable(&entities.Post{},&entities.Events{}).GetErrors()
	//if 0 < len(errs) {
	//	panic(errs)
	//}

<<<<<<< HEAD
	//errs := dbconn.CreateTable(&entities.Vote{}).GetErrors()
	//if 0 < len(errs) {
	//	panic(errs)
	//}

	postRepo := repository.NewPostGormRepo(dbconn)
	postserv := service.NewPostService(postRepo)
=======
	postRepo := postRepo.NewPostGormRepo(dbconn)
	postserv := postServ.NewPostService(postRepo)
>>>>>>> 22a1904d57f4055e35f6cb753c2113699a2fb359
	postHandler := handler.NewPostHandler(temp,postserv)

	eventRep := eventRepo.NewEventRepository(dbconn)
	eventserv := eventServ.NewEventService(eventRep)
	eventHandle := handler.NewEventHandler(temp,eventserv)





	fs := http.FileServer(http.Dir("ui/assets/"))
	http.Handle("/assets/",http.StripPrefix("/assets",fs))
	http.HandleFunc("/upost",postHandler.PostNew)
	http.HandleFunc("/posts",postHandler.Posts)
	http.HandleFunc("/",index)
	//http.HandleFunc("/newevent",newEvnet)

	http.HandleFunc("/events",eventHandle.Events)
	http.HandleFunc("/newevent",eventHandle.EventNew)
	http.ListenAndServe(":8181",nil)
}





