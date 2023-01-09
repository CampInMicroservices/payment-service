package db

import (
	"context"
	"time"
)

type Payment struct {
	ID        int64     `json:"id" db:"id"`
	BookingID int64     `json:"booking_id" db:"booking_id"`
	Price     float64   `json:"price" db:"price"`
	Paid      bool      `json:"paid" db:"paid"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreatePaymentParam struct {
	BookingID int64
	Price     float64
	Paid      bool
}

type UpdatePaymentParam struct {
	Paid bool
}

type ListPaymentParam struct {
	Offset int32
	Limit  int32
}

func (store *Store) GetPaymentByID(ctx context.Context, id int64) (payment Payment, err error) {

	const query = `SELECT * FROM "payments" WHERE "id" = $1`
	err = store.db.GetContext(ctx, &payment, query, id)

	return
}

func (store *Store) GetAllPayments(ctx context.Context, arg ListPaymentParam) (payments []Payment, err error) {

	const query = `SELECT * FROM "payments" OFFSET $1 LIMIT $2`
	payments = []Payment{}
	err = store.db.SelectContext(ctx, &payments, query, arg.Offset, arg.Limit)

	return
}

func (store *Store) CreatePayment(ctx context.Context, arg CreatePaymentParam) (Payment, error) {

	const query = `
	INSERT INTO "payments" ("booking_id", "price", "paid")
	VALUES ($1, $2, $3)
	RETURNING "id", "booking_id", "price", "paid", "created_at"
	`
	row := store.db.QueryRowContext(ctx, query, arg.BookingID, arg.Price, arg.Paid)

	var payment Payment
	err := row.Scan(
		&payment.ID,
		&payment.BookingID,
		&payment.Price,
		&payment.Paid,
		&payment.CreatedAt,
	)

	return payment, err
}

func (store *Store) UpdatePayment(ctx context.Context, arg UpdatePaymentParam, id int64) (Payment, error) {

	const query = `
	UPDATE "payments"
	SET "paid" = $1
	WHERE "id" = $2
	RETURNING "id", "booking_id", "price", "paid", "created_at"
	`
	row := store.db.QueryRowContext(ctx, query, arg.Paid, id)

	var payment Payment
	err := row.Scan(
		&payment.ID,
		&payment.BookingID,
		&payment.Price,
		&payment.Paid,
		&payment.CreatedAt,
	)

	return payment, err
}
