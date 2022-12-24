// This is an example of a repository
package db

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/ukama/ukama/systems/common/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const TaiNotUpdatedErr = "more recent tai for imsi exist"

// declare interface so that we can mock it
type HlrRecordRepo interface {
	Add(netName string, record *HlrRecord) error
	Get(id int) (*HlrRecord, error)
	GetByImsi(imsi string) (*HlrRecord, error)
	GetHlrRecordByUserUuid(userUuid uuid.UUID) ([]*HlrRecord, error)
	Update(imsi string, record *HlrRecord) error
	Delete(imsi string, nestedFunc ...func(*gorm.DB) error) error
	DeleteByUserId(user uuid.UUID, nestedFunc ...func(*gorm.DB) error) error
	UpdateTai(imis string, tai Tai) error
}

type hlrRecordRepo struct {
	db sql.Db
}

func NewHlrRecordRepo(db sql.Db) *hlrRecordRepo {
	return &hlrRecordRepo{
		db: db,
	}
}

func (r *hlrRecordRepo) Add(netName string, rec *HlrRecord) error {
	net, err := makeUserNetworkExist(r.db.GetGormDb(), netName)
	if err != nil {
		return err
	}
	rec.Network = net
	d := r.db.GetGormDb().Create(rec)
	return d.Error
}

func (r *hlrRecordRepo) Update(imsiToUpdate string, rec *HlrRecord) error {
	d := r.db.GetGormDb().Where("imsi=?", imsiToUpdate).Updates(rec)
	return d.Error
}

func (r *hlrRecordRepo) Get(id int) (*HlrRecord, error) {
	var hss HlrRecord
	result := r.db.GetGormDb().Preload(clause.Associations).First(&hss, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &hss, nil
}

func (r *hlrRecordRepo) GetByImsi(imsi string) (*HlrRecord, error) {
	var hlr HlrRecord
	result := r.db.GetGormDb().Preload(clause.Associations).Where("imsi=?", imsi).First(&hlr)
	if result.Error != nil {
		return nil, result.Error
	}

	return &hlr, nil
}

func (r *hlrRecordRepo) GetHlrRecordByUserUuid(userUuid uuid.UUID) ([]*HlrRecord, error) {
	var records []*HlrRecord
	result := r.db.GetGormDb().Preload(clause.Associations).Where("user_uuid=?", userUuid).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	return records, nil
}

func (r *hlrRecordRepo) Delete(imsi string, nestedFunc ...func(*gorm.DB) error) error {
	return r.db.ExecuteInTransaction2(func(tx *gorm.DB) *gorm.DB {
		return tx.Where(&HlrRecord{Imsi: imsi}).Delete(&HlrRecord{})
	}, nestedFunc...)
}

func (r *hlrRecordRepo) DeleteByUserId(user uuid.UUID, nestedFunc ...func(*gorm.DB) error) error {
	return r.db.ExecuteInTransaction2(func(tx *gorm.DB) *gorm.DB {
		return tx.Where(&HlrRecord{UserUuid: user}).Delete(&HlrRecord{})
	}, nestedFunc...)

}

// ReplaceTai removes all TAI record for IMSI and adds new ones
func (r *hlrRecordRepo) UpdateTai(imsi string, tai Tai) error {
	var imsiM HlrRecord
	return r.db.GetGormDb().Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&HlrRecord{}).Where("imsi=?", imsi).First(&imsiM).Error
		if err != nil {
			return errors.Wrap(err, "error getting imsi")
		}

		var count int64
		err = tx.Model(&tai).Where("imsi_id = ? and device_updated_at >= ?", imsiM.ID, tai.DeviceUpdatedAt).Count(&count).Error
		if err != nil {
			return errors.Wrap(err, "error getting tai count")
		}
		if count > 0 {
			return fmt.Errorf(TaiNotUpdatedErr)
		}

		err = tx.Where("imsi_id=?", imsiM.ID).Delete(&Tai{}).Error
		if err != nil {
			return errors.Wrap(err, "error deleting tai")
		}

		tai.ImsiID = imsiM.ID
		err = tx.Create(&tai).Error
		if err != nil {
			return errors.Wrap(err, "error adding tai")
		}
		return nil
	})
}
