package setup

import "github.com/svarlamov/bintrad/models"

func setupUsers() error {
	sasha := models.User{
		Username:        "181001",
		FirstName:       "Sasha",
		LastName:        "Varlamov",
		Grade:           11,
		Gender:          "M",
		Passkey:         "12345",
		StudentId:       181001,
		Email:           "181001@hkis.edu.hk",
		StartingBalance: 1000000.1,
	}
	return sasha.Create()
}
