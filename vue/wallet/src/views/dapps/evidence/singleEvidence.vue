<template>
  <div class="singleevid">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/dapps/evidence' }"
          >存证管理</el-breadcrumb-item
        >
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
                <el-input
                  v-model="eviName"
                  placeholder="请输入存证名称"
                ></el-input>
              </div>
            </div>
            <div class="assets-info-row">
              <span class="row-name">存证文件 </span>
              <div class="row-line sign-row-line">
                <el-input
                  placeholder="上传签名文件"
                  v-model="upfileName"
                  :disabled="true"
                >
                </el-input>
                <el-upload
                  class="add-file-right-input"
                  ref="upload"
                  :action="actionUrl()"
                  :on-error="uploadError"
                  :on-success="uploadSuccess"
                  accept="*"
                  :on-change="uploadChange"
                  :file-list="fileList"
                  :show-file-list="false"
                  :headers="headers()"
                  @click="clearOldVal"
                >
                </el-upload>
                <a
                  href="javascript:;"
                  class="cm-btn-160px cm-btn-bg009F72 upload-btn"
                >
                  上传</a
                >
              </div>
              <!-- <div class="row-line">
                <el-input
                  v-model="input"
                  placeholder="请上传存证文件"
                ></el-input>
                <a
                  href="javascript::"
                  class="cm-btn-94px cm-btn-bg009F72 upload-btn"
                  >上传</a
                >
              </div> -->
            </div>
            <div class="assets-info-row">
              <span class="row-name">存证信息 </span>
              <div class="row-line row-textarea">
                <el-input
                  type="textarea"
                  :rows="3"
                  placeholder="请输入存证信息"
                  resize="none"
                  v-model="eviDesc"
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
                  placeholder="请输入备注信息"
                  resize="none"
                  v-model="eviBankup"
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
                <div class="">{{ wallet.threshold | walType }}钱包</div>
              </div>
              <template v-if="wallet.threshold > 0">
                <div>
                  已签名账户数/需签名数：
                  <span class="trx-count">0/{{ wallet.threshold }}</span>
                </div>
              </template>
            </div>
            <div class="assets-info-row assets-info-txSigned">
              <span class="row-name">密钥密码 </span>
              <div class="assets-info-desc">
                输入密钥密码进行签名，单签钱包，完成签名后即可向区块链提交交易
              </div>
              <div class="assets-info-input">
                <el-input
                  v-model="pwd"
                  placeholder="请输入密码"
                  show-password
                ></el-input>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="singleevid-btns">
        <a
          href="javascript:;"
          class="cm-btn-295px cm-btn-border009F72 evidence-btn"
          @click="$router.go(-1)"
          >返回</a
        >
        <a
          href="javascript:;"
          class="cm-btn-295px cm-btn-border009F72 evidence-btn"
          @click="saveUpload"
          >上传存证</a
        >
      </div>
    </div>
  </div>
</template>
<script>
import { saveEvidence } from "@/server/dapps/evidence";
import { localCache } from "@/utils/utils";

export default {
  data() {
    return {
      pwd: "",
      eviName: "",
      eviDesc: "",
      eviBankup: "",
      pwd: "",
      fileList: [],
      upfileName: "",
      wallet: {},
    };
  },
  created() {
    let wallet = localCache.get("wallet");
    if (wallet) {
      this.wallet = wallet;
    }
  },
  methods: {
    async saveUpload() {
      let { eviName, eviDesc, eviBankup, pwd, upfileName, wallet } = this;
      if (!eviName) {
        this.$message.error("请输入存证名称");
        return;
      }

      let reg = /^[\u4E00-\u9FA5a-zA-Z0-9_]{1,}$/;

      if (!reg.test(eviName)) {
        this.$message.error("存证名称只允许输入字母数字中文及_");
        return;
      }

      // if (!upfileName) {
      //   this.$message.error("请输入上传存证文件");
      //   return;
      // }

      if (!eviDesc) {
        this.$message.error("请输入上传存证信息");
        return;
      }
      // if (!eviBankup) {
      //   this.$message.error("请输入备注信息");
      //   return;
      // }

      if (!pwd) {
        this.$message.error("请输入密码");
        return;
      }
      let commit = wallet.threshold > 1 ? false : true;
      let saveRes = await saveEvidence.call(this, {
        account: wallet.name,
        commit,
        file: upfileName,
        info: eviDesc,
        memo: eviBankup,
        name: eviName,
        password: pwd,
      });
      if (saveRes) {
        this.$router.push(`/dapps/evidence/evidSuccess?name=${saveRes.name}`);
      }
    },
    clearOldVal() {
      let uploadFilesArr = this.$refs.upload.uploadFiles; //上传文件列表
      if (uploadFilesArr.length == 0) {
      } else {
        this.$refs.upload.uploadFiles = [];
      }
    },

    actionUrl() {
      return window.location.origin + `/api/claims/${this.wallet.name}/upload`;
    },
    headers() {
      return {
        Authorization: localCache.get("authorization") || "",
      };
    },
    uploadError(error) {
      this.$message.error(err);
    },
    uploadChange(file, fileList) {
      this.fileList.splice(0, 1, file);
    },
    uploadSuccess(res) {
      if (res.code == 200) {
        this.upfileName = res.data;
      } else if (res.code == 3002) {
        this.$message.error(res.msg);
      }
    },
  },
};
</script>
<style>
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
.singleevid .row-name {
  width: 100px;
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
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
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
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
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
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

.sign-row-line {
  padding: 0 0 0px 10px;
  border-radius: 5px;
  position: relative;
}

.add-file-right-input {
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0px;
  bottom: 0;
  right: 0;
  left: 0;
  z-index: 2;
  opacity: 0;
}
.add-file-right-input .el-upload {
  width: 100%;
  height: 100%;
}
.info-row-left {
  display: flex;
  margin: 0 36px 0 0;
  font-size: 14px;
}
.more-assets-row {
  font-size: 16px;
  margin: 0 36px 20px 0;
}
.trx-count {
  color: #62f7d4;
}
</style>