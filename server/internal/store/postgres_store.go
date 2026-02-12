package store

import (
	"context"
	"database/sql"
	"time"

	"final/internal/models"
)

type PostgresStore struct {
	DB *sql.DB
}

// NewPostgresStore creates a PostgreSQL-backed store.
func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{DB: db}
}

//////////////////////////////
// USERS
//////////////////////////////

func (ps *PostgresStore) CreateUser(u models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	q := `
		INSERT INTO users (name, email, password_hash, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`
	err := ps.DB.QueryRowContext(ctx, q, u.Name, u.Email, u.PasswordHash, u.Role).
		Scan(&u.ID, &u.CreatedAt)
	return u, err
}

func (ps *PostgresStore) GetUserByEmail(email string) (models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var u models.User
	q := `SELECT id, name, email, password_hash, role, created_at FROM users WHERE email=$1`
	err := ps.DB.QueryRowContext(ctx, q, email).
		Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.Role, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return models.User{}, false
	}
	if err != nil {
		return models.User{}, false
	}
	return u, true
}

func (ps *PostgresStore) GetUserByID(id int) (models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var u models.User
	q := `SELECT id, name, email, password_hash, role, created_at FROM users WHERE id=$1`
	err := ps.DB.QueryRowContext(ctx, q, id).
		Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.Role, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return models.User{}, false
	}
	if err != nil {
		return models.User{}, false
	}
	return u, true
}

func (ps *PostgresStore) ListUsers() []models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := ps.DB.QueryContext(ctx, `SELECT id, name, email, password_hash, role, created_at FROM users ORDER BY id`)
	if err != nil {
		return []models.User{}
	}
	defer rows.Close()

	out := []models.User{}
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.Role, &u.CreatedAt); err != nil {
			return []models.User{}
		}
		out = append(out, u)
	}
	return out
}

func (ps *PostgresStore) UpdateUser(u models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Keep CreatedAt unchanged by reading it first
	var created time.Time
	err := ps.DB.QueryRowContext(ctx, `SELECT created_at FROM users WHERE id=$1`, u.ID).Scan(&created)
	if err != nil {
		return models.User{}, ErrNotFound
	}

	q := `UPDATE users SET name=$1, email=$2, password_hash=$3, role=$4 WHERE id=$5`
	_, err = ps.DB.ExecContext(ctx, q, u.Name, u.Email, u.PasswordHash, u.Role, u.ID)
	if err != nil {
		return models.User{}, err
	}

	u.CreatedAt = created
	return u, nil
}

func (ps *PostgresStore) DeleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := ps.DB.ExecContext(ctx, `DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return ErrNotFound
	}
	return nil
}

//////////////////////////////
// BOOKS
//////////////////////////////

func (ps *PostgresStore) CreateBook(b models.Book) models.Book {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	q := `INSERT INTO books (title, price, author) VALUES ($1,$2,$3) RETURNING id, created_at`
	err := ps.DB.QueryRowContext(ctx, q, b.Title, b.Price, b.Author).
		Scan(&b.ID, &b.CreatedAt)
	if err != nil {
		return models.Book{}
	}
	return b
}

func (ps *PostgresStore) ListBooks() []models.Book {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := ps.DB.QueryContext(ctx, `SELECT id, title, price, author, created_at FROM books ORDER BY id`)
	if err != nil {
		return []models.Book{}
	}
	defer rows.Close()

	out := []models.Book{}
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Price, &b.Author, &b.CreatedAt); err != nil {
			return []models.Book{}
		}
		out = append(out, b)
	}
	return out
}

func (ps *PostgresStore) GetBook(id int) (models.Book, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var b models.Book
	err := ps.DB.QueryRowContext(ctx, `SELECT id, title, price, author, created_at FROM books WHERE id=$1`, id).
		Scan(&b.ID, &b.Title, &b.Price, &b.Author, &b.CreatedAt)
	if err == sql.ErrNoRows {
		return models.Book{}, false
	}
	if err != nil {
		return models.Book{}, false
	}
	return b, true
}

func (ps *PostgresStore) UpdateBook(b models.Book) (models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var created time.Time
	err := ps.DB.QueryRowContext(ctx, `SELECT created_at FROM books WHERE id=$1`, b.ID).Scan(&created)
	if err != nil {
		return models.Book{}, ErrNotFound
	}

	_, err = ps.DB.ExecContext(ctx, `UPDATE books SET title=$1, price=$2, author=$3 WHERE id=$4`,
		b.Title, b.Price, b.Author, b.ID)
	if err != nil {
		return models.Book{}, err
	}

	b.CreatedAt = created
	return b, nil
}

func (ps *PostgresStore) DeleteBook(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := ps.DB.ExecContext(ctx, `DELETE FROM books WHERE id=$1`, id)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return ErrNotFound
	}
	return nil
}

//////////////////////////////
// ORDERS
//////////////////////////////

func (ps *PostgresStore) CreateOrder(o models.Order) models.Order {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := ps.DB.BeginTx(ctx, nil)
	if err != nil {
		return models.Order{}
	}
	defer tx.Rollback()

	// If client did not set status, keep same behavior as memory store
	if o.Status == "" {
		o.Status = "processing"
	}

	err = tx.QueryRowContext(ctx,
		`INSERT INTO orders (user_id, total, status) VALUES ($1,$2,$3) RETURNING id, created_at`,
		o.UserID, o.Total, o.Status,
	).Scan(&o.ID, &o.CreatedAt)
	if err != nil {
		return models.Order{}
	}

	// Save items (no unit_price in your model)
	for _, it := range o.Items {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO order_items (order_id, book_id, quantity) VALUES ($1,$2,$3)`,
			o.ID, it.BookID, it.Quantity,
		)
		if err != nil {
			return models.Order{}
		}
	}

	if err := tx.Commit(); err != nil {
		return models.Order{}
	}

	return o
}

func (ps *PostgresStore) ListOrders() []models.Order {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := ps.DB.QueryContext(ctx, `SELECT id, user_id, total, status, created_at FROM orders ORDER BY id DESC`)
	if err != nil {
		return []models.Order{}
	}
	defer rows.Close()

	out := []models.Order{}
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.Total, &o.Status, &o.CreatedAt); err != nil {
			return []models.Order{}
		}
		o.Items = ps.getOrderItems(ctx, o.ID)
		out = append(out, o)
	}
	return out
}

// If your Store interface requires it, this method is here.
// It returns orders for a specific user.
func (ps *PostgresStore) ListOrdersByUser(userID int) []models.Order {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := ps.DB.QueryContext(ctx,
		`SELECT id, user_id, total, status, created_at FROM orders WHERE user_id=$1 ORDER BY id DESC`,
		userID,
	)
	if err != nil {
		return []models.Order{}
	}
	defer rows.Close()

	out := []models.Order{}
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.Total, &o.Status, &o.CreatedAt); err != nil {
			return []models.Order{}
		}
		o.Items = ps.getOrderItems(ctx, o.ID)
		out = append(out, o)
	}
	return out
}

func (ps *PostgresStore) GetOrder(id int) (models.Order, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var o models.Order
	err := ps.DB.QueryRowContext(ctx,
		`SELECT id, user_id, total, status, created_at FROM orders WHERE id=$1`,
		id,
	).Scan(&o.ID, &o.UserID, &o.Total, &o.Status, &o.CreatedAt)

	if err == sql.ErrNoRows {
		return models.Order{}, false
	}
	if err != nil {
		return models.Order{}, false
	}

	o.Items = ps.getOrderItems(ctx, o.ID)
	return o, true
}

func (ps *PostgresStore) UpdateOrder(o models.Order) (models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := ps.DB.BeginTx(ctx, nil)
	if err != nil {
		return models.Order{}, err
	}
	defer tx.Rollback()

	// Keep CreatedAt unchanged
	var created time.Time
	err = tx.QueryRowContext(ctx, `SELECT created_at FROM orders WHERE id=$1`, o.ID).Scan(&created)
	if err != nil {
		return models.Order{}, ErrNotFound
	}

	_, err = tx.ExecContext(ctx, `UPDATE orders SET user_id=$1, total=$2, status=$3 WHERE id=$4`,
		o.UserID, o.Total, o.Status, o.ID)
	if err != nil {
		return models.Order{}, err
	}

	// Replace items
	_, err = tx.ExecContext(ctx, `DELETE FROM order_items WHERE order_id=$1`, o.ID)
	if err != nil {
		return models.Order{}, err
	}

	for _, it := range o.Items {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO order_items (order_id, book_id, quantity) VALUES ($1,$2,$3)`,
			o.ID, it.BookID, it.Quantity,
		)
		if err != nil {
			return models.Order{}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return models.Order{}, err
	}

	o.CreatedAt = created
	return o, nil
}

func (ps *PostgresStore) DeleteOrder(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := ps.DB.ExecContext(ctx, `DELETE FROM orders WHERE id=$1`, id)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return ErrNotFound
	}
	return nil
}
func (ps *PostgresStore) UpdateOrderStatus(id int, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := ps.DB.ExecContext(ctx, `UPDATE orders SET status=$1 WHERE id=$2`, status, id)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return ErrNotFound
	}
	return nil
}

func (ps *PostgresStore) getOrderItems(ctx context.Context, orderID int) []models.OrderItem {
	rows, err := ps.DB.QueryContext(ctx,
		`SELECT book_id, quantity FROM order_items WHERE order_id=$1 ORDER BY id`,
		orderID,
	)
	if err != nil {
		return []models.OrderItem{}
	}
	defer rows.Close()

	items := []models.OrderItem{}
	for rows.Next() {
		var it models.OrderItem
		if err := rows.Scan(&it.BookID, &it.Quantity); err != nil {
			return []models.OrderItem{}
		}
		items = append(items, it)
	}
	return items
}
