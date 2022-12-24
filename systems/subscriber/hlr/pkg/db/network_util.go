package db

import (
	"github.com/pkg/errors"
	"github.com/ukama/ukama/systems/common/sql"
	"gorm.io/gorm"
)

func makeUserNetworkExist(db *gorm.DB, netName string) (*Network, error) {
	net := Network{
		Name: netName,
	}
	d := db.First(&network, "name = ?", netName)
	if d.Error != nil {
		if sql.IsNotFoundError(d.Error) {
			d2 := db.Create(&net)
			if d2.Error != nil {
				return nil, errors.Wrap(d2.Error, "error adding the org")
			}

		} else {
			return nil, errors.Wrap(d.Error, "error finding the org")
		}
	}
	return &net, nil
}
