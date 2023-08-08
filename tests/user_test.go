package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/williaminfante/go_test_starter/entity"
	"github.com/williaminfante/go_test_starter/service"
)


func (t * SuiteTest) TestCreateUser() {
	_, err := service.CreateUser(entity.User{
		Name: "username1",
		Email: "username1@gmail.com",
	})
	assert.NoError(t.T(), err)

	_, err = service.CreateUser(entity.User{
		Name: "username2",
		Email: "username2@gmail.com",
	})
	assert.NoError(t.T(), err)

	_, err = service.CreateUser(entity.User{
		Name: "username3",
		Email: "username3@gmail.com",
	})
	assert.NoError(t.T(), err)
	service.ClearAllUsers()
}


func (t *SuiteTest) TestGetAllUsers() {
	users, err := service.GetAllUsers()

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), 0, len(users))

	_, err = service.CreateUser(entity.User{
		Name: "test_rat1",
		Email: "test_rat@gmail.com",
	})

	assert.NoError(t.T(), err)

	user, err := service.GetUserByEmail("test_rat@gmail.com")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "test_rat@gmail.com", user.Email)
	assert.Equal(t.T(), "test_rat1", user.Name)
}