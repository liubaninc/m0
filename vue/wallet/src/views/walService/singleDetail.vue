<template>
  <div class="sgdetail">
    <div class="com-breadcrumb sgdetail-title">钱包详情</div>
    <div class="cm-module-bg wallet-create-main">
      <div class="wallet-ct-success">
        <img
          src="../../assets/images/wallet/success_icon.png"
          alt=""
          srcset=""
        />
        钱包详情
      </div>
      <div class="wallet-form wallet-form-mesg">
        <div class="wallet-form-title wallet-form-success">
          <img
            src="../../assets/images/wallet/wallet_info_icon.png"
            alt=""
            srcset=""
          />
          钱包信息
        </div>
        <div class="wallet-form-info">
          <div class="form-info-row">
            <div class="info-row-name">钱包名称</div>
            <div>{{ walletInfo.name }}</div>
          </div>
          <div class="form-info-row">
            <div class="info-row-name">钱包类型</div>
            <div>{{ walletInfo.threshold | walType }}钱包</div>
          </div>
          <div class="form-info-row">
            <div class="info-row-name">钱包地址</div>
            <div class="sign-address-row public-sign-more">
              <div class="cm-text-overflow sign-id-row">
                {{ walletInfo.address }}
              </div>
              <a
                href="javascript:;"
                class="iconfont m0-copy info-copy-btn"
                v-clipboard:copy="walletInfo.address"
                v-clipboard:success="onCopy"
                v-clipboard:error="onError"
              ></a>
            </div>
          </div>
        </div>
      </div>
      <div class="wallet-form">
        <div class="wallet-form-title wallet-form-success">
          <img
            src="../../assets/images/wallet/wallet_key_icon.png"
            alt=""
            srcset=""
          />
          关联密钥信息
        </div>
        <div class="wallet-form-info">
          <template v-if="walletInfo.threshold > 0">
            <div class="form-info-row">
              <div class="info-row-name">多签公钥</div>
              <div class="sign-address-row public-sign-more">
                <div class="cm-text-overflow sign-id-row">
                  {{ walletInfo.public_key }}
                </div>
                <a
                  href="javascript:;"
                  class="iconfont m0-copy info-copy-btn"
                  v-clipboard:copy="walletInfo.public_key"
                  v-clipboard:success="onCopy"
                  v-clipboard:error="onError"
                ></a>
              </div>
            </div>
            <div class="form-info-row">
              <div class="info-row-name">参与公钥</div>
              <div class="sign-address-row">
                <template
                  v-if="walletInfo.multi_sig && walletInfo.multi_sig.length"
                >
                  <div
                    class="public-sign-more"
                    v-for="(sign, index) in walletInfo.multi_sig"
                    :key="index"
                  >
                    <div class="cm-text-overflow sign-id-row">{{ sign }}</div>

                    <a
                      href="javascript:;"
                      class="iconfont m0-copy info-copy-btn"
                      v-clipboard:copy="sign"
                      v-clipboard:success="onCopy"
                      v-clipboard:error="onError"
                    ></a>
                  </div>
                </template>
              </div>
            </div>
            <div class="form-info-row">
              <div class="info-row-name">签名数</div>
              <div>{{ walletInfo.threshold }}</div>
            </div>
          </template>
          <template v-else>
            <div class="form-info-row">
              <div class="info-row-name">公钥</div>
              <div class="info-row-pbcontext">
                <div class="cm-text-overflow public-key">
                  {{ walletInfo.public_key }}
                </div>
                <a
                  href="javascript:;"
                  class="iconfont m0-copy info-copy-btn"
                  v-clipboard:copy="walletInfo.public_key"
                  v-clipboard:success="onCopy"
                  v-clipboard:error="onError"
                ></a>
              </div>
            </div>
            <div class="form-info-row">
              <div class="info-row-name">签名算法</div>
              <div class="public-key">{{ walletInfo.algo }}</div>
            </div>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { queryWalletInfo } from "@/server/wallet";
import { localCache } from "@/utils/utils";

export default {
  data() {
    return {
      walletInfo: {},
    };
  },
  created() {
    let wallet = localCache.get("wallet");
    if (wallet) {
      this.getWalletDetailByName(wallet.name);
    }
  },
  methods: {
    onCopy(text) {
      if (text) {
        this.$message("复制成功");
      }
    },
    onError(e) {
      console.log(e);
    },
    async getWalletDetailByName(name) {
      let walletInfo = await queryWalletInfo({
        name,
      });
      if (walletInfo) {
        this.walletInfo = walletInfo;
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
/* .layout-header {
  height: 90px !important;
  line-height: 90px;
} */
.wallet-warpper {
  width: 90%;
  margin: 0 auto;
  color: #fff;
}
.sgdetail .wallet-create-main {
  padding: 40px 0 0;
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
  /* background: url("../../assets/images/wallet/copy_icon.png"); */
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
/** */
.sgdetail-title {
  font-size: 18px;
  margin: 0 0 20px;
  font-weight: 700;
}
/* .pb-key-row,
.wallet-address-row,
.more-public-key {
  width: 80%;
} */
.public-sign-more {
  width: 100%;
  display: flex;
  align-items: center;
}
.sign-address-row {
  width: 80%;
}
.sign-id-row {
  width: 90%;
  line-height: 30px;
}

.info-row-pbcontext {
  display: flex;
  width: 80%;
}
.public-key {
  width: 90%;
}
</style>