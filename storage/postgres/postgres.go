package postgres

import(
    "database/sql"
    _ "github.com/lib/pq"
)

type Postgres struct{
    user string
    password string
    dbname string
}

func NewPostgres(user, password, dbname string) *Postgres {
    return &Postgres{
        user:       user,
        password:   password,
        dbname:     dbname,
    }
}

func (pg *Postgres) Connect() (*sql.DB, error) {
    connStr := "user=" + pg.user + " password=" + pg.password + " dbname=" + pg.dbname + " sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}

