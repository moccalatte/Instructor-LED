package repository

import (
	"database/sql"
	"final-project-kelompok-1/utils/common"
	"fmt"
)

type CsvRepository interface {
	CsvStart() ([]string, error)
}

type csvRepository struct {
	db *sql.DB
}

func (c *csvRepository) CsvStart() ([]string, error) {
	var allsessionId []string
	rows, err := c.db.Query(common.CsvGetIdSession)
	if err != nil {
		fmt.Println("Error in repo for query csvgetidsession ", err.Error())
		return nil, fmt.Errorf("failed to get data from db : %s", err.Error())
	}

	for rows.Next() {
		var sessionId string
		err = rows.Scan(&sessionId)
		if err != nil {
			return nil, fmt.Errorf("failed to scan data : %s", err.Error())
		}
		allsessionId = append(allsessionId, sessionId)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during query : %s", err.Error())
	}
	return allsessionId, nil
}

func NewCsv(db *sql.DB) CsvRepository {
	return &csvRepository{db: db}
}
