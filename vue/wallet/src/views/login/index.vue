<template>
  <div class="login">
    <div class="cm-width1200">
      <div class="login-header">
        <!-- <div class="login-header-logo">
          <img src="../../assets/images/logo/login_logo.png" />
        </div> -->
        <img
          class="login-header-logo"
          src="../../assets/images/logo/logo.png"
        />
        <div class="login-wallet-name">
          <span class="logo-name">{{ pageContext.logoName }}</span
          >登录
        </div>
      </div>
      <div class="login-main">
        <div class="login-main-warpper">
          <div class="main-wp-left">
            <div class="wp-left-title">
              <img src="../../assets/images/login/wl_icon.png" />
              {{ pageContext.walName }}
            </div>
            <div class="wp-left-context">
              {{ pageContext.description }}
            </div>

            <div class="main-m0-bg">
              <img src="../../assets/images/login/bg.png" alt="" srcset="" />
            </div>
          </div>
          <div class="main-wp-right">
            <div class="wp-right-title">
              欢迎登录 <span class="wp-name">{{ pageContext.logoName }}</span>
            </div>
            <div class="wp-right-form">
              <div class="form-row">
                <div class="right-form-name">用户名</div>
                <div class="right-form-input">
                  <el-input
                    v-model="username"
                    placeholder="请输入用户名"
                  ></el-input>
                </div>
              </div>
              <div class="form-row">
                <div class="right-form-name">密码</div>
                <div class="right-form-input">
                  <el-input
                    v-model="pwd"
                    placeholder="请输入密码"
                    show-password
                  ></el-input>
                </div>
              </div>
              <div class="form-row">
                <div class="right-form-name">验证码</div>
                <div class="right-form-input yzm-input">
                  <el-input v-model="yzm" placeholder="请输入验证码"></el-input>
                </div>
                <a
                  href="javascript:;"
                  class="right-form-yzm"
                  @click="refreshImg"
                >
                  <img :src="imgCode" alt="" srcset="" />
                </a>
              </div>
              <div class="form-row">
                <div class="right-form-name"></div>
                <div class="right-form-input">
                  <a
                    href="javascript:;"
                    class="cm-btn-322px cm-btn-bg009F72 login-btn"
                    @click="login"
                  >
                    登录
                  </a>
                </div>
              </div>
              <div class="form-row">
                <div class="right-form-name label-form-box"></div>
                <div class="login-input">
                  <div class="login-input-box">
                    <label>
                      <input type="checkbox" v-model="checked" />
                      {{ day }}天内自动登录
                    </label>
                  </div>
                  <a
                    href="javascript:;"
                    class="login-to-register"
                    @click="toRegister"
                    >立即注册</a
                  >
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { getImgCode, loginUser } from "@/server/user/login";
import { queryPageContext } from "@/server/pageConfig";
import { localCache } from "@/utils/utils";

export default {
  data() {
    return {
      username: "",
      pwd: "",
      yzm: "",
      imgCode: "",
      captchaId: "",
      checked: false,
      day: 3,
      pageContext: {},
    };
  },
  async created() {
    this.getyzmImg();
    let pageContext = await queryPageContext();
    if (pageContext) {
      this.pageContext = pageContext;
    }
  },
  methods: {
    toRegister() {
      this.$router.push("/register");
    },
    async login() {
      let { username, pwd, yzm } = this;
      if (!username) {
        this.$message.error("用户名为空");
        console.log("======用户名为空=====");
        return;
      }
      if (!pwd) {
        this.$message.error("密码为空");
        console.log("======密码为空=====");
        return;
      }
      if (!yzm) {
        this.$message.error("验证码为空");
        console.log("======验证码为空=====");
        return;
      }
      let exp_duration = this.checked ? this.day * 24 * 60 : 0;
      let loginRes = await loginUser.call(this, {
        name: username,
        password: pwd,
        captchaId: this.captchaId,
        captcha: yzm,
        exp_duration,
      });
      if (loginRes) {
        localCache.set("loginUser", loginRes);
        this.$router.push("/wallet");
      } else {
        this.refreshImg();
      }
    },
    refreshImg() {
      this.getyzmImg();
    },
    async getyzmImg() {
      let { imageUrl, captchaId } = await getImgCode();
      this.imgCode = imageUrl && imageUrl;
      this.captchaId = captchaId && captchaId;
    },
  },
};
</script>
<style scoped>
.login {
  height: calc(100vh);
  background: rgba(26, 44, 66, 1);
  min-width: 1200px;
}
.login-header {
  -moz-user-select: none;
  -khtml-user-select: none;
  user-select: none;
  height: 54px;
  display: flex;
  align-items: center;
  justify-items: center;

  font-family: "ArialMT", "Arial", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 16px;
  color: #bac0c6;
}
.login-main {
  padding: 12% 0 0;
}
.login-main-warpper {
  display: flex;
  width: 80%;
  margin: 0 auto;
  color: #768ca8;
  justify-content: space-around;
}
.main-wp-right {
  /* width: 898px; */
  width: 48%;
  padding: 36px 30px;
  background: #fff;
  border-radius: 16px;
  width: 400px;
  /* height: 288px; */
}
.wp-left-title {
  font-family: "PingFangSC-Medium", "PingFang SC Medium", "PingFang SC",
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 26px;
  color: #62f7d4;
}
.wp-left-title img {
  width: 40px;
  height: 40px;
  vertical-align: text-bottom;
}
.wp-left-context {
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #768ca8;
  width: 80%;
  line-height: 20px;
  margin: 30px 0 0;
}
.wp-right-title {
  font-size: 16px;
  font-weight: 700;
  color: rgb(51, 51, 51);
  margin: 0 0 28px;
  display: flex;
  align-items: center;
}
.form-row {
  display: flex;
  align-items: center;
  margin: 0 0 12px;
}
.form-row:last-child {
  margin-bottom: 0;
}
.right-form-input {
  width: 96%;
}
.wp-name {
  font-size: 22px;
  color: #2faf94;
  margin: 0 10px;
}
.right-form-name {
  width: 70px;
  font-family: "PingFangSC-Regular", "PingFang SC", sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #425263;
}
.login-btn {
  color: #fff;
  font-size: 14px;
  width: 100%;
}
.login-input {
  flex: 1;
  display: flex;
  align-items: center;
  align-content: space-between;
  justify-content: space-between;
}
.login-input-box {
  display: flex;
  align-items: center;
  font-size: 14px;
}
.login-input-box input {
  margin: 0 2px 0 0;
  vertical-align: bottom;
}
.yzm-input {
  width: 57%;
}
.right-form-yzm {
  width: 40%;
  height: 40px;
  margin: 0 0 0 10px;
  background: #ccc;
  display: block;
}
.right-form-yzm img {
  width: 100%;
  height: 100%;
}
.login-to-register {
  color: #009f72;
  font-size: 14px;
}

.main-wp-left {
  position: relative;
  width: 48%;
}
.main-m0-bg {
  position: absolute;
  left: -13%;
  top: 24%;
  transform: rotate(343deg);
  z-index: 1;
  /* opacity: 0.6; */
}
.main-m0-bg img {
  width: 604px;
  height: 398px;
}
.label-form-box {
  width: 60px;
}

.login-header-logo {
  width: 26px;
  height: 26px;
}
.login-wallet-name {
  margin: 0 0 0 10px;
}
.logo-name::after {
  content: "";
  display: inline-block;
  background: #bac0c6;
  width: 1px;
  height: 20px;
  vertical-align: sub;
  margin: 0 16px;
}
</style>