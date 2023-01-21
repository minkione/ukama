package db_test

import (
	extsql "database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/ukama/ukama/systems/subscriber/hlr/pkg/client"
	int_db "github.com/ukama/ukama/systems/subscriber/hlr/pkg/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Iccid = "0123456789012345678912"

var sub = int_db.Hlr{
	Iccid:          Iccid,
	Imsi:           Imsi,
	Op:             []byte("0123456789012345"),
	Key:            []byte("0123456789012345"),
	Amf:            []byte("800"),
	AlgoType:       1,
	UeDlAmbrBps:    2000000,
	UeUlAmbrBps:    2000000,
	Sqn:            1,
	CsgIdPrsent:    false,
	CsgId:          0,
	DefaultApnName: "ukama",
}

var sim = client.SimCardInfo{
	Iccid:          Iccid,
	Imsi:           Imsi,
	Op:             []byte("0123456789012345"),
	Key:            []byte("0123456789012345"),
	Amf:            []byte("800"),
	AlgoType:       1,
	UeDlAmbrBps:    2000000,
	UeUlAmbrBps:    2000000,
	Sqn:            1,
	CsgIdPrsent:    false,
	CsgId:          0,
	DefaultApnName: "ukama",
}

var tai = int_db.Tai{
	PlmnId:          "00101",
	Tac:             101,
	DeviceUpdatedAt: time.Unix(int64(1639144056), 0),
}

func TestHlrRecordRepo_Add(t *testing.T) {

	t.Run("Add", func(t *testing.T) {
		// Arrange
		var db *extsql.DB
		var err error

		db, mock, err := sqlmock.New() // mock sql.DB
		assert.NoError(t, err)

		//row := sqlmock.NewRows([]string{"iccid", "imsi", "op", "amf", "key", "algo_type", "ue_dl_ambr_bps", "ue_ul_ambr_bps", "sqn", "csg_id_prsent", "csg_id", "default_apn_name", "network_id", "package_id"}).

		mock.ExpectBegin()

		mock.ExpectQuery(regexp.QuoteMeta(`INSERT`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sub.Iccid, sub.Imsi, sub.Op, sub.Amf, sub.Key, sub.AlgoType, sub.UeDlAmbrBps, sub.UeUlAmbrBps, sub.Sqn, sub.CsgIdPrsent, sub.CsgId, sub.DefaultApnName, sub.NetworkID, sub.PackageId).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		mock.ExpectCommit()

		dialector := postgres.New(postgres.Config{
			DSN:                  "sqlmock_db_0",
			DriverName:           "postgres",
			Conn:                 db,
			PreferSimpleProtocol: true,
		})
		gdb, err := gorm.Open(dialector, &gorm.Config{})
		assert.NoError(t, err)

		r := int_db.NewHlrRecordRepo(&UkamaDbMock{
			GormDb: gdb,
		})

		assert.NoError(t, err)

		// Act
		err = r.Add(&sub)

		// Assert
		assert.NoError(t, err)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)

	})

}

func TestHlrRecordRepo_Update(t *testing.T) {

	t.Run("UpdatePackage", func(t *testing.T) {
		// Arrange
		var db *extsql.DB
		var err error
		PackageId := "073d8584-5884-4f05-9686-4a11e3a97ea3"
		db, mock, err := sqlmock.New() // mock sql.DB
		assert.NoError(t, err)

		mock.ExpectBegin()

		mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).
			WithArgs(sqlmock.AnyArg(), PackageId, Imsi).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		dialector := postgres.New(postgres.Config{
			DSN:                  "sqlmock_db_0",
			DriverName:           "postgres",
			Conn:                 db,
			PreferSimpleProtocol: true,
		})
		gdb, err := gorm.Open(dialector, &gorm.Config{})
		assert.NoError(t, err)

		r := int_db.NewHlrRecordRepo(&UkamaDbMock{
			GormDb: gdb,
		})

		assert.NoError(t, err)

		// Act
		err = r.UpdatePackage(Imsi, PackageId)

		// Assert
		assert.NoError(t, err)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)

	})

}

func TestHlrRecordRepo_Get(t *testing.T) {

	t.Run("ReadByICCID", func(t *testing.T) {
		// Arrange
		var db *extsql.DB
		var err error

		db, mock, err := sqlmock.New() // mock sql.DB
		assert.NoError(t, err)

		hrow := sqlmock.NewRows([]string{"iccid", "imsi", "op", "amf", "key", "algo_type", "ue_dl_ambr_bps", "ue_ul_ambr_bps", "sqn", "csg_id_prsent", "csg_id", "default_apn_name", "network_id", "package_id"}).
			AddRow(sub.Iccid, sub.Imsi, sub.Op, sub.Amf, sub.Key, sub.AlgoType, sub.UeDlAmbrBps, sub.UeDlAmbrBps, sub.Sqn, sub.CsgIdPrsent, sub.CsgId, sub.DefaultApnName, sub.NetworkID, sub.PackageId)

		mock.ExpectQuery(`^SELECT.*hlrs.*`).
			WithArgs(Iccid).
			WillReturnRows(hrow)

		dialector := postgres.New(postgres.Config{
			DSN:                  "sqlmock_db_0",
			DriverName:           "postgres",
			Conn:                 db,
			PreferSimpleProtocol: true,
		})
		gdb, err := gorm.Open(dialector, &gorm.Config{})
		assert.NoError(t, err)

		r := int_db.NewHlrRecordRepo(&UkamaDbMock{
			GormDb: gdb,
		})

		assert.NoError(t, err)

		// Act
		hlr, err := r.GetByIccid(Iccid)

		// Assert
		assert.NoError(t, err)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)

		if assert.NotNil(t, hlr) {
			assert.EqualValues(t, hlr.Iccid, Iccid)
		}

	})

	t.Run("ReadByImsi", func(t *testing.T) {
		// Arrange
		var db *extsql.DB
		var err error

		db, mock, err := sqlmock.New() // mock sql.DB
		assert.NoError(t, err)

		hrow := sqlmock.NewRows([]string{"iccid", "imsi", "op", "amf", "key", "algo_type", "ue_dl_ambr_bps", "ue_ul_ambr_bps", "sqn", "csg_id_prsent", "csg_id", "default_apn_name", "network_id", "package_id"}).
			AddRow(sub.Iccid, sub.Imsi, sub.Op, sub.Amf, sub.Key, sub.AlgoType, sub.UeDlAmbrBps, sub.UeDlAmbrBps, sub.Sqn, sub.CsgIdPrsent, sub.CsgId, sub.DefaultApnName, sub.NetworkID, sub.PackageId)

		mock.ExpectQuery(`^SELECT.*hlrs.*`).
			WithArgs(Imsi).
			WillReturnRows(hrow)

		dialector := postgres.New(postgres.Config{
			DSN:                  "sqlmock_db_0",
			DriverName:           "postgres",
			Conn:                 db,
			PreferSimpleProtocol: true,
		})
		gdb, err := gorm.Open(dialector, &gorm.Config{})
		assert.NoError(t, err)

		r := int_db.NewHlrRecordRepo(&UkamaDbMock{
			GormDb: gdb,
		})

		assert.NoError(t, err)

		// Act
		hlr, err := r.GetByImsi(Imsi)

		// Assert
		assert.NoError(t, err)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)

		if assert.NotNil(t, hlr) {
			assert.EqualValues(t, hlr.Imsi, Imsi)
		}

	})

}
