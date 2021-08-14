package models

import (
	"time"
)

//User is the user models
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//Room is the room model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//restriction is the restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

//Reservation is the reservation model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room
	Processed int
}

//RoomRestriction for room restriction models
type RoomRestriction struct {
	ID              int
	StartDate       time.Time
	EndDate         time.Time
	RoomID          int
	ReservationID   int
	RestrictionID   int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Room            Room
	Reservation     Reservation
	RoomRestriction Reservation
}

//MailData holds email message
type MailData struct {
	To       string
	From     string
	Subject  string
	Content  string
	Template string
}
