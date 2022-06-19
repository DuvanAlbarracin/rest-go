package main

import (
	"database/sql"
	"log"
	"rest-go/part4/metroRail/dbUtils"
	"time"

    "github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type trainModel struct{
    ID int
    DriverName string
    OperatingStatus bool
}

type stationModel struct{
    ID int
    Name string
    OpeningTime time.Time
    ClosingTime time.Time
}

type scheduleModel struct{
    ID int
    TrainId int
    StationId int
    Arrivaltime time.Time
}

func (t *trainModel) Register(container *restful.Container){
    ws := new(restful.WebService)
    ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
    ws.Route(ws.GET("{train-id}").To(t,getTrain))
    ws.Route(ws.POST("").To(t.createTrain))
    ws.Route(ws.DELETE("{train-id}").To(t.deleteTrain))
    container.Add(ws)
}

func main(){
   db, errDb :=sql.Open("sqlite3", "./railapi.db")
   if errDb != nil{
       log.Fatal("Error opening the database...", errDb)
   }
   dbutils.Initialize(db)
}
