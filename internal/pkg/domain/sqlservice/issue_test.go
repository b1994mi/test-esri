package sqlservice

import (
	"log"
	"os"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Test_issueService_FindWithImageMeteranCategoryBy(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             1000 * time.Millisecond, // default
			LogLevel:                  logger.Error,            // default
			Colorful:                  true,                    // default
			IgnoreRecordNotFoundError: true,
		},
	)})
	if err != nil {
		panic("failed to connect db")
	}
	svc := &issueService{
		db: db,
	}
	got, err := svc.FindWithImageMeteranCategoryBy(map[string]interface{}{}, 1, 10)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(got)
}
