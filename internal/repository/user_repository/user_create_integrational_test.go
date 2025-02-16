package user_repository

import (
	"avito-shop/internal/database"
	"avito-shop/internal/entity"
	"context"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// Test_Create тест на получение всех данных
func Test_Create(t *testing.T) {
	err := godotenv.Load(filepath.Join("../../../", ".env"))

	assert.NoError(t, err)
	ctx := context.Background()

	// создаем соединение к бд для теста
	db, err := database.New(database.WithConn())
	// проверяем, что соединение было создано без ошибки
	assert.NoError(t, err)
	defer func() {
		if err := db.Close(ctx); err != nil {
			t.Errorf("Ошибка при закрытии базы данных: %v", err)
		}
	}()

	// каждый тест запускаем отдельной транзакцией в БД
	tx, err := db.Begin(ctx)
	assert.NoError(t, err)

	// после теста транзакцию откатываем, чтобы в Бд ничего не сохранилось
	defer tx.Rollback(ctx)

	// инициализация репозитория
	repo := New(tx)

	// вызов тестируемого метода
	result, err := repo.CreateUser(ctx, "alex", "password_hash")
	assert.NoError(t, err)

	// проверка, что репозиторий возвращает корректные данные
	assert.Greater(t, result.ID, 0)
	assert.Equal(t, "alex", result.Username)
	assert.Equal(t, "password_hash", result.PasswordHash)

	// получение данных из бд
	dataInDB := getData(ctx, t, tx)

	// проверка, что в базу вставлены корректные данные
	assert.Equal(t, "alex", dataInDB[0].Username)
	assert.Greater(t, len(dataInDB[0].PasswordHash), 0)
}

// Функция для получения вставленных данных
func getData(ctx context.Context, t *testing.T, db database.DataBase) []entity.User {
	var users []entity.User

	rows, err := db.Query(
		ctx,
		`SELECT "id", "username", "password_hash" FROM "users"`,
	)
	assert.NoError(t, err)

	defer rows.Close()

	for rows.Next() {
		user := entity.User{}
		var id int

		err = rows.Scan(
			&id,
			&user.Username,
			&user.PasswordHash,
		)

		assert.NoError(t, err)

		user.ID = id

		users = append(users, user)
	}

	return users
}
