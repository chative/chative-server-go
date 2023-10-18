package dbengine

import (
	"sync"
	"time"

	"chative-server-go/internal/config"
	"chative-server-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once     sync.Once
	err      error
	dbEngine *gorm.DB
)

func GetDbEngine(c config.Config) (*gorm.DB, error) {
	once.Do(func() {
		dbEngine, err = gorm.Open(postgres.Open(c.Dsn), &gorm.Config{})
		if err != nil {
			return
		}
		sqlDB, err := dbEngine.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.SetMaxIdleConns(8)
		sqlDB.SetMaxOpenConns(8)
		sqlDB.SetConnMaxLifetime(time.Minute)
		sqlDB.SetConnMaxIdleTime(time.Second * 30)
		go func() {
			for {
				time.Sleep(time.Second * 25)
				sqlDB.Ping()
			}
		}()
		dbEngine.AutoMigrate(&models.Invitation{}, &models.Team{}, &models.FriendRelation{},
			&models.InternalAccount{}, &models.ShareConversationCnf{}, &models.AskNewFriend{},
			&models.Conversation{}, &models.UserProfile{},
			&models.Group{}, &models.GroupMember{},
			&models.WebauthnUser{},
			&models.ReportLog{}, &models.UserFindPath{},
			&models.Reminder{},
		)
	})
	return dbEngine, err
}

func GetDB() (*gorm.DB, error) {
	return dbEngine, err
}
