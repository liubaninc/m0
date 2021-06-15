<template>
  <div class="trxdetail">
    <div class="detail-warpper">
      <el-breadcrumb separator="/" class="wallet-breadcrumb">
        <el-breadcrumb-item
          :to="{ path: `/trx?address=${address}&coin=${coin}` }"
          >交易记录</el-breadcrumb-item
        >
        <el-breadcrumb-item class="breadcrumb-cur-page"
          >交易详情
        </el-breadcrumb-item>
      </el-breadcrumb>
      <div class="cm-module-bg detail-main">
        <div class="detail-main-assets">
          <div class="assets-title">资产信息</div>
          <div class="cm-submodule-bg assets-info">
            <div class="assets-info-row">
              <span class="row-name">交易哈希 </span>
              <div class="cm-row-bg row-line row-line-hash">
                <div class="cm-text-overflow trx-hash-row">
                  {{ trxInfo.hash }}
                </div>
                <span
                  class="iconfont m0-copy copy-btn"
                  v-clipboard:copy="trxInfo.hash"
                  v-clipboard:success="onCopy"
                  v-clipboard:error="onError"
                ></span>
              </div>
            </div>
            <div class="assets-info-row2">
              <div class="info-row2-line">
                <span class="row-name">资产名称 </span>
                <div class="cm-row-bg row2-line-context">
                  {{ trxInfo.assets }}
                </div>
              </div>
              <div class="info-row2-line">
                <span class="row-name">交易类型 </span>
                <div class="cm-row-bg row2-line-context">
                  {{ trxInfo.type }}
                </div>
              </div>
            </div>
            <div class="assets-info-row2">
              <div class="info-row2-line">
                <span class="row-name">交易状态 </span>
                <div class="cm-row-bg row2-line-context">
                  {{ trxInfo.height | trxStatus }}
                </div>
              </div>
              <div class="info-row2-line">
                <span class="row-name">交易大小 </span>
                <div class="cm-row-bg row2-line-context">
                  {{ trxInfo.size | formateSize }}
                </div>
              </div>
            </div>
            <div class="assets-info-row2">
              <div class="info-row2-line">
                <span class="row-name">交易确认数 </span>
                <div class="cm-row-bg row2-line-context">
                  {{ trxInfo.confirmed }}
                </div>
              </div>
              <div class="info-row2-line">
                <span class="row-name">交易时间 </span>
                <div class="cm-row-bg row2-line-context">
                  {{ trxInfo.time | dateFormat("YYYY-mm-dd HH:MM:SS") }}
                </div>
              </div>
            </div>
            <div class="assets-info-row2">
              <div class="info-row2-line">
                <span class="row-name">所属区块 </span>
                <div class="cm-row-bg row2-line-context">
                  {{ trxInfo.height }}
                  <span
                    class="iconfont m0-copy copy-btn"
                    v-clipboard:copy="trxInfo.height"
                    v-clipboard:success="onCopy"
                    v-clipboard:error="onError"
                  ></span>
                </div>
              </div>
            </div>
          </div>

          <div class="trx-detail-main">
            <div class="assets-title trx-assets-title">交易详情</div>
            <div class="trx-item">
              <template v-for="(mesg, index) in trxInfo.utxo_msgs">
                <el-collapse :key="index" v-model="activeNames">
                  <el-collapse-item
                    :title="'交易 ' + (index + 1)"
                    :name="'MSG' + (index + 1)"
                  >
                    <el-tabs v-model="activeName">
                      <el-tab-pane
                        label="区块链交易信息"
                        :name="'MSG' + (index + 1) + '_first'"
                      >
                        <div class="trx-fee">
                          手续费:
                          <template v-if="trxInfo.fee">
                            {{ trxInfo.fee | formateAmount }}
                          </template>
                          <template v-else>---</template>
                        </div>
                        <div class="trx-info">
                          <div class="trx-info-input">
                            <div class="trx-info-name">inputs</div>
                            <el-table
                              :data="mesg.inputs || []"
                              lazy
                              style="width: 100%"
                            >
                              <el-table-column label="address">
                                <template slot-scope="scope">
                                  <el-tooltip
                                    class="item"
                                    effect="dark"
                                    :content="scope.row.address"
                                    placement="top-start"
                                  >
                                    <span class="cm-text-overflow table-row">{{
                                      scope.row.address
                                    }}</span>
                                  </el-tooltip>
                                </template>
                              </el-table-column>
                              <el-table-column label="amount">
                                <template slot-scope="scope">
                                  {{ scope.row.amount | formateAmount }}
                                </template>
                              </el-table-column>
                            </el-table>
                          </div>
                          <div class="trx-info-input">
                            <div class="trx-info-name">outputs</div>
                            <el-table
                              :data="mesg.outputs || []"
                              style="width: 100%"
                            >
                              <el-table-column label="address">
                                <template slot-scope="scope">
                                  <el-tooltip
                                    class="item"
                                    effect="dark"
                                    :content="scope.row.address"
                                    placement="top-start"
                                  >
                                    <span class="cm-text-overflow table-row">{{
                                      scope.row.address
                                    }}</span>
                                  </el-tooltip>
                                </template>
                              </el-table-column>
                              <el-table-column label="amount">
                                <template slot-scope="scope">
                                  {{ scope.row.amount | formateAmount }}
                                </template>
                              </el-table-column>
                            </el-table>
                          </div>
                        </div>
                      </el-tab-pane>
                      <el-tab-pane
                        label="区块链交易请求"
                        :name="'MSG' + (index + 1) + '_second'"
                      >
                        <template v-if="mesg.requests && mesg.requests.length">
                          <template v-for="request in mesg.requests">
                            <json-viewer :value="request"></json-viewer>
                          </template>
                        </template>
                        <template v-else>
                          <span class="block-request">暂无交易请求</span>
                        </template>
                      </el-tab-pane>
                      <el-tab-pane
                        label="区块链交易备注信息"
                        :name="'MSG' + (index + 1) + '_third'"
                      >
                        <span class="block-request">
                          {{ mesg.desc || "暂无备注信息" }}</span
                        >
                      </el-tab-pane>
                    </el-tabs>
                  </el-collapse-item>
                </el-collapse>
              </template>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { queryTrxDetail } from "@/server/transaction";
export default {
  data() {
    return {
      trxInfo: {},
      activeNames: ["MSG1"],
      activeName: "MSG1_first",
      address: "",
      coin: "",
    };
  },
  created() {
    let { hash, address, coin } = this.$route.query;
    this.coin = coin;
    this.address = address;
    this.getTrxInfo(hash);
  },
  filters: {
    formateAmount(amount) {
      if (!amount) return amount;
      let amt = /^\d*/.exec(amount)[0];
      let names = amount.replace(/^\d*/, "$1").split("$1");
      return amt + (names.length ? " (" + names[1] + ")" : "");
    },
  },
  methods: {
    async getTrxInfo(hash) {
      let trxInfo = await queryTrxDetail({ hash });
      this.trxInfo = trxInfo;
    },
    onCopy(text) {
      if (text) {
        this.$message("复制成功");
      }
    },
    onError(e) {
      console.log(e);
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
  width: 80px;
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
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
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
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
  margin: 0 0 20px;
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

.row-line-hash {
}
</style>