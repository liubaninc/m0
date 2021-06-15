<template>
  <div class="sgdetail">
    <div class="sgdetail-title">钱包详情</div>
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
            <div>FAFA12</div>
          </div>
          <div class="form-info-row">
            <div class="info-row-name">钱包类型</div>
            <div>多签钱包</div>
          </div>
          <div class="form-info-row">
            <div class="info-row-name">钱包地址</div>
            <div class="cm-text-overflow info-row-address">
              0dfds3j5jj6jfjs7jjc3jwj2j5j6j7jvj7j8jbhj3
            </div>
            <a href="javascript:;" class="iconfont m0-copy info-copy-btn"></a>
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
          <div class="form-info-row">
            <div class="info-row-name">多签公钥</div>
            <div>0dfds3j5jj6jfjs7jjc3jwj2j5j6j7jvj7j8jbhj3</div>
          </div>
          <div class="form-info-row">
            <div class="info-row-name">参与公钥</div>
            <div class="info-row-address">
              <div class="cm-text-overflow">
                0dfds3j5jj6jfjs7jjc3jwj2j5j6j7jvj7j8jbhj3
              </div>
              <div class="cm-text-overflow">
                0dfds3j5jj6jfjs7jjc3jwj2j5j6j7jvj7j8jbhj3
              </div>
            </div>
          </div>
          <div class="form-info-row">
            <div class="info-row-name">签名数</div>
            <div>2</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      active: 2,
      input: "",
    };
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
  font-size: 16px;
  margin: 0 0 10px;
}
</style>