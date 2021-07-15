<template>
  <div class="receive">
    <el-breadcrumb separator="/" class="wallet-breadcrumb">
      <el-breadcrumb-item :to="{ path: '/assets' }"
        >资产管理</el-breadcrumb-item
      >
      <el-breadcrumb-item class="breadcrumb-cur-page"
        >接收转账
      </el-breadcrumb-item>
    </el-breadcrumb>
    <div class="cm-module-bg wallet-create-main">
      <div class="wallet-ct-success">
        <img
          src="../../assets/images/wallet/success_icon.png"
          alt=""
          srcset=""
        />
        接收转账
      </div>
      <div class="wallet-form wallet-form-mesg">
        <div class="wallet-form-title wallet-form-success">
          <img
            src="../../assets/images/wallet/wallet_info_icon.png"
            alt=""
            srcset=""
          />
          收款信息
        </div>
        <div class="wallet-form-info">
          <div class="form-info-row">
            <div class="info-row-name">接收地址</div>
            <div class="cm-text-overflow recive-address">{{ address }}</div>
            <a
              href="javascript:;"
              class="iconfont m0-copy info-copy-btn"
              v-clipboard:copy="address"
              v-clipboard:success="onCopy"
              v-clipboard:error="onError"
            ></a>
          </div>
          <div class="form-info-row">
            <div class="info-row-name">资产名称</div>
            <div class="cm-text-overflow info-row-address">{{ name }}</div>
          </div>
          <div class="form-info-row">
            <div class="info-row-name">收款二维码</div>
            <div class="info-row-address receive-qrcode" ref="receiveQRCode">
              <!-- <img
                src="https://d1icd6shlvmxi6.cloudfront.net/gsc/4RVIM8/16/bb/ad/16bbadf439ee459fb018cc7d27e3ac6f/images/接收转账/u1761.png?token=295218a431659f3b7116a1e5f1c9a2f18707fb69809646afecb6a9fd7c6052ee"
                alt=""
                srcset=""
              /> -->
            </div>
          </div>
        </div>
      </div>
      <div class="receive-btns">
        <a
          href="javascript:;"
          class="cm-btn-295px cm-btn-bg4acb9b"
          @click="$router.go(-1)"
          >返回资产列表</a
        >
      </div>
    </div>
  </div>
</template>
<script>
import QRCode from 'qrcodejs2'

export default {
  data() {
    return {
      name: '',
      address: '',
    }
  },
  created() {
    let { name, address } = this.$route.query
    this.name = name
    this.address = address
  },
  mounted() {
    this.$nextTick(() => {
      this.bindQRCode(this.address)
    })
  },
  methods: {
    bindQRCode(text) {
      new QRCode(this.$refs.receiveQRCode, {
        text,
        width: 200,
        height: 200,
        colorDark: '#333333',
        colorLight: '#ffffff',
        correctLevel: QRCode.CorrectLevel.L,
      })
    },
    onCopy(text) {
      if (text) {
        this.$message('复制成功')
      }
    },
    onError(e) {
      console.log(e)
    },
  },
}
</script>
<style>
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
  /* justify-content: space-between; */
  width: 70%;
  margin: 0 auto;
}
.wl-success-btn {
  padding: 0;
}

/** receive start*/
.receive-qrcode img {
  width: 210px;
  height: 200px;
}
.receive-btns {
  display: flex;
  justify-content: center;
}
/**end receive*/

.recive-address {
  width: 50%;
}
</style>