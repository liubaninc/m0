<template>
  <div class="opearation">
    <el-breadcrumb separator="/" class="wallet-breadcrumb">
      <el-breadcrumb-item :to="{ path: '/mycontract' }"
        >合约管理</el-breadcrumb-item
      >
      <el-breadcrumb-item class="breadcrumb-cur-page"
        >多签交易签名
      </el-breadcrumb-item>
    </el-breadcrumb>
    <div class="cm-submodule-bg opearation-main">
      <div class="main-warpper">
        <span class="iconfont m0-chenggong success-icon"></span>
        <div class="main-wp-text">
          <template v-if="mode == 'deploy'">
            智能合约部署中，部署成功后会在合约列表内展示合约最新状态
          </template>
          <template v-else>
            智能合约升级中，升级成功后会在合约列表内展示合约最新状态
          </template>
        </div>
        <a
          href="javascript:;"
          class="cm-btn-bg4acb9b cm-btn-225px confirm-btn"
          @click="$router.push('/mycontract')"
          >确定</a
        >
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      mode: 'deploy',
    }
  },
  created() {
    let { mode } = this.$route.query
    if (mode) {
      this.mode = mode
    }
  },
}
</script>
<style scoped>
.opearation-main {
  height: calc(100vh - 157px);
  display: flex;
  align-items: center;
}
.confirm-btn {
  margin: 0 auto;
}
.main-warpper {
  width: 60%;
  margin: 0 auto;
  text-align: center;
}
.success-icon {
  color: #59f4d4;
  font-weight: 700;
  font-size: 40px;
}
.main-wp-text {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-size: 18px;
  color: #ffffff;
  text-align: center;
  margin: 20px 0 100px;
  line-height: 30px;
}
</style>