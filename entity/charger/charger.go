package charger

import (
	"time"
	"github.com/avegao/gocondi"
	"fmt"
	"database/sql"
)

type Charger struct {
	ID        uint64
	Name      string
	Host      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (c Charger) getTableName() string {
	return "openevse.chargers"
}

func New() Charger {
	return Charger{}
}

func FindOneById(id uint64, getDeleted ...bool) (c *Charger, err error) {
	const logTag = "Charger.FindOneById"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.WithField("id", id).Debugf("%s - START", logTag)

	db, err := container.GetDefaultDatabase()
	if err != nil {
		return
	}

	c = new(Charger)

	query := fmt.Sprintf("SELECT id, name, host, created_at, updated_at, deleted_at FROM %s WHERE id = $1", c.getTableName())

	if len(getDeleted) == 0 || !getDeleted[0] {
		query += " AND deleted_at IS NULL"
	}

	logger.
		WithField("query", query).
		WithField("id", id).
		Debugf("%s", logTag)

	if err = db.QueryRow(query, id).Scan(
		&c.ID,
		&c.Name,
		&c.Host,
		&c.CreatedAt,
		&c.UpdatedAt,
		&c.DeletedAt,
	); err != nil {
		c = nil

		if err == sql.ErrNoRows {
			err = nil
		} else {
			logger.WithError(err).Debugf("%s - STOP", logTag)
		}

		return
	}

	logger.WithField("charger", c).Debugf("%s - END", logTag)
	return
}
