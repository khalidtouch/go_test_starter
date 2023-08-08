package tests

import (
	"testing"

	"os"

	"github.com/stretchr/testify/suite"
	"github.com/williaminfante/go_test_starter/config"
	"github.com/williaminfante/go_test_starter/entity"
	"gorm.io/gorm"
)

type (
	SuiteTest struct {
		suite.Suite 
		db *gorm.DB 
	}
)


func TestSuite(t *testing.T) {
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "IDeyTellYou555!")
	os.Setenv("DB_DATABASE", "go_mysql_suite_test")
	
	defer os.Unsetenv("DB_HOST")
	defer os.Unsetenv("DB_PORT")
	defer os.Unsetenv("DB_USER")
	defer os.Unsetenv("DB_PASS")
	defer os.Unsetenv("DB_DATABASE")


	suite.Run(t, new(SuiteTest))
}


func getModels() []interface{} {
	return []interface{}{&entity.User{}}
}

func (t *SuiteTest) SetupSuite() {
	config.ConnectGorm() 
	t.db = config.GetDb() 

	for _, val := range getModels() {
		t.db.AutoMigrate(val)
	}
}


//@After
func (t *SuiteTest) TearDownSuite() {
	database, _ := t.db.DB() 
	defer database.Close() 

	//drop table 
	for _, val := range getModels() {
		t.db.Migrator().DropTable(val)
	}
}


//@Before 
func (t *SuiteTest) SetupTest() {

}


//@After 
func (t *SuiteTest) TearDownTest() {

}