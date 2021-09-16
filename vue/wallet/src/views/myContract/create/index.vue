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
                  <el-input
                    v-model="name"
                    placeholder="请输入合约名称"
                  ></el-input>
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
              <div class="col-upload" v-if="genWay == 1">
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
              <div class="col-template" v-else>
                <div class="contract-file">
                  <div class="upload-name">合约参数</div>
                  <div class="upload-input contract-params">
                    <el-input
                      type="text"
                      v-model="tempArgs"
                      placeholder="请输入合约参数(选填)"
                    >
                    </el-input>
                  </div>
                </div>
                <div class="trx-info contract-template">
                  <div class="upload-name">选择模板</div>
                  <div class="trx-info-input template-lists">
                    <el-table
                      :data="contractTemps"
                      style="width: 100%"
                      :header-cell-style="headerCell"
                      :lazy="true"
                    >
                      <el-table-column width="50">
                        <template slot-scope="scope">
                          <input
                            type="radio"
                            :value="scope.row.id"
                            v-model="tempId"
                            class="radio-radius"
                          />
                        </template>
                      </el-table-column>
                      <el-table-column
                        show-overflow-tooltip
                        prop="name"
                        label="合约模板名称"
                        width="180"
                      >
                      </el-table-column>

                      <el-table-column
                        show-overflow-tooltip
                        prop="description"
                        label="描述"
                      >
                      </el-table-column>
                      <el-table-column label="操作" width="100">
                        <template slot-scope="scope">
                          <el-button
                            type="text"
                            size="small"
                            class="opt-detail"
                            @click="toTempDetail(scope.row)"
                          >
                            详情
                          </el-button>
                        </template>
                      </el-table-column>
                    </el-table>
                    <div
                      class="pagination-main"
                      v-if="contractTemps && contractTemps.length"
                    >
                      <el-pagination
                        class="pagination"
                        @size-change="handleSizeChange"
                        @current-change="handleCurrentChange"
                        :current-page="page.pageNum"
                        :page-sizes="[10, 20, 30, 40]"
                        :page-size="page.pageSize"
                        layout="total, sizes, prev, pager, next, jumper"
                        :total="page.total"
                        background
                      >
                      </el-pagination>
                    </div>
                  </div>
                </div>
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
            >创建智能合约</a
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
            <a
              href="javascript:;"
              class="cm-btn-138px cm-btn-bg4acb9b"
              @click="toPublishContract"
              >立即部署合约</a
            >
          </div>
        </div>
      </div>
      <div class="dialog-yy"></div>
    </div>

    <div class="tmp-dialog" v-if="dialogIsShow">
      <div class="dialog-main">
        <div class="dialog-main-title">模板详情</div>
        <div class="dialog-main-desc">
          <div class="mian-desc-row">
            <div class="row-name">合约模板名称</div>
            <div class="row-text">{{ curTempDetial.name }}</div>
          </div>
          <div class="mian-desc-row">
            <div class="row-name">合约描述</div>
            <div class="row-text">
              {{ curTempDetial.description }}
            </div>
          </div>
          <div class="mian-desc-row">
            <div class="row-name">接口描述</div>
            <div class="row-text"></div>
          </div>
          <div class="main-desc-function">
            <el-table
              :data="curTempDetial.functions"
              style="width: 100%"
              :header-cell-style="tempHeaderCell"
              :lazy="true"
            >
              <el-table-column
                show-overflow-tooltip
                prop="name"
                label="函数名称"
                width="180"
              >
              </el-table-column>

              <el-table-column show-overflow-tooltip prop="args" label="参数">
              </el-table-column>
              <el-table-column
                show-overflow-tooltip
                prop="description"
                label="简介"
              >
              </el-table-column>
            </el-table>
          </div>
        </div>
        <div class="dialog-main-btns">
          <a
            href="javascript:;"
            class="cm-btn-bg009F72 cm-btn-138px confirm-btn"
            @click.stop.prevent="dialogIsShow = !1"
            >确定</a
          >
        </div>
      </div>
      <div class="dialog-yy" @click.stop.prevent="dialogIsShow = !1"></div>
    </div>
  </div>
</template>
<script>
import { createContract } from '@/server/contract'
import { queryContractMarket, queryContractInfo } from '@/server/markets'

import { localCache } from '@/utils/utils'

export default {
  data() {
    return {
      name: '',
      version: '',
      description: '',
      contractFile: {},
      args: '',
      contractGenTypes: [
        { id: 1, text: '上传合约' },
        { id: 2, text: '合约模板' },
      ],
      genWay: 1,
      isCtractFnished: false,
      dialogIsShow: false,
      contractDetail: {},
      contractTemps: [],
      tempId: '',
      tempArgs: '',
      curTempDetial: {},

      wallet: {},
      page: {
        pageSize: 10,
        total: 0,
        pageNum: 1,
      },
    }
  },
  watch: {
    genWay(newVal, oldVal) {
      if (newVal == 1) {
        this.tempId = ''
        this.tempArgs = ''
      } else {
        this.contractFile = {}
        this.args = ''
      }
    },
    dialogIsShow(nVal, oVal) {
      if (!nVal) {
        this.curTempDetial = {}
      }
    },
  },
  created() {
    let wallet = localCache.get('wallet')
    if (wallet) {
      this.wallet = wallet
    }
    this.getContractTemps(this.page.pageNum, this.page.pageSize)
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
    tempHeaderCell() {
      return {
        background: '#ebedf0!important',
        fontSize: '12px!important',
        lineHeight: '40px',
        padding: '0px',
        borderBottom: '0',
      }
    },
  },
  methods: {
    toTempDetail(item) {
      // this.curTempDetial = item
      this.getTempInfo(item.id)
      this.dialogIsShow = true
    },
    toPublishContract() {
      this.$router.replace(
        `/myContract/deploy?id=${this.contractDetail.id}&mode=deploy`
      )
    },
    async toCreateContract() {
      let {
        name,
        version,
        description,
        contractFile,
        args,
        genWay,
        tempId,
        tempArgs,
        wallet,
      } = this

      console.log(
        '===============>',
        /^[a-zA-Z_]{1}[0-9a-zA-Z_.]+[0-9a-zA-Z_]/.test(name),
        name
      )

      if (!/^[a-zA-Z_]{1}[0-9a-zA-Z_.]+[0-9a-zA-Z_]/.test(name)) {
        this.$message.error(`合约名称只能以大小写字母开头且名称长度大于2个字符`)
        return
      }

      if (!/^([1-9]\d|[1-9])(.([1-9]\d|\d)){2,}$/.test(version)) {
        this.$message.error(`合约版本号只支持数字和.且至少三位数`)
        return
      }

      if (genWay == 1) {
        if (!(contractFile && contractFile.raw)) {
          this.$message.error(`请上传合约文件`)
          return
        }
        // if (!args) {
        //   this.$message.error(`合约参数不能为空`)
        //   return
        // }
      } else {
        // if (!tempArgs) {
        //   this.$message.error(`合约参数不能为空`)
        //   return
        // }
      }

      this.loading = this.$loading({
        lock: true,
        text: '合约创建中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)',
      })

      let resInfo = await createContract({
        account_name: wallet.address,
        name,
        version,
        type: genWay,
        description,
        args: args || tempArgs,
        file: contractFile.raw || '',
        template_id: tempId || '',
      })
      if (resInfo) {
        this.loading.close()
        // this.$message.success(`合约创建成功`)
        this.isCtractFnished = true
        this.contractDetail = resInfo
      } else {
        this.loading.close()
        this.isCtractFnished = false
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
    async getContractTemps(pageNum = 1, pageSize = 10) {
      let { list, page_num, page_size, total } = await queryContractMarket({
        page_num: pageNum,
        page_size: pageSize,
      })
      if (list) {
        this.contractTemps = list
        this.page.pageSize = page_size
        this.page.total = total
        this.page.pageNum = page_num
      }
    },
    handleSizeChange(pageSize) {
      this.getContractTemps(this.page.pageNum, pageSize)
    },
    handleCurrentChange(pageNum) {
      this.getContractTemps(pageNum, this.page.pageSize)
    },
    async getTempInfo(id) {
      let curTempDetial = await queryContractInfo({ id })
      if (curTempDetial) {
        this.curTempDetial = curTempDetial
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
  width: 400px;
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
.main-desc-function {
  overflow: auto;
  width: 100%;
}

.cur-tab {
  border: 1px solid #4ecc9b !important;
  color: #4ecc9b;
  background: #3a526e;
}
</style>