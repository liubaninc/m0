<template>
  <div class="assetsdetail">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item :to="{ path: '/assets' }"
          >资产管理</el-breadcrumb-item
        >
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >资产详情
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg detail-main">
        <div class="detail-main-wp">
          <div class="detail-main-assets">
            <div class="assets-title">我的账户资产</div>
            <div class="cm-submodule-bg assets-info">
              <div class="assets-info-row">
                <span class="row-name assets-row-name">我的资产地址 </span>
                <div class="cm-row-bg row-line">
                  {{ myAccountAssets.address }}
                </div>
              </div>
              <div class="assets-info-row">
                <span class="row-name assets-row-name">我的资产余额 </span>
                <div class="cm-row-bg row-line">{{ accountAssetBalance }}</div>
              </div>
            </div>
          </div>
          <div class="detail-main-assets">
            <div class="assets-title">资产发行信息</div>
            <div class="cm-submodule-bg assets-info">
              <div class="assets-info-row">
                <span class="row-name assets-row-name">资产ID </span>
                <el-tooltip
                  effect="dark"
                  :content="assetInfo.ref_tx"
                  placement="top"
                >
                  <div class="cm-text-overflow cm-row-bg row-line">
                    {{ assetInfo.ref_tx }}
                  </div>
                </el-tooltip>
              </div>
              <div class="assets-info-row">
                <span class="row-name assets-row-name">资产名称 </span>
                <div class="cm-row-bg row-line">{{ assetInfo.denom }}</div>
              </div>
              <div class="assets-info-row">
                <span class="row-name assets-row-name">发行总量 </span>
                <div class="cm-row-bg row-line">
                  {{ assetInfo.mint_amount * 1 + assetInfo.burn_amount * 1 }}
                </div>
              </div>
              <div class="assets-info-row">
                <span class="row-name assets-row-name">流通总量 </span>
                <div class="cm-row-bg row-line">
                  {{ assetInfo.mint_amount }}
                </div>
              </div>
              <div class="assets-info-row">
                <span class="row-name assets-row-name">资产发行人 </span>
                <div class="cm-row-bg row-line">
                  {{ assetInfo.initiator }}
                </div>
              </div>
              <div class="assets-info-row">
                <span class="row-name assets-row-name">首次发行时间 </span>
                <div class="cm-row-bg row-line">
                  {{ assetInfo.time | dateFormat("YYYY-mm-dd HH:MM:SS") }}
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="assets-detail-btns">
          <a
            href="javascript::"
            class="cm-btn-bg009F72 cm-btn-303px"
            @click="$router.go(-1)"
            >返回</a
          >
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { queryAssetInfo, queryAssetsByAddress } from "@/server/assets";
export default {
  data() {
    return {
      assetInfo: {},
      myAccountAssets: {},
    };
  },
  computed: {
    accountAssetBalance() {
      let balance = this.myAccountAssets;
      if (balance.coins && balance.coins.length) {
        return balance.coins[0].amount;
      }
      return "---";
    },
  },
  created() {
    let { denom, address } = this.$route.query;
    if (denom) {
      this.getAssetsInfoByName(denom);
      this.getMyAssetsInfo(address, denom);
    }
  },
  methods: {
    async getAssetsInfoByName(denom) {
      let assetInfo = await queryAssetInfo({ denom });
      if (assetInfo) {
        this.assetInfo = assetInfo;
      }
    },
    async getMyAssetsInfo(address, coin) {
      let myAccountAssets = await queryAssetsByAddress({ address, coin });
      this.myAccountAssets = myAccountAssets;
    },
  },
};
</script>
<style scoped>
.detail-warpper {
  margin: 0 auto;
  color: #fff;
}

/*detail start* */
.assets-title {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
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
  width: 120px;
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
}
.assetsdetail .row-line {
  min-height: 45px;
  width: 60%;
  /* display: flex; */
  /* align-items: center; */
  /* justify-content: space-between; */
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  /* padding: 10px 20px; */
  padding: 0 20px;
  line-height: 45px;
}
/*end detail* */

/**assetsdetail start */
.assets-row-name {
  width: 110px;
}

.assets-detail-btns {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 0 20px;
}
.assetsdetail .detail-main-assets {
  padding: 0 40px;
}
.detail-main-wp {
  padding: 30px 0;
}
/**end  assetsdetail*/
</style>