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
        <div class="wallet-types">
          <div
            class="wallet-ty-item"
            :class="walletType == 1 ? 'cur-tab' : ''"
            @click="walletType = 1"
          >
            <div class="item-image">
              <img
                src="../../assets/images/wallet/single_wt_icon.png"
                alt=""
                srcset=""
              />
              <p class="cm-font-009F72 item-name">单签钱包</p>
            </div>
            <div class="item-desc">
              单签钱包需要您创建密钥，密钥归您个人所有；单签钱包内资产由一方密钥持有支配
            </div>
          </div>
          <div
            class="wallet-ty-item"
            :class="walletType == 2 ? 'cur-tab' : ''"
            @click="walletType = 2"
          >
            <div class="item-image">
              <img
                src="../../assets/images/wallet/more_wt_icon.png"
                alt=""
                srcset=""
              />
              <p class="cm-font-009F72 item-name">多签钱包</p>
            </div>
            <div class="item-desc">
              创建多签钱包页面-文案更新：
              多签钱包需要关联多个已创建的密钥，多签钱包内的资产由多方共同持有和支配。
              温馨提示：如需进行多签签名，钱包内的每把公钥都需独立进行“创建多签钱包”的操作。
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
            @click="goNextStep"
            >下一步</a
          >
        </div>
      </div>
    </div>
    <!-- </div> -->
  </div>
</template>
<script>
export default {
  data() {
    return {
      active: 0,
      walletType: 1,
    };
  },
  methods: {
    goNextStep() {
      if (this.walletType < 1) {
        this.$message.error("请选择创建钱包的类型");
        return;
      }
      if (this.walletType == 1) {
        this.$router.push(`/wallet/walSingleCreate`);
      } else {
        this.$router.push(`/wallet/walMoreCreate`);
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
/* .layout-header {
  height: 90px !important;
  line-height: 90px;
} */
.wallet-warpper {
  /* width: 80%; */
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
.wallet-types {
  display: flex;
  justify-content: center;
  margin: 36px 0 56px;
}
.wallet-ty-item {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 348px;
  height: 300px;
  margin: 0 20px 0 0;
  border: 1px solid rgba(132, 132, 132, 1);
  cursor: pointer;
  transition: all 0.6s ease-in-out;
  border-radius: 5px;
}
.wallet-ty-item:hover {
  border-color: #62f7d4;
  transition: all 0.6s ease-in-out;
  background: rgba(50, 68, 83, 1);
}
.wallet-ty-item:last-child {
  margin-right: 0;
}

.item-image {
  text-align: center;
  margin: 0 0 20px;
}
.item-image img {
  width: 38px;
  height: 38px;
  margin: 0 auto 10px;
}
.item-name {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 16px;
  text-align: center;
}

.item-desc {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: #cbcbcc;
  width: 70%;
  line-height: 20px;
  height: 42%;
}
.cur-tab {
  border-color: #62f7d4;
  transition: all 0.6s ease-in-out;
  background: rgba(50, 68, 83, 1);
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
</style>