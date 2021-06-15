<!-- 区块详情 -->
<template>
  <div class="assetsDetail width1200">
    <div class="breadcrumb">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/index' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: '/block/list' }">区块</el-breadcrumb-item>
        <el-breadcrumb-item>区块详情</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="Detail-message">
      <div class="title">
        区块高度 / {{blockInfo.height}}
      </div>
      <el-row>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">区块时间</span>
            <span>{{blockInfo.time}}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">区块哈希</span>
            <span class="cm-word-split cm-block-split">{{blockInfo.hash}}</span>
            <i
                class="el-icon-document-copy link ml30"
                v-clipboard:copy="blockInfo.hash"
                v-clipboard:success="onCopy"
                v-clipboard:error="onError"
              ></i>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content">
            <span class="label">区块大小</span>
            <span>{{blockInfo.size | formateSize}}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content">
            <span class="label">区块提议人</span>
            <span class="cm-word-split cm-block-split">{{blockInfo.proposer}}
            </span>
             <i
                class="el-icon-document-copy link ml30"
                v-clipboard:copy="blockInfo.proposer"
                v-clipboard:success="onCopy"
                v-clipboard:error="onError"
              ></i>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">交易数量</span>
            <span>{{blockInfo.tx_num}}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">前向区块链哈希</span>
            <span class="cm-word-split cm-block-split link"  @click="prevBlock(blockInfo.prev_hash)">{{ blockInfo.prev_hash }}</span>
            <i
              class="el-icon-document-copy link ml30"
              v-clipboard:copy="blockInfo.prev_hash"
              v-clipboard:success="onCopy"
              v-clipboard:error="onError"
            ></i>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="table-tlt">
      <i class="el-icon-s-operation"></i>交易列表
    </div>
    <div class="Detail-table">
      <el-table
        :data="blockInfo.txs"
        :stripe="true"
        :border="true"
        :header-cell-style="{
          background: '#f0f0f0',
          color: 'rgba(0, 0, 0, 0.85)',
          fontWeight: 500,
          textAlign: 'center',
        }"
      >
        <el-table-column align="center" label="交易哈希">
          <template slot-scope="scope">
            <el-tooltip class="item" effect="dark" :content="scope.row.hash" placement="top-start">
              <span class="trx-hash-split link" @click="toDetail(scope.row.hash)">{{
              scope.row.hash
            }}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column align="center" prop="time" label="交易时间">
        </el-table-column>
        <el-table-column
          align="center"
          label="交易大小"
        >
        <template slot-scope="scope">
            {{scope.row.size | formateSize}}
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="交易类型"
        >
         <template slot-scope="scope">
            {{scope.row.type}}
          </template>
        </el-table-column>
      </el-table>
      <!-- <el-pagination
        class="pagination"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="page.currentPage"
        :page-sizes="[10, 20, 30, 40]"
        :page-size="page.pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="page.total"
        background
      >
      </el-pagination> -->
    </div>
  </div>
</template>

<script>
import {queryBlocks} from '@server/block/queryBlock';
import {queryTxList} from '@server/transaction/queryTx';
export default {
  data() {
    return {
      blockInfo: {},
      trxLists:[],
      page: {
        total: 0, // 总页数
        currentPage: 1, // 当前页数
        pageSize: 10, // 每页显示多少条,
      },
    };
  },
  async created() {
    let { name } = this.$route.query;
    this.queryBlockInfo(name);
    // this.queryTrxLists();
  },
  methods: {
    async prevBlock(hash){
      let blockInfo = await this.queryBlockInfo(hash);
      this.replaceUrlParamVal('name',blockInfo.height)
    },
    toAssessDetail(hash){
       this.$router.push({
        path: "/assets/detail",
        query: {
          name: hash,
        },
      });
    },
    // 切换第几页
    handleCurrentChange(page) {
      // this.queryTrxLists(page,this.page.pageSize);
    },
    // 每页几条
    handleSizeChange(sizeNum) {
      // this.queryTrxLists(this.page.currentPage, sizeNum);
    },
    toDetail(name) {
      this.$router.push({
        path: "/block/tradeInfo",
        query: {
          name,
        },
      });
    },
    async queryBlockInfo(height){
      let blockInfo = height && await queryBlocks(`/blocks/${height}`) || {};
      // this.page.total = blockInfo.total;
      // this.page.currentPage = blockInfo.page_num;
      // this.page.pageSize = blockInfo.page_size;
      // this.trxLists = blockInfo.txs;
      this.blockInfo = blockInfo;
      return blockInfo;
    },
    async queryTrxLists(curPageNum=1,pageSize=10){
      let {page_num,page_size,page_total,txs,total} = curPageNum && await queryTxList('/transactions',{page_num:curPageNum,page_size: pageSize}) || {}
      this.page.total = total;
      this.page.currentPage = page_num;
      this.page.pageSize = page_size;
      this.trxLists = txs;
    },
    onCopy: function(e) {
      this.$message({
        message: `复制成功`
      });
      // console.log("You just copied: " + e.text);
    },
    onError: function(e) {
      this.$message({
        message: '复制失败'
      })
      // console.log("Failed to copy texts");
    },
    replaceUrlParamVal(paramName,replaceWith) {
      var oUrl = window.location.href.toString();
      var re = eval('/('+ paramName+'=)([^&]*)/gi');
      var nUrl = oUrl.replace(re,paramName+'='+replaceWith);
    　window.location.href=nUrl
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
.trx-hash-split{
  width: 90%;
   overflow: hidden;
  text-overflow:ellipsis;
  white-space: nowrap;
  display: inline-block;
  vertical-align: middle;
}
.cm-block-split{
  width: 50%;
}

</style>
