package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Represents record in HSS db
type Hlr struct {
	gorm.Model

	Iccid string `gorm:"index:hlr_unique_idx,unique,where:deleted_at is null;not null;size:22;check:iccid_checker,iccid ~ $$^\\d+$$"`
	//IMSI might not be unique as same IMSI might be authorized to use multiple network of Org which means multiple enetry for the IMSI in HLR or may be use many to many relattion here.
	// IMSI Sim ID  (International mobile subscriber identity) https://www.netmanias.com/en/post/blog/5929/lte/lte-user-identifiers-imsi-and-guti
	Imsi string `gorm:"index:hlr_unique_idx,unique,where:deleted_at is null;not null;size:15;check:hlr_checker,imsi ~ $$^\\d+$$"`
	// Pre Shared Key. This is optional and configured in operator’s DB in Authentication center and USIM. https://www.3glteinfo.com/lte-security-architecture/
	Op []byte `gorm:"size:16;"`
	// Pre Shared Key. Configured in operator’s DB in Authentication center and USIM
	Amf []byte `gorm:"size:2;"`
	// Key from the SIM
	Key            []byte `gorm:"size:16;"`
	AlgoType       uint32
	UeDlAmbrBps    uint32
	UeUlAmbrBps    uint32
	Sqn            uint64
	CsgIdPrsent    bool
	CsgId          uint32
	DefaultApnName string
	NetworkID      uuid.UUID `gorm:"not null;type:uuid"`
	Tai            Tai
	PackageId      string `gorm:"not null;type uuid"`
}

// Tracking Area Identity (TAI)
// Assumption: one IMIS can have only one tracking area
type Tai struct {
	gorm.Model
	HlrID           uint      `gorm:"uniqueIndex:tai_hlr_unique_idx;not null"`
	PlmnId          string    `gorm:"size:6;uniqueIndex:tai_hlr_unique_idx;not null"` // Public Land Mobile Network Identity (MCC+MNC)
	Tac             uint32    `gorm:"uniqueIndex:tai_hlr_unique_idx,where:deleted_at is null;not null"`
	DeviceUpdatedAt time.Time // time when it was updated on the device
}

type Guti struct {
	CreatedAt       time.Time // do not set it directly, it will be overridden
	DeviceUpdatedAt time.Time // time when it was updated on the device
	Imsi            string    `gorm:"uniqueIndex;not null;size:15;check:hlr_checker,imsi ~ $$^\\d+$$"`
	PlmnId          string    `gorm:"uniqueIndex:idx_guti;not null;size:6"`
	Mmegi           uint32    `gorm:"uniqueIndex:idx_guti;not null"`
	Mmec            uint32    `gorm:"uniqueIndex:idx_guti;not null"`
	MTmsi           uint32    `gorm:"uniqueIndex:idx_guti;not null"`
}
