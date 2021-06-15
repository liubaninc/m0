<!-- 节点列表 -->
<template>
  <div class="node width1200">
    <div class="node-separator breadcrumb">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/index' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>节点</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="node-table">
      <el-table
        :data="tableData"
        :stripe="true"
        :border="true"
        :header-cell-style="{
          background: '#f0f0f0',
          color: 'rgba(0, 0, 0, 0.85)',
          fontWeight: 500,
          textAlign: 'center',
        }"
      >
        <el-table-column align="center" prop="peerIP" label="节点IP">
        </el-table-column>
        <el-table-column
          align="center"
          prop="peerAlais"
          label="节点名称"
          width="200"
        >
        </el-table-column>
        <el-table-column align="center" label="节点类型">
          <template slot-scope="scope">
            <el-tag v-if="scope.row.peerType == 0">共识节点</el-tag>
            <el-tag v-if="scope.row.peerType == 1" type="success"
              >数据节点</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          prop="peerStartIme"
          label="节点启动时间"
          width="280"
        >
        </el-table-column>
        <el-table-column align="center" label="节点状态">
          <template slot-scope="scope">
            <el-tag v-if="scope.row.peerStatus == 0" type="success"
              >正常</el-tag
            >
            <el-tag v-if="scope.row.peerStatus == 1" type="warning"
              >异常</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column align="center" prop="peerVersion" label="节点版本">
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
import { nodeList } from "@/server/node/list.js";
export default {
  methods: {},
  data() {
    return {
      tableData: [
        {
          peerIP: 10,
          peerAlais: "0xfdafjaifdjafdnafafa",
          peerType: 1,
          peerStartIme: "2020-12-12 12:14:12",
          peerStatus: 0,
          peerVersion: 10,
        },
        {
          peerIP: 10,
          peerAlais: "0xfdafjaifdjafdnafafa",
          peerType: 1,
          peerStartIme: "2020-12-12 12:14:12",
          peerStatus: 0,
          peerVersion: 10,
        },
        {
          peerIP: 10,
          peerAlais: "0xfdafjaifdjafdnafafa",
          peerType: 0,
          peerStartIme: "2020-12-12 12:14:12",
          peerStatus: 1,
          peerVersion: 10,
        },
      ],
      page: {
        total: 0, // 总页数
        currentPage: 1, // 当前页数
        pageSize: 10, // 每页显示多少条,
      },
    };
  },
  created() {
    this.getList(this.page);
  },
  methods: {
    // 切换第几页
    handleCurrentChange: function(page) {
      this.page.currentPage = page;
      this.getList(this.page);
    },
    // 每页几条
    handleSizeChange: function(sizeNum) {
      this.page.pageSize = sizeNum;
      this.getList(this.page);
    },
    // 分页获取数据
    getList(page, params) {
      let url = "/peers";
      let paramsData = Object.assign(
        {
          page_num: page.currentPage,
          page_size: page.pageSize,
        },
        params
      );
      nodeList(url, paramsData).then((response) => {
        this.tableData = response.peers;
        this.page.total = response.total;
      });
    },
  },
};
</script>
<style scoped></style>
