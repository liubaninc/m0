<!-- 交易列表 -->
<template>
  <div class="assets width1200">
    <div class="breadcrumb">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/index' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>交易</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="contract-table">
      <el-table
        :data="trxLists"
        :stripe="true"
        :border="true"
        :header-cell-style="{
          background: '#f0f0f0',
          color: 'rgba(0, 0, 0, 0.85)',
          fontWeight: 500,
          textAlign: 'center',
        }"
      >
        <el-table-column align="center" width="300" label="交易哈希">
          <template slot-scope="scope">
            <el-tooltip class="item" effect="dark" :content="scope.row.hash" placement="top-start"">
              <span class="cm-word-split split-width link" @click="toTradeDetail(scope.row.hash)">{{
              scope.row.hash
            }}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column align="center" width="300" label="交易时间">
          <template slot-scope="scope">
            {{
              scope.row.time
            }}
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          prop="assetIssueTotal"
          label="交易大小"
        >
        <template slot-scope="scope" >
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
        <el-table-column align="center" label="交易状态">
            <template slot-scope="scope">
              <!-- <span :class="scope.row.status?'trx-success':'trx-fail'">{{scope.row.status | trxStatus}}</span> -->
              <span>{{scope.row.status | trxStatus}}</span>
            </template>
        </el-table-column>
        <el-table-column align="center" label="所属区块">
          <template slot-scope="scope">
            <el-link
              type="primary"
              @click="toBlockDetail(scope.row.height)"
              >{{ scope.row.height }}</el-link
            >
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
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
      </el-pagination>
    </div>
  </div>
</template>

<script>
import {queryTxList} from '@server/transaction/queryTx'
export default {
  data() {
    return {
      page: {
        total: 0, 
        currentPage: 1,
        pageSize: 10
      },
      trxLists:[]
    };
  },
  created() {
    this.queryTrxLists();
  },
  methods: {
    async queryTrxLists(currentPage = 1,pageSize = 10){
      let {total,page_num,page_size,txs } = await queryTxList("/transactions",{
        page_num:currentPage,
        page_size: pageSize
      });

      this.page.total = total;
      this.page.currentPage = page_num;
      this.page.pageSize = page_size;
      this.trxLists = txs;
    },
    // 切换第几页
    handleCurrentChange(page) {
      this.queryTrxLists(page,this.page.pageSize)
    },
    // 每页几条
    handleSizeChange(sizeNum) {
      this.queryTrxLists(this.page.page,sizeNum)
    },
    toTradeDetail(name) {
      this.$router.push({
        path: "/block/tradeInfo",
        query: {
          name,
        },
      });
    },
    toBlockDetail(height){
      this.$router.push({
        path: "/block/info",
        query: {
          name: height,
        },
      });
    }
  },
};
</script>
<style scoped></style>
