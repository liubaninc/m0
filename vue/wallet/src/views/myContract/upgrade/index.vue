<template>
  <div class="create">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/myContract' }"
          >我的合约</el-breadcrumb-item
        >
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >创建智能合约
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg create-main">
        <div class="contract-info">
          <div class="info-title">合约信息</div>
          <div class="cm-submodule-bg info-col">
            <div class="info-col-row2">
              <div class="col-row2-item">
                <div class="row2-item-name">合约名称</div>
                <div class="row2-item-text">
                  {{ name }}
                  <!-- <el-input
                    v-model="name"
                    placeholder="请输入合约名称"
                    aria-readonly
                    disabled
                  ></el-input> -->
                </div>
              </div>
              <div class="col-row2-item contract-version">
                <div class="row2-item-name">合约版本</div>
                <div class="row2-item-text">
                  <el-input
                    v-model="version"
                    placeholder="请输入数字及'.',如1.0.0"
                  ></el-input>
                </div>
              </div>
            </div>
            <div class="info-col-row">
              <div class="row2-item-name">合约描述</div>
              <div class="row-item-text">
                <el-input
                  type="textarea"
                  placeholder="请输入合约描述(选填)"
                  v-model="description"
                  resize="none"
                  maxlength="50"
                  show-word-limit
                >
                </el-input>
              </div>
            </div>
          </div>
        </div>
        <div class="contract-info">
          <div class="info-title">合约文件</div>
          <div class="cm-submodule-bg info-col">
            <div class="info-col-tab">
              <div class="col-title">选择生成方式</div>
              <div class="col-tabs">
                <template v-for="type in contractGenTypes">
                  <a
                    href="javascript:;"
                    :class="genWay == type.id ? 'cur-tab' : ''"
                    @click="genWay = type.id"
                    >{{ type.text }}</a
                  >
                </template>
              </div>
            </div>
            <div class="info-col-pannel">
              <div class="col-upload">
                <div class="contract-file">
                  <div class="upload-name">上传合约</div>
                  <div class="upload-desc">上传格式为.wasm，不超过5M</div>
                  <div class="upload-input">
                    <el-input
                      type="text"
                      placeholder="选择上传合约"
                      v-model="contractFile.name"
                      readonly
                      disabled
                    >
                    </el-input>
                    <a
                      href="javascript:;"
                      class="cm-btn-bg4acb9b cm-btn-94px upload-btn"
                      >上传</a
                    >
                    <el-upload
                      class="el-upload-mask"
                      action="*"
                      :auto-upload="false"
                      :show-file-list="false"
                      :on-change="uploadContract"
                      accept=".wasm"
                    >
                    </el-upload>
                  </div>
                </div>
                <div class="contract-file">
                  <div class="upload-name">合约参数</div>
                  <div class="upload-input contract-params">
                    <el-input
                      type="text"
                      v-model="args"
                      placeholder="请输入合约参数(选填)"
                    >
                    </el-input>
                  </div>
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
                <div class="row-line">{{ wallet.threshold | walType }}</div>
              </div>
              <template v-if="wallet.threshold > 0">
                <div class="assets-info-row">
                  <span class="row-name qm-num-name">
                    已签名账户数/需签名数
                  </span>
                  <div class="row-line account-count">
                    0/{{ wallet.threshold }}
                  </div>
                </div>
              </template>
            </div>
            <div class="assets-info-row assets-info-txSigned">
              <span class="row-name">密钥密码 </span>
              <div class="assets-info-desc">
                <template v-if="wallet.threshold > 0">
                  输入密钥密码进行签名，多签钱包需关联密钥账户进行签名，满足需签名数后才可向区块链提交交易
                </template>
                <template v-else>
                  输入密钥密码进行签名，单签钱包，完成签名后即可向区块链提交交易
                </template>
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
            class="cm-btn-295px cm-btn-border009F72"
            @click="$router.go(-1)"
            >返回</a
          >
          <a
            href="javascript:;"
            class="cm-btn-295px cm-btn-bg4acb9b create-contract-btn"
            @click="toCreateContract"
            >立即升级部署</a
          >
        </div>
      </div>
    </div>
    <div class="dialog" v-if="isCtractFnished">
      <div class="dialog-main del-dialog">
        <div class="dialog-main-title">
          <span class="main-title-name">创建智能合约</span>
          <!-- <a href="javascript:;" class="iconfont m0-guanbi"></a> -->
        </div>
        <div class="dialog-main-content">
          <div class="dialog-main-desc">合约创建成功,是否立即部署?</div>
          <div class="dialog-main-btns">
            <a
              href="javascript:;"
              class="cm-btn-138px cm-btn-border009F72"
              @click="$router.push(`/myContract`)"
              >返回合约列表</a
            >
            <a href="javascript:;" class="cm-btn-138px cm-btn-bg4acb9b"
              >立即部署合约</a
            >
          </div>
        </div>
      </div>
      <div class="dialog-yy"></div>
    </div>
  </div>
</template>
<script>
import {
  createContract,
  queryContractDetail,
  queryContractTrxByHash,
  publishContract,
} from '@/server/contract'

import { localCache, compareVersion } from '@/utils/utils'

export default {
  data() {
    return {
      name: '',
      version: '',
      description: '',
      contractFile: {},
      args: '',
      contractGenTypes: [{ id: 1, text: '上传合约' }],
      genWay: 1,
      oldVersion: '',
      wallet: {},
      pwd: '',
      isCtractFnished: false,
      commit: true,
    }
  },
  created() {
    let wallet = localCache.get('wallet')
    let { id } = this.$route.query
    if (wallet) {
      this.wallet = wallet
      this.getContractDetail(id)
    }
  },
  computed: {
    headerCell() {
      return {
        background: 'rgba(118, 140, 168, 1)!important',
        fontSize: '12px!important',
        lineHeight: '40px',
        padding: '0px',
      }
    },
  },
  methods: {
    async toPublishContract(id, pwd, commit) {
      let rePublish = await publishContract({
        id,
        mode: 'upgrade',
        password: pwd,
        commit,
        accountName: this.wallet.name,
      })
      if (rePublish) {
        this.loading.close()
        if (commit) {
          this.$router.push(`/myContract`)
        } else {
          this.$router.push(
            `/mycontract/deploy/moreDeploy?hash=${rePublish.hash}`
          )
        }
      } else {
        this.loading.close()
      }
      // this.$router.replace(
      //   `/myContract/deploy?id=${this.contractDetail.id}&mode=upgrade`
      // )
    },
    async toCreateContract() {
      let {
        name,
        version,
        oldVersion,
        description,
        contractFile,
        args,
        wallet,
        pwd,
      } = this

      if (!/^[a-zA-Z_]{1}[0-9a-zA-Z_.]+[0-9a-zA-Z_]/.test(name)) {
        this.$message.error(`合约名称只能以大小写字母开头且名称长度大于2个字符`)
        return
      }
      if (!/^([1-9]\d|[1-9])(.([1-9]\d|\d)){2,}$/.test(version)) {
        this.$message.error(`合约版本号只支持数字和.且至少三位数`)
        return
      }
      let versionMark = compareVersion(version, oldVersion)
      if (versionMark < 1) {
        this.$message.error(`新版本号必须高于老版本号`)
        return
      }

      if (!(contractFile && contractFile.raw)) {
        this.$message.error(`请上传合约文件`)
        return
      }

      if (!pwd) {
        this.$message.error(`请输入密码`)
        return
      }

      this.loading = this.$loading({
        lock: true,
        text: '合约升级中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)',
      })

      let resInfo = await createContract({
        account_name: wallet.address,
        name,
        version,
        type: 1,
        description,
        args,
        file: contractFile.raw || '',
        mode: 'upgrade',
      })
      if (resInfo) {
        // this.loading.close()
        if (wallet.threshold > 1) {
          // if (wallet.multi_sig.length < wallet.threshold - 1) {
          //   this.commit = false
          // } else {
          //   this.commit = true
          // }
          this.commit = false
        }
        this.toPublishContract(resInfo.id, pwd, this.commit)
        // this.isCtractFnished = true
        this.contractDetail = resInfo
      } else {
        this.loading.close()
        // this.isCtractFnished = false
      }
    },
    uploadContract(file) {
      if (file) {
        const isLt5M = file.size / 1024 / 1024 < 5
        if (!isLt5M) {
          this.$message.error('合约文件不能超过 5MB')
          return
        }
        this.contractFile = file
      }
    },
    async getContractDetail(id) {
      let contractInfo = await queryContractDetail({ id })
      if (contractInfo) {
        this.name = contractInfo.name
        this.version = contractInfo.version
        this.oldVersion = contractInfo.version
      }
    },
  },
}
</script>
<style scoped>
.create-main {
  padding: 0 0 30px;
}
.contract-info {
  padding: 30px 40px 0;
}
.info-col {
  padding: 35px 20px 30px 20px;
  border-radius: 5px;
}
.info-title {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-size: 14px;
  color: #ffffff;
  margin: 0 0 10px;
}
.info-col-row2 {
  display: flex;
  margin: 0 0 20px;
}
.col-row2-item {
  display: flex;
  align-items: center;
  justify-items: center;
  width: 40%;
}
.row2-item-text {
  flex: 1;
}
.row2-item-name {
  width: 100px;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-size: 14px;
  color: #ffffff;
}
.contract-version {
  margin: 0 0 0 6%;
}
.info-col-row {
  display: flex;
  width: 86%;
  align-items: center;
}
.row-item-text {
  flex: 1;
}

.col-title {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-size: 14px;
  color: #ffffff;
}
.col-tabs {
  display: flex;
  background: #768ca8;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-size: 14px;
  width: 200px;
}
.col-tabs a {
  display: block;
  width: 200px;
  height: 36px;
  text-align: center;
  line-height: 36px;
  border: 1px solid rgb(158 168 179);
}
.col-tabs a:last-child {
  border-left: 0;
}
.col-tabs a:first-child {
  border-right: 0;
}
.col-title {
  margin: 0 0 20px;
}

.col-upload {
  padding: 18px 0 0;
}
.contract-file {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-size: 14px;
  margin: 0 0 20px;
}
.upload-name {
  color: #ffffff;
}
.upload-desc {
  color: #768ca8;
  margin: 10px 0;
}
.upload-input {
  display: flex;
  width: 90%;
  position: relative;
}
.upload-btn {
  margin: 0 0 0 20px;
}
.contract-params {
  margin: 18px 0 0;
}
.col-template {
  padding: 18px 0 0;
}

.contract-template .upload-name {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-size: 14px;
  margin: 0 0 20px;
  color: #ffffff;
}
.opt-detail {
  color: #4ecc9b;
}

.singleevid-btns {
  display: flex;
  justify-content: center;
  margin: 18px 0 0;
}
.create-contract-btn {
  margin: 0 0 0 26px;
}

.el-upload-mask {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 10;
  width: 100%;
  height: 100%;
}

.del-dialog {
  display: flex;
  flex-direction: column;
}
.dialog-main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-content: center;
  align-items: center;
}
.dialog-main-desc {
  flex: 1;
  min-height: 190px;
  display: flex;
  align-items: center;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-size: 14px;
  color: #ffffff;
}
.dialog-main-btns {
  display: flex;
}
.dialog-main-btns a {
  margin: 0 0 0 20px;
  font-size: 14px;
}
.dialog-main-btns a:first-child {
  margin-left: 0;
}

.radio-radius {
  cursor: pointer;
}

.tmp-dialog .dialog-main-desc {
  padding: 20px 14px;
  flex: 1;
  min-height: 190px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-size: 14px;
  color: #ffffff;
}
.mian-desc-row {
  display: flex;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  margin: 0 0 20px 0;
  align-items: center;
}
.row-name {
  width: 100px;
  color: #768ca8;
}
.row-text {
  color: #ffffff;
  flex: 1;
  line-height: 20px;
}
.dialog-main-btns {
  display: flex;
  align-content: center;
  justify-content: center;
}
.confirm-btn {
  font-size: 16px;
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
  padding: 12px 0 18px 20px;
  border-radius: 5px;
}

.create .row-name {
  width: 100px;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
}
.create .row-line {
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
.row-line-btns {
  padding: 0;
}
.assets-info-txSigned {
  flex-direction: column;
  align-items: flex-start !important;
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
.assets-info-row {
  display: flex;
  align-items: center;
  margin: 0 0 20px;
}
.transfer-trx-type {
  display: flex;
}
.transfer-trx-type .row-name {
  width: 160px;
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
.cur-tab {
  border: 1px solid #4ecc9b !important;
  color: #4ecc9b;
  background: #3a526e;
}
</style>