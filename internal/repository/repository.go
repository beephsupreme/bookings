package repository

import (
	"github.com/beephsupreme/bookings/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) error
	SearchAvailabilityByDates(start, end time.Time, roomID int) (bool, error)
}
