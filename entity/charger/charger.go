package charger

import (
	"time"
	"github.com/avegao/gocondi"
	"fmt"
	"database/sql"
	pb "github.com/avegao/iot-openevse-service/resource/grpc"
)

const tableName = "openevse.chargers"

type chargerInterface interface {
	getTableName() string
	ToGrpcResponse() *pb.Charger
}

type Charger struct {
	chargerInterface
	ID        uint64
	Name      string
	Host      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (c Charger) ToGrpcResponse() *pb.Charger {
	var deletedAt string

	if c.DeletedAt != nil {
		deletedAt = c.DeletedAt.Format(time.RFC3339)
	}

	return &pb.Charger{
		Id: c.ID,
		Name: c.Name,
		Host: c.Host,
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
		UpdatedAt: c.UpdatedAt.Format(time.RFC3339),
		DeletedAt: deletedAt,
	}
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
		logger.WithError(err).Debugf("%s - STOP", logTag)
		return
	}

	c = new(Charger)

	query := fmt.Sprintf("SELECT id, name, host, created_at, updated_at, deleted_at FROM %s WHERE id = $1", tableName)

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

func FindAll(getDeleted ...bool) (chargers []Charger, err error) {
	const logTag = "Charger.FindOneById"

	container := gocondi.GetContainer()
	logger := container.GetLogger()
	logger.Debugf("%s - START", logTag)

	db, err := container.GetDefaultDatabase()
	if err != nil {
		logger.WithError(err).Debugf("%s - STOP", logTag)
		return chargers, err
	}

	query := fmt.Sprintf("SELECT id, name, host, created_at, updated_at, deleted_at FROM %s", tableName)

	if len(getDeleted) == 0 || !getDeleted[0] {
		query += " WHERE deleted_at IS NULL"
	}

	logger.
		WithField("query", query).
		Debugf("%s", logTag)

	rows, err := db.Query(query)
	if err != nil {
		logger.WithError(err).Debugf("%s - STOP", logTag)
		return chargers, err
	}

	for rows.Next() {
		c := Charger{}

		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.Host,
			&c.CreatedAt,
			&c.UpdatedAt,
			&c.DeletedAt,
		)

		if err != nil {
			logger.WithError(err).Debugf("%s - STOP", logTag)
			return
		}

		chargers = append(chargers, c)
	}

	logger.WithField("chargers", chargers).Debugf("%s - END", logTag)
	return
}