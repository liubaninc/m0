package model

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"
	"gorm.io/gorm"
)

type Model struct {
	logger log.Logger
	DB     *gorm.DB
}

func New(conn gorm.Dialector, logger log.Logger) *Model {
	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database %v, error %v", conn.Name(), err))
	}
	autoMigrate(db)

	return &Model{
		DB:     db,
		logger: logger.With("module", "model"),
	}
}

type Tabler interface {
	TableName() string
}

type TimeNormal struct { // 内嵌方式（推荐）
	time.Time
}

func (t TimeNormal) MarshalJSON() ([]byte, error) {
	// tune := fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

// Value insert timestamp into mysql need this function.
func (t TimeNormal) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *TimeNormal) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeNormal{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func autoMigrate(db *gorm.DB) {

	db.AutoMigrate(&BlockChain{})
	db.AutoMigrate(&BlockChainChart{})
	db.AutoMigrate(&BlockChainTPSChart{})
	db.AutoMigrate(&InvokeRequest{})
	db.AutoMigrate(&MsgOutput{})
	db.AutoMigrate(&MsgInput{})
	db.AutoMigrate(&MsgUTXO{})
	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&Block{})
	db.AutoMigrate(&Address{})
	db.AutoMigrate(&Asset{})
	db.AutoMigrate(&Contract{})
	db.AutoMigrate(&Peer{})
	db.AutoMigrate(&Validator{})
	db.AutoMigrate(&ContractUpgrade{})
	db.AutoMigrate(&ContractInvoke{})
	db.AutoMigrate(&Events{})
	// wallet
	db.AutoMigrate(&MTransaction{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Claim{})
	db.AutoMigrate(&MContract{})
	db.AutoMigrate(&MContractTemplate{})
	db.AutoMigrate(&MContractTemplateFunction{})
}
