package db_test

import (
	"fmt"
	"github.com/drossan/core-api/utils"
	"testing"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestFormRepository_CreateOrUpdate(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewFormRepository(database)

	form := &model.Form{
		Title: "Test Form",
	}

	err := repo.CreateOrUpdate(form)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, form.ID)
}

func TestFormRepository_GetAll(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewFormRepository(database)

	form1 := &model.Form{
		Title: "Test Form 1",
	}
	form2 := &model.Form{
		Title: "Test Form 2",
	}

	err := repo.CreateOrUpdate(form1)
	assert.Nil(t, err)
	err = repo.CreateOrUpdate(form2)
	assert.Nil(t, err)

	forms, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Len(t, forms, 2)
}

func TestFormRepository_Paginate(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewFormRepository(database)

	for i := 1; i <= 25; i++ {
		form := &model.Form{
			Title: fmt.Sprintf("Test Form %d", i),
		}
		err := repo.CreateOrUpdate(form)
		assert.Nil(t, err)
	}

	forms, total, err := repo.Paginate(1, 10)
	assert.Nil(t, err)
	assert.Len(t, forms, 10)
	assert.Equal(t, 25, total)
}

func TestFormRepository_Delete(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewFormRepository(database)

	form := &model.Form{
		Title: "Test Form",
	}

	err := repo.CreateOrUpdate(form)
	assert.Nil(t, err)

	err = repo.Delete(form)
	assert.Nil(t, err)

	forms, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Len(t, forms, 0) // Verifica que la base de datos esté vacía
}
