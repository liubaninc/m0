<template>
  <div class="trxdetail dcontract">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: `/mycontract` }"
          >我的合约</el-breadcrumb-item
        >
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >智能合约详情
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg detail-main">
        <div class="detail-main-assets">
          <div class="assets-title">智能合约信息</div>
          <div class="cm-submodule-bg assets-info">
            <div class="assets-info-row2">
              <div class="info-row2-line">
                <span class="row-name">合约名称 </span>
                <div class="cm-text-overflow cm-row-bg row2-line-context">
                  {{ contractInfo.name || '---' }}
                </div>
              </div>
              <div class="info-row2-line">
                <span class="row-name">版本 </span>
                <div class="cm-row-bg row2-line-context">
                  {{ contractInfo.version || '---' }}
                </div>
              </div>
            </div>
            <div class="assets-info-row2">
              <div class="info-row2-line">
                <span class="row-name">合约状态 </span>
                <div class="cm-row-bg row2-line-context">
                  {{ contractInfo.status | contractStatusText }}
                </div>
              </div>
              <div class="info-row2-line">
                <span class="row-name">部署链 </span>
                <div class="cm-row-bg row2-line-context">
                  {{ contractInfo.alliance_name || '---' }}
                </div>
              </div>
            </div>
            <div class="assets-info-row2">
              <div class="info-row2-line">
                <span class="row-name">创建时间 </span>
                <div class="cm-row-bg row2-line-context">
                  {{ contractInfo.created_at || '---' }}
                </div>
              </div>
            </div>
            <div class="assets-info-row">
              <span class="row-name">描述 </span>
              <div class="cm-row-bg row-line-context">
                {{ contractInfo.description || '---' }}
              </div>
            </div>
          </div>

          <div class="trx-detail-main">
            <div class="assets-title">更多</div>
            <div class="cm-submodule-bg trx-item">
              <el-tabs v-model="activeName">
                <el-tab-pane label="交易记录" name="trxRecord">
                  <div class="trx-info">
                    <div class="trx-info-input">
                      <el-table
                        :data="trxLists"
                        lazy
                        style="width: 100%"
                        :header-cell-style="headerCell"
                        :empty-text="`暂无交易信息`"
                      >
                        <el-table-column label="交易哈希" show-overflow-tooltip>
                          <template slot-scope="scope">
                            <span class="cm-text-overflow table-row">{{
                              scope.row.hash
                            }}</span>
                          </template>
                        </el-table-column>
                        <el-table-column
                          label="区块高度"
                          width="140"
                          show-overflow-tooltip
                        >
                          <template slot-scope="scope">
                            {{ scope.row.height }}
                          </template>
                        </el-table-column>
                        <el-table-column
                          label="交易时间"
                          width="210"
                          show-overflow-tooltip
                        >
                          <template slot-scope="scope">
                            {{ scope.row.time }}
                          </template>
                        </el-table-column>
                      </el-table>
                    </div>
                  </div>
                  <div
                    class="pagination-main detail-pages"
                    v-if="trxLists.length"
                  >
                    <el-pagination
                      class="pagination"
                      @size-change="handleSizeChange"
                      @current-change="handleCurrentChange"
                      :current-page="page.pageNum"
                      :page-sizes="[10, 20, 30, 40]"
                      :page-size="page.pageSize"
                      layout="total, sizes, prev, pager, next, jumper"
                      :total="page.pageTotal"
                      background
                    ></el-pagination>
                  </div>
                </el-tab-pane>
                <el-tab-pane label="历史版本及下载" name="downLoadHistory">
                  <div class="trx-info">
                    <div class="trx-info-input">
                      <el-table
                        :data="contractVersions"
                        lazy
                        style="width: 100%"
                        :header-cell-style="headerCell"
                      >
                        <el-table-column
                          label="合约版本"
                          width="100"
                          show-overflow-tooltip
                        >
                          <template slot-scope="scope">
                            <span class="cm-text-overflow table-row">{{
                              scope.row.version
                            }}</span>
                          </template>
                        </el-table-column>
                        <el-table-column label="更新时间" show-overflow-tooltip>
                          <template slot-scope="scope">
                            {{ scope.row.updated_at }}
                          </template>
                        </el-table-column>
                        <el-table-column label="更新原因" show-overflow-tooltip>
                          <template slot-scope="scope">
                            <span>{{ scope.row.description }}</span>
                          </template>
                        </el-table-column>
                        <el-table-column label="操作" width="120">
                          <template slot-scope="scope">
                            <a
                              href="javascript:;"
                              class="cm-font-009F72"
                              @click="toDownLoadConFile(scope.row)"
                              >下载</a
                            >
                          </template>
                        </el-table-column>
                      </el-table>
                    </div>
                  </div>
                  <!-- <template>
                      <span class="block-request">暂无交易请求</span>
                    </template> -->
                </el-tab-pane>
              </el-tabs>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import {
  queryContractDetail,
  queryContractVersions,
  queryContractTrxByName,
} from '@/server/contract'
import { downLoadFile } from '@/utils/utils'
export default {
  data() {
    return {
      activeNames: ['trxRecord'],
      activeName: 'trxRecord',
      page: {
        pageSize: 10,
        pageTotal: 0,
        pageNum: 1,
      },
      contractInfo: {},
      trxLists: [],
      contractVersions: [],
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
  async created() {
    let { id } = this.$route.query
    if (id) {
      this.id = id
      let info = await this.getContractInfoById(id)
      if (info) {
        this.getContractTrxLists(
          info.name,
          this.page.pageNum,
          this.page.pageSize
        )
        this.getContractVersion(info.name)
      }
    }
  },
  methods: {
    toDownLoadConFile(item) {
      let origin = window.location.origin
      downLoadFile(
        `${process.env.VUE_APP_PRO_BASE_URL}/mcontract/download/${item.id}`
      )
    },
    handleSizeChange(pageSize) {
      this.getContractTrxLists(
        this.contractInfo.name,
        this.page.pageNum,
        pageSize
      )
    },
    handleCurrentChange(pageNum) {
      this.getContractTrxLists(
        this.contractInfo.name,
        pageNum,
        this.page.pageSize
      )
    },
    async getContractInfoById(id) {
      let contractInfo = await queryContractDetail({ id })
      if (contractInfo) {
        this.contractInfo = contractInfo
        return contractInfo
      }
    },
    async getContractVersion(contractName) {
      let contractVersions = await queryContractVersions({
        contractName,
      })
      if (contractVersions) {
        this.contractVersions = contractVersions
      }
    },
    async getContractTrxLists(name, pageNum, pageSize) {
      let { page_num, page_size, total, txs } = await queryContractTrxByName({
        name,
        page_num: pageNum,
        page_size: pageSize,
      })
      if (txs) {
        this.trxLists = txs
        this.page.pageSize = page_size
        this.page.pageTotal = total
        this.page.pageNum = page_num
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
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
}
.row-line {
  min-height: 45px;
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
  /* background: url("../../assets/images/wallet/copy_icon.png"); */
  background-size: 100%;
  cursor: pointer;
  color: #22ac95;
}
/*end detail* */

/**trxdetail start*/
.assets-info-row2 {
  display: flex;
  margin-bottom: 20px;
}
.info-row2-line {
  display: flex;
  align-items: center;
  justify-items: center;
  width: 33%;
  margin: 0 20px 0 0;
}
.info-row2-line:nth-child(2n) {
  margin-right: 0;
}
.row2-line-context {
  min-height: 45px;
  padding: 10px 20px;
  display: flex;
  align-items: center;
  width: 74%;
  justify-content: space-between;
}
.trx-detail-main {
  padding: 17px 0 0;
}
.trx-item {
  /* margin: 0 0 20px; */
  padding: 20px 20px 30px;
}

.trx-hash-row {
  width: 90%;
}

.trx-info {
  display: flex;
}
.trx-info-input {
  flex: 1;
}
.trx-info-name {
  font-size: 22px;
}
/**end trxdetail */

.row-line-context {
  min-height: 45px;
  padding: 10px 20px;
  width: 60%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  line-height: 20px;
}
</style>