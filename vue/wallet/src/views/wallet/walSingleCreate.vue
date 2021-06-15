<template>
  <div class="wallet">
    <!-- <div class="cm-width1200"> -->
    <div class="wallet-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/wallet' }"
          >我的钱包</el-breadcrumb-item
        >
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >创建钱包
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg wallet-create-main">
        <div class="wallet-steps">
          <el-steps :active="active" align-center finish-status="wait">
            <el-step title="选择钱包类型"></el-step>
            <el-step title="创建钱包"></el-step>
            <el-step title="创建成功"></el-step>
          </el-steps>
        </div>
        <div class="wallet-form">
          <div class="wallet-form-title">
            <img
              src="../../assets/images/wallet/single_wt_icon.png"
              alt=""
              srcset=""
            />
            创建单签钱包
          </div>
          <div class="wallet-form-rows">
            <div class="form-rows-item">
              <div class="form-rows-name">请设置钱包名称</div>
              <div class="form-rows-desc">
                钱包名称是用来区分钱包的标签，该信息不会在区块链上保存
              </div>
              <div class="form-rows-input">
                <el-input
                  placeholder="请设置钱包名称"
                  v-model="walletName"
                  clearable
                >
                </el-input>
              </div>
            </div>
            <div class="form-rows-item">
              <div class="form-rows-name">密钥算法</div>
              <div class="form-rows-input">
                <template v-for="(alg, index) in algorithms">
                  <a
                    :key="alg.id"
                    href="javascript:;"
                    class="cm-btn-autopx wallet-btn-default"
                    :class="curAlg == index ? 'wallet-btn-cur' : ''"
                    @click="chooseAlg(alg, index)"
                    >{{ alg.text }}</a
                  >
                </template>
                <!-- <a
                  href="javascript:;"
                  class="cm-btn-303px wallet-btn-default wallet-btn-cur"
                  >SM2</a
                > -->
              </div>
            </div>
            <div class="form-rows-item">
              <div class="form-rows-name">请设置密钥密码</div>
              <div class="form-rows-desc">
                请设置至少10位字母数字混合的密码，密钥是根据你输入的密码生成的管理资产的加密凭证。请妥善保管、备份密码，忘记密码将导致钱包资产的损失。
              </div>
              <div class="form-rows-input">
                <el-input
                  placeholder="请设置密钥密码"
                  v-model="pwd"
                  show-password
                  clearable
                >
                </el-input>
              </div>
              <div class="form-rows-input rpsw-row">
                <el-input
                  placeholder="请确认密钥密码"
                  v-model="repwd"
                  show-password
                  clearable
                >
                </el-input>
              </div>
            </div>
          </div>
        </div>
        <div class="wallet-btns">
          <a
            href="javascript:;"
            class="cm-btn-200px cm-btn-border009F72 cm-btn-back"
            @click="$router.go(-1)"
            >返回</a
          >
          <a
            href="javascript:;"
            class="cm-btn-200px cm-btn-bg4acb9b cm-btn-back"
            @click="createWallet"
            >创建钱包</a
          >
        </div>
      </div>
      <!-- </div> -->
    </div>
  </div>
</template>
<script>
import { createWallet } from "@/server/wallet";
import { localCache } from "@/utils/utils";

export default {
  data() {
    return {
      algorithms: [
        {
          id: 10001,
          text: "SECP256K1",
        },
        {
          id: 10002,
          text: "SM2",
        },
      ],
      active: 1,
      walletName: "",
      curAlg: 0,
      pwd: "",
      repwd: "",
    };
  },
  methods: {
    chooseAlg(alg, index) {
      this.curAlg = index;
    },
    async createWallet() {
      let { walletName, pwd, repwd, curAlg, algorithms } = this;
      if (!walletName) {
        return this.$message.error("请输入钱包名称");
      }
      if (curAlg < 0) {
        return this.$message.error("请选择签名算法");
      }

      if (!/^(?=.*[0-9])(?=.*[a-zA-Z]).{10,30}$/.test(pwd)) {
        return this.$message.error("请设置至少10位字母数字混合的密码");
      }
      if (!repwd) {
        return this.$message.error("请输入确认密码");
      }
      if (pwd != repwd) {
        return this.$message.error("俩次输入的密码不一致");
      }

      let walletInfo = await createWallet.call(this, {
        name: walletName,
        password: pwd,
        algo: algorithms[curAlg].text,
      });
      if (walletInfo) {
        localCache.set("walletInfo", walletInfo);
        this.$router.push(`/wallet/walCreateSuccess`);
      }
    },
  },
};
</script>
<style scoped>
.wallet {
  background-color: #1b2c42;
  /* min-height: calc(100vh - 133px); */
}
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
.form-rows-item {
  margin: 0 0 30px;
}
.form-rows-item:last-child {
  margin-bottom: 0;
}
.wallet-form-title {
  display: flex;
  align-items: center;
  justify-items: center;
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 16px;
  color: #62f7d4;
  margin: 0 0 20px;
}
.wallet-form-title img {
  width: 22px;
  height: 22px;
  margin: 0 12px 0 0;
}
.form-rows-name {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  margin: 10px 0;
}
.form-rows-desc {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #768ca8;
  margin: 0px 0 10px;
  line-height: 20px;
}

.form-rows-input {
  height: 45px;
  line-height: 45px;
  display: flex;
  align-items: center;
  justify-content: space-between;
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

.wallet-btns {
  display: flex;
  justify-content: center;
}
.cm-btn-back {
  margin-right: 50px;
}
.cm-btn-back:last-child {
  margin-right: 0;
}
.rpsw-row {
  margin-top: 6px;
}
</style>