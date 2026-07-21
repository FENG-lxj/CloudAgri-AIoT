<template>
  <div class="box">
    <div class="left_box">
      <div class="top">
        <img src="../img/Device_Management/logo.jpg" alt="" class="logo" />
        <div class="logo_text">螭耀科技</div>
        <div class="vertical_line"></div>
        <div class="logo_text_two">云端田园</div>
      </div>
      <img src="../img/login/left_img.png" alt="" class="left_img" />
    </div>
    <div class="right_box">
      <div class="login_text">登录</div>
      <div class="form">
        <div class="input">
          <input
            class="input__inner"
            type="text"
            placeholder="手机号"
            v-model="Phone"
            oninput="if(value.length > 11)value = value.slice(0, 11)"
          />
        </div>
        <div class="input">
          <input
            class="input__inner"
            type="number"
            placeholder="验证码"
            v-model="Captcha"
            oninput="if(value.length > 6)value = value.slice(0, 6)"
          />
        </div>
        <span class="Vertical">|</span>
        <button class="Send_Captcha">
          <p v-if="countdown == 0" @click="Send_Captcha_btn">发送验证码</p>
          <p v-if="countdown != 0">{{ countdown }}秒后重试</p>
        </button>
        <button class="btn" @click="submit_btn">登录</button>
      </div>
    </div>
  </div>
</template>
<script>
import axios from "axios";
export default {
  data() {
    return {
      //Phone: "18607229637",
      Phone: "17548750710",
      Captcha: null,
      // Captcha: "123456",
      countdown: 0,
    };
  },
  methods: {
    startCountdown() {
      this.timer = setInterval(() => {
        if (this.countdown > 0) {
          this.countdown--;
          // console.log(this.countdown);
        }
      }, 1000);
    },


    // 发送验证码
    Send_Captcha_btn() {
      this.countdown = 60;
      this.startCountdown();
      axios
        .post("http://127.0.0.1:20201/IoTA/users/captcha", {
          phone: this.Phone,
        })
        .then((res) => {
          this.Captcha = res.data.data.code;
          console.log("success:", res.data.code);
          this.$store.commit("setPhone", this.Phone);
          if (res.data.code == 200) {
            this.Captcha_success();
          } else {
            this.Captcha_error();
          }
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
          this.Captcha_error();
        });
    },

    submit_btn() {
      this.startCountdown();
      axios
        .post("http://127.0.0.1:20201/IoTA/users/login", {
          phone: this.Phone,
          phoneCaptcha: this.Captcha,
        })
        .then((res) => {
          this.$store.commit("setphoneCaptcha", res.data.data.token);
          if (res.data.code == 200) {
            this.submit_success();
            this.countdown = 0;
            this.$router.replace("home");
          } else {
            this.submit_error();
          }
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
          this.submit_error();
        });
    },


    Captcha_success() {
      this.$notify({
        title: "验证码请求成功",
        message: this.Captcha,
        type: "success",
      });
    },
    Captcha_error() {
      this.$notify.error({
        title: "验证码请求失败",
      });
    },
    submit_success() {
      this.$notify({
        title: "登录成功",
        type: "success",
      });
    },
    submit_error() {
      this.$notify.error({
        title: "登录失败",
        message: "账号或验证码错误。请重试!",
      });
    },
  },
  // mounted() {
  //   this.fetchData(); // 在组件挂载后调用fetchData方法
  // },
};
</script>
<style scoped>
.box {
  position: fixed;
  height: 100%;
  width: 100%;
  top: 0;
  left: 0;
}
.top {
  position: fixed;
  /* background: #000; */
  height: 5vw;
  width: 100%;
}
.logo {
  position: absolute;
  border-radius: 15vw;
  height: 3.5vw;
  width: 3.5vw;
  top: 0.8vw;
  left: 0.8vw;
  cursor: pointer;
}
.logo_text {
  position: absolute;
  left: 5vw;
  top: 1.5vw;
  color: white;
  font-size: 1.6vw;
  cursor: pointer;
}
.logo_text_two {
  position: absolute;
  left: 12.7vw;
  top: 1.5vw;
  color: white;
  font-size: 1.6vw;
  cursor: pointer;
}
.vertical_line {
  position: absolute;
  border-left: 1.5px solid rgb(255, 255, 255);
  height: 2vw;
  left: 12vw;
  top: 1.7vw;
}
.left_box {
  position: absolute;
  width: 73vw;
  height: 100vh;
  background: #495ad4;
}
.left_img {
  position: absolute;
  top: 2vw;
  left: 7vw;
  width: 45vw;
  height: 45vw;
}
.right_box {
  position: absolute;
  width: 26vw;
  height: 50vh;
  left: 60vw;
  top: 26vh;
  border-radius: 1vw;
  box-shadow: 0px 6px 20px #202547;
  background: #fafafa;
}
.right_box p {
  position: absolute;
  color: #495ad4;
  left: 4vw;
  font-size: 2.3vw;
}
.login_text {
  font-family: "钉钉进步体 Regular", sans-serif;
  font-size: 3vw;
  color: #495ad4;
  text-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2); /* 添加阴影效果 */
  margin: 35px 0 0 40px;
}
.form {
  /* border: 5px solid #55ff00; */
  width: 21vw;
  height: 20vw;
  margin: 25px auto 0px auto;
}

.input {
  display: flex;
  height: 60px;
  margin: 20px auto 0 auto;
  align-items: center;
  border-radius: 20px;
  padding: 0 20px;
  background-color: #f9f9f9;
  box-shadow: inset 1px 1px 7px rgba(95, 95, 95, 0.3);
  transition: 300ms ease-in-out;
}

.input:focus {
  background-color: white;
  transform: scale(1.05);
  box-shadow: 13px 13px 100px #c2c2c2, -13px -13px 100px #ffffff;
}

.input__inner {
  flex: 1;
  border: none;
  outline: none;
  width: 100%;
  background-color: transparent;
  font-size: 30px;
  line-height: 30px;
  color: #333;
  -moz-appearance: textfield; /* Firefox */
  appearance: textfield; /* Chrome */
}
/* 隐藏输入框右边的上下调整箭头 */
.input__inner::-webkit-inner-spin-button,
.input__inner::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
.input__inner::-webkit-outer-spin-button,
.input__inner::-webkit-inner-spin-button {
  display: none;
}
.btn {
  margin: 30px auto 0px auto;
  height: 50px;
  width: 80%;
  color: #fff;
  background-color: #495ad4;
  border-radius: 20px;
  text-align: center; /* 水平居中 */
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
  padding: 10px 20px; /* 按钮的内边距 */
  font-size: 2vw;
  box-shadow: 2px 3px 3px #65656582;
  border: none; /* 去掉边框 */
}
.btn:active {
  background-color: #7284d4;
  border-color: #495ad4;
}
.Vertical {
  position: absolute;
  left: 13.4vw;
  top: 13.5vw;
  font-size: 4vh;
  font-weight: 100;
  color: #c7bdbd;
  text-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2); /* 添加阴影效果 */
}
.Send_Captcha {
  position: absolute;
  left: 10.3vw;
  top: 12.2vw;
  width: 24vw;
  border: none;
  padding: 0;
  margin: 0;
  font: inherit;
  color: inherit;
  background-color: transparent;
  cursor: pointer;
  text-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2); /* 添加阴影效果 */
}
.Send_Captcha p {
  color: #495ad4;
  font-size: 1.7vw;
}
.Send_Captcha p:active {
  color: #4959d4bd;
}
</style>