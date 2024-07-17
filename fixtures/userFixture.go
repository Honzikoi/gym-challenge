package fixtures

/*
func CreateUserFixture() error {
	password, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password: ", err)
	}

	user := models.Users{
		Username: "test",
		Email:    "test@test.test",
		Password: string(password),
	}

	result := database.DB.Db.Create(&user)
	if result.Error != nil {
		log.Fatal("Failed to create user fixture: ", result.Error)
	}

	log.Println("User fixture created successfully")
	return nil
}
*/
