//go:build integration
package repository

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/GabriellGds/go-orders/internal/models"
)

func Test_PingDb(t *testing.T) {
	err := testDb.Ping()
	if err != nil {
		t.Error("can't ping database")
	}
}

func Test_CreateUserRepository(t *testing.T) {
	tests := []struct {
		user models.User
		id   int
	}{
		{
			user: models.User{
				Name:      "Gabriel",
				Email:     "gabriel@test.com",
				Password:  "12345678",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			id: 1,
		},
		{
			user: models.User{
				Name:      "carlos",
				Email:     "carlos@test.com",
				Password:  "12345678",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			id: 2,
		},
		{
			user: models.User{
				Name:      "marcos",
				Email:     "marcos@test.com",
				Password:  "12345678",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			id: 3,
		},
		{
			user: models.User{
				Name:      "maria",
				Email:     "maria@test.com",
				Password:  "12345678",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			id: 4,
		},
		{
			user: models.User{
				Name:      "",
				Email:     "rafa@test.com",
				Password:  "12345678",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			id: 5,
		},
	}
	ctx := context.Background()

	for _, e := range tests {
		e.user.HashPassword(e.user.Password)

		id, err := testRepo.CreateUserRepository(ctx, e.user)
		if err != nil {
			t.Error("CreateUserRepository returned an error: ", err)
		}

		if id != e.id {
			t.Errorf("CreateUserRepository returned wrong id; expected %d, but got %d", e.id, id)
		}
	}
}

func Test_ListUserRepository(t *testing.T) {
	ctx := context.Background()
	users, err := testRepo.ListUserRepository(ctx)
	if err != nil {
		t.Error("ListUserRepository reports an error: ", err)
	}

	if len(users) != 5 {
		t.Errorf("ListUsersRepository reports wrong size; expected 5, but got %d", len(users))
	}
}

func Test_FindUserRepository(t *testing.T) {
	tests := []struct {
		id    int
		email string
	}{
		{id: 1, email: "gabriel@test.com"},
		{id: 2, email: "carlos@test.com"},
		{id: 3, email: "marcos@test.com"},
		{id: 4, email: "maria@test.com"},
		{id: 5, email: "rafa@test.com"},
	}

	ctx := context.Background()

	for _, e := range tests {
		user, err := testRepo.FindUserRepository(ctx, e.id)
		if err != nil {
			t.Error("error FindUserRepository: ", err)
		}

		if user.Email != e.email {
			t.Errorf("wrong email returned by FindUserRepository; expected %s, but got %s", e.email, user.Email)
		}
	}
}

func Test_FindUserByEmailRepository(t *testing.T) {
	tests := []struct {
		id    int
		email string
	}{
		{id: 1, email: "gabriel@test.com"},
		{id: 2, email: "carlos@test.com"},
		{id: 3, email: "marcos@test.com"},
		{id: 4, email: "maria@test.com"},
		{id: 5, email: "rafa@test.com"},
	}

	ctx := context.Background()

	for _, e := range tests {
		user, err := testRepo.FindUserByEmailRepository(ctx, e.email)
		if err != nil {
			t.Error("error findUserByEmailRepositor", err)
		}

		if user.ID != e.id {
			t.Errorf("wrong id returned by findUserEmailRepository, expected %d but got %d", e.id, user.ID)
		}
	}
}

func Test_UpdateUserRepository(t *testing.T) {
	tests := []struct {
		id   int
		user *models.User
	}{
		{id: 1, user: &models.User{Name: "ceo"}},
		{id: 2, user: &models.User{Name: "carol"}},
		{id: 3, user: &models.User{Name: "tulo"}},
		{id: 4, user: &models.User{Name: "dry"}},
		{id: 5, user: &models.User{Name: "wet"}},
	}

	ctx := context.Background()

	for _, e := range tests {
		err := testRepo.UpdateUserRepository(ctx, e.id, e.user)
		if err != nil {
			t.Error("error UpadteUserRepository: ", err)
		}

		user, _ := testRepo.FindUserRepository(ctx, e.id)
		if user.Name != e.user.Name {
			t.Errorf("expected updated record to have name %s, but got %s", e.user.Name, user.Name)
		}
	}
}

func Test_DeleteUserRepository(t *testing.T) {
	tests := []struct {
		id int
	}{
		{id: 1},
		{id: 2},
		{id: 3},
		{id: 4},
		{id: 5},
	}
	ctx := context.Background()

	for _, e := range tests {
		err := testRepo.DeleteUserRepository(ctx, e.id)
		if err != nil {
			t.Error("error DeleteUsereRepository: ", err)
		}

		_, err = testRepo.FindUserRepository(ctx, e.id)
		if err == nil {
			t.Errorf("retrieved user id %d, who should have been deleted", e.id)
		}
	}
}

func Test_CreateItemRepository(t *testing.T) {
	tests := []struct {
		id   int
		item models.Item
	}{
		{id: 1, item: models.Item{Name: "hamburguer", Price: 15.0}},
		{id: 2, item: models.Item{Name: "hot dog", Price: 25.0}},
		{id: 3, item: models.Item{Name: "panqueca", Price: 32.5}},
		{id: 4, item: models.Item{Name: "churrasco", Price: 49.90}},
		{id: 5, item: models.Item{Name: "chips", Price: 112.58}},
	}
	ctx := context.Background()

	for _, e := range tests {
		item, err := testRepo.CreateItemRepository(ctx, &e.item)
		if err != nil {
			t.Error("error CreateitemRepository: ", err)
		}

		if item.ID != e.id {
			t.Errorf("CreateItemRepository returned wrong id; expected %d, but got %d", e.id, item.ID)
		}
	}
}

func Test_FindItemRepository(t *testing.T) {
	tests := []struct {
		id   int
		name string
	}{
		{id: 1, name: "hamburguer"},
		{id: 2, name: "hot dog"},
		{id: 3, name: "panqueca"},
		{id: 4, name: "churrasco"},
		{id: 5, name: "chips"},
	}
	ctx := context.Background()

	for _, e := range tests {
		item, err := testRepo.FindItemRepository(ctx, e.id)
		if err != nil {
			t.Error("error FindItemRepository: ", err)
		}

		if item.Name != e.name {
			t.Errorf("wrong name returned by findItemRepository, expected %s, but got %s", e.name, item.Name)
		}
	}
}

func Test_ItemsRepository(t *testing.T) {
	ctx := context.Background()

	items, err := testRepo.ItemsRepository(ctx)
	if err != nil {
		t.Error("error ItemsRepository")
	}

	if len(items) != 5 {
		t.Errorf("ItemsRepository reports wrong size; expected 5, but got %d", len(items))
	}
}

func Test_DeleteItemRepository(t *testing.T) {
	tests := []struct {
		id int
	}{
		{id: 1},
		{id: 2},
		{id: 3},
		{id: 4},
		{id: 5},
	}
	ctx := context.Background()

	for _, e := range tests {
		err := testRepo.DeleteItemRepository(ctx, e.id)
		if err != nil {
			t.Error("error DeleteItemRepository: ", err)
		}

		_, err = testRepo.FindItemRepository(ctx, e.id)
		if err == nil {
			t.Errorf("retrieved item id %d, who should have been deleted", e.id)
		}
	}
}


func Test_CreateOrderRepository(t *testing.T) {
	tests := []struct {
		order   *models.Order
		orderID int
	}{
		{
			order: &models.Order{
				UserID: 1,
				Items: []models.OrderItems{
					{ItemID: 1, Quantity: 2},
					{ItemID: 2, Quantity: 2},
				},
			},
			orderID: 1,
		},
		{
			order: &models.Order{
				UserID: 1,
				Items: []models.OrderItems{
					{ItemID: 3, Quantity: 2},
					{ItemID: 2, Quantity: 1},
				},
			},
			orderID: 2,
		},
		{
			order: &models.Order{
				UserID: 2,
				Items: []models.OrderItems{
					{ItemID: 3, Quantity: 2},
					{ItemID: 4, Quantity: 5},
				},
			},
			orderID: 3,
		},
	}

	ctx := context.Background()

	for _, e := range tests {
		orderID, err := testRepo.CreateOrderRepository(ctx, e.order)
		if err != nil {
			t.Error("error createOrderRepository: ", err)
		}

		if orderID != e.orderID {
			t.Errorf("CreateOrderRepository returned wrong id; expected %d, but got %d", e.orderID, orderID)
		}
	}
}

func FindOrderRepository(t *testing.T) {
	tests := []struct {
		orderID int
		order   models.Order
	}{
		{
			orderID: 1,
			order: models.Order{
				ID:     1,
				UserID: 1,
				Items: []models.OrderItems{
					{ItemID: 1, Quantity: 2},
					{ItemID: 2, Quantity: 2},
				},
			},
		},
		{
			orderID: 2,
			order: models.Order{
				ID:     2,
				UserID: 1,
				Items: []models.OrderItems{
					{ItemID: 3, Quantity: 2},
					{ItemID: 2, Quantity: 1},
				},
			},
		},
		{
			orderID: 3,
			order: models.Order{
				ID:     3,
				UserID: 2,
				Items: []models.OrderItems{
					{ItemID: 3, Quantity: 2},
					{ItemID: 4, Quantity: 5},
				},
			},
		},
	}
	ctx := context.Background()

	for _, e := range tests {
		order, err := testRepo.FindOrderRepository(ctx, e.orderID)
		if err != nil {
			t.Error("error FindOrderRepository: ", err)
		}

		if !reflect.DeepEqual(order, e.order) {
			t.Errorf("FindOrderRepository returned wrong order, expected %v, but got %v", e.order, order)
		}
	}
}

func DeleteOrderRepository(t *testing.T) {
	tests := []struct {
		id int
	}{
		{id: 1},
		{id: 2},
		{id: 3},
	}
	ctx := context.Background()

	for _, e := range tests {
		err := testRepo.DeleteOrderRepository(ctx, e.id)
		if err != nil {
			t.Error("error DeleteOrderRepository: ", err)
		}

		_, err = testRepo.FindOrderRepository(ctx, e.id)
		if err == nil {
			t.Errorf("retrieved order id %d, who should have been deleted", e.id)
		}
	}

}
