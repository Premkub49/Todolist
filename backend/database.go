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

func createTask(task *Task) error {
	_, err := db.Exec(
		"INSERT INTO userlist (listname,deadline,detail,username) VALUES ($1,$2,$3,$4)", task.Listname, task.Deadline, task.Detail, task.Username,
	)
	if err != nil {
		return err
	}
	return nil
}

func getUserTask(username string) ([]Task, error) {
	rows, err := db.Query(
		"SELECT * FROM userlist WHERE username = $1", username,
	)
	if err != nil {
		return nil, err
	}
	var Tasks []Task
	defer rows.Close()
	for rows.Next() {
		var task Task
		if err = rows.Scan(&task.ID, &task.Listname, &task.Deadline, &task.Detail, &task.Username); err != nil {
			return nil, err
		}
		Tasks = append(Tasks, task)
	}
	return Tasks, nil
}

func deleteTask(id int) error {
	_, err = db.Exec(
		"DELETE FROM userlist WHERE id = $1", id,
	)
	if err != nil {
		return err
	}
	return nil
}

func updateTask(task *Task) error {
	if task.Listname != "" {
		_, err = db.Exec(
			"UPDATE userlist SET listname = $1 WHERE id = $2", task.Listname, task.ID,
		)
		if err != nil {
			return err
		}
	}
	if task.Deadline != "" {
		_, err = db.Exec(
			"UPDATE userlist SET deadline = $1 WHERE id = $2", task.Deadline, task.ID,
		)
		if err != nil {
			return err
		}
	}
	_, err = db.Exec(
		"UPDATE userlist SET detail = $1 WHERE id = $2", task.Detail, task.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
