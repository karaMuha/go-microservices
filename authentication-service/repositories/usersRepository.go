package repositories

import (
	"authentication/models"
	"database/sql"
	"net/http"
	"time"
)

type UsersRepository struct {
	dbHandler *sql.DB
}

func NewUsersRepository(dbHandler *sql.DB) UsersRepositoryInterface {
	return &UsersRepository{
		dbHandler: dbHandler,
	}
}

func (ur UsersRepository) QueryCreateUser(user *models.User, hashedPassword string) *models.ResponseError {
	query := `
		INSERT INTO
			users(email, first_name, last_name, user_password, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING id`
	row := ur.dbHandler.QueryRow(query, user.Email, user.FirstName, user.LastName, hashedPassword, user.CreatedAt, user.UpdatedAt)

	var userId string
	err := row.Scan(&userId)

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (ur UsersRepository) QueryGetAllUsers() ([]*models.User, *models.ResponseError) {
	query := `
		SELECT
			email,
			first_name,
			last_name,
			is_active,
			created_at,
			updated_at
		FROM
			users`
	rows, err := ur.dbHandler.Query(query)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	users := make([]*models.User, 0)
	var id, email, firstName, lastName string
	var active bool
	var createdAt, updatedAt time.Time

	for rows.Next() {
		err := rows.Scan(&id, &email, &firstName, &lastName, &active, &createdAt, &updatedAt)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		user := &models.User{
			ID:        id,
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
			Active:    active,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return users, nil
}

func (ur UsersRepository) QueryGetUserByEmail(email string) (*models.User, *models.ResponseError) {
	query := `
		SELECT
			id,
			first_name,
			last_name,
			user_password,
			is_active,
			created_at,
			updated_at
		FROM
			users`
	row := ur.dbHandler.QueryRow(query)

	var user models.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Active, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	user.Email = email
	return &user, nil

}

func (ur UsersRepository) QueryUpdateUser(user *models.User) *models.ResponseError {
	query := `
		UPDATE
			users
		SET
			email = $1,
			first_name = $2,
			last_name = $3,
			is_active = $4,
			updated_at = $5
		WHERE
			id = $6`
	row, err := ur.dbHandler.Exec(query, user.Email, user.FirstName, user.LastName, user.Active, time.Now(), user.ID)

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	rowsAffected, err := row.RowsAffected()

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if rowsAffected == 1 {
		return &models.ResponseError{
			Message: "User not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}

func (ur UsersRepository) QueryDeleteUser(id string) *models.ResponseError {
	query := `
		DELETE FROM
			users
		WHERE
			id = $1`
	row, err := ur.dbHandler.Exec(query, id)

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	rowsAffected, err := row.RowsAffected()

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if rowsAffected == 1 {
		return &models.ResponseError{
			Message: "User not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}
