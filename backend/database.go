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

func getUserTask(task *Task) ([]Task, error) {
	sql := "SELECT * FROM userlist WHERE username = $1"
	if task.Listname != "" {
		sql += " AND listname LIKE '%" + task.Listname + "%'"
	}
	if task.Deadline != "" {
		sql += " AND deadline = '" + task.Deadline + "'"
	}
	rows, err := db.Query(
		sql, task.Username,
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

func updateTask(task *Task) (Task, error) {
	sql := "UPDATE userlist SET"
	update := false
	if task.Listname != "" {
		sql += " listname = '" + task.Listname + "'"
		update = true
	}
	if task.Deadline != "" {
		sql += " deadline = '" + task.Deadline + "'"
		update = true
	}
	if task.Detail != "" {
		sql += " detail = '" + task.Detail + "'"
		update = true
	}
	if update {
		sql += " WHERE id = "
		_, err := db.Exec(
			sql+"$1", task.ID,
		)
		if err != nil {
			return Task{}, err
		}
	}
	row := db.QueryRow(
		"SELECT * FROM userlist WHERE id = $1", task.ID,
	)
	editTask := Task{}
	err = row.Scan(&editTask.ID, &editTask.Listname, &editTask.Deadline, &editTask.Deadline, &editTask.Username)
	if err != nil {
		return Task{}, err
	}
	return editTask, nil
}
