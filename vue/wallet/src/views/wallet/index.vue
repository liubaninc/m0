<template>
  <div class="wallet">
    <!-- <div class="cm-width1200"> -->
    <div class="wallet-warpper">
      <div class="cm-module-bg wallet-top">
        <div class="wallet-tp-left">
          <div class="cm-font-009F72 tp-left-name">M0区块链数字钱包</div>
          <div class="cm-text-overflow3 tp-left-desc">
            {{ pageContext.walletDesc }}
          </div>
          <a
            href="javascript:;"
            class="cm-btn-190px cm-btn-bg009F72 wallet-create-btn"
            @click="createWallet"
            >创建数字钱包</a
          >
        </div>
        <div class="wallet-tp-right">
          <img src="../../assets/images/wallet/wallet_icon.png" />
        </div>
      </div>
      <div class="wallet-list">
        <div class="wallet-list-title">
          <div class="title-left">
            <img src="../../assets/images/wallet/wallet_img.png" alt="" />
            <span>我的钱包 </span>
          </div>
          <div class="title-right">
            已有钱包?立即<a
              href="javascript:;"
              class="cm-font-009F72"
              @click="importWallet"
              >导入钱包</a
            >
          </div>
        </div>
        <template v-if="walletLists && walletLists.length">
          <div class="wallet-list-main">
            <div
              class="cm-module-bg list-main-card"
              v-for="wallet in walletLists"
              :key="wallet.name"
              @click="loadWalletInfo(wallet)"
            >
              <div class="cm-text-overflow card-name">{{ wallet.name }}</div>
              <template v-if="wallet.threshold > 0">
                <div class="card-tag">多签</div>
              </template>
              <template v-else>
                <div class="card-tag">单签</div>
              </template>
              <a
                href="javascript:;"
                class="cm-btn-160px cm-btn-bg009F72 loading-btn"
                >载入</a
              >
            </div>
          </div>
        </template>
        <template v-else>
          <div class="cm-module-bg wallet-no-data">暂无钱包</div>
        </template>
      </div>
      <!-- </div> -->
    </div>
  </div>
</template>
<script>
import { queryWalletLists } from "@/server/wallet";
import { localCache } from "@/utils/utils";
import { queryPageContext } from "@/server/pageConfig";

export default {
  data() {
    return {
      walletLists: [],
      page_num: 1,
      page_total: 0,
      page_size: 10,
      pageContext: {},
    };
  },
  async created() {
    this.getWallets(this.page_num, this.page_size);
    let pageContext = await queryPageContext();
    if (pageContext) {
      this.pageContext = pageContext;
    }
  },
  methods: {
    createWallet() {
      this.$router.push(`/wallet/createTypes`);
    },
    loadWalletInfo(wallet) {
      localCache.set("wallet", wallet);
      this.$router.push(`/assets`);
    },
    importWallet() {
      this.$router.push(`/wallet/importTypes`);
    },
    async getWallets(pageNum, pageSize) {
      let {
        accounts,
        page_num,
        page_size,
        page_total,
      } = await queryWalletLists({
        page_num: pageNum,
        page_size: pageSize,
      });
      this.walletLists = accounts;
      this.page_num = page_num;
      this.page_total = page_total;
      this.page_size = page_size;
    },
  },
};
</script>
<style scoped>
.wallet {
  background-color: #1b2c42;
  /* min-height: calc(100vh - 97px); */
  padding: 20px 0 0;
}
.wallet-warpper {
  width: 90%;
  margin: 0 auto;
  color: #fff;
}
.wallet-top {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-radius: 5px;
  /* padding: 30px 40px; */
  padding: 30px 50px 30px 60px;
}
.tp-left-name {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 24px;
}
.tp-left-desc {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: #cbcbcc;
  width: 70%;
  line-height: 20px;
  margin: 20px 0;
  height: 62px;
}
.wallet-tp-right img {
  width: 303px;
  height: 203px;
}
.wallet-create-btn {
  font-size: 14px;
}
.wallet-list {
  width: 100%;
  margin: 40px 0 0;
}
.wallet-list-main {
  padding: 20px 0 0;
  display: grid;
  grid-template-columns: repeat(3, 23%);
  /* grid-gap: calc((100% - 23% * 3) / 2); */
  column-gap: calc((100% - 23% * 3) / 2);
}

.wallet-list-title {
  display: flex;
  align-items: center;

  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 16px;
  color: #ffffff;
  justify-content: space-between;
  line-height: 50px;
}
.title-left img {
  width: 30px;
  height: 20px;
  vertical-align: sub;
  margin: 0 10px 0 0;
}
.title-right {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
}

.list-main-card {
  width: 274px;
  height: 196px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  cursor: pointer;
  border: 1px solid transparent;
  transition: all 0.6s ease-in-out;
  border-radius: 5px;

  margin: 0 0px 20px 0;
}
.list-main-card:hover {
  border-color: #62f7d4;
  transition: all 0.6s ease-in-out;
}
.card-name {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
  text-align: center;
  /* line-height: 16px; */
  width: 80%;
}
.card-tag {
  width: 67px;
  height: 25px;
  text-align: center;
  line-height: 25px;
  font-size: 10px;
  color: #ffffff;
  background: url("../../assets/images/wallet/tag_icon.png");
  background-size: 100%;
  background-repeat: no-repeat;
  margin: 20px 0 30px;
}
.loading-btn {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
}

.wallet-no-data {
  height: 300px;
  text-align: center;
  line-height: 300px;
  font-size: 14px;
}
.wallet-tp-left {
  width: 80%;
}
</style>