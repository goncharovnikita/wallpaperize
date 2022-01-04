package repo

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/goncharovnikita/wallpaperize/back/models"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	db *sql.DB
}

func NewSQLite(db *sql.DB) *SQLite {
	return &SQLite{
		db: db,
	}
}

func (r *SQLite) Prepare() error {
	m := `
	CREATE TABLE IF NOT EXISTS images (
		data TEXT
	);
	`

	_, err := r.db.Exec(m)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLite) SetImages(images []*models.DBImage) error {
	preparedValues := make([]string, 0, len(images))
	imageDatas := make([][]byte, 0, len(images))

	for _, image := range images {
		preparedValues = append(preparedValues, "(?)")
		imageDatas = append(imageDatas, image.Data)
	}

	stmt, err := r.db.Prepare(
		fmt.Sprintf(`INSERT INTO images (data) VALUES (%s)`, strings.Join(preparedValues, ",")),
	)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(imageDatas); err != nil {
		return err
	}

	return nil
}

func (r *SQLite) GetImages(limit int) ([]*models.DBImage, error) {
	stmt, err := r.db.Prepare(
		`SELECT data FROM images limit ?`,
	)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*models.DBImage, 0)

	for rows.Next() {
		img := models.DBImage{}

		if err := rows.Scan(&img.Data); err != nil {
			return nil, err
		}

		result = append(result, &img)
	}

	return result, nil
}
