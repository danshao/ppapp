package main

import (
	"log"
	"os"
	"strconv"

	"github.com/danshao/ppapp/db"
	"github.com/danshao/ppapp/models"
	"github.com/gin-gonic/gin"
)

var MAXINT = strconv.Itoa(int(^uint(0) >> 1))

func main() {
	db.InitDb()
	router := gin.Default()

	router.GET("/items", func(c *gin.Context) {

		level := c.DefaultQuery("level", "%")
		job := c.Query("job")
		limit := c.DefaultQuery("limit", MAXINT)
		strength := c.DefaultQuery("strength", MAXINT)
		dexterity := c.DefaultQuery("dexterity", MAXINT)
		intelligence := c.DefaultQuery("intelligence", MAXINT)
		vitality := c.DefaultQuery("vitality", MAXINT)

		db := db.SharedConnection()
		items := []models.Item{}

		switch job {
		case "Barbarian":
			db.Order("strength desc, vitality desc").Limit(limit).Where("level LIKE ? AND strength <= ? AND dexterity <= ? AND intelligence <= ? AND vitality <= ?", level, strength, dexterity, intelligence, vitality).Find(&items)
			break
		case "Hunter":
			db.Order("dexterity desc, vitality desc").Limit(limit).Where("level LIKE ? AND strength <= ? AND dexterity <= ? AND intelligence <= ? AND vitality <= ?", level, strength, dexterity, intelligence, vitality).Find(&items)
			break
		case "Mage":
			db.Order("intelligence desc, vitality desc").Limit(limit).Where("level LIKE ? AND strength <= ? AND dexterity <= ? AND intelligence <= ? AND vitality <= ?", level, strength, dexterity, intelligence, vitality).Find(&items)
			break
		default:
			db.Limit(limit).Where("level LIKE ? AND strength <= ? AND dexterity <= ? AND intelligence <= ? AND vitality <= ?", level, strength, dexterity, intelligence, vitality).Find(&items)
			break
		}

		c.JSON(200, items)
	})

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Println("Running api server in dev mode")
	}

	router.Run(":8080")
}
