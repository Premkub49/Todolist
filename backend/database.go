package main

func createUser(user *User) error {
	_, err := db.Exec(
		"INSERT INTO userdata (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password,
	)
	if err != nil {
		return err
	}
	return nil
}
func getUser(user *User) (User, error) {
	row := db.QueryRow(
		"SELECT * FROM userdata WHERE username = $1;", user.Username,
	)
	selectUser := User{}
	err = row.Scan(&selectUser.Username, &selectUser.Email, &selectUser.Password)
	if err != nil {
		return User{}, err
	}
	return selectUser, nil
}
