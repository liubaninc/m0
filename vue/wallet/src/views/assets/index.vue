<template>
  <div class="com-breadcrumb assets">
    <div class="trx-title">资产管理</div>
    <div class="trx-condition assets-condition-row">
      <div class="trx-condition-row">
        <span>资产</span>
        <el-select v-model="curAssets" placeholder="全部" @change="changeAsset">
          <el-option key="全部" label="全部" value="全部"> </el-option>
          <el-option
            v-for="assets in assetSelectLists"
            :key="assets.denom"
            :label="assets.denom"
            :value="assets.denom"
          >
          </el-option>
        </el-select>
      </div>
      <div class="trx-condition-btns">
        <template v-if="wallet.threshold > 0">
          <a
            href="javascript:;"
            class="cm-btn-border009F72 cm-btn-225px"
            @click="$router.push('/assets/signature')"
            ><span class="iconfont m0-bianjimian edit-icon"></span
            >多签交易签名</a
          >
        </template>
        <a
          href="javascript:;"
          class="cm-btn-bg4acb9b cm-btn-225px publish-assets"
          @click="publishAssets"
          >发行资产</a
        >
      </div>
    </div>
    <div class="assets-cnd">
      <el-checkbox v-model="isVisibility" @change="hideZeroAssets"
        >隐藏持有为0的资产
      </el-checkbox>
    </div>
    <div class="trx-list">
      <template v-if="assetLists && assetLists.length">
        <template v-for="assets in assetLists">
          <div
            class="cm-module-bg trx-list-row assets-row-item"
            v-if="assets.denom != '全部'"
          >
            <div class="list-row-top">
              <div class="assets-left">
                <div class="assets-title">资产名称</div>
                <div class="assets-name">{{ assets.denom }}</div>
              </div>
              <div class="assets-left">
                <div class="assets-title">当前持有</div>
                <div class="assets-name">{{ assets.amount }}</div>
              </div>
              <div class="top-items">
                <div class="tp-item-tab" @click="receiveAssets(assets)">
                  <img
                    src="../../assets/images/detail/dapps/cz_icon.png"
                    alt=""
                    srcset=""
                  />
                  接收转账
                </div>
                <div class="tp-item-tab" @click="trxOutassets(assets)">
                  <img
                    src="../../assets/images/detail/dapps/cz_icon.png"
                    alt=""
                    srcset=""
                  />
                  资产转出
                </div>
                <div class="tp-item-tab tp-more-tab">
                  <img
                    src="../../assets/images/detail/dapps/cz_icon.png"
                    alt=""
                    srcset=""
                  />
                  更多

                  <div class="wt-el-menu">
                    <a
                      href="javascript:;"
                      class="menu-item"
                      @click="toAssetDetail(assets)"
                      >资产详情</a
                    >
                    <a
                      href="javascript:;"
                      class="menu-item"
                      @click="addAssets(assets)"
                      >增发资产</a
                    >
                    <a
                      href="javascript:;"
                      class="menu-item"
                      @click="burnAssets(assets)"
                      >销毁资产</a
                    >
                    <a
                      href="javascript:;"
                      class="menu-item"
                      @click="toTrxLists(assets)"
                      >交易记录</a
                    >
                    <div class="el-menu-triangle"></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </template>
      </template>
      <template v-else>
        <div class="cm-module-bg assets-nodata">暂无资产信息</div>
      </template>

      <div v-if="assetLists && assetLists.length" class="pagination-main">
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
</template>
<script>
import { queryAssetLists, queryAssetsListPage } from "@/server/assets";
import { localCache } from "@/utils/utils";
export default {
  data() {
    return {
      isVisibility: false,
      curAssets: "全部",
      assetSelectLists: {},
      assetLists: [],
      wallet: {},
      page: {
        pageSize: 10,
        total: 0,
        pageNum: 1,
      },
    };
  },
  created() {
    let wallet = localCache.get("wallet");
    if (wallet) {
      this.wallet = wallet;
      this.getAssetsList(wallet.address);
      this.getAssets(wallet.address, this.page.pageNum, this.page.pageSize);
      this.timer = setInterval(() => {
        this.getAssets(wallet.address, this.page.pageNum, this.page.pageSize);
      }, 1000 * 60);
    }
  },
  destroyed() {
    clearInterval(this.timer);
  },
  methods: {
    handleSizeChange(pageSize) {
      this.getAssets(this.wallet.address, this.page.pageNum, pageSize);
    },
    handleCurrentChange(pageNum) {
      this.getAssets(this.wallet.address, pageNum, this.page.pageSize);
    },
    burnAssets(assets) {
      if (assets) {
        this.$router.push(`/assets/destoryAssets?assetsName=${assets.denom}`);
      }
    },
    addAssets(assets) {
      if (assets) {
        this.$router.push(`/assets/addSessets?assetsName=${assets.denom}`);
      }
    },
    publishAssets(assets) {
      this.$router.push(`/assets/publishAsset`);
    },
    toAssetDetail(assets) {
      if (assets.denom) {
        this.$router.push(
          `/assets/detail?denom=${assets.denom}&address=${this.wallet.address}`
        );
      }
    },
    trxOutassets(asset) {
      let { wallet } = this;
      if (wallet.threshold > 0) {
        this.$router.push(`/assets/transferOutMore?coin=${asset.denom}`);
      } else {
        this.$router.push(`/assets/transferOut?coin=${asset.denom}`);
      }
    },
    receiveAssets(assets) {
      this.$router.push(
        `/assets/receive?address=${this.wallet.address}&name=${assets.denom}`
      );
    },
    toTrxLists(assets) {
      this.$router.push(
        `/trx?address=${this.wallet.address}&coin=${assets.denom}`
      );
    },
    hideZeroAssets(mark) {
      if (mark) {
        this.assetLists = this.assetSelectLists.filter(
          (assets) => assets.amount != 0
        );
      } else {
        this.assetSelectLists = this.assetSelectLists;
      }
    },
    changeAsset(name) {
      this.assetLists = this.getAssetsRowList(name);
    },
    getAssetsRowList(name) {
      if (name == "全部") return this.assetSelectLists;
      return this.assetSelectLists.filter((assets) => assets.denom == name);
    },
    async getAssets(address, pageNum, pageSize) {
      let { coins, page_num, page_size, total } = await queryAssetsListPage({
        address,
        page_num: pageNum,
        page_size: pageSize,
      });
      if (coins) {
        this.assetLists = coins;
        this.page.pageNum = page_num;
        this.page.total = total;
        this.page.pageSize = pageSize;
      }
    },
    async getAssetsList(address) {
      let assetLists = await queryAssetLists({
        address,
      });
      if (assetLists) {
        this.assetSelectLists = assetLists.coins;
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
  color: #22ac95;
  cursor: pointer;
}
.trx-btn-info img {
  width: 20px;
  height: 20px;
  margin: 0 auto 18px;
  transform: translate(-16px, 0px);
}
.trx-hash {
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

/**asset start */
.top-items {
  display: flex;
  justify-content: space-between;
  align-content: space-between;
  width: 25%;
  height: 125px;
  text-align: center;
  align-items: center;
}
.tp-item-tab {
  text-align: center;
  cursor: pointer;
  color: #22ac95;
  width: 120px;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.tp-item-tab img {
  width: 22px;
  height: 22px;
  display: block;
  margin: 0 auto 10px;
}
.trx-condition-btns {
  display: flex;
}
.assets-condition-row {
  justify-content: space-between;
  /* margin: 0px 0 15px; */
  align-items: center;
  height: 85px;
}
.publish-assets {
  margin: 0 0 0 20px;
}
.assets-cnd .el-checkbox__input.is-checked + .el-checkbox__label {
  color: #fff;
}
.assets-cnd .el-checkbox {
  color: #fff;
}
.assets-cnd {
  padding: 10px 0;
}

/** wt-el-menu start*/
.wt-el-menu {
  width: 166px;
  background-color: rgba(247, 247, 247, 1);
  border-radius: 11px;
  box-shadow: 3px 4px 3px rgb(73 98 128);
  position: absolute;
  top: 110px;
  right: 6px;
  z-index: 999;
  display: none;
}
.tp-more-tab {
  position: relative;
}
.tp-more-tab:hover .wt-el-menu {
  display: block;
}
.menu-item {
  height: 40px;
  line-height: 40px;
  display: block;
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: #2dae95;
  text-align: center;
  border-bottom: 1px solid #ccc;
  position: relative;
}
.menu-item:nth-child(4) {
  border: 0;
}
.el-menu-triangle {
  width: 0;
  height: 0;
  border-left: 10px solid transparent;
  border-right: 10px solid transparent;
  border-bottom: 10px solid #fff;
  position: absolute;
  top: -9px;
  right: 19px;
}
.assets-row-item {
  padding: 0px 30px;
  height: auto;
}
/** end wt-el-menu*/
.edit-icon {
  margin: 0 10px 0 0;
  color: #63f7d4;
  font-size: 14px;
}
/**asset start */
.assets-nodata {
  min-height: 200px;
  text-align: center;
  line-height: 200px;
}
/**end asset */
</style>