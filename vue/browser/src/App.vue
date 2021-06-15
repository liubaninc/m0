<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
import "@/assets/style/common.css";
export default {
  name: "App",
  components: {},
};
</script>

<style>
/* 此处全局样式 */
.width1200 {
  margin: 0 auto;
  width: 1200px;
}
.breadcrumb {
  font-weight: 700;
  font-size: 16px;
  margin-bottom: 30px;
  margin-top: 20px;
}
.pagination {
  margin: 15px 0;
  text-align: right;
}
.content {
  padding-bottom: 60px;
}
.oneLine {
  width: 300px;
  display: inline-block;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  vertical-align: middle;
}
.oneLine180 {
  width: 180px !important;
}

.cm-word-split {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  display: inline-block;
  vertical-align: middle;
}
.split-width {
  width: 90% ;
}
.link{
    color: #409eff;
    cursor: pointer;
}

.trx-success{
  color: #409eff;
}
.trx-fail{
  color: #f56c6c;
}

.el-pagination{
  padding: 32px 5px 0!important;
  text-align: center!important;;
}
.Detail-message .title{
  font-size: 16px!important;
}
.trx-info .el-table{
  font-size: 16px!important;
}
.trx-info .el-table thead{
  color:#303133!important;
  font-weight: 400;
}

/* 全部头部信息 */
.Detail-message {
  font-size: 14px;
}
.Detail-message .title {
  font-size: 16px;
  margin: 15px 0;
}
.bg-purple-light {
  background: #fafafa;
}
.grid-content {
  border-radius: 4px;
  height: 50px;
  line-height: 50px;
  padding: 0 15px 0 30px;
  margin-right: 10px;
}
.grid-content:hover {
  background: #e6f7ff;
}
.grid-content .label {
  width: 150px;
  display: inline-block;
}
.grid-content .link {
  color: #409eff;
}
.grid-content .ml30 {
  margin-left: 30px;
}

</style>
