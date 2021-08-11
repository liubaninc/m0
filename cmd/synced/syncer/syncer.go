package syncer

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	peertypes "github.com/liubaninc/m0/x/peer/types"
	permissiontypes "github.com/liubaninc/m0/x/permission/types"
	validatortypes "github.com/liubaninc/m0/x/validator/types"
	pkitypes "github.com/liubaninc/m0/x/pki/types"
	"net"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	wasmtypes "github.com/liubaninc/m0/x/wasm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/tendermint/libs/bytes"

	"github.com/liubaninc/m0/cmd/synced/model"
	msdk "github.com/liubaninc/m0/sdk"
	"github.com/tendermint/tendermint/libs/log"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"gorm.io/gorm"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const TIME_FORMAT = "2006-01-02 15:04:05.999"

var (
	Local, _ = time.LoadLocation("Asia/Shanghai")

	actions = map[string]string {
		"CreatePeerID": "新增节点",
		"UpdatePeerID": "更新节点",
		"DeletePeerID": "删除节点",
		"SetPermission": "设置用户",
		"AddRootCert": "上传根证书",
		"AddCert": "上传证书",
		"RevokeRootCert": "注销根证书",
		"RevokeCert": "注销证书",
		"FreezeCert": "冻结证书",
		"UnfreezeCert": "解冻证书",
		"CreateValidator": "新增共识节点",
		"LeaveValidator": "删除共识节点",
		"Deploy": "部署合约",
		"Upgrade": "升级合约",
		"Invoke": "调用合约",
		"Freeze": "解冻合约",
		"Unfreeze": "冻结合约",
		"ProposeDeployContract":"申请合约",
		"ApproveDeployContract":"批准合约",
	}
)

type Syncer struct {
	logger     log.Logger
	db         *gorm.DB
	client     msdk.Client
	blockchain model.BlockChain
}

func New(db *gorm.DB, client msdk.Client, logger log.Logger) *Syncer {
	return &Syncer{
		db:     db,
		logger: logger.With("module", "sync"),
		client: client,
	}
}

func (synced *Syncer) Run() {
	curHeight := int64(1)
	result := synced.db.FirstOrInit(&synced.blockchain)
	if result.Error != nil {
		panic(result.Error)
	} else if synced.blockchain.BlockNum == 0 {
		// none
		// TODO Genesis
	} else {
		// 一致性检查
		var block model.Block
		result := synced.db.Last(&block)
		if result.Error != nil {
			panic(result.Error)
		}
		if synced.blockchain.BlockNum != block.Height {
			panic(fmt.Errorf("database mismatch"))
		}

		// 一致性检查
		resultBlock, err := synced.client.GetBlock(block.Height)
		if err != nil {
			panic(fmt.Errorf("blockchain mismatch %v", err))
		}
		if block.PrevHash != resultBlock.Block.Header.LastBlockID.Hash.String() ||
			block.Hash != resultBlock.BlockID.Hash.String() {
			panic(fmt.Errorf("mismatch hash"))
		}
		curHeight = block.Height + 1
	}

	// 遇到错误,休眠延长
	var count int64
	sleepFun := func() {
		count++
		synced.logger.Debug("sleep", "duration", time.Duration(count)*time.Second)
		time.Sleep(time.Duration(count) * time.Second)
	}

	go func() {
		var latestPeersTime time.Time
		var latestResultBlock *coretypes.ResultBlock
		for {
			// 获取链上最新区块
			if latestResultBlock == nil || latestResultBlock.Block == nil || curHeight > latestResultBlock.Block.Height {
				if resultBlock, err := synced.client.GetBlockLatest(); err != nil {
					synced.logger.Error("GetBlockLatest", "error", err)
				} else {
					latestResultBlock = resultBlock
				}
			}

			// 无可更新区块
			if latestResultBlock == nil || latestResultBlock.Block == nil || curHeight > latestResultBlock.Block.Height {
				time.Sleep(time.Second)
				continue
			}

			// 获取区块
			resultBlock, err := synced.client.GetBlock(curHeight)
			if err != nil {
				synced.logger.Error("GetBlock", "height", curHeight, "error", err)
				sleepFun()
				continue
			}

			tx := synced.db.Begin()
			if tx.Error != nil {
				synced.logger.Error("database tx begin error", "height", curHeight, "error", tx.Error)
				sleepFun()
				continue
			}

			n, err := synced.processValidators(curHeight, tx)
			if err != nil {
				synced.logger.Error("processValidators error", "height", curHeight, "error", tx.Error)
				sleepFun()
				continue
			}
			synced.blockchain.ValidatorNum = int64(n)

			if latestPeersTime.IsZero() || time.Now().Sub(latestPeersTime) > time.Minute*1 {
				n, err := synced.processPeers(tx)
				if err != nil {
					synced.logger.Error("write peer", "height", curHeight, "error", err)
					tx.Rollback()
					sleepFun()
					continue
				}
				synced.blockchain.PeerNum = int64(n)
				latestPeersTime = time.Now()
			}
			// 写入区块
			if err := synced.processBlock(resultBlock, tx); err != nil {
				synced.logger.Error("write block", "height", curHeight, "error", err)
				tx.Rollback()
				sleepFun()
				continue
			}
			if result := tx.Save(&synced.blockchain); result.Error != nil {
				synced.logger.Error("write blockchain", "height", curHeight, "error", result.Error)
				tx.Rollback()
				sleepFun()
				continue
			}
			if result := tx.Commit(); result.Error != nil {
				synced.logger.Error("database tx commit error", "height", curHeight, "error", tx.Error)
				tx.Rollback()
				sleepFun()
				continue
			}
			curHeight++
			count = 0
		}
	}()
}

func (synced *Syncer) processPeers(db *gorm.DB) (int, error) {
	db.Where("1 = 1").Delete(&model.Peer{})
	status, err := synced.client.GetStatus()
	if err != nil {
		return 0, nil
	}
	nodeInfo := status.NodeInfo
	ip := nodeInfo.ListenAddr[len("tcp://"):]
	ip = ip[:strings.Index(ip, ":")]
	if addr := net.ParseIP(ip); addr != nil {
		ip = addr.String()
	} else if addr, err := net.ResolveIPAddr("ip", ip); err == nil {
		ip = addr.IP.String()
	}
	peer := &model.Peer{
		Name:    nodeInfo.Moniker,
		IP:      ip,
		Version: fmt.Sprintf("V%d", nodeInfo.ProtocolVersion.App),
		NodeID:  string(nodeInfo.DefaultNodeID),
		Status:  0,
		Type:    1,
	}
	synced.updatePeerStatus(peer, db)
	if result := db.Save(peer); result.Error != nil {
		return 0, result.Error
	}

	netInfo, err := synced.client.GetNetInfo()
	if err != nil {
		return 0, nil
	}
	for _, p := range netInfo.Peers {
		nodeInfo = p.NodeInfo
		peer := &model.Peer{
			Name:    nodeInfo.Moniker,
			IP:      p.RemoteIP,
			Version: fmt.Sprintf("V%d", nodeInfo.ProtocolVersion.App),
			NodeID:  string(nodeInfo.DefaultNodeID),
			Time:    p.ConnectionStatus.Duration.String(),
			Status:  0,
			Type:    1,
		}
		synced.updatePeerStatus(peer, db)
		if result := db.Save(peer); result.Error != nil {
			return 0, result.Error
		}
	}
	return netInfo.NPeers + 1, nil
}

func (synced *Syncer) processValidators(height int64, db *gorm.DB) (int, error) {
	db.Where("1 = 1").Delete(&model.Validator{})
	result, err := synced.client.GetValidators(height, 1, 1000)
	if err != nil {
		return 0, err
	}
	for _, v := range result.Validators {
		validator := &model.Validator{
			Name:   v.Address.String(),
			PubKey: hex.EncodeToString(v.PubKey.Bytes()),
			Pow:    v.VotingPower,
			NodeID: hex.EncodeToString(v.PubKey.Address()),
		}
		if result := db.Save(validator); result.Error != nil {
			return 0, result.Error
		}
	}
	return result.Total, nil
}

func (synced *Syncer) processBlock(resultBlock *coretypes.ResultBlock, db *gorm.DB) error {
	synced.logger.Debug("processBlock ...", "height", resultBlock.Block.Height)
	// collecting
	txNum := 0
	msgNum := 0
	assetNum := 0
	contractNum := 0

	// 区块
	block := model.Block{
		Height:   resultBlock.Block.Height,
		Size:     resultBlock.Block.Size(),
		Hash:     resultBlock.BlockID.Hash.String(),
		Time:     resultBlock.Block.Time.In(Local).Format(TIME_FORMAT),
		PrevHash: resultBlock.Block.LastBlockID.Hash.String(),
		Proposer: resultBlock.Block.ProposerAddress.String(),
		TxNum:    len(resultBlock.Block.Txs),
	}
	txNum += block.TxNum

	c := synced.client.WithHeight(block.Height)
	// 交易
	assetsMap := map[string]bool{}
	addressesMap := map[string]bool{}
	contractsMap := map[string]int{}
	contractHashs := map[string][]string{}
	hashMap := map[string]*model.Transaction{}
	for _, tx := range resultBlock.Block.Txs {
		if err := synced.processTxEvents(tx.Hash(), block.Time, db); err != nil {
			return err
		}

		mtx, err := synced.processTx(tx.Hash(), block.Time)
		if err != nil {
			return err
		}
		hashMap[mtx.Hash] = mtx
		block.TxList = append(block.TxList, mtx)
		msgNum += mtx.MsgNum

		// 地址金额可能发生变动，需重新获取
		addresses := strings.Split(mtx.Addresses, ",")
		for _, addr := range addresses {
			if len(addr) == 0 {
				continue
			}
			if _, ok := addressesMap[addr]; !ok {
				addressesMap[addr] = true
			}
		}

		// 资产详情可能发生变动，需重新获取
		if strings.Contains(mtx.Type, "资产发行") || strings.Contains(mtx.Type, "资产销毁") {
			assets := strings.Split(mtx.Assets, ",")
			for _, asset := range assets {
				if len(asset) == 0 {
					continue
				}
				if _, ok := assetsMap[asset]; !ok {
					assetsMap[asset] = true
				}
			}
		}

		// 合约详情可能发生变动，需重新获取
		contracts := strings.Split(mtx.Contracts, ",")
		exists := map[string]bool{}
		for _, contract := range contracts {
			if len(contract) == 0 {
				continue
			}
			if _, ok := exists[contract]; !ok {
				contractsMap[contract] += 1
				exists[contract] = true
				if _, ok := contractHashs[contract]; !ok {
					contractHashs[contract] = make([]string, 0)
				}
				contractHashs[contract] = append(contractHashs[contract], mtx.Hash)
			}
		}

		if h := mtx.Height - 10; h > 0 {
			db.Unscoped().Delete(&model.MTransaction{
				Height: h,
			})
		}
	}

	// 更新关联地址
	for address := range addressesMap {
		var addr model.Address
		if result := db.FirstOrInit(&addr, map[string]interface{}{"address": address}); result.Error != nil {
			return result.Error
		}
		if addr.ID == 0 {
			acct, err := c.GetAccount(address)
			if err != nil {
				return err
			}
			var acc authtypes.BaseAccount
			if err := proto.Unmarshal(acct.Account.Value, &acc); err != nil {
				return err
			}
			addr.Address = acc.GetAddress().String()
			addr.AccountNumber = acc.GetAccountNumber()
			addr.Sequence = acc.GetSequence()
		}
		res, err := c.GetAccountBalances(addr.Address, nil, 0, 1000, false)
		if err != nil {
			return err
		}
		addr.Balance = res.GetBalances().String() + ","
		if result := db.Save(&addr); result.Error != nil {
			return result.Error
		}
	}

	assetsMap["m0token"] = true
	// 更新关联资产
	for asset := range assetsMap {
		ast, err := c.GetToken(asset)
		if err != nil {
			return err
		}
		if ast.Token == nil {
			continue
		}
		var item model.Asset
		if result := db.FirstOrInit(&item, map[string]interface{}{"name": asset}); result.Error != nil {
			return result.Error
		}
		if item.ID == 0 {
			item.Name = ast.Token.Name
			item.Initiator = ast.Token.Issuer
			item.RefTxID = ast.Token.IssueTx
			item.Time = block.Time
			assetNum += 1
		}
		item.MintAmount = ast.Token.Supply
		t, _ := sdk.ParseCoinsNormalized(ast.Token.Supply)
		c, _ := sdk.ParseCoinsNormalized(ast.Token.Circulating)
		item.BurnAmount = t.Sub(c).String()
		if result := db.Save(&item); result.Error != nil {
			return result.Error
		}
	}

	// 更新关联合约
	for contract, n := range contractsMap {
		var item model.Contract
		if strings.HasPrefix(contract, ":") {
			res, err := c.GetContract(strings.TrimLeft(contract, ":"))
			if err != nil {
				return err
			}

			if result := db.FirstOrInit(&item, map[string]interface{}{"name": res.Contract.Name}); result.Error != nil {
				return result.Error
			}
			for index, hash := range contractHashs[contract] {
				upgrade := model.ContractUpgrade{
					Name:    item.Name,
					Version: fmt.Sprintf("v%v", item.Deploy+int64(index)+1),
					Time:    block.Time,
					Hash:    hash,
				}
				if result := db.Save(&upgrade); result.Error != nil {
					return result.Error
				}
			}

			if item.ID == 0 {
				item.Name = res.Contract.Name
				item.Initiator = res.Contract.Initiator
				// item.Number = res.Contract.Number
				item.Runtime = res.Contract.Desc.Runtime
				item.VmCompiler = res.Contract.Desc.VmCompiler
				item.ContractType = res.Contract.Desc.ContractType
				item.Time = block.Time
				contractNum += 1
			}
			item.Digest = hex.EncodeToString(res.Contract.Desc.Digest)
			item.Deploy += int64(n)
			item.Version = fmt.Sprintf("v%v", item.Deploy)
		} else {
			if result := db.Find(&item, map[string]interface{}{"name": contract}); result.Error != nil {
				return result.Error
			}
			item.Invoke += int64(n)
			for _, hash := range contractHashs[contract] {
				invoke := model.ContractInvoke{
					Name:    item.Name,
					Version: item.Version,
					Time:    block.Time,
					Hash:    hash,
					Size:    hashMap[hash].Size,
					Height:  block.Height,
				}
				if result := db.Save(&invoke); result.Error != nil {
					return result.Error
				}
			}
		}
		if result := db.Save(&item); result.Error != nil {
			return result.Error
		}
	}

	if result := db.Create(&block); result.Error != nil {
		return result.Error
	}
	synced.blockchain.BlockNum = block.Height
	synced.blockchain.TxNum += int64(txNum)
	synced.blockchain.MsgNum += int64(msgNum)
	synced.blockchain.ContractNum += int64(contractNum)
	synced.blockchain.AssetNum += int64(assetNum)
	day := resultBlock.Block.Time.Format("2006-01-02")
	minute := resultBlock.Block.Time.Format("2006-01-02 15:04:05")

	var blockchainChart model.BlockChainChart
	if result := db.FirstOrCreate(&blockchainChart, model.BlockChainChart{Time: day}); result.Error != nil {
		return result.Error
	}
	blockchainChart.BlockNum += 1
	blockchainChart.TxNum += txNum
	blockchainChart.MsgNum += msgNum
	blockchainChart.AssetNum += assetNum
	blockchainChart.ContractNum += contractNum
	if result := db.Save(&blockchainChart); result.Error != nil {
		return result.Error
	}

	if txNum > 0 {
		var blockchainTPSChart model.BlockChainTPSChart
		if result := db.FirstOrCreate(&blockchainTPSChart, model.BlockChainTPSChart{Time: minute}); result.Error != nil {
			return result.Error
		}
		blockchainTPSChart.TxNum += txNum
		blockchainTPSChart.MsgNum += msgNum
		if result := db.Save(&blockchainTPSChart); result.Error != nil {
			return result.Error
		}
		db.Where("tx_num < ? AND msg_num < ?", txNum, msgNum).Delete(&model.BlockChainTPSChart{})
	}

	synced.logger.Debug("processBlock", "height", resultBlock.Block.Height)
	return nil
}

func (synced *Syncer) processTxEvents(hash []byte, time string, db *gorm.DB) error {
	synced.logger.Debug("processTxEvents ...", "hash", hex.EncodeToString(hash))
	resultTx, err := synced.client.GetTx(hex.EncodeToString(hash))
	if err != nil {
		return err
	}

	var e model.Events
	e.Hash = resultTx.Hash.String()
	e.Height = resultTx.Height
	e.Time = time

	attrs := make([]sdk.Attribute, 0)
	for _, event := range resultTx.TxResult.Events {
		if event.Type == "message" {
			for _, attr := range event.Attributes {
				switch key := string(attr.Key); key {
				case "sender":
					e.Operator = string(attr.Value)
				case "module":
					e.Route = string(attr.Value)
				case "action":
					e.Type = actions[string(attr.Value)]
				default:
				}
			}
		} else {
			for _, attr := range event.Attributes {
				attrs = append(attrs, sdk.NewAttribute(string(attr.Key), string(attr.Value)))
			}
		}
	}
	bts, _ := json.Marshal(attrs)
	e.Detail = string(bts)
	db.Save(&e)
	return nil
}

func (synced *Syncer) processTx(hash []byte, time string) (*model.Transaction, error) {
	synced.logger.Debug("processTx ...", "hash", hex.EncodeToString(hash))
	resultTx, err := synced.client.GetTx(hex.EncodeToString(hash))
	if err != nil {
		return nil, err
	}

	stx, err := synced.client.TxConfig.TxDecoder()(resultTx.Tx)
	if err != nil {
		return nil, err
	}
	txBuilder, err := synced.client.TxConfig.WrapTxBuilder(stx)
	if err != nil {
		return nil, err
	}
	tx := txBuilder.GetTx()

	mtx := &model.Transaction{
		Hash:     resultTx.Hash.String(),
		Size:     resultTx.TxResult.Size(),
		Memo:     tx.GetMemo(),
		Fee:      tx.GetFee().String(),
		Status:   resultTx.Height > 0,
		Time:     time,
		Height:   resultTx.Height,
		MsgNum:   len(tx.GetMsgs()),
		UTXOMsgs: make([]*model.MsgUTXO, len(tx.GetMsgs())),
	}

	types := make([]string, mtx.MsgNum)
	assetsMap := map[string]bool{}
	var assets []string
	addressesMap := map[string]bool{}
	var addresses []string
	contractsMap := map[string]bool{}
	var contracts []string
	for index, msg := range tx.GetMsgs() {
		var umsg *model.MsgUTXO
		switch msg := msg.(type) {
		case *utxotypes.MsgIssue:
			umsg = ProcessMsgIssue(msg)
		case *utxotypes.MsgDestroy:
			umsg = ProcessMsgDestroy(msg)
		case *utxotypes.MsgSend:
			umsg = ProcessMsgSend(msg)
		case *wasmtypes.MsgDeploy:
			umsg = ProcessMsgDeploy(msg)
		case *wasmtypes.MsgUpgrade:
			umsg = ProcessMsgUpgrade(msg)
		case *wasmtypes.MsgInvoke:
			umsg = ProcessMsgInvoke(msg)
		case *wasmtypes.MsgFreeze:
			umsg = &model.MsgUTXO{
				Type: "冻结合约",
			}
		case *wasmtypes.MsgUnfreeze:
			umsg = &model.MsgUTXO{
				Type: "解冻合约",
			}
		case *wasmtypes.MsgUndeploy:
			umsg = &model.MsgUTXO{
				Type: "删除合约",
			}
		case *peertypes.MsgCreatePeerID:
			umsg = &model.MsgUTXO{
				Type: "新增节点",
			}
		case *peertypes.MsgDeletePeerID:
			umsg = &model.MsgUTXO{
				Type: "移除节点",
			}
		case *permissiontypes.MsgSetPermission:
			umsg = &model.MsgUTXO{
				Type: "新增用户",
			}
		case *validatortypes.MsgCreateValidator:
			umsg = &model.MsgUTXO{
				Type: "新增共识节点",
			}
		case *validatortypes.MsgLeaveValidator:
			umsg = &model.MsgUTXO{
				Type: "移除共识节点",
			}
		case *pkitypes.MsgAddRootCert:
			umsg = &model.MsgUTXO{
				Type: "新增根证书",
			}
		case *pkitypes.MsgAddCert:
			umsg = &model.MsgUTXO{
				Type: "新增证书",
			}
		case *pkitypes.MsgFreezeCert:
			umsg = &model.MsgUTXO{
				Type: "冻结证书",
			}
		case *pkitypes.MsgRevokeCert:
			umsg = &model.MsgUTXO{
				Type: "注销证书",
			}
		default:
			return nil, fmt.Errorf("not support route %v, type %v", msg.Route(), msg.Type())
		}
		mtx.UTXOMsgs[index] = umsg
		types[index] = umsg.Type
		taddresses := strings.Split(umsg.Addresses, ",")
		for _, addr := range taddresses {
			if len(addr) == 0 {
				continue
			}
			if _, ok := addressesMap[addr]; !ok {
				addressesMap[addr] = true
				addresses = append(addresses, addr)
			}
		}
		tassets := strings.Split(umsg.Assets, ",")
		for _, asset := range tassets {
			if len(asset) == 0 {
				continue
			}
			if _, ok := assetsMap[asset]; !ok {
				assetsMap[asset] = true
				assets = append(assets, asset)
			}
		}
		tcontracts := strings.Split(umsg.Contracts, ",")
		for _, contract := range tcontracts {
			if len(contract) == 0 {
				continue
			}
			if _, ok := contractsMap[contract]; !ok {
				contractsMap[contract] = true
				contracts = append(contracts, contract)
			} else if strings.HasPrefix(contract, ":") {
				contracts = append(contracts, contract)
			}
		}
	}
	if len(addresses) > 0 {
		mtx.Addresses = "," + strings.Join(addresses, ",") + ","
	}
	if len(assets) > 0 {
		mtx.Assets = "," + strings.Join(assets, ",") + ","
	}
	if len(contracts) > 0 {
		mtx.Contracts = "," + strings.Join(contracts, ",") + ","
	}
	mtx.Type = strings.Join(types, ",")

	synced.logger.Debug("processTx", "hash", hex.EncodeToString(hash))
	return mtx, nil
}

func ProcessMsgIssue(msg *utxotypes.MsgIssue) *model.MsgUTXO {
	mmsg := &model.MsgUTXO{
		InputNum:  len(msg.Inputs),
		OutputNum: len(msg.Outputs),
		Desc:      msg.Desc,
	}

	// 关联地址
	var addresses []string
	addressesMap := map[string]bool{}
	// 关联资产
	var assets []string
	assetsMap := map[string]bool{}
	for _, input := range msg.Inputs {
		mmsg.Inputs = append(mmsg.Inputs, &model.MsgInput{
			RefTxID:      input.RefTx,
			RefMsgOffset: input.RefMsg,
			RefOffset:    input.RefOffset,
			Address:      input.FromAddr,
			Amount:       input.Amount.String(),
			FrozenHeight: input.FrozenHeight,
		})
		if has := assetsMap[input.Amount.Denom]; !has {
			assetsMap[input.Amount.Denom] = true
			assets = append(assets, input.Amount.Denom)
		}
		if has := addressesMap[input.FromAddr]; !has {
			addressesMap[input.FromAddr] = true
			addresses = append(addresses, input.FromAddr)
		}
	}
	for _, output := range msg.Outputs {
		mmsg.Outputs = append(mmsg.Outputs, &model.MsgOutput{
			Address:      output.ToAddr,
			Amount:       output.Amount.String(),
			FrozenHeight: output.FrozenHeight,
		})
		if has := assetsMap[output.Amount.Denom]; !has {
			assetsMap[output.Amount.Denom] = true
			assets = append(assets, output.Amount.Denom)
		}
		if has := addressesMap[output.ToAddr]; !has {
			addressesMap[output.ToAddr] = true
			addresses = append(addresses, output.ToAddr)
		}
	}

	if len(addresses) > 0 {
		mmsg.Addresses = "," + strings.Join(addresses, ",") + ","
	}
	if len(assets) > 0 {
		mmsg.Assets = "," + strings.Join(assets, ",") + ","
	}
	mmsg.Type = "资产发行"
	return mmsg
}
func ProcessMsgSend(msg *utxotypes.MsgSend) *model.MsgUTXO {
	mmsg := &model.MsgUTXO{
		InputNum:  len(msg.Inputs),
		OutputNum: len(msg.Outputs),
		Desc:      msg.Desc,
	}

	// 关联地址
	var addresses []string
	addressesMap := map[string]bool{}
	// 关联资产
	var assets []string
	assetsMap := map[string]bool{}
	for _, input := range msg.Inputs {
		mmsg.Inputs = append(mmsg.Inputs, &model.MsgInput{
			RefTxID:      input.RefTx,
			RefMsgOffset: input.RefMsg,
			RefOffset:    input.RefOffset,
			Address:      input.FromAddr,
			Amount:       input.Amount.String(),
			FrozenHeight: input.FrozenHeight,
		})
		if has := assetsMap[input.Amount.Denom]; !has {
			assetsMap[input.Amount.Denom] = true
			assets = append(assets, input.Amount.Denom)
		}
		if has := addressesMap[input.FromAddr]; !has {
			addressesMap[input.FromAddr] = true
			addresses = append(addresses, input.FromAddr)
		}
	}
	for _, output := range msg.Outputs {
		mmsg.Outputs = append(mmsg.Outputs, &model.MsgOutput{
			Address:      output.ToAddr,
			Amount:       output.Amount.String(),
			FrozenHeight: output.FrozenHeight,
		})
		if has := assetsMap[output.Amount.Denom]; !has {
			assetsMap[output.Amount.Denom] = true
			assets = append(assets, output.Amount.Denom)
		}
		if has := addressesMap[output.ToAddr]; !has {
			addressesMap[output.ToAddr] = true
			addresses = append(addresses, output.ToAddr)
		}
	}

	if len(addresses) > 0 {
		mmsg.Addresses = "," + strings.Join(addresses, ",") + ","
	}
	if len(assets) > 0 {
		mmsg.Assets = "," + strings.Join(assets, ",") + ","
	}
	mmsg.Type = "资产转移"
	return mmsg
}
func ProcessMsgDestroy(msg *utxotypes.MsgDestroy) *model.MsgUTXO {
	mmsg := &model.MsgUTXO{
		InputNum:  len(msg.Inputs),
		OutputNum: len(msg.Outputs),
		Desc:      msg.Desc,
	}

	// 关联地址
	var addresses []string
	addressesMap := map[string]bool{}
	// 关联资产
	var assets []string
	assetsMap := map[string]bool{}
	for _, input := range msg.Inputs {
		mmsg.Inputs = append(mmsg.Inputs, &model.MsgInput{
			RefTxID:      input.RefTx,
			RefMsgOffset: input.RefMsg,
			RefOffset:    input.RefOffset,
			Address:      input.FromAddr,
			Amount:       input.Amount.String(),
			FrozenHeight: input.FrozenHeight,
		})
		if has := assetsMap[input.Amount.Denom]; !has {
			assetsMap[input.Amount.Denom] = true
			assets = append(assets, input.Amount.Denom)
		}
		if has := addressesMap[input.FromAddr]; !has {
			addressesMap[input.FromAddr] = true
			addresses = append(addresses, input.FromAddr)
		}
	}
	for _, output := range msg.Outputs {
		mmsg.Outputs = append(mmsg.Outputs, &model.MsgOutput{
			Address:      output.ToAddr,
			Amount:       output.Amount.String(),
			FrozenHeight: output.FrozenHeight,
		})
		if has := assetsMap[output.Amount.Denom]; !has {
			assetsMap[output.Amount.Denom] = true
			assets = append(assets, output.Amount.Denom)
		}
		if has := addressesMap[output.ToAddr]; !has {
			addressesMap[output.ToAddr] = true
			addresses = append(addresses, output.ToAddr)
		}
	}

	if len(addresses) > 0 {
		mmsg.Addresses = "," + strings.Join(addresses, ",") + ","
	}
	if len(assets) > 0 {
		mmsg.Assets = "," + strings.Join(assets, ",") + ","
	}
	mmsg.Type = "资产销毁"
	return mmsg
}

func ProcessMsgDeploy(msg *wasmtypes.MsgDeploy) *model.MsgUTXO {
	mmsg := &model.MsgUTXO{
		InputNum:           len(msg.Inputs),
		OutputNum:          len(msg.Outputs),
		ContractRequestNum: 1,
		Desc:               msg.Desc,
	}

	// 关联地址
	var addresses []string
	addressesMap := map[string]bool{}
	// 关联资产
	var assets []string
	assetsMap := map[string]bool{}
	// 关联合约
	var contracts []string
	contractsMap := map[string]bool{}
	for _, input := range msg.Inputs {
		mmsg.Inputs = append(mmsg.Inputs, &model.MsgInput{
			RefTxID:      bytes.HexBytes(input.RefTx).String(),
			RefMsgOffset: input.RefMsg,
			RefOffset:    input.RefOffset,
			Address:      input.FromAddr,
			Amount:       input.Amount.String(),
			FrozenHeight: input.FrozenHeight,
		})
		if has := assetsMap[input.Amount.Denom]; !has {
			assetsMap[input.Amount.Denom] = true
			assets = append(assets, input.Amount.Denom)
		}
		if has := addressesMap[input.FromAddr]; !has {
			addressesMap[input.FromAddr] = true
			addresses = append(addresses, input.FromAddr)
		}
	}
	for _, output := range msg.Outputs {
		mmsg.Outputs = append(mmsg.Outputs, &model.MsgOutput{
			Address:      output.ToAddr,
			Amount:       output.Amount.String(),
			FrozenHeight: output.FrozenHeight,
		})
		if has := assetsMap[output.Amount.Denom]; !has {
			assetsMap[output.Amount.Denom] = true
			assets = append(assets, output.Amount.Denom)
		}
		if has := addressesMap[output.ToAddr]; !has {
			addressesMap[output.ToAddr] = true
			addresses = append(addresses, output.ToAddr)
		}
	}

	mmsg.ContractRequests = append(mmsg.ContractRequests, &model.InvokeRequest{
		ModuleName:   msg.ContractDesc.ContractType,
		ContractName: msg.ContractName,
		Args:         msg.Args,
	})
	contractName := ":" + msg.ContractName
	if has := contractsMap[contractName]; !has {
		contractsMap[contractName] = true
		contracts = append(contracts, contractName)
	}
	if len(addresses) > 0 {
		mmsg.Addresses = "," + strings.Join(addresses, ",") + ","
	}
	if len(assets) > 0 {
		mmsg.Assets = "," + strings.Join(assets, ",") + ","
	}
	if len(contracts) > 0 {
		mmsg.Contracts = "," + strings.Join(contracts, ",") + ","
	}
	mmsg.Type = "合约部署"
	return mmsg
}
func ProcessMsgUpgrade(msg *wasmtypes.MsgUpgrade) *model.MsgUTXO {
	mmsg := &model.MsgUTXO{
		InputNum:           len(msg.Inputs),
		OutputNum:          len(msg.Outputs),
		ContractRequestNum: 1,
		Desc:               msg.Desc,
	}

	// 关联地址
	var addresses []string
	addressesMap := map[string]bool{}
	// 关联资产
	var assets []string
	assetsMap := map[string]bool{}
	// 关联合约
	var contracts []string
	contractsMap := map[string]bool{}
	for _, input := range msg.Inputs {
		mmsg.Inputs = append(mmsg.Inputs, &model.MsgInput{
			RefTxID:      bytes.HexBytes(input.RefTx).String(),
			RefMsgOffset: input.RefMsg,
			RefOffset:    input.RefOffset,
			Address:      input.FromAddr,
			Amount:       input.Amount.String(),
			FrozenHeight: input.FrozenHeight,
		})
		if has := assetsMap[input.Amount.Denom]; !has {
			assetsMap[input.Amount.Denom] = true
			assets = append(assets, input.Amount.Denom)
		}
		if has := addressesMap[input.FromAddr]; !has {
			addressesMap[input.FromAddr] = true
			addresses = append(addresses, input.FromAddr)
		}
	}
	for _, output := range msg.Outputs {
		mmsg.Outputs = append(mmsg.Outputs, &model.MsgOutput{
			Address:      output.ToAddr,
			Amount:       output.Amount.String(),
			FrozenHeight: output.FrozenHeight,
		})
		if has := assetsMap[output.Amount.Denom]; !has {
			assetsMap[output.Amount.Denom] = true
			assets = append(assets, output.Amount.Denom)
		}
		if has := addressesMap[output.ToAddr]; !has {
			addressesMap[output.ToAddr] = true
			addresses = append(addresses, output.ToAddr)
		}
	}

	mmsg.ContractRequests = append(mmsg.ContractRequests, &model.InvokeRequest{
		ContractName: msg.ContractName,
	})
	contractName := ":" + msg.ContractName
	if has := contractsMap[contractName]; !has {
		contractsMap[contractName] = true
		contracts = append(contracts, contractName)
	}

	if len(addresses) > 0 {
		mmsg.Addresses = "," + strings.Join(addresses, ",") + ","
	}
	if len(assets) > 0 {
		mmsg.Assets = "," + strings.Join(assets, ",") + ","
	}
	if len(contracts) > 0 {
		mmsg.Contracts = "," + strings.Join(contracts, ",") + ","
	}
	mmsg.Type = "合约升级"
	return mmsg
}

func ProcessMsgInvoke(msg *wasmtypes.MsgInvoke) *model.MsgUTXO {
	mmsg := &model.MsgUTXO{
		InputNum:           len(msg.Inputs),
		OutputNum:          len(msg.Outputs),
		ContractRequestNum: len(msg.ContractRequests),
		Desc:               msg.Desc,
	}

	// 关联地址
	var addresses []string
	addressesMap := map[string]bool{}
	// 关联资产
	var assets []string
	assetsMap := map[string]bool{}
	// 关联合约
	var contracts []string
	contractsMap := map[string]bool{}
	for _, input := range msg.Inputs {
		mmsg.Inputs = append(mmsg.Inputs, &model.MsgInput{
			RefTxID:      bytes.HexBytes(input.RefTx).String(),
			RefMsgOffset: input.RefMsg,
			RefOffset:    input.RefOffset,
			Address:      input.FromAddr,
			Amount:       input.Amount.String(),
			FrozenHeight: input.FrozenHeight,
		})
		if has := assetsMap[input.Amount.Denom]; !has {
			assetsMap[input.Amount.Denom] = true
			assets = append(assets, input.Amount.Denom)
		}
		if has := addressesMap[input.FromAddr]; !has {
			addressesMap[input.FromAddr] = true
			addresses = append(addresses, input.FromAddr)
		}
	}
	for _, output := range msg.Outputs {
		mmsg.Outputs = append(mmsg.Outputs, &model.MsgOutput{
			Address:      output.ToAddr,
			Amount:       output.Amount.String(),
			FrozenHeight: output.FrozenHeight,
		})
		if has := assetsMap[output.Amount.Denom]; !has {
			assetsMap[output.Amount.Denom] = true
			assets = append(assets, output.Amount.Denom)
		}
		if has := addressesMap[output.ToAddr]; !has {
			addressesMap[output.ToAddr] = true
			addresses = append(addresses, output.ToAddr)
		}
	}

	for _, request := range msg.ContractRequests {
		mmsg.ContractRequests = append(mmsg.ContractRequests, &model.InvokeRequest{
			ModuleName:   request.ModuleName,
			ContractName: request.ContractName,
			MethodName:   request.MethodName,
			Args:         request.Args,
			Amount:       request.Amount.String(),
		})
		contractName := request.ContractName
		if request.ModuleName == "kernel" &&
			(request.MethodName == "Deploy" || request.MethodName == "Upgrade") {
			var args map[string][]byte
			if err := json.Unmarshal([]byte(request.Args), &args); err != nil {
				panic(err)
			}
			if request.MethodName == "Deploy" {
				contractName = ":" + string(args["contract_name"])
			} else {
				contractName = ":" + string(args["contract_name"])
			}
		}
		if has := contractsMap[contractName]; !has {
			contractsMap[contractName] = true
			contracts = append(contracts, contractName)
		}
	}
	if len(addresses) > 0 {
		mmsg.Addresses = "," + strings.Join(addresses, ",") + ","
	}
	if len(assets) > 0 {
		mmsg.Assets = "," + strings.Join(assets, ",") + ","
	}
	if len(contracts) > 0 {
		mmsg.Contracts = "," + strings.Join(contracts, ",") + ","
	}
	mmsg.Type = "合约调用"
	return mmsg
}

func (synced *Syncer) updatePeerStatus(peer *model.Peer, db *gorm.DB) {
	var validator model.Validator
	if result := db.Find(&validator, map[string]interface{}{
		"node_id": peer.NodeID,
	}); result.RowsAffected > 0 {
		peer.Type = 0
		return
	}
}
