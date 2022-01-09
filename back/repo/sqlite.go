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
	sqlArgs := make([]interface{}, 0, len(images))

	for _, image := range images {
		preparedValues = append(preparedValues, "(?)")
		sqlArgs = append(sqlArgs, image.Data)
	}

	q := fmt.Sprintf(`INSERT INTO images (data) VALUES %s`, strings.Join(preparedValues, ","))

	stmt, err := r.db.Prepare(
		q,
	)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(sqlArgs...); err != nil {
		return err
	}

	return nil
}

func (r *SQLite) GetRandomImages(limit int) ([]*models.DBImage, error) {
	stmt, err := r.db.Prepare(
		`SELECT data FROM images ORDER BY random() LIMIT ?`,
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

func (r *SQLite) ImagesCount() (int, error) {
	var result int

	if err := r.db.QueryRow("SELECT COUNT(*) FROM images", nil).Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}

func (r *SQLite) RemoveFirstImages(count int) error {
	if _, err := r.db.Exec(`
	DELETE FROM images WHERE rowid IN (
		SELECT rowid FROM images
		ORDER BY rowid
		LIMIT ?
	)
	`, count); err != nil {
		return err
	}

	return nil
}
