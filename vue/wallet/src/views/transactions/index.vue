<template>
  <div class="com-breadcrumb trx">
    <div class="trx-title">交易记录</div>
    <div class="trx-condition">
      <div class="trx-condition-row">
        <span>资产</span>
        <el-select v-model="assetValue" @change="assetsChange">
          <template v-if="assets.length">
            <el-option label="全部" value="全部"> </el-option>
            <el-option
              v-for="coin in assets"
              :key="coin.denom"
              :label="coin.denom"
              :value="coin.denom"
            >
            </el-option>
          </template>
        </el-select>
      </div>
      <div class="trx-condition-row">
        <span>交易类型 </span>
        <el-select v-model="trxType" @change="typeChange">
          <el-option
            v-for="item in trxTypeLists"
            :key="item.id"
            :label="item.label"
            :value="item.value"
          >
          </el-option>
        </el-select>
      </div>
    </div>
    <div class="trx-list">
      <template v-if="trxs.length">
        <template v-for="trx in trxs">
          <div class="cm-module-bg trx-list-row" :key="trx.hash">
            <div class="list-row-top">
              <div class="top-assets assets-left">
                <div class="assets-title">资产名称</div>
                <div class="assets-name">{{ trx.assets }}</div>
              </div>
              <div class="top-assets">
                <div class="assets-title">交易类型</div>
                <div class="assets-name">{{ trx.type }}</div>
              </div>
              <div class="top-assets trx-status">
                <div class="assets-title">交易状态</div>
                <div class="assets-name">{{ trx.height | trxStatus }}</div>
              </div>
              <div class="top-assets trx-btn-info" @click="toTrxDetail(trx)">
                <img
                  src="../../assets/images/detail/dapps/cz_icon.png"
                  alt=""
                  srcset=""
                />
                <div>交易详情</div>
              </div>
            </div>
            <div class="list-row-bottom">
              <div>
                <span class="trx-hash-name">交易哈希</span>
                <el-tooltip effect="dark" :content="trx.hash" placement="top">
                  <span class="cm-text-overflow trx-hash">
                    {{ trx.hash }}</span
                  ></el-tooltip
                >
              </div>
              <div>
                <span class="trx-hash-name">交易时间</span>
                <span>{{ trx.time | dateFormat("YYYY-mm-dd HH:MM:SS") }} </span>
              </div>
            </div>
          </div>
        </template>
      </template>
      <template v-else>
        <div class="cm-module-bg trx-list-row trx-no-data">暂无交易记录</div>
      </template>

      <div v-if="trxs && trxs.length" class="pagination-main">
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
        >
        </el-pagination>
      </div>
    </div>
  </div>
</template>
<script>
import { queryTrLists } from "@/server/transaction";
import { queryAssetsByAddress } from "@/server/assets";
import { trxTypes } from "@/config/asset";
export default {
  data() {
    return {
      address: "",
      coin: "",
      trxTypeLists: [],
      assetValue: "全部",
      trxType: "全部",
      assets: [],
      trxs: [],
      page: {
        pageSize: 10,
        pageTotal: 0,
        pageNum: 1,
      },
    };
  },
  created() {
    let { address, coin } = this.$route.query;
    this.trxTypeLists = trxTypes();
    this.address = address;
    if (coin) {
      this.coin = coin;
      this.assetValue = coin;
      this.getTrxListByAddress(address, coin);
      this.getAddressAllAssets(address);
      this.timer = setInterval(() => {
        this.getTrxListByAddress(address, coin);
      }, 1000 * 60);
    } else {
      this.getTrxListByAddress(address, "");
      this.getAddressAllAssets(address);
      this.timer = setInterval(() => {
        this.getTrxListByAddress(address, "");
      }, 1000 * 60);
    }
  },
  destroyed() {
    clearInterval(this.timer);
  },
  methods: {
    handleSizeChange(pageSize) {
      this.getTrxListByAddress(
        this.address,
        name,
        this.page.pageNum,
        pageSize,
        this.trxType
      );
    },
    handleCurrentChange(pageNum) {
      this.getTrxListByAddress(
        this.address,
        name,
        pageNum,
        this.page.pageSize,
        this.trxType
      );
    },
    assetsChange(name) {
      this.getTrxListByAddress(
        this.address,
        name,
        this.page.pageNum,
        this.page.pageSize,
        this.trxType
      );
    },
    typeChange(type) {
      this.getTrxListByAddress(
        this.address,
        this.assetValue,
        this.page.pageNum,
        this.page.pageSize,
        type
      );
    },
    toTrxDetail(trx) {
      this.$router.push(
        `/trx/detail?hash=${trx.hash}&address=${this.address}&coin=${this.coin}`
      );
    },
    async getAddressAllAssets(addrs) {
      let { coins, address } = await queryAssetsByAddress({ address: addrs });
      if (coins) {
        this.assets = coins;
      }
    },
    async getTrxListByAddress(
      address,
      coin,
      pageNum = this.pageNum,
      pageSize = this.pageSize,
      type = ""
    ) {
      if (coin == "全部") coin = "";
      if (type == "全部") type = "";
      let { page_num, page_size, total, txs } = await queryTrLists({
        address,
        coin,
        type,
        page_num: pageNum,
        page_size: pageSize,
      });
      if (txs) {
        this.page.pageNum = page_num;
        this.page.pageSize = page_size;
        this.page.pageTotal = total;
        this.trxs = txs;
      } else {
        this.trxs = [];
      }
    },
  },
};
</script>
<style >
.trx-title {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
}
.trx-condition {
  display: flex;
}

.trx-condition-row {
  margin: 30px 0 15px;
  margin-right: 40px;
}
.trx-condition-row :last-child {
  margin-right: 0;
}
.trx-condition-row span {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  margin: 0 30px 0 0;
}
.trx-list {
  padding: 16px 0 0;
}
.trx-list-row {
  border-radius: 5px;
  height: 155px;
  display: flex;
  flex-direction: column;
  padding: 26px 30px;
  justify-content: center;
  margin: 0 0 20px;
}
.trx-list-row:last-child {
  margin-bottom: 0;
}
.list-row-top {
  display: flex;
  align-items: center;
  align-content: center;
}
.top-assets {
  width: 20%;
}
.assets-left {
  width: 40%;
}
.assets-title {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #768ca8;
}
.assets-name {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
  margin: 14px 0 0;
}
.list-row-bottom {
  display: flex;
  justify-content: space-between;
  margin: 30px 0 0;
  color: #cccccc;
}
.trx-btn-info {
  text-align: right;
  color: #62f7d4;
  cursor: pointer;
}
.trx-btn-info img {
  width: 20px;
  height: 20px;
  margin: 0 auto 18px;
  transform: translate(-16px, 0px);
}
.trx-hash-name {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 12px;
  color: #768ca8;
  margin: 0 14px 0 0;
}
.trx-status {
  text-align: center;
}
.trx-hash {
  display: inline-block;
  width: 80%;
  vertical-align: middle;
}
.trx-no-data {
  text-align: center;
  font-size: 16px;
}
</style>