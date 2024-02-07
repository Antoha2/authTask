package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/antoha2/auth/config"
	authService "github.com/antoha2/auth/service"
	authRepository "github.com/antoha2/auth/service/authRepository"
	web "github.com/antoha2/auth/transport"
)

func main() {

	Run()

}

func initDb(cfg *config.Config) (*gorm.DB, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Dbname,
		cfg.DB.Sslmode,
	)

	// Prep config
	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf(" failed to parse config: %v", err)
	}

	// Make connections
	dbx, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return nil, fmt.Errorf(" failed to create connection db: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbx,
	}), &gorm.Config{})

	err = dbx.Ping()
	if err != nil {
		return nil, fmt.Errorf(" error to ping connection pool: %v", err)
	}
	log.Printf("(auth) Подключение к базе данных на http://127.0.0.1:%d\n", cfg.DB.Port)
	return gormDB, nil
}

func Run() {

	cfg := config.GetConfig()
	gormDB, err := initDb(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	authRep := authRepository.NewAuthPostgres(gormDB)

	authSer := authService.NewAuthService(authRep)

	Tran := web.NewWeb(*authSer)

	go Tran.StartHTTP()
	go Tran.StartGRPC(authSer)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	Tran.Stop()

}
