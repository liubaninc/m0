<template>
  <div class="detail">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/dapps/evidence' }"
          >存证管理</el-breadcrumb-item
        >
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >存证详情
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg detail-main">
        <div class="detail-main-assets">
          <div class="assets-title">资产信息</div>
          <div class="cm-submodule-bg assets-info">
            <div class="detail-assets-row">
              <span class="row-name">交易哈希 </span>
              <div class="cm-row-bg row-line">
                <span class="cm-text-overflow txt-row-line">{{
                  eviDetail.hash
                }}</span>
                <span
                  class="iconfont m0-copy copy-btn"
                  v-clipboard:copy="eviDetail.hash"
                  v-clipboard:success="onCopy"
                  v-clipboard:error="onError"
                ></span>
              </div>
            </div>
            <div class="detail-assets-row">
              <span class="row-name">存证名称 </span>
              <div class="cm-row-bg row-line">
                {{ eviDetail.name }}
              </div>
            </div>
            <div class="detail-assets-row">
              <span class="row-name">备注信息 </span>
              <div class="cm-row-bg row-line row-textarea">
                {{ eviDetail.memo }}
              </div>
            </div>
          </div>
        </div>

        <div class="detail-main-assets">
          <div class="assets-title">存证文件</div>
          <div class="cm-submodule-bg assets-info">
            <div class="detail-assets-row">
              <span class="row-name">文件名 </span>
              <div class="cm-row-bg row-line">
                <span class="cm-text-overflow txt-row-line">
                  {{ eviDetail.file || '---' }}</span
                >
              </div>
            </div>
            <div class="detail-assets-row">
              <span class="row-name">文件大小 </span>
              <div class="cm-row-bg row-line">
                {{ eviDetail.size | formateSize }}
              </div>
            </div>
            <div class="detail-assets-row">
              <span class="row-name"> </span>
              <div class="row-line row-line-btns">
                <a
                  href="javascript:;"
                  class="cm-btn-225px cm-btn-border009F72"
                  @click="downLoadSign(eviDetail)"
                  >下载链上文件</a
                >
                <a
                  href="javascript:;"
                  class="cm-btn-225px cm-btn-border009F72"
                  @click="showVerDialog(eviDetail)"
                  >校验我的文件</a
                >
              </div>
            </div>
          </div>
        </div>

        <div class="detail-main-assets">
          <div class="assets-title">存证数据</div>
          <div class="cm-submodule-bg assets-info">
            <div class="detail-assets-row">
              <span class="row-name">存证数据 </span>
              <div class="cm-row-bg row-line row-textarea">
                {{ eviDetail.info }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <a
      href="javascript:;"
      class="cm-btn-200px cm-btn-bg009F72 dl-reback-btn"
      @click="$router.go(-1)"
      >返回</a
    >
    <!--vaild dialog start-->
    <div
      class="cm-warpper-bg vaild-bg"
      @click="isShowVerDia = false"
      v-if="isShowVerDia"
    >
      <div class="vaild-dialog" @click.stop.prevent="">
        <div class="dialog-title">校验存证文件</div>
        <div class="dialog-row">
          <span class="dialog-row-title">上传文件 </span>
          <!-- <div class="dialog-row-input">
            <el-input v-model="input" placeholder="请输入内容"></el-input>
          </div>
          <a
            href="javascript:;"
            class="cm-btn-94px cm-btn-bg009F72 dialog-upload-btn"
            >上传</a
          > -->
          <div class="dialog-row-input" @click.stop="">
            <el-input
              placeholder="上传签名文件"
              v-model="upfileName"
              :disabled="true"
            >
            </el-input>
            <el-upload
              class="add-file-right-input"
              ref="upload"
              :action="actionUrl"
              :on-error="uploadFalse"
              :on-success="uploadSuccess"
              accept="*"
              :on-change="uploadChange"
              :file-list="fileList"
              :show-file-list="false"
              :headers="headers"
              :data="uploadParams"
              @click.stop.prevent="clearOldVal"
            >
            </el-upload>
            <a
              href="javascript:;"
              class="cm-btn-160px cm-btn-bg009F72 sigin-upload-btn"
            >
              上传</a
            >
          </div>
        </div>
        <div class="dialog-row">
          <span class="dialog-row-title"> </span>
          <a
            href="javascript:;"
            class="cm-btn-225px cm-btn-border009F72"
            @click.stop.prevent="isShowVerDia = false"
            >取消</a
          >
          <a
            href="javascript:;"
            class="cm-btn-225px cm-btn-border009F72 immediate-btn"
            @click.stop.prevent="verificateFile"
            >立即校验</a
          >
        </div>
      </div>
    </div>
    <!--end vaild dialog-->
    <!--vaild-success start-->
    <div class="cm-warpper-bg vaild-success" v-if="isShowSuccess">
      <div class="success-dialog">
        <img src="../../../assets/images/wallet/success_icon.png" alt="" />
        <p class="success-dia-desc">校验成功，校验文件为存证文件</p>
        <a
          href="javascript:;"
          class="cm-btn-200px cm-btn-bg009F72 confirm-btn"
          @click.stop.prevent="confirmDia"
          >确定</a
        >
      </div>
    </div>
    <div class="cm-warpper-bg vaild-success" v-if="isShowError">
      <div class="success-dialog">
        <img src="../../../assets/images/wallet/error_icon.png" alt="" />
        <p class="success-dia-desc">校验失败，校验文件非存证文件</p>
        <a
          href="javascript:;"
          class="cm-btn-200px cm-btn-bg009F72 confirm-btn"
          @click.stop.prevent="confirmDia"
          >确定</a
        >
      </div>
    </div>
    <!--end vaild-success-->
  </div>
</template>
<script>
import { queryEvidenceDetail, verifyFile } from '@/server/dapps/evidence'
import { localCache } from '@/utils/utils'

export default {
  data() {
    return {
      wallet: {},
      loginUser: {},
      eviDetail: {},
      isShowVerDia: false,
      upfileName: '',
      fileList: [],
      isShowSuccess: false,
      isShowError: false,
      curFileMd5: '',
    }
  },
  computed: {
    uploadParams() {
      return {
        verify: true,
      }
    },
    actionUrl() {
      return window.location.origin + `/api/claims/${this.wallet.name}/upload`
    },
    headers() {
      return {
        Authorization: localCache.get('authorization') || '',
      }
    },
  },
  created() {
    let { name } = this.$route.query
    let wallet = localCache.get('wallet')
    let loginUser = localCache.get('loginUser')
    if (wallet && name) {
      this.wallet = wallet
      this.loginUser = loginUser
      this.getEviDetail(wallet.name, name)
    }
  },
  methods: {
    confirmDia() {
      this.isShowVerDia = false
      this.isShowSuccess = false
      this.isShowError = false
      this.$router.go(-1)
    },
    clearOldVal() {
      let uploadFilesArr = this.$refs.upload.uploadFiles //上传文件列表
      if (uploadFilesArr.length == 0) {
      } else {
        this.$refs.upload.uploadFiles = []
      }
    },
    uploadFalse(error) {
      this.$message.error(error)
    },
    uploadChange(file, fileList) {
      this.fileList.splice(0, 1, file)
    },
    uploadSuccess(res) {
      if (res.code == 200) {
        let { data } = res
        if (data) {
          this.upfileName = data.file
          this.curFileMd5 = data.md5
        }
      } else if (res.code == 3002) {
        this.$message.error(res.msg)
      }
    },
    async getEviDetail(account, name) {
      let eviDetail = await queryEvidenceDetail({
        account,
        name,
      })
      if (eviDetail) {
        this.eviDetail = eviDetail
      }
    },
    downLoadSign(ceviDetail) {
      if (!ceviDetail.file) {
        this.$message.error('暂无存证文件,无法下载')
        return
      }
      let { wallet, eviDetail, loginUser } = this
      if (wallet && eviDetail && loginUser) {
        let origin = window.location.origin
        let elink = document.createElement('a')
        elink.download = wallet.address
        elink.style.display = 'none'
        elink.href = `${process.env.VUE_APP_PRO_BASE_URL}/claims/download/${loginUser.name}/${wallet.name}/${eviDetail.file}`
        document.body.appendChild(elink)
        elink.click()
        document.body.removeChild(elink)
      }
    },
    showVerDialog(eviDetail) {
      if (!eviDetail.file) {
        this.$message.error('暂无存证文件,无法校验')
        return
      }
      this.isShowVerDia = true
    },
    async verificateFile() {
      let { upfileName, wallet, eviDetail, curFileMd5 } = this
      if (!upfileName) {
        this.$message.error('请上传认证文件')
        return
      }

      let verInfo = await verifyFile({
        account: wallet.name,
        name: eviDetail.name,
        md5: curFileMd5,
      })
      if (verInfo) {
        this.isShowSuccess = true
      } else {
        this.isShowError = true
      }
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
  /* min-height: 100vh; */
}
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
.detail-assets-row {
  display: flex;
  align-items: center;
  margin: 0 0 20px;
}
.detail-assets-row:last-child {
  margin-bottom: 0;
}
.row-name {
  width: 80px;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
}
.row-line {
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
.dl-reback-btn {
  /* margin: 30px 0 0 20%; */
  margin: 30px auto;
}
/*end detail* */

/**vaild-dialog start */

.vaild-dialog {
  width: 633px;
  background-color: #425263;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 50px 60px;
  border-radius: 5px;
}
.dialog-title {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
  text-align: center;
  margin: 0 0 50px;
}
.dialog-row {
  display: flex;
  height: 40px;
  line-height: 40px;
  margin: 0 0 40px;
}
.dialog-row:last-child {
  margin-bottom: 0;
}
.dialog-row-input {
  /* width: 338px; */
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  padding: 10px 0px;
  position: relative;
}
.dialog-row-title {
  display: inline-block;
  width: 80px;
  height: 100%;
  font-size: 14px;
}
.dialog-upload-btn {
  margin: 0 0 0 20px;
}
.immediate-btn {
  margin: 0 0 0 20px;
}
/**end vaild-dialog */

/**vaild-success start */
.vaild-success {
}
.success-dialog {
  width: 633px;
  background-color: #425263;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 40px 20px 60px;
}
.success-dialog img {
  width: 54px;
  height: 54px;
}
.success-dia-desc {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 16px;
  color: #ffffff;
  margin: 40px 0 50px;
}
.confirm-btn {
  font-size: 16px;
}
/**end vaild-success */
.txt-row-line {
  width: 90%;
  display: inline-block;
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

.sigin-upload-btn {
  margin: 0 0 0 20px;
}
</style>