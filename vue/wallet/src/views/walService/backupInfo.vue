<template>
  <div class="backupinfo">
    <div class="">
      <div class="backupinfo-title">密钥备份</div>
      <div class="cm-module-bg wallet-create-main">
        <div class="wallet-ct-success">
          <img
            src="../../assets/images/wallet/success_icon.png"
            alt=""
            srcset=""
          />
          私钥查看备份
          <p class="evid-desc">私钥信息关联你的账户资产，请妥善保管</p>
        </div>
        <div class="wallet-form">
          <div class="wallet-form-title wallet-form-success">
            <img
              src="../../assets/images/wallet/wallet_key_icon.png"
              alt=""
              srcset=""
            />
            钱包信息
          </div>
          <div class="wallet-form-info">
            <div class="form-info-row">
              <div class="info-row-name">钱包名称</div>
              <div class="">{{ backupInfo.name }}</div>
            </div>
            <div class="form-info-row">
              <div class="info-row-name">钱包类型</div>
              <div>{{ backupInfo.threshold | walType }}钱包</div>
            </div>
            <div class="form-info-row">
              <div class="info-row-name">钱包地址</div>
              <div class="cm-text-overflow public-key">
                {{ backupInfo.address }}
              </div>
              <a
                href="javascript:;"
                class="iconfont m0-copy info-copy-btn"
                v-clipboard:copy="backupInfo.address"
                v-clipboard:success="onCopy"
                v-clipboard:error="onError"
              ></a>
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
            我的密钥信息
          </div>
          <div class="wallet-form-info">
            <div class="form-info-row">
              <div class="info-row-name">公钥</div>
              <div class="cm-text-overflow public-key">
                {{ backupInfo.public_key }}
              </div>
              <a
                href="javascript:;"
                class="iconfont m0-copy info-copy-btn"
                v-clipboard:copy="backupInfo.public_key"
                v-clipboard:success="onCopy"
                v-clipboard:error="onError"
              ></a>
            </div>
            <div class="form-info-row">
              <div class="info-row-name">签名算法</div>
              <div>{{ backupInfo.algo }}</div>
            </div>
            <div class="form-info-row">
              <div class="info-row-name">私钥(HEX)</div>
              <div class="cm-text-overflow public-key">
                {{ backupInfo.private_key }}
              </div>
              <a
                href="javascript:;"
                class="iconfont m0-copy info-copy-btn"
                v-clipboard:copy="backupInfo.private_key"
                v-clipboard:success="onCopy"
                v-clipboard:error="onError"
              ></a>
            </div>
            <div class="form-info-row">
              <div class="info-row-name">私钥助记词</div>
              <div class="cm-text-overflow public-key">
                {{ backupInfo.mnemonic }}
              </div>
              <a
                href="javascript:;"
                class="iconfont m0-copy info-copy-btn"
                v-clipboard:copy="backupInfo.mnemonic"
                v-clipboard:success="onCopy"
                v-clipboard:error="onError"
              ></a>
            </div>
          </div>
        </div>
        <div class="wallet-btns cz-btns">
          <a
            href="javascript:;"
            class="cm-btn-295px cm-btn-border009F72"
            @click="goBank"
            >返回</a
          >
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { localCache } from "@/utils/utils";

export default {
  data() {
    return {
      backupInfo: {},
    };
  },
  created() {
    let backupInfo = localCache.get("backupInfo");
    if (backupInfo) {
      this.backupInfo = backupInfo;
    }
  },
  methods: {
    goBank() {
      localCache.remove("backupInfo");
      this.$router.go(-1);
    },
    onCopy(text) {
      if (text) {
        this.$message("复制成功");
      }
    },
    onError(e) {
      console.log(e);
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
.wallet-create-main {
  padding: 40px 0 118px;
  border-radius: 5px;
}
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
.cz-btns {
  padding: 0;
  justify-content: center;
}
.evid-desc {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
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

/**backupinfo start */
.backupinfo-title {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
  margin: 0 0 20px;
}
/**end backupinfo */
.public-key {
  width: 80%;
  line-height: 26px;
}
</style>