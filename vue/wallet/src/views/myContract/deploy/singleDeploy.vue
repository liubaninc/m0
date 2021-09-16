<template>
  <div class="single">
    <div class="single-main">
      <span class="iconfont m0-chenggong success-icon"></span>
      <div class="single-main-desc">
        智能合约部署中，部署成功后会在合约列表内展示合约最新状态
      </div>
      <a href="javascript:;" class="cm-btn-bg4acb9b cm-btn-295px">确定</a>
    </div>
  </div>
</template>

<style scoped>
.single {
  display: flex;
  justify-content: center;
  height: calc(100vh - 117px);
  width: 50%;
  margin: 0 auto;
  overflow: hidden;
}
.single-main {
  display: flex;
  align-items: center;
  align-content: center;
  flex-direction: column;
  margin: 18vh auto 0;
}
.success-icon {
  color: #63f7d4;
  font-weight: 700;
  font-size: 40px;
}
.single-main-desc {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-size: 18px;
  color: #ffffff;
  margin: 20px 0 180px;
}
</style>