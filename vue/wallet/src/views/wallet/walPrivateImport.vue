<template>
  <div class="wallet">
    <div class="cm-width1200">
      <div class="wallet-warpper">
        <el-breadcrumb separator="/" class="wallet-breadcrumb">
          <el-breadcrumb-item :to="{ path: '/wallet' }"
            >我的钱包</el-breadcrumb-item
          >
          <el-breadcrumb-item class="breadcrumb-cur-page"
            >导入钱包
          </el-breadcrumb-item>
        </el-breadcrumb>
        <div class="cm-module-bg wallet-create-main">
          <div class="wallet-steps">
            <el-steps :active="active" align-center finish-status="wait">
              <el-step title="选择导入类型"></el-step>
              <el-step title="导入钱包"></el-step>
              <el-step title="导入成功"></el-step>
            </el-steps>
          </div>
          <div class="wallet-form">
            <div class="wallet-form-title">
              <img
                src="../../assets/images/wallet/single_wt_icon.png"
                alt=""
                srcset=""
              />
              私钥导入钱包
            </div>
            <div class="wallet-form-rows">
              <div class="form-rows-item">
                <div class="form-rows-name">请设置钱包名称</div>
                <div class="form-rows-desc">
                  钱包名称是用来区分钱包的标签，该信息不会在区块链上保存
                </div>
                <div class="form-rows-input">
                  <el-input
                    placeholder="请输入内容"
                    v-model="walletName"
                    clearable
                  >
                  </el-input>
                </div>
              </div>
              <div class="form-rows-item">
                <div class="form-rows-name">请输入私钥(HEX)</div>
                <div class="form-rows-desc">请输入HEX格式的私钥文本</div>
                <div class="form-rows-input">
                  <el-input
                    placeholder="请输入内容"
                    v-model="privateKey"
                    clearable
                  >
                  </el-input>
                </div>
              </div>
              <div class="form-rows-item">
                <div class="form-rows-name">密钥算法</div>
                <div class="form-rows-input">
                  <a
                    v-for="(alg, index) in algorithms"
                    :key="alg.id"
                    href="javascript:;"
                    class="cm-btn-autopx wallet-btn-default"
                    :class="curAlg == index ? 'wallet-btn-cur' : ''"
                    @click="curAlg = index"
                    >{{ alg.text }}</a
                  >
                </div>
              </div>
              <div class="form-rows-item">
                <div class="form-rows-name">请设置密钥密码</div>
                <div class="form-rows-desc">
                  请设置至少10位字母数字混合的密码，密钥是根据你输入的密码生成的管理资产的加密凭证。请妥善保管、备份密码，忘记密码将导致钱包资产的损失。
                </div>
                <div class="form-rows-input">
                  <el-input
                    placeholder="请输入密码"
                    v-model="pwd"
                    clearable
                    show-password
                  >
                  </el-input>
                </div>
                <div class="form-rows-input rpsw-row">
                  <el-input
                    placeholder="请输入确认密码"
                    v-model="repwd"
                    clearable
                    show-password
                  >
                  </el-input>
                </div>
              </div>
            </div>
          </div>
          <div class="wallet-btns">
            <a
              href="javascript:;"
              class="cm-btn-200px cm-btn-border009F72"
              @click="$router.go(-1)"
              >返回</a
            >
            <a
              href="javascript:;"
              class="cm-btn-200px cm-btn-bg009F72"
              @click="importPrivateWallet"
              >导入钱包</a
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { createSingleAcc } from '@/server/account/account'
import { localCache } from '@/utils/utils'
export default {
  data() {
    return {
      algorithms: [
        {
          id: 10001,
          text: 'SECP256K1',
        },
        {
          id: 10002,
          text: 'SM2',
        },
      ],
      active: 1,
      input: '',
      walletName: '',
      privateKey: '',
      pwd: '',
      repwd: '',
      curAlg: 0,
    }
  },
  methods: {
    async importPrivateWallet() {
      let { walletName, pwd, repwd, privateKey, curAlg, algorithms } = this
      if (!walletName) {
        this.$message.error('请设置钱包名称')
        return
      }
      if (!privateKey) {
        this.$message.error('请输入私钥')
        return
      }
      if (curAlg < 0) {
        this.$message.error('请选择算法')
        return
      }
      if (!pwd) {
        this.$message.error('请设置密码')
        return
      }
      if (pwd != repwd) {
        this.$message.error('俩次输入的密码不一致')
        return
      }
      let walletInfo = await createSingleAcc.call(this, {
        name: walletName,
        password: pwd,
        algo: algorithms[curAlg].text,
        private_key: privateKey,
      })
      if (walletInfo) {
        localCache.set('walletInfo', walletInfo)
        this.$router.push(`/wallet/walCreateSuccess?name=importwal`)
      }
    },
  },
}
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
  padding: 20px 0 60px;
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
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
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
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  margin: 10px 0;
}
.form-rows-desc {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
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

.wallet-btns {
  display: flex;
  justify-content: center;
  width: 70%;
  margin: 0 auto;
}
.wallet-btns a {
  margin: 0 50px 0 0;
}
.wallet-btns a:last-child {
  margin-right: 0;
}
.wallet-btn-cur {
  border-color: #62f7d4;
  color: #62f7d4;
}

.rpsw-row {
  margin-top: 6px;
}
</style>