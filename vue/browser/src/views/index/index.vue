<template>
  <div class="home">
    <div class="home-top">
      <div class="width1200 home-main-warpper">
        <h1 class="bytom-common-fb8622">
          <span>区块链浏览器</span><span></span>
        </h1>
        <div class="main-browser">
          <div class="main-browser-info">
            <div class="browser-info-item">
              <img src="../../assets/index/block_icon.png" />
              <p>区块高度</p>
              <p class="block-altitude">{{ chainInfo.block_num }}</p>
            </div>
            <div class="browser-info-item">
              <img src="../../assets/index/charge_icon.png" />
              <p>交易总量</p>
              <p class="block-altitude">{{ chainInfo.tx_num }}</p>
            </div>
            <div class="browser-info-item">
              <img src="../../assets/index/hour_icon.png" />
              <p>节点总数</p>
              <p class="block-altitude">{{ chainInfo.peer_num }}</p>
            </div>
            <div class="browser-info-item">
              <img src="../../assets/index/address_icon.png" />
              <p>资产总数</p>
              <p class="block-altitude">{{ chainInfo.asset_num }}</p>
            </div>
            <div class="browser-info-item">
              <img src="../../assets/index/save_icon.png" />
              <p>合约总数</p>
              <p class="block-altitude">{{ chainInfo.contract_num }}</p>
            </div>
            <div class="browser-info-item">
              <img src="../../assets/index/tps_icon.png" />
              <p>共识节点</p>
              <p class="block-altitude">{{ chainInfo.validator_num }}</p>
            </div>
          </div>
          <div class="main-block-change">
            <div class="main-block-title">7天区块高度</div>
            <div ref="blockCharts" class="block-charts"></div>
          </div>
        </div>
      </div>
    </div>
    <div class="home-list">
      <div class="width1200">
        <div class="table-tlt">
          <i class="el-icon-s-operation"></i>区块&交易
        </div>
        <div class="block-charge-table">
          <div class="block-table">
            <div class="detail-title">
              <span>最新区块</span>
              <el-link
                type="primary"
                class="block-more"
                @click="toMore('block')"
                >查看更多</el-link
              >
            </div>
            <el-table
              :data="blocks"
              :stripe="true"
              :border="true"
              :header-cell-style="{
                background: '#f0f0f0',
                color: 'rgba(0, 0, 0, 0.85)',
                fontWeight: 500,
                textAlign: 'center',
              }"
            >
              <el-table-column align="center" label="区块高度" width="80">
                <template slot-scope="scope">
                  <el-link
                    type="primary"
                    @click="toDetail(scope.row.height, 'block')"
                    >{{ scope.row.height }}</el-link
                  >
                </template>
              </el-table-column>
              <el-table-column
                width="252"
                align="center"
                prop="time"
                label="区块时间"
              >
              </el-table-column>
              <el-table-column
                align="center"
                prop="tx_num"
                label="交易数"
                width="80"
              >
              </el-table-column>
              <el-table-column align="center" label="出块人">
                <template slot-scope="scope">
                  <el-tooltip
                    class="item"
                    effect="dark"
                    :content="scope.row.proposer"
                    placement="top-end"
                  >
                    <span class="cm-word-split split-width">{{
                      scope.row.proposer
                    }}</span>
                  </el-tooltip>
                </template>
              </el-table-column>
            </el-table>
          </div>
          <div class="charge-table">
            <div class="detail-title">
              <span>最新交易</span>
              <el-link type="primary" class="block-more" @click="toMore('trx')"
                >查看更多</el-link
              >
            </div>
            <el-table
              :data="txs"
              :stripe="true"
              :border="true"
              :header-cell-style="{
                background: '#f0f0f0',
                color: 'rgba(0, 0, 0, 0.85)',
                fontWeight: 500,
                textAlign: 'center',
              }"
            >
              <el-table-column align="center" label="交易Hash">
                <template slot-scope="scope">
                  <el-tooltip
                    class="item"
                    effect="dark"
                    :content="scope.row.hash"
                    placement="top-start"
                  >
                    <span
                      class="cm-word-split split-width link"
                      @click="toDetail(scope.row.hash, 'trx')"
                      >{{ scope.row.hash }}</span
                    >
                  </el-tooltip>
                </template>
              </el-table-column>
              <el-table-column
                width="200"
                align="center"
                prop="time"
                label="交易时间"
              >
              </el-table-column>
              <el-table-column align="center" label="交易大小" width="80">
                <template slot-scope="scope">
                  <span>{{ scope.row.size | formateSize }}</span>
                </template>
              </el-table-column>
              <el-table-column
                align="center"
                prop="confirmed"
                label="确认数"
                width="80"
              >
              </el-table-column>
            </el-table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { createCategoryCharts } from '@/utils/echarts/category'
import { queryBlockChainInfo, queryBlocks } from '@server/block/queryBlock'
import { queryTxInfo, queryTxList } from '@server/transaction/queryTx'
export default {
  data() {
    return {
      chainInfo: {},
      blocks: [],
      txs: [],
    }
  },
  async created() {
    let chainInfo = (await queryBlockChainInfo()) || {}
    this.chainInfo = chainInfo
    let { blocks } = (await queryBlocks()) || []
    this.blocks = this.spliceFilter(blocks, 0, 6)
    let { txs } = (await queryTxList()) || {}
    this.txs = this.spliceFilter(txs, 0, 6)
  },
  mounted() {
    this.initCharts()
  },
  methods: {
    async initCharts() {
      let charts = await this.queryEcharts()
      let xAxis = []
      let series = []
      charts &&
        charts.reverse().forEach((item) => {
          xAxis.push(item['time'])
          series.push(item['number'])
        })
      createCategoryCharts(this.$refs.blockCharts, {
        xAxis,
        series,
      })
    },
    async queryEcharts() {
      let { blocks: blockHeight } = (await queryTxInfo()) || {}
      return blockHeight
    },
    toDetail(id, name) {
      let path = name == 'block' ? '/block/info' : '/block/tradeInfo'
      this.$router.push({
        path,
        query: {
          name: id,
          from: 'index',
        },
      })
    },
    toMore(name) {
      let path = name == 'block' ? '/block/list' : '/block/tradeList'
      this.$router.push({
        path,
        query: {
          from: 'index',
        },
      })
    },
    spliceFilter(data, start, end) {
      return data && data.length ? data.slice(start, end) : []
    },
  },
}
</script>

<style lang="scss">
.el-main {
  padding: 0 !important;
}
.home-top {
  width: 100%;
  // height: 352px;
  background: #fff url(../../assets/index/index_header.png) top -68px center/1920px
    no-repeat;
}
.bytom-common-fb8622 {
  font-size: 24px;
  font-weight: 400;
  color: #fff;
  padding: 40px 0 20px;
}
.main-browser {
  display: flex;
  align-content: space-between;
  justify-content: space-between;
}
.main-browser-info,
.main-block-change {
  flex: 0 0 576px;
  // height: 280px;
  background-color: rgb(255, 255, 255);
  box-shadow: rgb(0 0 0 / 8%) 0px 4px 8px;
  // padding: 40px 0;
}
.main-browser-info {
  padding: 40px 0 20px;
}
.main-block-change {
  // flex: 0 0 376px;
  padding: 40px 0 20px;
}
.main-browser-info {
  display: grid;
  grid-template-columns: 33.33% 33.33% 33.33%;
  grid-template-rows: calc(100wh / 2);
}
.browser-info-item {
  text-align: center;
  padding: 0 0 10px 0;
}
.browser-info-item img {
  width: 32px;
}
.browser-info-item p {
  color: rgb(179, 179, 179);
  font-size: 14px;
  line-height: 20px;
  margin: 9px 0px;
}
.browser-info-item .block-altitude {
  color: rgb(36, 36, 36);
  font-size: 24px;
}

.home-list {
  padding: 30px 0 0;
}
.detail-title {
  display: flex;
  justify-content: space-between;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  font-variant: tabular-nums;
  line-height: 1.5;
  list-style: none;
  -webkit-font-feature-settings: 'tnum';
  font-feature-settings: 'tnum';
  padding: 16px 0;
}
.el-table thead {
  background: #fafafa;
}
.table-tlt {
  margin: 20px 0 0;
  font-size: 18px;
}
.table-tlt i {
  color: #409eff;
  font-weight: 700;
  margin-right: 10px;
}
.el-icon-document-copy {
  cursor: pointer;
}
.block-charge-table {
  display: flex;
  justify-content: space-between;
}
.block-table,
.charge-table {
  flex: 0 0 576px;
}
.block-more:hover {
  text-decoration: none;
}

.block-charts {
  width: 100%;
  height: 100%;
}
.main-block-title {
  font-size: 22px;
  font-weight: 700;
  padding: 0 0 10px 20px;
}
</style>
