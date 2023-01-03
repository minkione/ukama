// This is an example of a repository
package db

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ukama/ukama/systems/common/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const TaiNotUpdatedErr = "more recent tai for imsi exist"

// declare interface so that we can mock it
type HlrRecordRepo interface {
	Add(network string, record *Hlr) error
	Get(id int) (*Hlr, error)
	GetByImsi(imsi string) (*Hlr, error)
	GetByIccid(iccid string) (*Hlr, error)
	Update(imsi string, record *Hlr) error
	UpdatePackage(imsi string, packageId string) error
	DeleteByIccid(iccid string, nestedFunc ...func(*gorm.DB) error) error
	Delete(imsi string, nestedFunc ...func(*gorm.DB) error) error
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

func (r *hlrRecordRepo) Add(network string, rec *Hlr) error {
	d := r.db.GetGormDb().Create(rec)
	return d.Error
}

func (r *hlrRecordRepo) Update(imsiToUpdate string, rec *Hlr) error {
	d := r.db.GetGormDb().Where("imsi=?", imsiToUpdate).Updates(rec)
	return d.Error
}

func (r *hlrRecordRepo) UpdatePackage(imsiToUpdate string, packageId string) error {
	rec := &Hlr{PackageId: packageId}
	d := r.db.GetGormDb().Where("imsi=?", imsiToUpdate).Updates(rec)
	return d.Error
}

func (r *hlrRecordRepo) Get(id int) (*Hlr, error) {
	var hss Hlr
	result := r.db.GetGormDb().Preload(clause.Associations).First(&hss, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &hss, nil
}

func (r *hlrRecordRepo) GetByImsi(imsi string) (*Hlr, error) {
	var hlr Hlr
	result := r.db.GetGormDb().Preload(clause.Associations).Where("imsi=?", imsi).First(&hlr)
	if result.Error != nil {
		return nil, result.Error
	}

	return &hlr, nil
}

func (r *hlrRecordRepo) GetByIccid(iccid string) (*Hlr, error) {
	var hlr Hlr
	result := r.db.GetGormDb().Preload(clause.Associations).Where("iccid=?", iccid).First(&hlr)
	if result.Error != nil {
		return nil, result.Error
	}

	return &hlr, nil
}

func (r *hlrRecordRepo) Delete(imsi string, nestedFunc ...func(*gorm.DB) error) error {
	return r.db.ExecuteInTransaction2(func(tx *gorm.DB) *gorm.DB {
		return tx.Where(&Hlr{Imsi: imsi}).Delete(&Hlr{})
	}, nestedFunc...)
}

func (r *hlrRecordRepo) DeleteByIccid(iccid string, nestedFunc ...func(*gorm.DB) error) error {
	return r.db.ExecuteInTransaction2(func(tx *gorm.DB) *gorm.DB {
		return tx.Where(&Hlr{Iccid: iccid}).Delete(&Hlr{})
	}, nestedFunc...)

}

// ReplaceTai removes all TAI record for IMSI and adds new ones
func (r *hlrRecordRepo) UpdateTai(imsi string, tai Tai) error {
	var imsiM Hlr
	return r.db.GetGormDb().Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&Hlr{}).Where("imsi=?", imsi).First(&imsiM).Error
		if err != nil {
			return errors.Wrap(err, "error getting imsi")
		}

		var count int64
		err = tx.Model(&tai).Where("hlr_id = ? and device_updated_at >= ?", imsiM.ID, tai.DeviceUpdatedAt).Count(&count).Error
		if err != nil {
			return errors.Wrap(err, "error getting tai count")
		}
		if count > 0 {
			return fmt.Errorf(TaiNotUpdatedErr)
		}

		err = tx.Where("hlr_id=?", imsiM.ID).Delete(&Tai{}).Error
		if err != nil {
			return errors.Wrap(err, "error deleting tai")
		}

		tai.HlrID = imsiM.ID
		err = tx.Create(&tai).Error
		if err != nil {
			return errors.Wrap(err, "error adding tai")
		}
		return nil
	})
}
