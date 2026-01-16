//go:generate go run generate.go

package ent

import (
	"context"
	"fmt"
	"os"

	"github.com/benchplus/goorm/ent/user"
	"github.com/benchplus/goorm/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type EntORM struct {
	client *Client
	ctx    context.Context
}

func New() *EntORM {
	return &EntORM{
		ctx: context.Background(),
	}
}

func (e *EntORM) Init(dsn string) error {
	client, err := Open("sqlite3", dsn)
	if err != nil {
		return err
	}
	e.client = client
	return nil
}

func (e *EntORM) Close() error {
	return e.client.Close()
}

func (e *EntORM) CreateTable() error {
	return e.client.Schema.Create(e.ctx)
}

func (e *EntORM) DropTable() error {
	// Delete all records (ENT doesn't provide direct table drop)
	_, _ = e.client.User.Delete().Exec(e.ctx)
	return nil
}

func (e *EntORM) Insert(userModel *models.User) error {
	u, err := e.client.User.
		Create().
		SetName(userModel.Name).
		SetEmail(userModel.Email).
		SetAge(userModel.Age).
		Save(e.ctx)
	if err != nil {
		return err
	}
	userModel.ID = u.ID
	return nil
}

func (e *EntORM) InsertBatch(users []*models.User) error {
	if len(users) == 0 {
		return nil
	}
	builders := make([]*UserCreate, len(users))
	for i, u := range users {
		builders[i] = e.client.User.
			Create().
			SetName(u.Name).
			SetEmail(u.Email).
			SetAge(u.Age)
	}
	createdUsers, err := e.client.User.CreateBulk(builders...).Save(e.ctx)
	if err != nil {
		return err
	}
	for i, u := range createdUsers {
		users[i].ID = u.ID
	}
	return nil
}

func (e *EntORM) GetByID(id int64) (*models.User, error) {
	u, err := e.client.User.Get(e.ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Age:   u.Age,
	}, nil
}

func (e *EntORM) GetByIDs(ids []int64) ([]*models.User, error) {
	users, err := e.client.User.Query().
		Where(user.IDIn(ids...)).
		All(e.ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*models.User, len(users))
	for i, u := range users {
		result[i] = &models.User{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Age:   u.Age,
		}
	}
	return result, nil
}

func (e *EntORM) Update(userModel *models.User) error {
	_, err := e.client.User.
		UpdateOneID(userModel.ID).
		SetName(userModel.Name).
		SetEmail(userModel.Email).
		SetAge(userModel.Age).
		Save(e.ctx)
	return err
}

func (e *EntORM) Delete(id int64) error {
	return e.client.User.DeleteOneID(id).Exec(e.ctx)
}

func (e *EntORM) Count() (int64, error) {
	count, err := e.client.User.Query().Count(e.ctx)
	return int64(count), err
}

func (e *EntORM) GetAll(limit, offset int) ([]*models.User, error) {
	users, err := e.client.User.Query().
		Limit(limit).
		Offset(offset).
		All(e.ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*models.User, len(users))
	for i, u := range users {
		result[i] = &models.User{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Age:   u.Age,
		}
	}
	return result, nil
}

// GetDSN 生成测试用的 DSN
func GetDSN() string {
	return fmt.Sprintf("file:%s?cache=shared&mode=memory&_fk=1", getTempFile())
}

func getTempFile() string {
	tmpfile, _ := os.CreateTemp("", "ent_*.db")
	tmpfile.Close()
	return tmpfile.Name()
}
