<template>
  <div class="download">
    <el-breadcrumb separator="/" class="wallet-breadcrumb">
      <el-breadcrumb-item :to="{ path: '/wallet' }"
        >钱包列表</el-breadcrumb-item
      >
      <el-breadcrumb-item class="breadcrumb-cur-page"
        >下载SDK
      </el-breadcrumb-item>
    </el-breadcrumb>
    <div class="cm-submodule-bg download-mian">
      <div class="download-title">下载中心</div>
      <div class="download-list">
        <div class="download-row">
          <img
            src="../../assets/images/download/sdk_icon.png"
            alt=""
            srcset=""
            class="row-icon"
          />
          <div class="row-desc">TK-SDK-Golang V1.0</div>
          <a
            href="javascript:;"
            class="cm-btn-225px cm-btn-bg4acb9b"
            @click="downGoSDK"
            >下载</a
          >
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { downLoadFile } from '@/utils/utils'
export default {
  methods: {
    downGoSDK() {
      downLoadFile(`${process.env.VUE_APP_PRO_BASE_URL}/mcontract/download/sdk`)
    },
  },
}
</script>
<style scoped>
.download-mian {
  height: calc(100vh - 157px);
}
.download-title {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-size: 14px;
  color: #62f7d4;
  width: 90%;
  margin: 0 auto;
  padding: 30px 0 50px;
}
.download-list {
}
.download-row {
  display: flex;
  align-items: center;
  padding: 0 30px 0 50px;
  background-color: rgba(255, 255, 255, 1);
  border-color: rgba(204, 204, 204, 1);
  border-radius: 3px;
  -moz-box-shadow: 5px 5px 5px rgba(188, 188, 188, 0.349019607843137);
  -webkit-box-shadow: 5px 5px 5px rgb(188 188 188 / 35%);
  box-shadow: 5px 5px 5px rgb(188 188 188 / 35%);
  width: 90%;
  margin: 0 auto;
}
.row-icon {
  width: 135px;
  height: 124px;
}
.row-desc {
  flex: 1;
  font-family: 'Arial-BoldMT', 'Arial Bold', 'Arial', sans-serif;
  font-weight: 700;
  font-size: 22px;
  color: #333333;
  padding: 10px 0 0;
}
</style>