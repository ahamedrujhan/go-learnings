package models

import "Event_Management/db"

type Register struct {
	Event_id int
	User_id  int
}

func (register *Register) Save() error {
	query := `
INSERT INTO registrations (event_id, user_id)
VALUES (?,?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(register.Event_id, register.User_id)
	return err
}

func (register *Register) Delete() error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(register.Event_id, register.User_id)
	return err
}
