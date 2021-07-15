<template>
  <div class="wallet">
    <!-- <div class="cm-width1200"> -->
    <div class="wallet-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/' }">我的钱包</el-breadcrumb-item>
        <el-breadcrumb-item class="breadcrumb-cur-page">
          <template v-if="curSucessType == 'importwal'"> 导入钱包 </template>
          <template v-else> 创建钱包 </template>
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg wallet-create-main">
        <div class="wallet-steps">
          <template v-if="curSucessType == 'importwal'">
            <el-steps :active="active" align-center finish-status="wait">
              <el-step title="选择导入类型"></el-step>
              <el-step title="导入钱包"></el-step>
              <el-step title="导入成功"></el-step>
            </el-steps>
          </template>
          <template v-else>
            <el-steps :active="active" align-center finish-status="wait">
              <el-step title="选择钱包类型"></el-step>
              <el-step title="创建钱包"></el-step>
              <el-step title="创建成功"></el-step>
            </el-steps>
          </template>
        </div>
        <div class="wallet-ct-success">
          <img
            src="../../assets/images/wallet/success_icon.png"
            alt=""
            srcset=""
          />
          <template v-if="curSucessType == 'importwal'">
            钱包导入成功
          </template>
          <template v-else>
            <template v-if="!type"> 单签钱包创建成功 </template>
            <template v-else> 多签钱包创建成功 </template>
            <template v-if="walletInfo.threshold > 0">
              <div class="more-tips">
                温馨提示：如需进行多签签名，钱包内的每把公钥都需独立进行“创建多签钱包”的操作。
              </div>
            </template>
          </template>
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
              <div class="info-row-name">钱包地址</div>
              <div class="cm-text-overflow info-row-address">
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
        <div class="wallet-form">
          <div class="wallet-form-title wallet-form-success">
            <img
              src="../../assets/images/wallet/wallet_key_icon.png"
              alt=""
              srcset=""
            />
            密钥信息
          </div>
          <div class="wallet-form-info">
            <template v-if="!type">
              <div class="form-info-row">
                <div class="info-row-name">公钥</div>
                <div class="cm-text-overflow info-row-address">
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
              <div class="form-info-row">
                <div class="info-row-name">签名算法</div>
                <div>{{ walletInfo.algo }}</div>
              </div>
              <div class="form-info-row">
                <div class="info-row-name">私钥(HEX)</div>
                <div class="cm-text-overflow info-row-address">
                  {{ walletInfo.private_key }}
                </div>
                <a
                  href="javascript:;"
                  class="iconfont m0-copy info-copy-btn"
                  v-clipboard:copy="walletInfo.private_key"
                  v-clipboard:success="onCopy"
                  v-clipboard:error="onError"
                ></a>
              </div>
              <div class="form-info-row" v-if="walletInfo.mnemonic">
                <div class="info-row-name">私钥助记词</div>
                <div class="cm-text-overflow info-row-address">
                  {{ walletInfo.mnemonic }}
                </div>
                <a
                  href="javascript:;"
                  class="iconfont m0-copy info-copy-btn"
                  v-clipboard:copy="walletInfo.mnemonic"
                  v-clipboard:success="onCopy"
                  v-clipboard:error="onError"
                ></a>
              </div>
            </template>
            <template v-else>
              <div class="form-info-row">
                <div class="info-row-name">公钥</div>
                <div class="cm-text-overflow info-row-address">
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
              <div class="form-info-row">
                <div class="info-row-name">参与公钥</div>
                <div class="info-row-con">
                  <template v-for="(psin, index) in walletInfo.multi_sig">
                    <div class="row-multi-sig">
                      <div :key="index" class="cm-text-overflow multi-sig-text">
                        {{ psin }}
                      </div>
                      <a
                        href="javascript:;"
                        class="iconfont m0-copy info-copy-btn"
                        v-clipboard:copy="psin"
                        v-clipboard:success="onCopy"
                        v-clipboard:error="onError"
                      ></a>
                    </div>
                  </template>
                </div>
              </div>
              <div class="form-info-row">
                <div class="info-row-name">签名数</div>
                <div class="info-row-address">{{ walletInfo.threshold }}</div>
              </div>
            </template>
          </div>
        </div>
        <div class="wallet-btns wl-success-btn">
          <a
            href="javascript:;"
            class="cm-btn-295px cm-btn-border009F72 cm-btn-rewallet"
            @click="backWalletList"
            >返回钱包列表</a
          >
          <a
            href="javascript:;"
            class="cm-btn-295px cm-btn-bg4acb9b"
            @click="goWallet"
            >进入钱包</a
          >
        </div>
      </div>
    </div>
    <!-- </div> -->
  </div>
</template>
<script>
import { localCache } from '@/utils/utils'
export default {
  data() {
    return {
      active: 2,
      walletInfo: {},
      type: null,
      curSucessType: '',
    }
  },
  created() {
    let walletInfo = localCache.get('walletInfo')
    let { type, name } = this.$route.query
    if (walletInfo) {
      this.walletInfo = walletInfo
      this.type = type && type

      this.curSucessType = name
    }
  },
  mounted() {
    if (window.history && window.history.pushState) {
      history.pushState(null, null, document.URL)
      window.addEventListener('popstate', this.goBack, false)
    }
  },
  destroyed() {
    window.removeEventListener('popstate', this.goBack, false)
  },
  methods: {
    goBack() {
      history.pushState(null, null, document.URL)
    },
    backWalletList() {
      localCache.remove('walletInfo')
      this.$router.replace(`/wallet`)
    },
    goWallet() {
      localCache.remove('walletInfo')
      localCache.set('wallet', this.walletInfo)
      this.$router.replace(`/assets?address=${this.walletInfo.address}`)
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
.wallet {
  background-color: #1b2c42;
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
  width: 64%;
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
  justify-content: center;
}
.wl-success-btn {
  padding: 0;
}

.cm-btn-rewallet {
  margin: 0 20px 0 0;
}
.row-multi-sig {
  display: flex;
  align-items: center;
}
.multi-sig-text {
  width: 80%;
  line-height: 30px;
}
.info-row-con {
  width: 80%;
}
.more-tips {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: #cbcbcc;
  width: 70%;
  margin: 15px auto 0;
}
</style>