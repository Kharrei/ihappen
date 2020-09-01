package app

import "time"

// A single event
type Activity struct {
	ID         int         `json:"id"`
	Title      string      `json:"title"`
	ImageURL   string      `json:"imageURL"`
	Theme      string      `json:"theme"`
	Details    string      `json:"details"`
	Schedule   []*Schedule `json:"schedule"`
	Host       *Host       `json:"host"`
	Paid       bool        `json:"paid"`
	Ticket     []*Ticket   `json:"tickets"`
	StartDate  time.Time   `json:"startDate"`
	EndDate    time.Time   `json:"endDate"`
	Slots      int         `json:"slots"`
	Location   string      `json:"location"`
	Venue      string      `json:"venue"`
	Activities []*Tag      `json:"activities"`
	Online     bool        `json:"online"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}

// A new event
type NewActivity struct {
	Title      string `json:"title"`
	//Image      []*ImageUpload `json:"attachment"`
	Theme     string    `json:"theme"`
	Details   string    `json:"details"`
	HostID    int       `json:"hostId"`
	Paid      bool      `json:"paid"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Slots     int       `json:"slots"`
	Location  string    `json:"location"`
	Venue     string    `json:"venue"`
	Online    bool      `json:"online"`
	CreatedAt time.Time `json:"createdAt"`
}

//UpdateActivity is a request to update an existing event
type UpdateActivity struct {
	ID         int    `route:"id"`
	Title      string `json:"title"`
	//Image      []*ImageUpload `json:"attachment"`
	Theme     string    `json:"theme"`
	Details   string    `json:"details"`
	HostID    int       `json:"hostId"`
	Paid      bool      `json:"paid"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Slots     int       `json:"slots"`
	Location  string    `json:"location"`
	Venue     string    `json:"venue"`
	Online    bool      `json:"online"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// A request to delete an existing event
type DeleteActivity struct {
	ID int `route:"id"`
}

// An event host
type Host struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Phone     int       `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateEditHost struct {
	ID    int    `route:"id"`
	Name  string `json:"name"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
}

type DeleteHost struct {
	ID int `json:"id"`
}

type Ticket struct {
	ID       int       `json:"id"`
	Name     string    `json:"string"`
	Details string `json:"string"`
	Deadline time.Time `json:"deadline"`
	Price    float64   `json:"price"`
	Currency   string      `json:"currency"`
}

type CreateEditTicket struct {
	ID       int       `route:"id"`
	Name     string    `json:"string"`
	Deadline time.Time `json:"deadline"`
	Price    float64   `json:"price"`
}

type DeleteTicket struct {
	ID int `json:"id"`
}

type Schedule struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Theme     string    `json:"theme"`
	Details   string    `json:"details"`
	Format    string    `json:"format"`
	Paid      bool      `json:"paid"`
	Currency  string    `json:"currency"`
	Charges   float64   `json:"charges"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Slots     int       `json:"slots"`
	Venue     string    `json:"venue"`
	Speaker   string    `json:"speaker"`
}

type NewSchedule struct {
	ActivityID int       `json:"activityId"`
	Title      string    `json:"title"`
	Theme      string    `json:"theme"`
	Details    string    `json:"details"`
	Format     string    `json:"format"`
	Paid       bool      `json:"paid"`
	Currency   string    `json:"currency"`
	Charges    float64   `json:"charges"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	Slots      int       `json:"slots"`
	Venue      string    `json:"venue"`
	Speaker    string    `json:"speaker"`
}

type UpdateSchedule struct {
	ID        int       `route:"id"`
	Title     string    `json:"title"`
	Theme     string    `json:"theme"`
	Details   string    `json:"details"`
	Format    string    `json:"format"`
	Paid      bool      `json:"paid"`
	Currency  string    `json:"currency"`
	Charges   float64   `json:"charges"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Slots     int       `json:"slots"`
	Venue     string    `json:"venue"`
	Speaker   string    `json:"speaker"`
}

type DeleteSchedule struct {
	ID int `route:"id"`
}

type Tag struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color" format:"upper"`
}

type CreateTag struct {
	Name  string `json:"name"`
	Color string `json:"color" format:"upper"`
}

type EditTag struct {
	ID    int    `route:"id"`
	Name  string `json:"name"`
	Color string `json:"color" format:"upper"`
}

type DeleteTag struct {
	ID int `route:"id"`
}

type AssignUnassignTag struct {
	ID         int `route:"id"`
	ActivityID int `route:"activityId"`
}

//ImageUpload is the input model used to upload/remove an image
type ImageUpload struct {
	BlobKey string           `json:"bkey"`
	Upload  *ImageUploadData `json:"upload"`
	Remove  bool             `json:"remove"`
}

//ImageUploadData is the input model used to upload a new logo
type ImageUploadData struct {
	FileName    string `json:"fileName"`
	ContentType string `json:"contentType"`
	Content     []byte `json:"content"`
}
