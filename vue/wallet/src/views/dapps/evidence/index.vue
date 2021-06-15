<template>
  <div class="dapps">
    <div class="dapps-container">
      <div class="com-breadcrumb dapps-con-title">存证服务</div>
      <div class="dapps-con-btns">
        <a
          href="javascript:;"
          class="cm-btn-bg009F72 cm-btn-138px"
          @click="toUpEvidence"
          >上传存证</a
        >
      </div>
      <div class="dapps-con-list">
        <template v-if="claims && claims.length">
          <template v-for="(clm, index) in claims">
            <div
              class="cm-module-bg con-list-row"
              :key="index"
              @click="toDetail(clm)"
            >
              <div class="list-row-left">
                <p class="row-left-title">{{ clm.name }}</p>
                <el-tooltip effect="dark" :content="clm.hash" placement="top">
                  <p class="cm-text-overflow row-left-trxhash">
                    交易哈希：{{ clm.hash }}
                  </p>
                </el-tooltip>
              </div>
              <div class="list-row-right">
                <img
                  src="../../../assets/images/detail/dapps/cz_icon.png"
                  alt=""
                  srcset=""
                />
                存证详情
              </div>
            </div>
          </template>
        </template>
        <template v-else>
          <div class="cm-submodule-bg con-list-nodata">暂无存证数据</div>
        </template>

        <div v-if="claims && claims.length" class="pagination-main">
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
  </div>
</template>
<script>
import { queryEvidenceList } from "@/server/dapps/evidence";
import { localCache } from "@/utils/utils";

export default {
  data() {
    return {
      wallet: {},
      claims: [],
      page: {
        pageNum: 1,
        pageSize: 10,
        pageTotal: 10,
      },
    };
  },
  created() {
    let wallet = localCache.get("wallet");
    let { page } = this;
    if (wallet) {
      this.wallet = wallet;
      this.getEviLists(wallet.name, page.pageNum, page.pageSize);
    }
  },
  methods: {
    handleSizeChange(pageSize) {
      this.getEviLists(this.wallet.name, this.page.pageNum, pageSize);
    },
    handleCurrentChange(pageNum) {
      this.getEviLists(this.wallet.name, pageNum, this.page.pageSize);
    },
    toUpEvidence() {
      this.$router.push(`/dapps/evidence/singleEvidence`);
    },
    toDetail(claim) {
      this.$router.push(`/dapps/evidence/detail?name=${claim.name}`);
    },
    async getEviLists(account, pageNum, pageSize) {
      let { claims, page_num, page_size, total } = await queryEvidenceList({
        account,
        page_num: pageNum,
        page_size: pageSize,
      });
      this.claims = claims;
      this.page.pageNum = page_num;
      this.page.pageSize = page_size;
      this.page.pageTotal = total;
    },
  },
};
</script>

<style>
.dapps-container {
  /* padding: 20px 30px; */
}
.dapps-con-title {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
}
.dapps-con-btns {
  margin: 10px 0 40px;
  display: flex;
  justify-content: flex-end;
}
.dapps-con-btns a {
  font-size: 14px;
}
.con-list-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 125px;
  padding: 0 30px;
  border-radius: 5px;
  cursor: pointer;
  margin: 0 0 20px;
}
.row-left-title {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 20px;
  color: #ffffff;
}
.row-left-trxhash {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #768ca8;
  margin: 30px 0 0;
  width: 500px;
}
.list-row-right {
  width: 60px;
  text-align: center;
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #62f7d4;
  cursor: pointer;
}
.list-row-right img {
  width: 20px;
  height: 20px;
  display: block;
  margin: 0 auto 10px;
}
.con-list-nodata {
  min-height: 200px;
  text-align: center;
  line-height: 200px;
  width: 100%;
  color: #768ca8;
  font-size: 16px;
}
</style>