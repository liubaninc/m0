<template>
  <div class="wallet">
    <div class="wallet-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/assets' }"
          >资产管理</el-breadcrumb-item
        >
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >{{ trxInfo.type }}
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg wallet-create-main">
        <div class="wallet-ct-success">
          <img
            src="../../../assets/images/wallet/success_icon.png"
            alt=""
            srcset=""
          />
          {{ trxInfo.type }}交易已提交

          <p class="transfer-desc">
            将签名文件发送给钱包内其他密钥账户进行签名
          </p>
        </div>
        <div class="wallet-form wallet-form-mesg">
          <div class="wallet-form-title wallet-form-success">
            <img
              src="../../../assets/images/wallet/wallet_info_icon.png"
              alt=""
              srcset=""
            />
            签名信息
          </div>
          <div class="wallet-form-info">
            <div class="form-info-row">
              <div class="info-row-name">已签名/需签名</div>
              <div class="qm-cont">
                {{ trxInfo.signatures && trxInfo.signatures.length }}/{{
                  wallet.threshold
                }}
              </div>
            </div>
            <div class="form-info-row">
              <div class="info-row-name">已签名账户</div>
              <template v-if="trxInfo.signatures && trxInfo.signatures.length">
                <div
                  class="cm-text-overflow info-row-address"
                  v-for="sign in trxInfo.signatures"
                >
                  {{ sign }}
                </div>
              </template>
            </div>
            <div class="form-info-row">
              <div class="info-row-name">签名文件</div>
              <div class="cm-text-overflow info-row-address qm-file">
                {{ trxInfo.hash }}
              </div>
            </div>
          </div>
        </div>
        <div class="wallet-btns wl-success-btn">
          <a
            href="javascript:;"
            class="cm-btn-295px cm-btn-border009F72"
            @click="$router.push('/assets')"
            >返回资产列表</a
          >
          <a
            href="javascript:;"
            class="cm-btn-295px cm-btn-border009F72"
            @click="downloadFileByHash(trxInfo.hash)"
            >下载签名文件</a
          >
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { queryTrxDetail } from "@/server/transaction";
import { localCache } from "@/utils/utils";

export default {
  data() {
    return {
      trxInfo: {},
      wallet: {},
    };
  },
  created() {
    let { hash } = this.$route.query;
    let wallet = localCache.get("wallet");
    if (hash) {
      this.wallet = wallet;
      this.getTrxInfo(hash);
    }
  },
  methods: {
    downloadFileByHash(hash) {
      let origin = window.location.origin;
      let elink = document.createElement("a");
      elink.download = hash;
      elink.style.display = "none";
      elink.href = `${origin}/api/download/${hash}`;
      console.log(elink.href);
      document.body.appendChild(elink);
      elink.click();
      document.body.removeChild(elink);
    },
    async getTrxInfo(hash) {
      let trxInfo = await queryTrxDetail({
        hash,
      });
      if (trxInfo) {
        this.trxInfo = trxInfo;
      }
    },
  },
};
</script>
<style>
.wallet {
  background-color: #1b2c42;
  /* min-height: 100vh; */
}
.wallet-warpper {
  /* width: 80%; */
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
/* .wallet-steps .el-step__icon {
  border: 0;
  width: 13px;
  height: 13px;
  background: #768ca8;
} */
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
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
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
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: #1a2c42;
}

.wallet-btn-default {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
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
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
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
.wl-success-btn {
  padding: 0;
}

.back-asset-list {
  justify-content: center;
}

.transfer-desc {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #778592;
  text-align: center;
  margin: 16px 0 0;
}

.qm-cont {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: #ff5a58;
}
.qm-file {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #2dae95;
}
</style>