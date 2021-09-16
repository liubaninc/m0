<template>
  <div class="singleevid">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/' }">存证管理</el-breadcrumb-item>
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >上传存证
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg detail-main">
        <div class="detail-main-assets">
          <div class="assets-title">存证信息</div>
          <div class="cm-submodule-bg assets-info">
            <div class="assets-info-row">
              <span class="row-name">存证名称 </span>
              <div class="row-line">
                <el-input v-model="input" placeholder="请输入内容"></el-input>
              </div>
            </div>
            <div class="assets-info-row">
              <span class="row-name">存证文件 </span>
              <div class="row-line">
                <el-input v-model="input" placeholder="请输入内容"></el-input>
                <a
                  href="javascript::"
                  class="cm-btn-94px cm-btn-bg009F72 upload-btn"
                  >上传</a
                >
              </div>
            </div>
            <div class="assets-info-row">
              <span class="row-name">存证信息 </span>
              <div class="row-line row-textarea">
                <el-input
                  type="textarea"
                  :rows="3"
                  placeholder="请输入内容"
                  resize="none"
                  v-model="textarea"
                  maxlength="50"
                  show-word-limit
                >
                </el-input>
              </div>
            </div>

            <div class="assets-info-row">
              <span class="row-name">备注信息 </span>
              <div class="row-line row-textarea">
                <el-input
                  type="textarea"
                  :rows="3"
                  placeholder="请输入内容"
                  resize="none"
                  v-model="textarea"
                  maxlength="50"
                  show-word-limit
                >
                </el-input>
              </div>
            </div>
          </div>
        </div>

        <div class="detail-main-assets">
          <div class="assets-title">交易签名</div>
          <div class="cm-submodule-bg assets-info">
            <div class="assets-info-row more-assets-row">
              <div class="info-row-left">
                <span class="row-name"> 当前钱包类型 </span>
                <div class="">单签钱包</div>
              </div>
              <div>
                已签名账户数/需签名数： <span class="trx-count">0/3</span>
              </div>
            </div>
            <div class="assets-info-row assets-info-txSigned">
              <span class="row-name">密钥密码 </span>
              <div class="assets-info-desc">
                输入密钥密码进行签名，单签钱包，完成签名后即可向区块链提交交易
              </div>
              <div class="assets-info-input">
                <el-input v-model="input" placeholder="请输入内容"></el-input>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="singleevid-btns">
        <a
          href="javascript:;"
          class="cm-btn-295px cm-btn-border009F72 evidence-btn"
          >返回</a
        >
        <a
          href="javascript:;"
          class="cm-btn-295px cm-btn-border009F72 evidence-btn"
          >上传存证</a
        >
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      input: '',
      textarea: '',
    }
  },
}
</script>
<style>
.detail-warpper {
  margin: 0 auto;
  color: #fff;
}

/*detail start* */
.assets-title {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
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
  /* align-items: center; */
  margin: 0 0 20px;
}
.assets-info-row:last-child {
  margin-bottom: 0;
}
.singleevid .row-name {
  width: 100px;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
}
.singleevid .row-line {
  min-height: 45px;
  /* line-height: 45px; */
  width: 60%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  padding: 10px 0px;
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
/*end detail* */

.upload-btn {
  margin: 0 0 0 24px;
}
.assets-info-txSigned {
  flex-direction: column;
  align-items: flex-start;
}
.assets-info-desc {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #768ca8;
  line-height: 40px;
}
.assets-info-input {
  width: 60%;
  margin: 2px 0 0;
}

.singleevid-btns {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 30px 0;
}
.evidence-btn {
  font-size: 14px;
  margin: 0 0 0 20px;
}
.evidence-btn:first-child {
  margin-left: 0;
}

.info-row-left {
  display: flex;
  margin: 0 36px 0 0;
}
.more-assets-row {
  font-size: 16px;
  margin: 0 36px 20px 0;
}
.trx-count {
  color: #62f7d4;
}
</style>