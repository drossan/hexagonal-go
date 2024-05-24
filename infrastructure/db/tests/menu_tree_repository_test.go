package db_test

import (
	"fmt"
	"github.com/drossan/core-api/utils"
	"testing"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestMenuTreeRepository_CreateOrUpdate(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewMenuTreeRepository(database)

	menu := &model.MenuTree{
		Title: "Test Menu",
	}

	err := repo.CreateOrUpdate(menu)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, menu.ID)
}

func TestMenuTreeRepository_GetAll(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewMenuTreeRepository(database)

	menu1 := &model.MenuTree{
		Title: "Test Menu 1",
	}
	menu2 := &model.MenuTree{
		Title: "Test Menu 2",
	}

	err := repo.CreateOrUpdate(menu1)
	assert.Nil(t, err)
	err = repo.CreateOrUpdate(menu2)
	assert.Nil(t, err)

	menus, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Len(t, menus, 2)
}

func TestMenuTreeRepository_Paginate(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewMenuTreeRepository(database)

	for i := 1; i <= 25; i++ {
		menu := &model.MenuTree{
			Title: fmt.Sprintf("Test Menu %d", i),
		}
		err := repo.CreateOrUpdate(menu)
		assert.Nil(t, err)
	}

	menus, total, err := repo.Paginate(1, 10)
	assert.Nil(t, err)
	assert.Len(t, menus, 10)
	assert.Equal(t, 25, total)
}

func TestMenuTreeRepository_Delete(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewMenuTreeRepository(database)

	menu := &model.MenuTree{
		Title: "Test Menu",
	}

	err := repo.CreateOrUpdate(menu)
	assert.Nil(t, err)

	err = repo.Delete(menu)
	assert.Nil(t, err)

	menus, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Len(t, menus, 0) // Verifica que la base de datos esté vacía
}
