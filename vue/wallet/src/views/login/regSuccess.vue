<template>
  <div class="regsuccess">
    <el-header class="layout-header">
      <Header />
    </el-header>
    <div class="cm-width1200 cm-module-bg reg-container">
      <div class="reg-container-top">
        <img src="../../assets/images/login/success_icon.png" />
        <div>恭喜您，注册成功</div>
      </div>
      <div class="reg-container-steps">
        <div class="steps-title">快速使用引导</div>
        <div class="steps-items">
          <div class="steps-card">
            <div class="card-border"></div>
            <div class="card-content">
              <img src="../../assets/images/login/lg_wallet.png" />
              <div class="content-name">创建钱包</div>
            </div>
          </div>
          <div class="steps-split"></div>
          <div class="steps-card">
            <div class="card-border"></div>
            <div class="card-content">
              <img src="../../assets/images/login/asset_create.png" />
              <div class="content-name">资产创建</div>
            </div>
          </div>
          <div class="steps-split"></div>
          <div class="steps-card">
            <div class="card-border"></div>
            <div class="card-content">
              <img src="../../assets/images/login/trx_mg.png" />
              <div class="content-name">交易管理</div>
            </div>
          </div>
          <div class="steps-split"></div>
          <div class="steps-card">
            <div class="card-border"></div>
            <div class="card-content">
              <img src="../../assets/images/login/apply_mg.png" />
              <div class="content-name">应用服务</div>
            </div>
          </div>
        </div>
      </div>
      <div class="reg-container-btns">
        <a
          href="javascript:;"
          class="cm-btn-426px cm-btn-bg4acb9b reg-btn"
          @click="$router.push('/wallet/createTypes')"
          >立即创建钱包</a
        >
        <a
          href="javascript:;"
          class="cm-font-009F72 reg-enter-all"
          @click="$router.push('/wallet')"
          >暂不创建，进入总览</a
        >
      </div>
    </div>
  </div>
</template>
<script>
import Header from "@/components/header/header";
export default {
  components: {
    Header,
  },
};
</script>
<style scoped>
.regsuccess {
  background-color: #1b2c42;
  height: 100vh;
}
.reg-container {
  padding: 70px 0 58px;
  margin: 20px auto 0;
}
/* .layout-header {
  height: 90px !important;
  line-height: 90px;
} */
.reg-container-top {
  text-align: center;
  font-size: 18px;
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
}
.reg-container-top img {
  width: 62px;
  height: 62px;
  margin: 0 0 20px;
}
.reg-container-steps {
  padding: 70px 0 100px;
}
.steps-title {
  text-align: center;
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #cccccc;
  margin: 0 0 56px;
}
.steps-items {
  display: flex;
  align-items: center;
  justify-content: center;
}
.steps-card {
  width: 170px;
  height: 152px;
  background-color: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  position: relative;
  border-radius: 5px;
  overflow: hidden;
}
.card-border {
  width: 100%;
  height: 5px;
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 16px;
  color: #cccccc;
  text-align: left;
  background-color: #000;
  position: absolute;
  top: 0;
}
.card-content {
  text-align: center;
}
.card-content img {
  width: 55px;
  height: 55px;
  margin: 0 auto 20px;
}
.content-name {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 20px;
  color: #425263;
  line-height: 16px;
}
.steps-split {
  width: 32px;
  height: 32px;
  background-image: url("../../assets/images/login/reg_more.png");
  background-size: 100%;
  background-repeat: no-repeat;
  margin: 0 18px;
}
.reg-container-btns {
  text-align: center;
}
.reg-btn {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  margin: 0 auto 40px;
}
.reg-enter-all {
}
</style>