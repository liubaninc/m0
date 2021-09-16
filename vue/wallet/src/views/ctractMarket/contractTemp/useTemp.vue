<template>
  <div class="contract-temp">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/ctractMarket' }"
          >合约市场</el-breadcrumb-item
        >
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >创建智能合约
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg detail-main">
        <div class="detail-main-assets">
          <div class="assets-title">合约信息</div>
          <div class="cm-submodule-bg contract-info">
            <div class="col-sm-2">
              <div class="col-sm-row contract-name">
                <div class="row-name">合约名称</div>
                <div class="row-input">
                  <el-input
                    v-model="contractName"
                    placeholder="请输入智能合约名称"
                  ></el-input>
                </div>
              </div>
              <div class="col-sm-row">
                <div class="row-name">合约版本</div>
                <div class="row-input">
                  <el-input
                    v-model="contractVersion"
                    placeholder="请输入智能合约版本"
                  ></el-input>
                </div>
              </div>
            </div>
            <div class="col-sm-1">
              <div class="col-sm-row">
                <div class="row-name">合约参数</div>
                <div class="row-input">
                  <el-input
                    v-model="args"
                    placeholder="请输入智能合约参数"
                  ></el-input>
                </div>
              </div>
            </div>
            <div class="col-sm-1">
              <div class="row-name">合约描述</div>
              <div class="row-input row-textarea">
                <el-input
                  type="textarea"
                  :rows="3"
                  :autosize="false"
                  resize="none"
                  placeholder="请输入合约描述"
                  v-model="contractDesc"
                  maxlength="50"
                  show-word-limit
                ></el-input>
              </div>
            </div>
          </div>
        </div>

        <div class="detail-main-assets">
          <div class="assets-title">模板详情</div>
          <div class="cm-submodule-bg temp-detail">
            <div class="temp-row">
              <div class="temp-row-name">合约模板名称</div>
              <div class="temp-row-text">
                {{ tempContractInfo.name }}
              </div>
            </div>
            <div class="temp-row">
              <div class="temp-row-name">合约描述</div>
              <div class="temp-row-text">
                {{ tempContractInfo.description }}
              </div>
            </div>
            <div class="temp-row">
              <div class="temp-row-name">接口描述</div>
              <div class="temp-row-text"></div>
            </div>
            <div class="trx-info-input template-lists">
              <el-table
                :data="tempContractInfo.functions"
                style="width: 100%"
                :header-cell-style="headerCell"
                :lazy="true"
              >
                <el-table-column
                  show-overflow-tooltip
                  prop="name"
                  label="函数名"
                  width="100"
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
          @click.prevent="createContractByTem"
          >创建智能合约</a
        >
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
  </div>
</template>
<script>
import { queryContractInfo } from '@/server/markets'
import { createContract } from '@/server/contract'
import { localCache } from '@/utils/utils'

export default {
  data() {
    return {
      wallet: {},
      tempContractInfo: {},
      contractName: '',
      contractVersion: '',
      contractDesc: '',
      isCtractFnished: false,
      contractDetail: {},
      args: '',
    }
  },
  created() {
    let wallet = localCache.get('wallet')
    if (wallet) {
      this.wallet = wallet
    }
    let { id } = this.$route.query
    if (id) {
      this.getContractDetail(id)
    }
  },
  computed: {
    headerCell() {
      return {
        // background: 'rgba(118, 140, 168, 1)!important',
        background: '#ebedf0!important',
        fontSize: '12px!important',
        lineHeight: '40px',
        padding: '0px',
        borderBottom: '0',
      }
    },
  },
  methods: {
    toPublishContract() {
      this.$router.replace(
        `/myContract/deploy?id=${this.contractDetail.id}&mode=deploy`
      )
    },
    async createContractByTem() {
      let {
        tempContractInfo,
        contractName,
        contractVersion,
        contractDesc,
        wallet,
        args,
      } = this

      if (!/^[a-zA-Z_]{1}[0-9a-zA-Z_.]+[0-9a-zA-Z_]/.test(contractName)) {
        this.$message.error(`合约名称只能以大小写字母开头且名称长度大于2个字符`)
        return
      }

      if (!/^([1-9]\d|[1-9])(.([1-9]\d|\d)){2,}$/.test(contractVersion)) {
        this.$message.error(`合约版本号只支持数字和.且至少三位数`)
        return
      }

      this.loading = this.$loading({
        lock: true,
        text: '创建合约中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)',
      })

      let resInfo = await createContract({
        account_name: wallet.address,
        name: contractName,
        version: contractVersion,
        type: 2,
        template_id: tempContractInfo.id,
        description: contractDesc,
        args,
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
    async getContractDetail(id) {
      let tempContractInfo = await queryContractInfo({ id })
      if (tempContractInfo) {
        this.tempContractInfo = tempContractInfo
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
.row-textarea {
  line-height: 20px;
  width: 100%;
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

.contract-info {
  padding: 20px 16px;
}
.row-name {
  width: 90px;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-size: 14px;
  color: #ffffff;
}
.col-sm-2 {
  display: flex;
}
.col-sm-row,
.col-sm-1 {
  display: flex;
  align-items: center;
  flex: 1;
}
.col-sm-1 {
  margin: 20px 0 0;
}
.row-input {
  flex: 1;
}
.contract-name {
  margin: 0 20px 0 0;
}
.temp-detail {
  padding: 40px 20px;
}
.temp-row {
  display: flex;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-size: 14px;
  color: #ffffff;
  margin: 0 0 20px;
  align-items: center;
}
.temp-row-name {
  width: 100px;
}
.temp-row-text {
  flex: 1;
  line-height: 20px;
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
</style>