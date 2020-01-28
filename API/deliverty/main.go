package main

import (
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"github.com/minas528/API/Event/repository"
	"github.com/minas528/API/Event/service"
	"github.com/minas528/API/deliverty/http/handler"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main()  {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:minpass@localhost:9090/elect?sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	//errs :=dbconn.CreateTable(&entity.Events{}).GetErrors()
	//if len(errs)>0 {
	//	panic(errs)
	//}

	eventRepo := repository.NewEventRepository(dbconn)
	eventServ := service.NewEventService(eventRepo)

	eventHanlder := handler.NewEventHandler(eventServ)

	router := httprouter.New()

	router.GET("/events",eventHanlder.GetEvents)
	router.GET("/events/:id",eventHanlder.GetSingleEvent)
	router.PUT("/events/:id",eventHanlder.PutEvent)
	router.POST("/events",eventHanlder.PostEvent)
	router.DELETE("/events/:id",eventHanlder.DeleteEvents)

	http.ListenAndServe(":8282",router)

}


