<!-- 交易详情 -->
<template>
  <div class="assetsDetail width1200">
    <div class="breadcrumb">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/index' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: '/block/tradeList' }"
          >交易</el-breadcrumb-item
        >
        <el-breadcrumb-item>交易详情</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="Detail-message">
      <div class="title">
        交易哈希 /
        <!-- <span class="cm-word-split hide-width">{{ trxInfo.hash }}</span> -->
        <span class="">{{ trxInfo.hash }}</span>
      </div>
      <el-row>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">交易时间</span>
            <span>{{ trxInfo.time }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">交易状态</span>
            <!-- <span :class="trxInfo.status ? 'trx-success' : 'trx-fail'">{{
              trxInfo.status | trxStatus
            }}</span> -->
            {{
              trxInfo.status | trxStatus
            }}
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content">
            <span class="label">交易大小</span>
            <span>{{ trxInfo.size | formateSize }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content">
            <span class="label">交易确认数</span>
            <span>{{ trxInfo.confirmed }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">交易数量</span>
            <span>{{ trxInfo.msg_num }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">所属区块</span>
            <span
              class="link"
              :title="trxInfo.height"
              @click="toDetail(trxInfo.height)"
              >{{ trxInfo.height }}</span
            >
            <i
              class="el-icon-document-copy link ml30"
              v-clipboard:copy="trxInfo.height"
              v-clipboard:success="onCopy"
              v-clipboard:error="onError"
            ></i>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="table-tlt"><i class="el-icon-s-operation"></i>交易详情</div>
    <div class="Detail-table trx-detail">
      <div class="trxd-message">
        <template v-for="(mesg, index) in trxInfo.utxo_msgs">
          <el-collapse
            :key="index"
            v-model="activeNames"
            @change="handleChange"
          >
            <el-collapse-item
              :title="'MSG' + (index + 1)"
              :name="'MSG' + (index + 1)"
            >
              <el-tabs v-model="activeName" @tab-click="handleClick">
                <el-tab-pane
                  label="区块链交易信息"
                  :name="'MSG' + (index + 1) + '_first'"
                >
                  <div class="trx-fee">
                    手续费: {{ trxInfo.fee | formateAmount }}
                  </div>
                  <div class="trx-info">
                    <div class="trx-info-input">
                      <div class="trx-info-name">
                        inputs
                      </div>
                      <el-table :data="mesg.inputs || []" style="width: 100%">
                        <el-table-column label="address">
                          <template slot-scope="scope">
                            <el-tooltip class="item" effect="dark" :content="scope.row.address" placement="top-start">
                         <span class="cm-word-split split-width">{{
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
                    <div class="trx-info-split"></div>
                    <div class="trx-info-input">
                      <div class="trx-info-name">
                        outputs
                      </div>
                      <el-table :data="mesg.outputs || []" style="width: 100%">
                        <el-table-column label="address">
                          <template slot-scope="scope">
                            <el-tooltip class="item" effect="dark" :content="scope.row.address" placement="top-start">
                              <span class="cm-word-split split-width">{{
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
                 <span class="block-request"> {{ mesg.desc || "暂无备注信息" }}</span>
                </el-tab-pane>
              </el-tabs>
            </el-collapse-item>
          </el-collapse>
        </template>
      </div>
    </div>
  </div>
</template>

<script>
import { queryTxList } from "@server/transaction/queryTx";

export default {
  data() {
    return {
      trxInfo: {},
      activeNames: ["MSG1"],
      activeName: "MSG1_first"
    };
  },
  created() {
    let { name: hash } = this.$route.query;
    this.queryTrxInfoByHash(hash);
  },
  filters: {
    formateAmount(amount) {
      if (!amount) return amount;
      let amt = /^\d*/.exec(amount)[0];
      let names = amount.replace(/^\d*/,"$1").split("$1");
      return amt + (names.length?" ("+names[1]+")":"");
    },
  },
  methods: {
    handleClick(tab, event) {
      // console.log(tab, event);
    },
    handleChange(val) {
      // console.log(val);
    },
    async queryTrxInfoByHash(hash) {
      let trxInfo =
        (hash && (await queryTxList(`/transactions/${hash}`))) || {};
      this.trxInfo = trxInfo;
    },

    toDetail(height) {
      this.$router.push({
        path: "/block/info",
        query: {
          name: height,
        },
      });
    },
    // 复制
    onCopy(e) {
      this.$message({
        message: `复制成功 : ${e.text}`,
      });
    },
    onError(e) {
      this.$message({
        message: `复制失败 : ${e.text}`,
      });
    },
  },
};
</script>
<style scoped>

.table-tlt {
  margin: 20px 0;
  font-size: 16px;
}
.table-tlt i {
  color: #409eff;
  font-weight: 700;
  margin-right: 10px;
}
.el-icon-document-copy {
  cursor: pointer;
}

.hide-width {
  width: 30%;
}
.trx-hash-spilt {
  width: 90%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  display: inline-block;
  vertical-align: middle;
}

.trx-detail {
  padding: 18px 0 0;
}
.trx-info {
  display: flex;
}
.trx-info-input {
  flex: 1;
}
.trx-info-split {
  width: 30px;
}
.trx-info-name {
  font-size: 22px;
}
.el-table td,
.el-table th.is-leaf {
  border-bottom: 0 !important;
}
/* .trx-info .el-table tr:last-child td{
} */
.trx-info .el-table::before {
  content: "";
  display: none !important;
}
.trx-fee {
  text-align: right;
  font-size: 16px;
  padding: 0 20px 0 0;
}

.table-tlt{
  margin-bottom:0 ;
}
.Detail-table .el-table{
  color: #1b1b1b!important;
}

.block-request{
  font-size: 14px;
}
</style>
