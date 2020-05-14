package databases

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Db *sqlx.DB
	Host string
	Port string
	User string
	Password string
	DBName string
}

func (sql *PostgresDB) Connect() {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		sql.Host, sql.Port, sql.User, sql.Password, sql.DBName)

	sql.Db = sqlx.MustConnect("postgres", conn)

	if err := sql.Db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connect Postgres OK")
}

func (sql *PostgresDB) Close() {
	sql.Db.Close()
}