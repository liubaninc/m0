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
              src="../../assets/images/wallet/more_wt_icon.png"
              alt=""
              srcset=""
            />
            创建多签钱包
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
              <div class="form-rows-name">选择我的可用密钥</div>
              <div class="form-rows-desc">
                选择我的账户中已创建的密钥信息，若无可用密钥，可通过创建单签钱包创建密钥。
              </div>
              <div class="form-rows-input">
                <el-select v-model="publicKeyName" placeholder="请选择我的秘钥">
                  <el-option
                    v-for="item in keyLists"
                    :key="item.name"
                    :label="`${item.public_key}--${item.name}`"
                    :value="item.name"
                  >
                  </el-option>
                </el-select>
              </div>
            </div>
            <div class="form-rows-item">
              <div class="form-rows-name">添加其他密钥</div>
              <div class="form-rows-desc">
                添加共同签名的密钥账户（公钥）信息，用于创建多方签署的钱包
              </div>
              <div class="form-rows-input">
                <el-input
                  placeholder="请输入密钥账户"
                  v-model="otherPrivateKey"
                  clearable
                >
                </el-input>
              </div>
              <template v-for="(rowInput, index) in keysInput">
                <div class="form-add-row" :key="index">
                  <el-input
                    class="form-rows-input"
                    placeholder="请输入密钥账户"
                    v-model="rowInput.value"
                    clearable
                  >
                  </el-input>
                  <a href="javascript:;" class="del-btn" @click="delRow(index)"
                    >删除</a
                  >
                </div>
              </template>
              <a
                href="javascript:;"
                class="add-private-account"
                @click="addOtherKey"
                >添加密钥账户
              </a>
            </div>
            <div class="form-rows-item">
              <div class="form-rows-name">所需签名数</div>
              <div class="form-rows-desc">账户交易时，需要签名数量</div>
              <div class="form-rows-input">
                <el-input
                  placeholder="请输入签名数"
                  v-model="qmCount"
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
            @click="toCreateMoreSignWallet"
            >创建钱包</a
          >
        </div>
      </div>
    </div>
    <!-- </div> -->
  </div>
</template>
<script>
import { queryWalletLists, createMoreSignWallet } from '@/server/wallet'
import { localCache } from '@/utils/utils'
export default {
  data() {
    return {
      active: 1,
      keyLists: [],
      publicKeyName: '',
      keysInput: [],
      qmCount: '',
      otherPrivateKey: '',
      walletName: '',
    }
  },
  created() {
    this.getKeyLists()
  },
  destroyed() {
    clearTimeout(this.timer && this.timer)
  },
  methods: {
    toCreateMoreSignWallet() {
      let {
        walletName,
        publicKeyName,
        qmCount,
        keysInput,
        otherPrivateKey,
        keyLists,
      } = this

      if (!walletName) {
        this.$message.error('请输入钱包名称')
        return
      }
      if (!publicKeyName) {
        this.$message.error('请选择公钥')
        return
      }

      if (!otherPrivateKey) {
        this.$message.error('请输入添加其他密钥')
        return
      }

      if (!qmCount) {
        this.$message.error('请输入所需签名数')
        return
      }

      if (!(qmCount >= 1)) {
        return
      }
      let publicKey = keyLists.filter((list) => list.name == publicKeyName)[0][
        'public_key'
      ]
      let multis = keysInput.map(({ value }) => value) || []

      this.timer = setTimeout(async () => {
        this.loading = this.$loading({
          lock: true,
          text: '创建中...',
          spinner: 'el-icon-loading',
          background: 'rgba(0, 0, 0, 0.7)',
        })
        let mulWalletInfo = await createMoreSignWallet.call(this, {
          name: walletName,
          related: publicKeyName,
          multi_sig: [publicKey, otherPrivateKey, ...multis],
          threshold: qmCount * 1,
        })
        if (mulWalletInfo) {
          clearTimeout(this.timer)
          this.loading.close()
          localCache.set('walletInfo', mulWalletInfo)
          this.$router.push(`/wallet/walCreateSuccess?type=2`)
        } else {
          this.loading.close()
        }
      }, 300)
    },
    addOtherKey() {
      this.keysInput.push({ value: '' })
    },
    delRow(index) {
      this.keysInput.splice(index, 1)
    },
    async getKeyLists() {
      // let { accounts } = await queryWalletLists({
      //   page_size: 1,
      //   page_num: 1000,
      // })
      let { accounts } = await queryWalletLists({
        page_size: 10000,
        page_num: 1,
      })
      if (accounts) {
        this.keyLists = accounts.filter((account) => account.threshold < 1)
      }
    },
  },
}
</script>
<style>
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
  width: 94%;
}
.form-rows-input .el-select {
  width: 100%;
}
.form-add-row {
  display: flex;
  align-items: center;
}
.del-btn {
  width: 30px;
  display: block;
  margin: 0 0 0 10px;
  color: #ff5a58;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
}

.add-private-account {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: #2bc9a6;
  margin: 10px 0 0;
  display: block;
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
  padding: 0 45px 0 0;
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