<template>
  <div class="publish sign">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/mycontract' }"
          >合约管理</el-breadcrumb-item
        >
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >多签交易签名
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg detail-main sigin-warpper">
        <div class="detail-main-assets">
          <div class="assets-title">签名文件</div>
          <div class="cm-submodule-bg transferout-publish">
            <div class="transferout-assets-info">
              <div class="assets-info-row transferout-info-row">
                <span class="row-name">签名文件 </span>
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
                    :on-error="uploadFalse"
                    :on-success="uploadSuccess"
                    accept="*"
                    :on-change="uploadChange"
                    :file-list="fileList"
                    :before-upload="uploadBefore"
                    :show-file-list="false"
                    :headers="headers()"
                    @click="clearOldVal"
                  >
                  </el-upload>
                  <a
                    href="javascript:;"
                    class="cm-btn-160px cm-btn-bg4acb9b upload-btn"
                  >
                    上传</a
                  >
                </div>
              </div>
            </div>
          </div>
        </div>

        <div v-if="Object.keys(trxInfo).length">
          <div class="detail-main-assets">
            <div class="assets-title">智能合约信息</div>
            <div class="cm-submodule-bg transferout-publish">
              <div class="transferout-assets-info">
                <div class="assets-info-row transferout-info-row">
                  <span class="row-name">合约名称 </span>
                  <div
                    class="
                      cm-row-bg
                      row-line
                      sign-row-line
                      assets-line
                      padd-left
                    "
                  >
                    {{ trxInfo.name }}
                  </div>
                </div>
                <div class="assets-info-row transferout-info-row">
                  <span class="row-name">合约操作 </span>
                  <div class="cm-row-bg row-line sign-row-line padd-left">
                    {{ trxInfo.mode | txtStatus }}
                  </div>
                </div>
              </div>
              <div class="transferout-assets-info send-person">
                <div class="assets-info-row transferout-info-row">
                  <span class="row-name">合约版本 </span>
                  <div
                    class="
                      cm-row-bg cm-text-overflow
                      sign-row-line
                      qm-file-input
                    "
                  >
                    {{ trxInfo.version }}
                  </div>
                </div>
                <div class="assets-info-row transferout-info-row">
                  <span class="row-name">合约参数 </span>
                  <div class="cm-row-bg row-line sign-row-line padd-left">
                    {{ trxInfo.args }}
                  </div>
                </div>
              </div>
              <div class="transferout-assets-info send-person">
                <span class="row-name">合约描述 </span>
                <div class="cm-row-bg contract-descript">
                  {{ trxInfo.description }}
                </div>
              </div>
              <div class="transferout-assets-info send-person">
                <div class="assets-info-row transferout-info-row">
                  <span class="row-name">发起人 </span>
                  <div
                    class="
                      cm-row-bg cm-text-overflow
                      sign-row-line
                      qm-file-input
                    "
                  >
                    {{ trxInfo.address }}
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="detail-main-assets">
            <div class="assets-title">交易签名</div>
            <div class="cm-submodule-bg assets-info">
              <div class="transfer-trx-type">
                <div class="assets-info-row">
                  <span class="row-name"> 当前钱包类型 </span>
                  <div class="row-line">
                    {{ wallet.threshold | walType }}
                  </div>
                </div>
                <div class="assets-info-row">
                  <span class="row-name qm-num-name">
                    已签名账户数/需签名数
                  </span>
                  <div class="row-line account-count">
                    {{ trxInfo.signatures.length }}/{{ wallet.threshold }}
                  </div>
                </div>
              </div>
              <div class="assets-info-row assets-info-txSigned">
                <span class="row-name">密钥密码 </span>
                <div class="assets-info-desc">
                  输入密钥密码进行签名，多签钱包需关联密钥账户进行签名，满足需签名数后才可向区块链提交交易
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

          <div class="singleevid-btns">
            <a
              href="javascript:;"
              class="cm-btn-295px cm-btn-border009F72 evidence-btn"
              @click="$router.go(-1)"
              >返回</a
            >
            <a
              href="javascript:;"
              class="cm-btn-295px cm-btn-bg4acb9b evidence-btn"
              @click="confirmSign"
              >确认签名</a
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { upLoadFile } from '@/server/upload'
import { queryContractTrxByHash, signContract } from '@/server/contract'
import { localCache } from '@/utils/utils'
import { tipStatus } from '@/utils/filters/status'

export default {
  data() {
    return {
      fileList: [],
      upfileName: '',
      pwd: '',
      trxInfo: {},
      wallet: {},
    }
  },
  created() {
    let wallet = localCache.get('wallet')
    this.wallet = wallet
  },
  methods: {
    clearOldVal() {
      let uploadFilesArr = this.$refs.upload.uploadFiles //上传文件列表
      if (uploadFilesArr.length == 0) {
      } else {
        this.$refs.upload.uploadFiles = []
      }
    },
    async confirmSign() {
      let { pwd, trxInfo, wallet } = this

      if (!pwd) {
        this.$message.error('请输入密码')
        return
      }
      this.loading = this.$loading({
        lock: true,
        text: `${tipStatus(trxInfo.mode)}`,
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)',
      })
      let commit =
        trxInfo.signatures.length == wallet.threshold - 1 ? true : false
      let resTrx = await signContract({
        commit,
        hash: trxInfo.hash,
        name: wallet.name,
        password: pwd,
      })
      if (resTrx) {
        this.loading.close()
        if (!commit) {
          this.$router.push(`/mycontract/deploy/moreDeploy?hash=${resTrx.hash}`)
          return
        }
        if (trxInfo.mode == 'deploy' || trxInfo.mode == 'upgrade') {
          this.$router.push(
            `/myContract/signature/opSuccess?mode=${trxInfo.mode}`
          )
        } else {
          // this.$router.push(
          //   `/myContract/deploy?id=${trxInfo.id}&mode=${trxInfo.mode}&hash=${resTrx.hash}`
          // )
          this.$router.push(`/mycontract`)
          return
        }
      } else {
        this.loading.close()
      }
    },
    actionUrl() {
      return `${process.env.VUE_APP_PRO_BASE_URL}/mcontract/tx/upload`
    },
    headers() {
      return {
        Authorization: localCache.get('authorization') || '',
      }
    },
    uploadBefore(file) {
      this.upfileName = file.name
    },
    uploadFalse(error) {
      this.$message.error(error)
    },
    uploadChange(file, fileList) {
      this.fileList.splice(0, 1, file)
    },
    uploadSuccess(res) {
      if (res.code == 200) {
        this.getTrxInfo(res.data.hash)
      } else {
        this.$message.error(res.msg)
      }
    },

    async getTrxInfo(hash) {
      let trxInfo = await queryContractTrxByHash({
        hash,
      })
      if (trxInfo) {
        this.trxInfo = trxInfo
      }
    },
  },
}
</script>
<style scoped>
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
.transferout-assets-info {
  display: flex;
  align-items: center;
  padding: 0px 0 0px 20px;
  border-radius: 5px;
}
.assets-info-row {
  display: flex;
  align-items: center;
  margin: 0 0 20px;
}
.transferout-info-row {
  margin-bottom: 0;
  margin-left: 20px;
  width: 80%;
}
.transferout-info-row:first-child {
  margin-left: 0;
}
.assets-info-row:last-child {
  margin-bottom: 0;
}
.publish .row-name {
  width: 100px;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
}
.publish .row-line {
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
/**transferout start */
.transfer-assets-name {
  width: 40%;
  height: 40px;
  line-height: 40px;
  padding: 0 0 0 20px;

  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  letter-spacing: normal;
  color: #ffffff;
}

.transferout .add-address-line {
  padding: 0 0 0 20px;
}
.transferout-publish {
  padding: 35px 0 30px 0px;
  border-radius: 5px;
}
/**end transferout  */
/***publish start */
.publish {
}
.publish-assets-desc {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #768ca8;
  padding: 0 0 16px;
}
.publish-assets-input {
  width: 60%;
}
.publish-info-row {
  padding: 0 0 0 20px;
}
.publish-del-btn {
  width: 100px;
  font-family: PingFangSC-Regular, 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: rgb(255, 90, 88);
  margin: 0 0 0 20px;
}

.transfer-trx-type {
  display: flex;
}
.transfer-trx-type .row-name {
}
.transfer-trx-type .assets-info-row {
  margin: 0;
}
.transferout .transfer-trx-type .row-line {
  width: auto;
}
.transfer-trx-type .row-line {
  padding: 0 20px;
}
.transfer-trx-type .account-count {
  color: #62f7d4;
}
.transfer-trx-type .qm-num-name {
  width: 230px;
}

/***end publish */
.send-person {
  margin: 10px 0 0;
}

.assets-info-target {
  margin: 0 0 10px;
}

.sigin-warpper {
  padding: 0 0 30px;
}
.assets-line {
  width: 300px !important;
}
.padd-left {
  padding-left: 10px !important;
}

.contract-descript {
  width: 80%;
  height: 125px;
  padding: 10px 20px;
  border-radius: 5px;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
}
</style>