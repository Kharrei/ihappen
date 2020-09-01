package actions

import (
	"context"
	"database/sql"
	"log"
)

type Activity struct {
	Model *models.Activity
}

func (input *Activity) Initialize() interface{}  {
	input.Model = new(models.Activity)
	return input.model
}

func (input *Activity) GetActivity(db *sql.DB) error {
	return db.QueryRow("SELECT title, imageUrl, theme, details, hostId, paid, startDate, endDate, slots, location, venue, online, createdAt, updatedAt FROM activities WHERE id=$1", input.Model.ID).Scan(&input.Model.Title, &input.Model.ImageURL, &input.Model.Theme, &input.Model.Details, &input.Model.hostID, &input.Model.Paid, &input.Model.StartDate, &input.Model.EndDate, &input.Model.Slots, &input.Model.Location, &input.Model.Venue, &input.Model.Online, &input.Model.CreatedAt, &input.Model.UpdatedAt)
}

type UpdateActivity struct {
	Model *models.UpdateActivity
}

func (input *UpdateActivity) Initialize() interface{} {
	input.Model = new(*models.UpdateActivity)
	return input.Model
}

func (input *UpdateActivity) Validate(db *sql.DB) error {
	stmt, err := db.Prepare("UPDATE activities SET title=?, categoryId=?, theme=?, details=?, hostId=?, paid=?, currency=?, startDate=?, endDate=?, slots=?, location=?, venue=?, online=?, updatedAt=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(input.Model.Title, input.Model.CategoryID, input.Model.Theme, input.Model.Details, input.Model.HostID, input.Model.Paid, input.Model.Currency, input.Model.StartDate, input.Model.EndDate, input.Model.Slots, input.Model.Location, input.Model.Venue, input.Model.Online, input.Model.UpdatedAt, input.Model.ID); err != nil {
		return err
	}

	return nil
}

type DeleteActivity struct {
	Model *models.DeleteActivity
}

func (input *DeleteActivity) Initialize() interface{}  {
	input.Model = new(models.DeleteActivity)
	return input.model
}

func (input *DeleteActivity) Validate(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM activities WHERE id=$1", input.Model.id)

	return err
}

type NewActivity struct {
	Model *models.UpdateActivity
}

func (input *NewActivity) Initialize() interface{} {
	input.Model = new(*models.UpdateActivity)
	return input.Model
}

func (input *NewActivity) Validate(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO activities(title, imageUrl, theme, details, hostId, paid, startDate, endDate, slots, location, venue, online, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(input.Model.Title, input.Model.ImageURL, input.Model.Theme, input.Model.Details, input.Model.HostID, input.Model.Paid, input.Model.StartDate, input.Model.EndDate, input.Model.Slots, input.Model.Location, input.Model.Venue, input.Model.Online, input.Model.UpdatedAt); err != nil {
		return err
	}

	return nil
}




