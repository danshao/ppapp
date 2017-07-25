package db

import (
	"math/rand"
	"time"

	"github.com/danshao/ppapp/models"
)

// initialize the database with `Users`` and `Items`` tables. Will drop tables if they already exist.
func InitDb() {
	rand.Seed(time.Now().UTC().UnixNano())
	db := SharedConnection()

	db.DropTableIfExists(&models.User{})
	db.CreateTable(&models.User{})

	db.DropTableIfExists(&models.Item{})
	db.CreateTable(&models.Item{})
	populateItemsTable(20, 3)
}

// Populates the Items table given the max number of levels defined by `levels`. The number of items per level is varied by `entropyItemsPerLevel`
func populateItemsTable(levels int, entropyItemsPerLevel int) {
	x := 1
	for x <= levels {
		item := generateItemForLevel(x)
		db := SharedConnection()
		db.Create(&item)

		if a := rand.Intn(entropyItemsPerLevel + 1); a == 1 {
			x++
		}
	}
}

// Generate an item with random attributes for a specified `level`
func generateItemForLevel(level int) models.Item {
	db := SharedConnection()
	var item models.Item
	invalidItem := true

	// loop through until we generate a valid item
	for invalidItem {
		strength := rand.Intn(4 * level)
		dexterity := rand.Intn(4 * level)
		intelligence := rand.Intn(4 * level)
		vitality := rand.Intn(4 * level)

		// first requirement, the sum of item attributes cannot exceed four times the level
		if strength+dexterity+intelligence+vitality <= (4 * level) {
			// second requirement, no duplicate items
			var count int
			db.Model(&models.Item{}).Where("STRENGTH = ? AND DEXTERITY = ? AND INTELLIGENCE = ? AND VITALITY = ?", strength, dexterity, intelligence, vitality).Count(&count)

			if count == 0 {
				invalidItem = false
				item = models.Item{Level: level, Strength: strength, Dexterity: dexterity, Intelligence: intelligence, Vitality: vitality}
			}
		}
	}

	return item
}
