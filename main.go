package main

import (

	"log"
	"net/http"
	"os"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type flight struct {

	Flightcode  string
	Origin string
	Destination string
	Price string
	Currency string

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
	var results []flight

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
