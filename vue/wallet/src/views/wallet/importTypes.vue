<template>
  <div class="wallet">
    <!-- <div class="cm-width1200"> -->
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
        <div class="wallet-types">
          <div
            class="wallet-ty-item"
            :class="walletType == 1 ? 'cur-tab' : ''"
            @click="walletType = 1"
          >
            <div class="item-image">
              <img
                src="../../assets/images/wallet/wallet_import_icon.png"
                alt=""
                srcset=""
              />
              <p class="cm-font-009F72 item-name">助记词导入</p>
            </div>
            <div class="item-desc">
              使用助记词导入钱包，助记词由单词组成，以空格隔开
            </div>
          </div>
          <div
            class="wallet-ty-item"
            :class="walletType == 2 ? 'cur-tab' : ''"
            @click="walletType = 2"
          >
            <div class="item-image">
              <img
                src="../../assets/images/wallet/private_key_icon.png"
                alt=""
                srcset=""
              />
              <p class="cm-font-009F72 item-name">私钥导入</p>
            </div>
            <div class="item-desc">使用私钥导入钱包，需输入明文私钥字符</div>
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
            @click="toImportWallet"
            >下一步</a
          >
        </div>
      </div>
      <!-- </div> -->
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      active: 0,
      walletType: 1,
    }
  },
  methods: {
    toImportWallet() {
      if (this.walletType < 1) return
      if (this.walletType == 1) {
        this.$router.push(`/wallet/walWordImport?type=1`)
      } else {
        this.$router.push(`/wallet/walPrivateImport?type=2`)
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
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 16px;
  text-align: center;
}

.item-desc {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: #cbcbcc;
  width: 70%;
  line-height: 20px;
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

.cur-tab {
  border-color: #62f7d4;
  transition: all 0.6s ease-in-out;
  background: rgba(50, 68, 83, 1);
}
</style>