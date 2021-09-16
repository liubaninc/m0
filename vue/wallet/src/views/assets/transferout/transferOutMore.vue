<template>
  <div class="transferout">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/assets' }"
          >资产管理</el-breadcrumb-item
        >
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >资产详情
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg detail-main">
        <div class="detail-main-assets">
          <div class="assets-title">资产信息</div>
          <div class="cm-submodule-bg assets-info">
            <div class="assets-info-row">
              <span class="row-name">资产名称 </span>
              <div class="cm-row-bg row-line transferout-row-line">
                {{ assetInfo.denom }}
              </div>
            </div>
            <div class="assets-info-row">
              <span class="row-name">资产地址 </span>
              <div
                class="cm-text-overflow cm-row-bg row-line transferout-row-line"
              >
                {{ wallet.address }}
              </div>
            </div>
            <div class="assets-info-row">
              <span class="row-name">资产余额 </span>
              <div class="row-line">
                {{ assetInfo.amount }}
              </div>
            </div>
          </div>
        </div>

        <div class="detail-main-assets">
          <div class="assets-title">转账详情</div>
          <div class="cm-submodule-bg transferout-publish">
            <div class="transferout-assets-info">
              <div class="assets-info-row transferout-info-row">
                <span class="row-name">目标地址 </span>
                <div class="row-line">
                  <el-input
                    v-model="to"
                    placeholder="请输入目标地址"
                  ></el-input>
                </div>
              </div>
              <div class="assets-info-row transferout-info-row">
                <span class="row-name">转出数量 </span>
                <div class="row-line">
                  <el-input
                    v-model="amount"
                    placeholder="请输入转出数量"
                  ></el-input>
                </div>
              </div>
              <a href="javascript:;" class="del-row"></a>
            </div>
            <template v-for="(rowInput, index) in keysInput">
              <div class="transferout-assets-info">
                <div class="assets-info-row transferout-info-row">
                  <span class="row-name">目标地址 </span>
                  <div class="row-line">
                    <el-input
                      v-model="rowInput.to"
                      placeholder="请输入目标地址"
                    ></el-input>
                  </div>
                </div>
                <div class="assets-info-row transferout-info-row">
                  <span class="row-name">转出数量 </span>
                  <div class="row-line">
                    <el-input
                      v-model="rowInput.amount"
                      placeholder="请输入转出数量"
                    ></el-input>
                  </div>
                </div>
                <a href="javascript:;" class="del-row" @click="delRow(index)"
                  >删除</a
                >
              </div>
            </template>
            <div class="assets-info-row add-address">
              <span class="row-name"></span>
              <div class="row-line">
                <a
                  href="javascript:;"
                  class="cm-btn-295px cm-btn-border009F72 evidence-btn"
                  @click="addRow"
                  >添加转账地址</a
                >
              </div>
            </div>
          </div>
        </div>

        <div class="detail-main-assets">
          <div class="assets-title">交易签名</div>
          <div class="cm-submodule-bg assets-info">
            <div class="transfer-trx-type">
              <div class="assets-info-row">
                <span class="row-name"> 当前钱包类型 </span>
                <div class="row-line">多签钱包</div>
              </div>
              <div class="assets-info-row">
                <span class="row-name"> 已签名账户数/需签名数 </span>
                <div class="row-line account-count">
                  0/{{ wallet.threshold }}
                </div>
              </div>
            </div>
            <div class="assets-info-row assets-info-txSigned">
              <span class="row-name">密钥密码 </span>
              <div class="assets-info-desc">
                输入密钥密码进行签名，单签钱包，完成签名后即可向区块链提交交易
              </div>
              <div class="assets-info-input">
                <el-input
                  v-model="pwd"
                  placeholder="请输入密码"
                  show-password
                ></el-input>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="singleevid-btns">
        <a
          href="javascript:;"
          class="cm-btn-295px cm-btn-border009F72 evidence-btn"
          @click="$router.go(-1)"
          >返回</a
        >
        <a
          href="javascript:;"
          class="cm-btn-295px cm-btn-bg4acb9b evidence-btn"
          @click="transferOutAsset"
          >转出资产</a
        >
      </div>
    </div>
  </div>
</template>
<script>
import { queryAssetsByAddress, transferOutAssets } from '@/server/assets'
import { localCache } from '@/utils/utils'

export default {
  data() {
    return {
      assetInfo: {},
      coin: '',
      to: '',
      amount: '',
      pwd: '',
      keysInput: [],
      wallet: {},
    }
  },
  created() {
    let { coin } = this.$route.query
    let wallet = localCache.get('wallet')
    if (coin && wallet) {
      this.coin = coin
      this.wallet = wallet
      this.getAssetsInfo(wallet.address, this.coin)
    }
  },
  methods: {
    async transferOutAsset() {
      let { wallet, pwd, amount, to, keysInput } = this

      if (!to) {
        this.$message.error(`目标地址不能为空`)
        return
      }

      if (!amount) {
        this.$message.error(`转出数量不能为空`)
        return
      }

      if (!pwd) {
        this.$message.error(`密码不能为空`)
        return
      }

      if (keysInput.length) {
        keysInput.forEach((item) => {
          let mnt = /[^0-9](.+)?/gi.exec(item['amount'])
          if (!mnt) {
            item['amount'] = '' + item['amount'] + this.coin
          }
        })
      }
      let trxs = await transferOutAssets.call(this, {
        from: wallet.name,
        tos: [{ to, amount: amount + this.coin }, ...keysInput],
        password: pwd,
        commit: wallet.threshold > 1 ? !1 : !0,
      })
      if (trxs) {
        this.$router.push(`/assets/transferOutIng?hash=${trxs.hash}`)
      }
    },
    addRow() {
      this.keysInput.push({ to: '', amount: '' })
    },
    delRow(index) {
      this.keysInput.splice(index, 1)
    },
    async getAssetsInfo(addrs, coin) {
      let { address, coins } = await queryAssetsByAddress({
        address: addrs,
        coin,
      })
      if (coins && coins.length) {
        this.assetInfo = coins[0]
      }
    },
  },
}
</script>
<style scoped>
.detail-warpper {
  margin: 0 auto;
  color: #fff;
}

/*detail start* */
.assets-title {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  height: 40px;
  line-height: 40px;
}
.detail-main-assets {
  padding: 20px 40px;
}
.assets-info {
  padding: 35px 0 30px 20px;
  border-radius: 5px;
}
.transferout-assets-info {
  display: flex;
  align-items: center;
  padding: 0px 0 0px 20px;
  border-radius: 5px;
}
.assets-info-row {
  display: flex;
  align-items: center;
  margin: 0 0 20px;
}
.transferout-info-row {
  margin-bottom: 0;
  margin-left: 20px;
  width: 80%;
}
.transferout-info-row:first-child {
  margin-left: 0;
}
.assets-info-row:last-child {
  margin-bottom: 0;
}
.transferout .row-name {
  width: 100px;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
}
.transferout .row-line {
  min-height: 45px;
  /* line-height: 45px; */
  width: 60%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  padding: 10px 0px;
}
.row-textarea {
  line-height: 20px;
}
.copy-btn {
  width: 16px;
  height: 16px;
  display: inline-block;
  /* background: url("../../../assets/images/wallet/copy_icon.png"); */
  background-size: 100%;
  cursor: pointer;
  color: #22ac95;
}
.row-line-btns {
  padding: 0;
}
/*end detail* */

.upload-btn {
  margin: 0 0 0 24px;
}
.assets-info-txSigned {
  flex-direction: column;
  align-items: flex-start;
}
.assets-info-desc {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #768ca8;
  line-height: 40px;
}
.assets-info-input {
  width: 60%;
  margin: 2px 0 0;
}

.singleevid-btns {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 30px 0;
}
.evidence-btn {
  font-size: 14px;
  margin: 0 0 0 20px;
}
.evidence-btn:first-child {
  margin-left: 0;
}
/**transferout start */
.transfer-assets-name {
  width: 40%;
  height: 40px;
  line-height: 40px;
  padding: 0 0 0 20px;

  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  letter-spacing: normal;
  color: #ffffff;
}

.transferout .add-address {
  padding: 0 0 0 20px;
}
.transferout-publish {
  padding: 35px 0 30px 0px;
  border-radius: 5px;
}
.transferout .transferout-row-line {
  padding: 0 0 0 20px;
  width: 38%;
}
.transferout-info-row .row-line {
  width: 80%;
}
.del-row {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ff5a58;
  display: inline-block;
  width: 106px;
  margin: 0 0 0 10px;
}
.transfer-trx-type {
  display: flex;
}
.transfer-trx-type .row-name {
  width: auto;
}
.transfer-trx-type .assets-info-row {
  margin: 0;
}
.transferout .transfer-trx-type .row-line {
  width: auto;
}
.transfer-trx-type .row-line {
  padding: 0 20px;
}
.transfer-trx-type .account-count {
  color: #62f7d4;
}
/**end transferout  */
</style>