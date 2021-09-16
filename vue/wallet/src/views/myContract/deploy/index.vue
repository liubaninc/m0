<template>
  <div class="publish">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/mycontract' }"
          >我的合约</el-breadcrumb-item
        >
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >部署智能合约
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg detail-main">
        <div class="detail-main-assets">
          <div class="assets-title">交易签名</div>
          <div class="cm-submodule-bg assets-info">
            <div class="transfer-trx-type">
              <div class="assets-info-row">
                <span class="row-name"> 当前钱包类型 </span>
                <div class="row-line">{{ wallet.threshold | walType }}</div>
              </div>
              <template v-if="wallet.threshold > 0">
                <div class="assets-info-row">
                  <span class="row-name qm-num-name">
                    已签名账户数/需签名数
                  </span>
                  <div class="row-line account-count">
                    {{
                      (trxInfo.signatures && trxInfo.signatures.length) || '0'
                    }}/{{ wallet.threshold }}
                  </div>
                </div>
              </template>
            </div>
            <div class="assets-info-row assets-info-txSigned">
              <span class="row-name">密钥密码 </span>
              <div class="assets-info-desc">
                <template v-if="wallet.threshold > 0">
                  输入密钥密码进行签名，多签钱包需关联密钥账户进行签名，满足需签名数后才可向区块链提交交易
                </template>
                <template v-else>
                  输入密钥密码进行签名，单签钱包，完成签名后即可向区块链提交交易
                </template>
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
          @click="handleContract"
          >{{ mode | btnStatus }}</a
        >
      </div>
    </div>
  </div>
</template>
<script>
import {
  publishContract,
  contractOperate,
  queryContractTrxByHash,
} from '@/server/contract'
import { localCache } from '@/utils/utils'
import { tipStatus } from '@/utils/filters/status'

export default {
  data() {
    return {
      pwd: '',
      wallet: {},
      id: '',
      mode: '',
      commit: true,
      hash: '',
      trxInfo: {},
    }
  },
  created() {
    let wallet = localCache.get('wallet')
    let { id, mode, hash } = this.$route.query
    if (wallet) {
      this.wallet = wallet
      this.id = id
      this.mode = mode && mode
      hash && this.getTrxInfo(hash, wallet.address)
    }
  },
  methods: {
    async handleContract() {
      let { id, pwd, wallet, trxInfo } = this
      if (!pwd) {
        this.$message.error(`请输入密码`)
        return
      }
      if (wallet.threshold > 1) {
        if (trxInfo.signatures) {
          if (trxInfo.signatures.length < wallet.threshold - 1) {
            this.commit = false
          } else {
            this.commit = true
          }
        } else {
          this.commit = false
        }
      }
      if (this.mode == 'deploy') {
        this.toPublishContract(id, pwd, this.commit)
      } else {
        this.toOptContract(id, pwd, this.commit, this.mode)
      }
    },
    async toOptContract(id, pwd, commit, mode = 'upgrade') {
      this.loading = this.$loading({
        lock: true,
        text: `${tipStatus(mode)}...`,
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)',
      })
      let optResult = await contractOperate({
        id,
        mode,
        password: pwd,
        commit,
        accountName: this.wallet.name,
      })
      if (optResult) {
        this.loading.close()
        if (commit) {
          this.$router.push(`/myContract`)
        } else {
          this.$router.push(
            `/mycontract/deploy/moreDeploy?hash=${optResult.hash}`
          )
        }
      } else {
        this.loading.close()
      }
    },
    async toPublishContract(id, pwd, commit) {
      this.loading = this.$loading({
        lock: true,
        text: '部署中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)',
      })

      let rePublish = await publishContract({
        id,
        mode: 'deploy',
        password: pwd,
        commit,
        accountName: this.wallet.name,
      })
      if (rePublish) {
        this.loading.close()
        if (commit) {
          this.$router.push(`/myContract`)
        } else {
          this.$router.push(
            `/mycontract/deploy/moreDeploy?hash=${rePublish.hash}`
          )
        }
      } else {
        this.loading.close()
      }
    },
    async getTrxInfo(hash, address) {
      let trxInfo = await queryContractTrxByHash({
        hash,
        address,
      })
      if (trxInfo) {
        this.trxInfo = trxInfo
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
.publish .row-name {
  width: 100px;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
}
.publish .row-line {
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
.row-line-btns {
  padding: 0;
}
/*end detail* */

.upload-btn {
  margin: 0 0 0 24px;
}
.assets-info-txSigned {
  flex-direction: column;
  align-items: flex-start !important;
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

.transferout .add-address-line {
  padding: 0 0 0 20px;
}
.transferout-publish {
  padding: 35px 0 30px 0px;
  border-radius: 5px;
}
/**end transferout  */
/***publish start */
.publish {
}
.publish-assets-desc {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #768ca8;
  padding: 0 0 16px;
}
.publish-assets-input {
  width: 60%;
}
.publish-info-row {
  padding: 0 0 0 20px;
}
.publish-del-btn {
  width: 100px;
  font-family: PingFangSC-Regular, 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: rgb(255, 90, 88);
  margin: 0 0 0 20px;
}

.transfer-trx-type {
  display: flex;
}
.transfer-trx-type .row-name {
  width: 160px;
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
.transfer-trx-type .qm-num-name {
  width: 230px;
}
/***end publish */
</style>