package models

import "github.com/Pedrogabriel152/Go/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := `SELECT * FROM todos WHERE id = $1`

	err = conn.QueryRow(row, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
