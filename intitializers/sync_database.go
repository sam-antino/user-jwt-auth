package intitializers

import (
	"user-jwt-auth/models/entities"
)

func SyncDatabase() {
	Db.AutoMigrate(&entities.Users{})
}
