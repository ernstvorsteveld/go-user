package repository

import (
	models "github.com/ernstvorsteveld/go-user"
)

var JohnUsername string = "john.doe"
var SanderUsername = "sander.kuiper"
var AnneliesUsername = "annelies.gerritsen"

func createUsers(userRepository UserRepository) map[string]models.User {
	userMap := make(map[string]models.User)

	user := models.NewUser(JohnUsername)
	userRepository.Create(user)
	userMap[JohnUsername] = user

	user = models.NewUser(SanderUsername)
	userRepository.Create(user)
	userMap[SanderUsername] = user

	user = models.NewUser(AnneliesUsername)
	userRepository.Create(user)
	userMap[AnneliesUsername] = user

	return userMap
}
