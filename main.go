package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/b1994mi/test-esri/internal/pkg/domain/sqlmodel"
	"github.com/b1994mi/test-esri/internal/pkg/route"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
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
		panic(fmt.Errorf("failed to connect db: %v", err))
	}

	err = migrateDB(db)
	if err != nil {
		panic(fmt.Errorf("failed to migrate db: %v", err))
	}

	r := route.NewRouter(
		route.SetupHandler(
			route.SetupUsecase(
				route.SetupSQLService(db),
			),
		),
	)

	port := ":5000"
	httpLn, err := net.Listen("tcp", port)
	if err != nil {
		panic(fmt.Errorf("failed to listen to port %v: %v", port, err))
	}

	httpServer := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      r,
	}

	go func() {
		log.Printf("running on port %v", port)
		err := httpServer.Serve(httpLn)
		if err != nil {
			log.Println(fmt.Errorf("failed to serve http: %v", err))
		}
	}()

	log.Println(waitExitSignal().String())

	ctx := context.Background()
	// Graceful shutdown.
	err = httpServer.Shutdown(ctx)
	if err != nil {
		log.Println(fmt.Errorf("trying to shutdown with error: %v", err))
	}
}

func waitExitSignal() os.Signal {
	ch := make(chan os.Signal, 3)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	return <-ch
}

func migrateDB(db *gorm.DB) error {
	// db.Migrator().DropTable(
	// 	&sqlmodel.User{},
	// 	&sqlmodel.Meteran{},
	// 	&sqlmodel.IssueCategory{},
	// 	&sqlmodel.IssueImage{},
	// 	&sqlmodel.Issue{},
	// )

	err := db.Migrator().AutoMigrate(
		&sqlmodel.User{},
		&sqlmodel.Meteran{},
		&sqlmodel.IssueCategory{},
		&sqlmodel.IssueImage{},
		&sqlmodel.Issue{},
	)
	if err != nil {
		return err
	}

	var m sqlmodel.User
	err = db.Where("username", "admin").Find(&m).Error
	if err != nil {
		return err
	}

	if m.ID != 0 {
		return nil
	}

	b := []byte("admin")
	password, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	tx := db.Begin()
	defer tx.Rollback()

	tx.Create(&sqlmodel.User{
		FullName:       "admin",
		Username:       "admin",
		Email:          "aaa@aaa.aaa",
		PasswordHash:   string(password),
		AvatarFileName: "aaa.jpg",
		IsDeleted:      false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})

	tx.Create(&[]sqlmodel.Meteran{{
		MeteranCode: "777",
		Address:     "Jl. Suatu Jalan 1, Kecamatan, Jakarta",
		Lat:         -6.244561398703197,
		Lon:         106.7903002090279,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, {
		MeteranCode: "888",
		Address:     "Jl. Suatu Jalan 2, Kecamatan, Jakarta",
		Lat:         -6.243025300450427,
		Lon:         106.82772304943019,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, {
		MeteranCode: "999",
		Address:     "Jl. Suatu Jalan 3, Kecamatan, Jakarta",
		Lat:         -6.199943785957455,
		Lon:         106.86124328503183,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}})

	tx.Create(&sqlmodel.IssueCategory{
		Name:      "Category 1",
		Class:     "IMPORTANT",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	tx.Commit()

	return nil
}
