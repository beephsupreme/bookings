package repository

import "github.com/beephsupreme/bookings/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}