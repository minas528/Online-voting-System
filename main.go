package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/minas528/Online-voting-System/delivery/http/handler"
	"github.com/minas528/Online-voting-System/post/repository"
	"github.com/minas528/Online-voting-System/post/service"
	"html/template"
	"net/http"
)





var temp = template.Must(template.ParseGlob("ui/templates/*"))



func main()  {

	dbconn,err := gorm.Open("postgres","postgres://postgres:minpass@localhost:9090/votes?sslmode=disable")
	if err != nil{
		panic(err)
	}

	defer dbconn.Close()

	//errs := dbconn.CreateTable(&entities.Post{}).GetErrors()
	//if 0 < len(errs) {
	//	panic(errs)
	//}

	postRepo := repository.NewPostGormRepo(dbconn)
	postserv := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(temp,postserv)





	fs := http.FileServer(http.Dir("ui/assets/"))
	http.Handle("/assets/",http.StripPrefix("/assets",fs))
	http.HandleFunc("/upost",postHandler.PostNew)
	http.HandleFunc("/posts",postHandler.Posts)
	//http.HandleFunc("/",index)
	http.ListenAndServe(":8181",nil)
}





