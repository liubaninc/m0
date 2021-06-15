<template>
  <div class="backup">
    <div class="detail-warpper">
      <div class="com-breadcrumb backup-title">密钥备份</div>
      <div class="cm-module-bg detail-main">
        <div class="detail-main-assets">
          <div class="assets-title">我的钱包信息</div>
          <div class="cm-submodule-bg assets-info">
            <div class="assets-info-row">
              <span class="row-name">钱包名称 </span>
              <div class="cm-row-bg row-line">
                {{ wallet.name }}
              </div>
            </div>
            <div class="assets-info-row">
              <span class="row-name">钱包类型 </span>
              <div class="cm-row-bg row-line">
                {{ wallet.threshold | walType }}钱包
              </div>
            </div>
            <div class="assets-info-row">
              <span class="row-name">钱包地址 </span>
              <div class="cm-row-bg row-line">
                {{ wallet.address }}
              </div>
            </div>
          </div>
        </div>

        <div class="detail-main-assets">
          <div class="assets-title">交易密码</div>
          <div class="cm-submodule-bg assets-info">
            <div class="assets-info-row trxpwd-info-row">
              <span class="row-name">密钥密码 </span>
              <div class="trxpwd-desc">
                输入密钥密码查看私钥信息，私钥信息关联你的账户资产，请妥善保管。
              </div>
              <div class="trxpwd-input">
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

      <a
        href="javascript:;"
        class="cm-btn-200px cm-btn-bg009F72 reback-btn"
        @click="toViewPrivateInfo"
        >查看密钥信息</a
      >
    </div>
  </div>
</template>
<script>
import { exportAcc } from "@/server/account/account";
import { localCache } from "@/utils/utils";
export default {
  data() {
    return {
      pwd: "",
    };
  },
  created() {
    let wallet = localCache.get("wallet");
    if (wallet) {
      this.wallet = wallet;
    }
  },
  methods: {
    toViewPrivateInfo() {
      let { wallet, pwd } = this;
      if (!pwd) {
        this.$message.error(`请输入密码`);
        return;
      }
      this.toViewBankUp(wallet.name, pwd);
    },
    async toViewBankUp(name, password) {
      let backupInfo = await exportAcc.call(this, {
        name,
        password,
      });
      if (backupInfo) {
        localCache.set("backupInfo", backupInfo);
        this.$router.push(`/walService/backupInfo`);
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
.detail-warpper {
  margin: 0 auto;
  color: #fff;
}

/*detail start* */
.assets-title {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
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
.assets-info-row {
  display: flex;
  align-items: center;
  margin: 0 0 20px;
}
.assets-info-row:last-child {
  margin-bottom: 0;
}
.row-name {
  width: 80px;
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
}
.row-line {
  min-height: 45px;
  width: 60%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  padding: 10px 20px;
}
.row-textarea {
  line-height: 20px;
}
.copy-btn {
  width: 16px;
  height: 16px;
  display: inline-block;
  /* background: url("../../assets/images/wallet/copy_icon.png"); */
  background-size: 100%;
  cursor: pointer;
  color: #22ac95;
}
.row-line-btns {
  padding: 0;
}
.reback-btn {
  margin: 30px auto;
}
/*end detail* */
/**backup start */
.backup-title {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
  margin: 0 0 20px;
}
.trxpwd-info-row {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}
.trxpwd-desc {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #768ca8;
  margin: 10px 0 20px;
}
.trxpwd-input {
  width: 70%;
}
/**end backup */
</style>