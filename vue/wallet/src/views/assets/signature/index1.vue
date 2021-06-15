<template>
  <div class="destory">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/' }">资产管理</el-breadcrumb-item>
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >多签交易签名
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg detail-main">
        <div class="detail-main-assets">
          <div class="assets-title">签名文件</div>
          <div class="cm-submodule-bg assets-info">
            <div class="assets-info-row">
              <span class="row-name">签名文件 </span>
              <div class="row-line">
                <el-input
                  v-model="input"
                  placeholder="请上传签名文件"
                  :disabled="true"
                ></el-input>
              </div>
              <a href="javascript:;" class="cm-btn-160px cm-btn-bg009F72"
                >上传</a
              >
            </div>
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
.detail-warpper {
  margin: 0 auto;
  color: #fff;
}

/*detail start* */
.assets-title {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  height: 40px;
  line-height: 40px;
}
.detail-main-assets {
  padding: 20px 40px;
}
.assets-info {
  padding: 35px 0 30px 20px;
  border-radius: 5px;
}
.assets-info-row {
  display: flex;
  align-items: center;
  margin: 0 0 20px;
}
.assets-info-row:last-child {
  margin-bottom: 0;
}
.row-name {
  width: 80px;
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
}
.destory .row-name {
  width: 94px;
}
.row-line {
  min-height: 45px;
  /* line-height: 45px; */
  width: 60%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  padding: 10px 20px;
}
.row-textarea {
  line-height: 20px;
}
.copy-btn {
  width: 16px;
  height: 16px;
  display: inline-block;
  /* background: url("../../../assets/images/wallet/copy_icon.png"); */
  background-size: 100%;
  cursor: pointer;
  color: #22ac95;
}
.row-line-btns {
  padding: 0;
}
.reback-btn {
  margin: 30px auto;
}
/*end detail* */

/**destory start */
.destory-assets-pwdtitle {
  margin: 0 0 18px;
}
.row-desc {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #768ca8;
  margin: 8px 0 0;
}
.destory-assets-input {
  width: 76%;
}

.destory-btns {
  display: flex;
}
.destory-info-row {
  margin: 0 0 0;
}
.row-line-input {
  padding: 0;
}
/**end destory */
</style>