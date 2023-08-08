package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/williaminfante/go_test_starter/entity"
	"github.com/williaminfante/go_test_starter/service"
)


func (t * SuiteTest) TestCreateUser() {
	service.ClearAllUsers()
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
		Email: "username2@gmail.com",
	})
	assert.Error(t.T(), err)

	service.ClearAllUsers()
}


func (t *SuiteTest) TestGetAllUsers() {
	_, err := service.GetAllUsers()

	assert.NoError(t.T(), err)
}