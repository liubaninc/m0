<template>
  <div class="evid">
    <div class="">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/mycontract' }"
          >我的合约</el-breadcrumb-item
        >
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >部署智能合约
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg wallet-create-main">
        <div class="wallet-ct-success">
          <img
            src="../../../assets/images/wallet/success_icon.png"
            alt=""
            srcset=""
          />
          交易已提交

          <template v-if="wallet.threshold > 0">
            <p class="evid-desc">将签名文件发送给钱包内其他密钥账户进行签名</p>
          </template>
        </div>
        <div class="wallet-form">
          <template v-if="wallet.threshold > 0">
            <div class="wallet-form-info">
              <div class="form-info-row">
                <div class="info-row-name">已签名/需签名</div>
                <div class="info-row-num">
                  {{
                    (trxInfo.signatures && trxInfo.signatures.length) || '0'
                  }}/{{ wallet.threshold }}
                </div>
              </div>
              <div class="form-info-row">
                <div class="info-row-name">已签名账户</div>
                <div class="info-row-address">
                  <template v-if="trxInfo.signatures">
                    <div
                      class="cm-text-overflow"
                      v-for="sign in trxInfo.signatures"
                    >
                      {{ sign }}
                    </div>
                  </template>
                  <template v-else>
                    <div class="cm-text-overflow">---</div>
                  </template>
                </div>
              </div>
              <div class="form-info-row">
                <div class="info-row-name">签名文件</div>
                <div class="cm-text-overflow info-row-file">
                  {{ trxInfo.hash || '---' }}
                </div>
              </div>
            </div>
          </template>
          <template v-else>
            <div class="wallet-form-title wallet-form-success">
              <img
                src="../../../assets/images/wallet/wallet_key_icon.png"
                alt=""
                srcset=""
              />
              存证信息
            </div>
            <div class="wallet-form-info">
              <div class="form-info-row">
                <div class="info-row-name">存证名称</div>
                <div class="cz-row">{{ eviDetail.name }}</div>
              </div>
              <div class="form-info-row">
                <div class="info-row-name">交易哈希</div>
                <el-tooltip
                  effect="dark"
                  :content="eviDetail.thash"
                  placement="top-end"
                >
                  <div class="cm-text-overflow cz-row">
                    {{ eviDetail.thash }}
                  </div>
                </el-tooltip>
              </div>
            </div>
          </template>
        </div>
        <div class="wallet-btns cz-btns">
          <a
            href="javascript:;"
            class="cm-btn-295px cm-btn-border009F72"
            @click="$router.push(`/mycontract`)"
            >返回合约列表</a
          >
          <a
            href="javascript:;"
            class="cm-btn-295px cm-btn-bg4acb9b download-btn"
            @click="downLoadSign"
            >下载签名文件</a
          >
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { localCache, downLoadFile } from '@/utils/utils'
import { queryEvidenceDetail } from '@/server/dapps/evidence'
import { queryContractTrxByHash } from '@/server/contract'

export default {
  data() {
    return {
      wallet: {},
      eviDetail: {},
      trxInfo: {},
      loginUser: {},
    }
  },
  async created() {
    let { name, hash } = this.$route.query
    let wallet = localCache.get('wallet')
    let loginUser = localCache.get('loginUser')
    if (wallet) {
      this.wallet = wallet
      this.loginUser = loginUser
      this.getTrxInfo(hash, wallet.address)
    }
  },
  methods: {
    downLoadSign() {
      downLoadFile(
        `${process.env.VUE_APP_PRO_BASE_URL}/mcontract/tx/download/${this.trxInfo.hash}`
      )
    },
    async getTrxInfo(hash, address, coin) {
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
<style>
.wallet {
  background-color: #1b2c42;
  /* min-height: 100vh; */
}
/* .layout-header {
  height: 90px !important;
  line-height: 90px;
} */
.wallet-warpper {
  width: 90%;
  margin: 0 auto;
  color: #fff;
}
.wallet-create-main {
  padding: 40px 0 118px;
  border-radius: 5px;
}
/**steps start */
.wallet-steps {
  width: 40%;
  margin: 0 auto;
}
.wallet-steps .el-steps .is-process {
  color: #fff !important;
}
.wallet-steps .el-step__icon-inner {
  display: none;
}
.wallet-steps .el-step.is-horizontal .el-step__line {
  top: 50%;
  transform: translate(0, -50%);
  background: #768ca8;
}
.wallet-steps .is-process .is-text {
  background: #fff;
}
/**end steps */
.wallet-form {
  width: 70%;
  margin: 0 auto;
  padding: 0 0 60px;
}

.wallet-form-info {
  background: #fff;
  border-radius: 5px;
  padding: 30px 20px;
}
.form-info-row {
  display: flex;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #1a2c42;
  margin: 0 0 20px;
  align-items: center;
}
.info-row-address {
  width: 40%;
  line-height: 20px;
}
.form-info-row:last-child {
  margin-bottom: 0;
}
.info-copy-btn {
  width: 20px;
  height: 20px;
  display: block;
  /* background: url("../../../assets/images/wallet/copy_icon.png"); */
  background-size: 100%;
  background-repeat: no-repeat;
  margin-left: 20px;
  color: #22ac95;
}

.wallet-form-mesg {
  padding-bottom: 30px;
}

.info-row-name {
  width: 100px;
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: #1a2c42;
}

.wallet-btn-default {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #cccccc;
  border-radius: 5px;
  border: 1px solid rgba(118, 140, 168, 1);
}
.wallet-btn-cur {
  border-color: #62f7d4;
  color: #62f7d4;
}

.wallet-ct-success {
  text-align: center;
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
  margin: 24px 0 40px;
}
.wallet-form-success {
  display: flex;
  align-items: center;
  margin: 0 0 10px;
}
.wallet-form-success img {
  width: 22px;
  height: 22px;
  margin-right: 10px;
}
.wallet-ct-success img {
  width: 40px;
  height: 40px;
  display: block;
  margin: 0 auto 12px;
}
.wallet-btns {
  display: flex;
  justify-content: space-between;
  width: 70%;
  margin: 0 auto;
}
.cz-btns {
  padding: 0;
  justify-content: center;
}
.evid-desc {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #778592;
  margin: 10px 0 0;
}
.info-row-num {
  color: #ff5a58;
}
.info-row-file {
  color: #2dae95;
}
.download-btn {
  margin: 0 0 0 24px;
}

.cz-row {
  width: 90%;
}
</style>