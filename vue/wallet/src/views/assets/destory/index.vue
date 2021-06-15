<template>
  <div class="destory">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/' }">资产管理</el-breadcrumb-item>
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >销毁资产
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg detail-main">
        <div class="detail-main-assets">
          <div class="assets-title">我的资产信息</div>
          <div class="cm-submodule-bg assets-info">
            <div class="assets-info-row">
              <span class="row-name">资产名称 </span>
              <div class="cm-row-bg row-line">
                <template v-if="assetDetail.coins">
                  {{ assetDetail.coins[0].denom }}
                </template>
              </div>
            </div>
            <div class="assets-info-row">
              <span class="row-name">资产地址 </span>
              <div class="cm-row-bg row-line">
                {{ assetDetail.address }}
              </div>
            </div>
            <div class="assets-info-row">
              <span class="row-name">资产余额 </span>
              <div class="cm-row-bg row-line">
                <template v-if="assetDetail.coins">
                  {{ assetDetail.coins[0].amount }}
                </template>
              </div>
            </div>
          </div>
        </div>

        <div class="detail-main-assets">
          <div class="assets-title">销毁去向</div>
          <div class="cm-submodule-bg assets-info">
            <div class="assets-info-row">
              <span class="row-name">目标地址 </span>
              <div class="cm-row-bg row-line">
                {{ burn }}
              </div>
            </div>
            <div class="assets-info-row">
              <span class="row-name">销毁数量 </span>
              <div class="row-line row-line-input">
                <el-input
                  v-model="amount"
                  placeholder="请输入销毁数量"
                ></el-input>
              </div>
            </div>
          </div>
        </div>

        <div class="detail-main-assets">
          <div class="assets-title">交易签名</div>
          <div class="cm-submodule-bg assets-info">
            <!-- <div class="assets-info-row destory-info-row">
              <span class="row-name">当前钱包类型 </span>
              <div class="row-line">{{ wallet.threshold | walType }}钱包</div>
            </div> -->
            <div class="destory-rows-types">
              <div class="assets-info-row destory-info-row">
                <span class="row-name wallet-type">当前钱包类型 </span>
                <div class="row-line">{{ wallet.threshold | walType }}钱包</div>
              </div>
              <template>
                <div
                  class="assets-info-row destory-info-row"
                  v-if="wallet.threshold > 0"
                >
                  <span class="row-name destory-qm-name"
                    >已签名账户数/需签名数
                  </span>
                  <div class="row-line destory-qm-count">
                    0/{{ wallet.threshold }}
                  </div>
                </div>
              </template>
            </div>

            <div class="destory-assets-pwdtitle">
              <p class="row-name">密钥密码</p>
              <p class="row-desc">
                输入密钥密码进行签名，单签钱包，完成签名后即可向区块链提交交易
              </p>
            </div>
            <div class="destory-assets-input">
              <el-input
                v-model="pwd"
                placeholder="请输入密码"
                show-password
              ></el-input>
            </div>
          </div>
        </div>
      </div>

      <div class="destory-btns">
        <a
          href="javascript:;"
          class="cm-btn-295px cm-btn-border009F72 reback-btn"
          @click="$router.go(-1)"
          >返回</a
        >
        <a
          href="javascript:;"
          class="cm-btn-295px cm-btn-bg4acb9b reback-btn"
          @click="burnAssets"
          >销毁资产</a
        >
      </div>
    </div>
    <!--vaild dialog start-->
    <div class="cm-warpper-bg vaild-bg" v-if="false">
      <div class="vaild-dialog">
        <div class="dialog-title">校验存证文件</div>
        <div class="dialog-row">
          <span class="dialog-row-title">上传文件 </span>
          <div class="dialog-row-input">
            <el-input v-model="input" placeholder="请输入内容"></el-input>
          </div>
          <a
            href="javascript:;"
            class="cm-btn-94px cm-btn-bg009F72 dialog-upload-btn"
            >上传</a
          >
        </div>
        <div class="dialog-row">
          <span class="dialog-row-title"> </span>
          <a href="javascript:;" class="cm-btn-225px cm-btn-border009F72"
            >取消</a
          >
          <a
            href="javascript:;"
            class="cm-btn-225px cm-btn-border009F72 immediate-btn"
            >立即校验</a
          >
        </div>
      </div>
    </div>
    <!--end vaild dialog-->
    <!--vaild-success start-->
    <div class="cm-warpper-bg vaild-success" v-if="false">
      <div class="success-dialog">
        <img src="../../../assets/images/wallet/success_icon.png" alt="" />
        <p class="success-dia-desc">校验成功，校验文件为存证文件</p>
        <a href="javascript:;" class="cm-btn-200px cm-btn-bg009F72 confirm-btn"
          >确定</a
        >
      </div>
    </div>
    <!--end vaild-success-->
  </div>
</template>
<script>
import { queryAssetsByAddress, burnAssetsByAdd } from "@/server/assets";
import { localCache } from "@/utils/utils";
export default {
  data() {
    return {
      wallet: {},
      assetDetail: {},
      burn: "tk12xnsae6sqwnjgtaefteyunxvea2uzevakpl6t8",
      pwd: "",
      amount: "",
      assetsName: "",
    };
  },
  created() {
    let { assetsName } = this.$route.query;
    let wallet = localCache.get("wallet");
    if (wallet) {
      this.wallet = wallet;
      this.assetsName = assetsName;

      this.getAssetsInfo(wallet.address, assetsName);
    }
  },
  methods: {
    async burnAssets() {
      let { amount, pwd, wallet, assetDetail, assetsName, burn } = this;
      if (!amount) {
        this.$message.error("请输入销毁数量");
        return;
      }
      if (!pwd) {
        this.$message.error("请输入密码");
        return;
      }

      amount = /[^0-9](.+)?/gi.exec(amount) ? amount : amount + assetsName;

      let commit = wallet.threshold > 1 ? false : true;
      let resBurn = await burnAssetsByAdd.call(this, {
        from: wallet.address,
        tos: [
          {
            amount,
          },
        ],
        commit,
        password: pwd,
      });

      if (resBurn) {
        if (wallet && wallet.threshold > 0) {
          this.$router.push(
            `/assets/publicIng?hash=${resBurn.hash}&coin=${assetsName}&address=${burn}`
          );
        } else {
          this.$router.push(
            `/assets/publicSuccess?hash=${resBurn.hash}&coin=${assetsName}&address=${burn}`
          );
        }
      }
    },
    async getAssetsInfo(address, coin) {
      let assetDetail = await queryAssetsByAddress({
        address,
        coin,
      });
      if (assetDetail) {
        this.assetDetail = assetDetail;
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
.destory .row-name {
  width: 94px;
}
.row-line {
  min-height: 45px;
  /* line-height: 45px; */
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
  /* background: url("../../../assets/images/wallet/copy_icon.png"); */
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

/**vaild-dialog start */

.vaild-dialog {
  width: 633px;
  background-color: #425263;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 50px 60px;
  border-radius: 5px;
}
.dialog-title {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
  text-align: center;
  margin: 0 0 50px;
}
.dialog-row {
  display: flex;
  height: 40px;
  line-height: 40px;
  margin: 0 0 40px;
}
.dialog-row:last-child {
  margin-bottom: 0;
}
.dialog-row-input {
  width: 338px;
}
.dialog-row-title {
  display: inline-block;
  width: 80px;
  height: 100%;
  font-size: 14px;
}
.dialog-upload-btn {
  margin: 0 0 0 20px;
}
.immediate-btn {
  margin: 0 0 0 20px;
}
/**end vaild-dialog */

/**vaild-success start */
.vaild-success {
}
.success-dialog {
  width: 633px;
  background-color: #425263;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 40px 20px 60px;
}
.success-dialog img {
  width: 54px;
  height: 54px;
}
.success-dia-desc {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 16px;
  color: #ffffff;
  margin: 40px 0 50px;
}
.confirm-btn {
  font-size: 16px;
}
/**end vaild-success */
/**destory start */
.destory-assets-pwdtitle {
  margin: 0 0 18px;
}
.row-desc {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #768ca8;
  margin: 8px 0 0;
}
.destory-assets-input {
  width: 76%;
}

.destory-btns {
  display: flex;
}
.destory-info-row {
  margin: 0 0 0;
}
.row-line-input {
  padding: 0;
}

.destory-rows-types {
  display: flex;
}

.destory-info-row .destory-qm-name {
  width: 240px;
}
.destory-info-row .destory-qm-count {
  color: #62f7d4;
}
.destory .wallet-type {
  width: 140px;
}
/**end destory */
</style>