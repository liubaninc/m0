<!-- 区块列表 -->
<template>
  <div class="assets width1200">
    <div class="breadcrumb">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/index' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>区块</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="contract-table">
      <el-table
        :data="blockLists"
        :stripe="true"
        :border="true"
        :header-cell-style="{
          background: '#f0f0f0',
          color: 'rgba(0, 0, 0, 0.85)',
          fontWeight: 500,
          textAlign: 'center',
        }"
      >
        <el-table-column align="center" width="124" label="区块高度">
          <template slot-scope="scope">
            <span class="link" @click="toDetail(scope.row.height)">{{
              scope.row.height
            }}</span>
          </template>
        </el-table-column>
        <el-table-column width="260" align="center" label="区块时间">
          <template slot-scope="scope">
            {{
              scope.row.time
            }}
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="区块大小"
          width="124"
        >
          <template  slot-scope="scope">
            <span>{{scope.row.size | formateSize}}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          prop="tx_num"
          label="交易数"
          width="124"
        >
        </el-table-column>
        <el-table-column align="center" label="区块提议人">
          <template slot-scope="scope">
            <el-tooltip class="item" effect="dark" :content="scope.row.proposer" placement="top-start">
              <span class="cm-word-split split-width">{{scope.row.proposer}}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column align="center" label="区块哈希">
          <template slot-scope="scope">
            <el-tooltip class="item" effect="dark" :content="scope.row.hash" placement="top-end">
              <span
                class="cm-word-split split-width link"
                @click="toDetail(scope.row.hash)"
                >{{ scope.row.hash }}</span
              >
            </el-tooltip>
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
import {queryBlocks} from '@server/block/queryBlock';
export default {
  data() {
    return {
      page: {
        total: 0, // 总页数
        currentPage: 1, // 当前页数
        pageSize: 10, // 每页显示多少条,
      },
      blockLists:[]
    };
  },
  created() {
    this.queryBlockLists();
  },
  methods: {
    async queryBlockLists(currentPage=1,pageSize=10){
      let {blocks,page_num,page_size,total} =  await queryBlocks('/blocks',{
        page_num:currentPage,
        page_size: pageSize
      });
      this.blockLists = blocks;
      this.page.total = total;
      this.page.currentPage = page_num;
      this.page.pageSize = page_size;
    },
    // 切换第几页
    handleCurrentChange(page) {
      this.queryBlockLists(page,this.page.pageSize);
    },
    // 每页几条
    handleSizeChange(sizeNum) {
      this.queryBlockLists(this.page.currentPage,sizeNum);
    },
    // 区块详情页
    toDetail(name) {
      this.$router.push({
        path: "/block/info",
        query: {
          name:name,
        },
      });
    },
  },
};
</script>
<style scoped>
</style>
