package main

import (

	"log"
	"net/http"
	"os"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
	"time"
)

type Flight struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	FlightCode string `json:"flightCode" bson:"flightCode"`
	Origin string `json:"origin" bson:"origin"`
  Destination string `json:"destination" bson:"destination"`
  Price float32 `json:"price" bson:"price"`
  Currency string `json:"currency" bson:"currency"`
  Date time.Time `json:"date" bson:"date"`
	Passengers int `json:"passengers" bson:"passengers"`
	Capacity int `json:"capacity" bson:"capacity"`
}

func main() {

	port := os.Getenv("PORT")
	router := gin.Default()


	if port == "" {
		port = "8081"
	}

	session, err := mgo.Dial("mongodb://admin:admin@ds147551.mlab.com:47551/vianca-db")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	cc := session.DB("vianca-db").C("vuelo")
	var results []Flight

	router.POST("/reserve", func(c *gin.Context) {
		var json Flight
		c.BindJSON(&json)
		fmt.Println(json)
		// fmt.Println(json.currency)
		fmt.Println(json.Currency)
	})

	router.GET("/search/origin/:ciudadori", func(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Next()
	ciudadori := c.Param("ciudadori")
	err = cc.Find(bson.M{"origin": ciudadori}).All(&results)
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")

  c.Next()

	if err != nil {
		log.Fatal(err)
	}

		c.JSON(http.StatusOK, gin.H{"code": "12345", "name": "Vianca Airlines", "thumbnail":"https://image.freepik.com/icones-gratuites/avion-noir_318-31722.jpg", "results": results})
	})

	router.Run(":" + port)

}
