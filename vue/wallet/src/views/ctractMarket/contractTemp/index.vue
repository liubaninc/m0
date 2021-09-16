<template>
  <div class="com-breadcrumb market">
    <div class="trx-title">合约市场</div>
    <div class="market-main">
      <div class="cm-sub-bg market-main-desc">
        <div class="market-img">
          <img
            src="../../../assets/images/market/market_icon.png"
            alt=""
            srcset=""
          />
        </div>
        <div class="market-desc">
          <div class="desc-name">智能合约商店</div>
          <div class="desc-text">
            快速选择合约模板，轻松创建、部署、管理智能合约，提升研发效率，降低业务成本
          </div>
        </div>
      </div>
      <div class="market-list">
        <template v-if="contractMarkets && contractMarkets.length">
          <template v-for="market in contractMarkets">
            <div class="market-item">
              <div class="item-top">
                <div class="cm-text-overflow item-name">
                  {{ market.name }}
                </div>
                <div class="cm-text-overflow2 item-desc">
                  {{ market.description }}
                </div>
              </div>
              <div class="item-btns">
                <a href="javascript:;" @click.prevent.stop="useTemp(market)"
                  ><span class="iconfont m0-xiazai icon-btn"></span>立即使用</a
                >
                <span class="split-line"></span>
                <a href="javascript:;" @click="toTempDetail(market)"
                  ><span class="iconfont m0-tishi icon-btn"></span>模板详情</a
                >
              </div>
            </div>
          </template>
        </template>
      </div>
      <!-- <div class="pagination-main">
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
      </div> -->

      <div class="dialog" v-if="dialogIsShow">
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
            <div class="trx-info-input template-lists">
              <el-table
                :data="curTempDetial.functions"
                style="width: 100%"
                :header-cell-style="headerCell"
                :lazy="true"
              >
                <el-table-column
                  show-overflow-tooltip
                  prop="name"
                  label="函数名"
                  width="140"
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
  </div>
</template>
<script>
import { queryContractMarket, queryContractInfo } from '@/server/markets'
export default {
  data() {
    return {
      contractMarkets: [],
      page: {
        pageSize: 100000,
        total: 0,
        pageNum: 1,
      },
      dialogIsShow: false,
      curTempDetial: {},
    }
  },
  created() {
    this.getContractMarkets()
  },
  watch: {
    dialogIsShow(nVal, oVal) {
      if (!nVal) {
        this.curTempDetial = {}
      }
    },
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
    useTemp(item) {
      this.$router.push(`/ctractMarket/useTemp?id=${item.id}`)
    },
    toTempDetail(item) {
      // this.curTempDetial = item
      this.dialogIsShow = true
      this.getTempInfo(item.id)
    },
    handleSizeChange(pageSize) {
      this.getContractMarkets(this.page.pageNum, pageSize)
    },
    handleCurrentChange(pageNum) {
      this.getContractMarkets(pageNum, this.page.pageSize)
    },
    async getContractMarkets(pageNum = 1, pageSize = 10) {
      let { list, page_num, page_size, total } = await queryContractMarket({
        page_num: pageNum,
        page_size: pageSize,
      })
      if (list) {
        this.contractMarkets = list
        this.page.pageSize = page_size
        this.page.total = total
        this.page.pageNum = page_num
      }
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
.trx-title {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
  margin: 0 0 20px;
}

.market-main-desc {
  display: flex;
  align-items: center;
  height: 140px;
  padding: 0 40px;
  border-radius: 11px;
}
.market-img {
  width: 60px;
  height: 60px;
}
.market-img img {
  width: 100%;
  height: 100%;
}
.market-desc {
  margin: 0 0 0 30px;
}
.desc-name {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 24px;
  color: #62f7d4;
}
.desc-text {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: #cbcbcc;
  margin: 16px 0 0;
}
.market-list {
  margin: 20px 0 0;
  display: flex;
  flex-wrap: wrap;
  /* display: grid;
  grid-template-columns: 45% 45%;
  grid-column-gap: 10%;
  grid-row-gap: 10%; */
}
.market-item {
  /* width: 586px; */
  /* width: 556px; */
  width: 48%;
  height: 193px;
  background-color: rgba(42, 64, 92, 1);
  border-radius: 11px;
  font-family: 'PingFangSC-Thin', 'PingFang SC Thin', 'PingFang SC', sans-serif;
  font-weight: 200;
  font-size: 12px;
  color: #2a405c;
  display: flex;
  flex-direction: column;
  margin: 0 20px 20px 0px;
  /* display: inline-block; */
}
.market-item:nth-child(2n) {
  margin-right: 0;
}
.item-top {
  flex: 1;
  padding: 36px 30px 0;
}
.item-name {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
}
.item-desc {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #b9b9b9;
  margin: 14px 0 0;
  line-height: 20px;
}
.item-btns {
  border-top: 1px solid #797979;
  height: 46px;
  line-height: 46px;
  display: flex;
  align-items: center;
  justify-items: center;
}
.item-btns a {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #62f7d4;
  flex: 1;
  text-align: center;
}
.icon-btn {
  vertical-align: bottom;
  margin: 0 10px 0 0;
}
.split-line {
  width: 1px;
  background: #6e7174;
  height: 60%;
}

.dialog-main-desc {
  padding: 20px 14px;
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
</style>