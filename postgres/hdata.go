package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/kens1n/vitechmssample/graph/model"
)

type HdataRepo struct {
	DB *pg.DB
}

func (u *HdataRepo) GetHdataDataByGuidAndHash(guid string, hash string) (string, error) {
	var hdata model.Hdata
	err := u.DB.Model(&hdata).Where("guid = ? AND hash = ?", guid, hash).First()
	if err != nil {
		return "", err
	}

	return hdata.Data, nil
}
