package database

import (
	m "inventory/model"
	u "inventory/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectMysql : function to connect with mysql db and auto migrate data and seed it
func ConnectMysql() {
	var err error

	dbName := u.GetEnvByKey("DB_NAME")
	dbUser := u.GetEnvByKey("DB_USER")
	dbPass := u.GetEnvByKey("DB_PASS")

	dsn := dbUser + ":" + dbPass + "@tcp(127.0.0.1:3306)/" + dbName + "?parseTime=True&loc=Asia%2FJakarta"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Info),
		CreateBatchSize: 1000,
	})
	if err != nil {
		panic(err.Error())
	}

	var models = []interface{}{
		&m.Item{},
		&m.User{},
	}

	if gin.Mode() != gin.ReleaseMode {
		for _, v := range models {
			err = DB.Migrator().DropTable(v)
			if err != nil {
				return
			}
		}
	}

	err = DB.AutoMigrate(models...)
	if err != nil {
		return
	}

	seedDataUsers(DB)
}

// seedDataUsers : function to seed data user
func seedDataUsers(db *gorm.DB) {
	// Set new data to be seeded
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("12345"), bcrypt.DefaultCost)
	user := &m.User{
		Name:     "Reza",
		Email:    "reza@gmail.com",
		Password: string(hashedPassword),
		Phone:    "081234567890",
		Address:  "Tangerang",
	}

	db.Create(user)
}
