<template>
  <div class="commonHeader">
    <router-link to="/wallet">
      <!-- <img
        class="cm-header-logo"
        src="../../assets/images/logo/wallet_logo.png"
      /> -->
      <div class="com-logo">
        <img class="cm-logo-icon" src="../../assets/images/logo/logo.png" />
        <span class="cm-logo-text">M0 Wallet</span>
      </div>
    </router-link>
    <div class="cm-header-user">
      <div>
        <a href="javascript:;" @click="$router.push(`/download`)">下载中心</a>
      </div>
      <div class="header-user-item">
        <!-- <span class="user-name"
          > -->
        <i class="iconfont m0-denglu user-icon"></i>{{ user.name }}
        <!-- </span> -->
        <div class="header-options">
          <a href="javascript:;" class="options-item" @click="loginOut">
            <i class="iconfont m0-tuichu login-out-icon"></i>退出
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { layoutUser } from '@/server/user/login'
import { localCache } from '@/utils/utils'
export default {
  data() {
    return {
      user: {},
    }
  },
  created() {
    let user = localCache.get('loginUser')
    if (user) {
      this.user = user
    }
  },
  methods: {
    async loginOut() {
      let outRes = await layoutUser({})
      localCache.remove('authorization')
      localCache.remove('loginUser')
      this.$router.replace('/login')
    },
  },
}
</script>
<style scoped>
.commonHeader {
  padding: 0;
  /* width: 1200px; */
  width: 98%;
  margin: 0 auto;
  display: flex;
  justify-items: center;
  align-items: center;
  height: 100%;
}
.cm-header-logo {
  width: 140px;
  height: 36px;
}
.cm-header-user {
  width: 90%;
  color: #fff;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: flex-end;
}

.header-user-item {
  width: 136px;
  height: 100%;
  position: relative;
  text-align: center;
  transition: all ease-in-out 0.6s;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}
.user-name {
  height: 100%;
  display: block;
  /* height: 100%;
  display: flex;
  align-items: center; */
}
.header-options {
  width: 136px;
  position: absolute;
  top: 53px;
  z-index: 299;
  background-color: rgba(247, 247, 247, 1);
  box-shadow: 3px 4px 3px rgb(73 98 128);
  font-size: 12px;
  display: none;
  transition: all ease-in-out 0.6s;
}

.options-item {
  display: block;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  color: #009f72;
  height: 40px;
  line-height: 40px;
}
.header-user-item:hover .header-options {
  display: block;
  transition: all ease-in-out 0.6s;
}

.user-icon {
  margin: 0 10px 0 0;
}
.commonHeader .commonHeader-a {
  display: flex;
  align-items: center;
  width: 150px;
}

.login-out-icon {
  vertical-align: bottom;
  margin: 0 6px 0 0;
}

.com-logo {
  display: flex;
  align-items: center;
  width: 150px;
  -moz-user-select: none;
  -khtml-user-select: none;
  user-select: none;
}
.cm-logo-icon {
  width: 26px;
  height: 26px;
}
.cm-logo-text {
  font-family: 'ArialMT', 'Arial', sans-serif;
  font-weight: 400;
  font-size: 16px;
  color: #ffffff;
  margin-left: 16px;
}
</style>
