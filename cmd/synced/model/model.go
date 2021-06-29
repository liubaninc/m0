package model

import (
	"fmt"

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

func autoMigrate(db *gorm.DB) {

	//创建数据库
	db.Exec("create database m0 default charset utf8 collate utf8_general_ci")

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
