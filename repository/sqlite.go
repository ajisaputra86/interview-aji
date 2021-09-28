package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/fajrirahmat/interview-aji/model"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose"
)

type sqlite struct {
	db *sql.DB
}

func NewSQLLiteConnection(dbpath string) (DB, error) {
	connURL := fmt.Sprintf("file:%s?cache=shared", dbpath)
	db, err := sql.Open("sqlite3", connURL)
	if err != nil {
		return nil, err
	}
	log.Println("PING: ", db.Ping())
	db.SetMaxOpenConns(5)

	log.Printf("Migrate DB: %v", goose.Up(db, "repository"))

	return &sqlite{
		db: db,
	}, nil
}

func (s *sqlite) Close() error {
	return s.db.Close()
}

func (s *sqlite) ListLocation(ctx context.Context) ([]*model.Location, error) {
	rows, _ := s.db.Query("select id, name from locations")
	defer rows.Close()
	var list []*model.Location
	for rows.Next() {
		l := &model.Location{}
		if err := rows.Scan(&l.Id, &l.Name); err != nil {
			log.Println(err.Error())
			continue
		}
		list = append(list, l)
	}
	return list, nil
}

func (s *sqlite) AddCheckin(ctx context.Context, request *model.CheckInOutRequest) ([]*model.CheckInOutRequest, error) {
	var data []*model.CheckInOutRequest
	_, err := s.db.Exec("INSERT INTO checkin (identifier, type, location_id) VALUES ($1, $2, $3)",
		request.Identifier, request.Type, request.LocationId)

	if err != nil {
		return data, err
	}

	return data, nil
}
