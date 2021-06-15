package model

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	logger log.Logger
	DB     *gorm.DB
}

func New(dbHost string, dbPort int, dbUser string, dbPassword string, dbName string, logger log.Logger) *Model {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database %v, error %v", dsn, err))
	}
	autoMigrate(db)

	return &Model{
		DB:     db,
		logger: logger.With("module", "model"),
	}
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
	// wallet
	db.AutoMigrate(&MTransaction{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Claim{})
}
