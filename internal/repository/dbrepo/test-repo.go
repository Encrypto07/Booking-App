package dbrepo

import (
	"errors"
	"time"

	"github.com/Encrypto07/Booking-App/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

//InsertReservation inserts a reservation into database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

//InsertRoomRestriction insert a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

//SearchAvailabilityByDates returns true if avaliability exists for roomID and false  for roomID if no availability exists
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

//SearchAvailabilityForAllRooms returns a slice of avaliable rooms if any for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

//GetRoomByID gets room by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("some error")
	}

	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	return 1, "", nil
}

//AllReservations returns a slice of all reservation
func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

//AllNewReservations returns a slice of new reservation
func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

//GetReservationByID returns one reservation by id
func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	var res models.Reservation
	return res, nil
}

//UpdateReservation updates a reservation in the database
func (m *testDBRepo) UpdateReservation(u models.Reservation) error {
	return nil
}

//DeleteReservation delete one reservationsby id
func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

//UpdateProcessedForReservation updates process for reservation by id
func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}
