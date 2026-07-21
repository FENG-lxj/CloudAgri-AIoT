<template>
  <div class="home">
    <div class="top" @mouseenter="UserBoxMouseOutTrigger">
      <img src="../img/Device_Management/logo.jpg" alt="" class="logo" @click="home"/>
      <div class="logo_text" @click="home">螭耀科技</div>
      <div class="vertical_line"></div>
      <div class="logo_text_two" @click="home">云端田园</div>
      <div class="Device_management" @click="Device_management_Btn">
        设备管理
      </div>
      <div class="Userimg_left"></div>
      <img
          src="../img/SuperAdmin/AI.jpg"
          class="UserImg"
          @mouseenter="Hoveringtrigger"
      />
    </div>
    <hr class="xian"/>
    <div
        v-if="showBox"
        class="UserBox"
        id="UserBox"
        @mouseenter="Hoveringtrigger"
        @mouseleave="UserBoxMouseOutTrigger"
    >
      <div class="UserBoxPro">
        <img
            src="../img/SuperAdmin/AI.jpg"
            alt=""
            class="UserBoxImg"
        />
        <div class="UserText">
          <div class="UserBoxUser">{{this.$store.state.Phone}}</div>
          <div class="UserBoxUserJi">超级管理员</div>
        </div>
      </div>
      <hr style="margin: 0;"/>
      <div class="Super_Management" @click="SuperAdmin">超级管理面板</div>
      <hr style="margin: 0;"/>
      <div class="ExitPro" id="Exit" @click="Exit">退出登录</div>
      <div style="position: relative; height: 0.7vw"></div>
    </div>
    <div class="Gradient">
      <img
          src="../img/980.jpeg"
          alt=""
          class="Gradient_Img"
      />
      <div class="GradientImg"></div>
    </div>
    <div class="Carousel">
      <el-carousel height="399px" class="yuan">
        <el-carousel-item v-for="(item, index) in carouselItems" :key="index">
          <img
              :src="item"
              alt="Carousel Image"
              @click="CarouselBtn(index)"
              class="Carousel_Img"
          />
        </el-carousel-item>
      </el-carousel>
    </div>
    <div class="horizontal-line">AIOT农业生态智能化管控系统</div>
    <div class="xian2"></div>
    <div class="FeatureCards">
      <div class="FarmCard" @click="FarmCardBtn"></div>
      <div class="pastureCard" @click="pastureCardBtn"></div>
      <div class="PredictionCard" @click="PredictionCardBtn"></div>
      <div class="AI_surveillance_Card" @click="AICardBtn"></div>
      <div class="AICard" @click="show_AICard_Modal"></div>
    </div>
    <div class="bottom">
      联系我们：XXXXXXXXXX
      <br /><br />
      备案号：XXXXXXXXXXXXXX
      <br /><br />
      版权所有 © 螭耀科技团队
    </div>
    <AI_Agro></AI_Agro>
  </div>
</template>

<script>
import AI_Agro from "@/AI_Agro/AI_Agro.vue";
export default {
  components: {
    AI_Agro,
  },
  data() {
    return {
      showBox: false,
      carouselItems: [
        require("@/img/Round-robin_diagram/agro.png"),
        require("@/img/Round-robin_diagram/Crop_storage.png"),
        require("@/img/Round-robin_diagram/AI_surveillance.png"),
        require("@/img/Round-robin_diagram/AI.png"),
        require("@/img/Round-robin_diagram/true.jpg"),
      ],
      isFlipped: true,
    };
  },
  methods: {
    home() {},
    show_AICard_Modal() {
      this.$store.commit("set_isModalVisible", true);
      document.body.style.overflow = "hidden";
    },
    closeModal() {
      this.isModalVisible = false;
      document.body.style.overflow = "auto";
    },
    school() {
      window.location.href = "https://www.imut.edu.cn/";
    },
    rotateCard() {
      this.isFlipped = true;
    },
    resetRotation() {
      this.isFlipped = false;
    },
    Hoveringtrigger() {
      this.showBox = true;
    },
    UserBoxMouseOutTrigger() {
      this.showBox = false;
    },
    Exit() {
      this.$router.push("/login");
    },
    CarouselBtn(index) {
      if (index == 0) {
        this.$router.push("/Agro_ecological");
      } else if (index == 1) {
        this.$router.push("/Storage_Main");
      } else if (index == 2) {
        this.$router.push("/AI_Vision_Main");
      } else if (index == 3) {
        this.$store.commit("set_isModalVisible", true);
        document.body.style.overflow = "hidden";
      } else if (index == 4) {
        this.$router.push("/YuceHome");
      }
    },
    Device_management_Btn() {
      this.$router.push("/Device_Management_Page");
    },
    FarmCardBtn() {
      this.$router.push("/Agro_ecological");
    },
    pastureCardBtn() {
      this.$router.push("/Storage_Main");
    },
    AICardBtn() {
      this.$router.push("/AI_Vision_Main");
    },
    PredictionCardBtn() {
      // this.$router.push("/ProductionForecasts");
      this.$router.push("/YuceHome");
    },
    SuperAdmin() {
      this.$router.push("/SuperAdmin");
    },
  },
};
</script>

<style scoped>
/* 遮罩层样式 */
.modal {
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.7); /* 半透明黑色背景 */
  z-index: 1000; /* 设置比其他元素更高的 z-index，使其位于其他元素之上 */
}

/* 内容区域样式 */
.modal-content {
  background: #fff; /* 白色背景 */
  padding: 20px;
  border-radius: 5px;
  position: relative;
}

/* 关闭按钮样式 */
.close {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 20px;
  cursor: pointer;
}
.home {
  position: fixed;
  height: 100%;
  width: 100%;
  top: 0;
  left: 0;
  overflow: auto;
  background: #f4f4f4;
}
.UserBox {
  display: block;
  position: absolute;
  margin-left: 82vw;
  width: 17vw;
  background: rgb(236, 236, 236);
  height: 12.3vw;
  z-index: 999;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  /* border-radius: 1vw; */
}
/* #UserBox{
  opacity: 0;
  transform: translateY(-20px);
  transition: opacity 0.3s ease-in-out, transform 0.3s ease-in-out;
} */
/* #UserBox:hover {
  opacity: 1;
  transform: translateY(0);
} */
.UserBoxPro{
  display: flex;
}
.UserBoxUser {
  position: relative;
  font-weight: bold;
  font-size: 1.3vw;
  margin: 14px 0 0 2px;
}
.UserBoxUserJi {
  position: relative;
  font-weight: bold;
  font-size: 1vw;
  margin: 0 0 0 4px;
}
.UserBoxImg {
  position: relative;
  width: 3.5vw;
  height: 3.5vw;
  border-radius: 3vw;
  margin: 10px 10px 10px 20px;
}
.Super_Management {
  height: 3.6vw;
  width: 17vw;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  cursor: pointer;
}
.Super_Management:hover {
  background-color: rgb(219, 219, 219); /* 在鼠标悬停时将颜色改变为红色 */
}
.ExitPro{
  height: 3.6vw;
  width: 17vw;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  cursor: pointer;
}
.ExitPro:hover {
  background-color: rgb(219, 219, 219); /* 在鼠标悬停时将颜色改变为红色 */
}
.top {
  height: 3.9vw;
  background: rgb(41, 152, 255);
}
.logo {
  position: absolute;
  border-radius: 15vw;
  height: 3vw;
  width: 3vw;
  top: 0.5vw;
  left: 0.8vw;
  cursor: pointer;
}
.logo_text {
  position: absolute;
  left: 4.5vw;
  top: 1.1vw;
  color: rgb(0, 0, 0);;
  font-size: 20px;
  cursor: pointer;
}
.logo_text_two {
  position: absolute;
  left: 10.6vw;
  top: 1.1vw;
  color: rgb(0, 0, 0);;
  font-size: 20px;
  cursor: pointer;
}
.vertical_line {
  position: absolute;
  border-left: 1px solid rgb(60, 60, 60);
  height: 2vw;
  left: 10.2vw;
  top: 1vw;
}
.xian{
  margin: 0;
  border: none; /* 清除默认的边框样式 */
  border-bottom: 1px solid rgb(60, 60, 60); /* 设置横线颜色为红色 */
}
.xian2{
  width: 75%; /* 设置线的宽度 */
  margin: 0 auto; /* 水平居中 */
  border-bottom: 2px solid rgb(60, 60, 60); /* 设置线的样式和颜色 */
  margin-top: 80px; /* 设置线与上方文字的距离 */
  box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2); /* 添加阴影效果 */
}
/* .logo {
  position: absolute;
  top: 0.2vw;
  left: 1vw;
  width: 3vw;
  border-radius: 2vw;
  cursor: pointer;
}
.logoName {
  position: absolute;
  top: 1vw;
  left: 4.5vw;
  cursor: pointer;
} */
.Device_management {
  position: absolute;
  left: 87vw;
  top: 1.1vw;
  font-size: 20px;
  /* font-weight: 800; */
  cursor: pointer;
}
.Userimg_left {
  position: absolute;
  border-left: 0.5px solid rgb(60, 60, 60);;
  height: 3.9vw;
  left: 93.5vw;
}
.UserImg {
  position: absolute;
  width: 3.3vw;
  left: 94.6vw;
  border-radius: 2vw;
  top: 0.3vw;
  cursor: pointer;
}
.phone {
  position: absolute;
  left: 91.5vw;
  top: 1.7vw;
}
.Gradient {
  height: 7vw;

  /* width: 100vw; */
  /* background: linear-gradient(to bottom, rgba(0, 0, 255, 1), rgba(0, 0, 255, 0)); */
}
.Gradient_Img {
  position: absolute;
  width: 100%;
  height: 18vw;
  z-index: -999;
}
.GradientImg {
  height: 18vw;
  width: 100%;
  background: linear-gradient(to bottom, rgb(41, 152, 255), transparent);
  z-index: 999;
}
/* 翻转END */
.horizontal-line {
  position: relative;
  top: 1.8vw;
  font-size: 3vw;
  width: 100%;
  height: 2vw;
  text-align: center; /* 文字水平居中 */
  text-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2); /* 添加阴影效果 */
}

.Carousel {
  /* overflow: hidden; */
  margin: 0 auto 0 auto;
  width: 1368px;
  height: 399px;
  border-radius: 20px;
  position: relative;
  top: 0;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}
.el-carousel__item {
  border-radius: 20px;
}
.el-carousel__item img {
  height: 100%;
  width: 100%;
  /* object-fit: cover; 保持图片比例并裁剪以填充整个区域 */
}
.yuan{
  border-radius: 20px;
}
.FeatureCards {
  position: relative;
  height: 650px;
  width: 100%;
}

.FarmCard:hover {
  transform: scale(1.02);
}
.FarmCard:active {
  transform: scale(0.97); /* 按下时缩小 */
}
.AI_surveillance_Card:active {
  transform: scale(0.97); /* 按下时缩小 */
}
.AICard:active {
  transform: scale(0.97); /* 按下时缩小 */
}
.pastureCard:active {
  transform: scale(0.97); /* 按下时缩小 */
}

.pastureCard:hover {
  transform: scale(1.02);
}

.AI_surveillance_Card:hover {
  transform: scale(1.02);
}

.AICard:hover {
  transform: scale(1.02);
}
.bottom {
  position: relative;
  display: flex;
  height: 9.5vw;
  background: #2b303b;
  color: white;
  font-size: 14px;
  align-items: center;
  justify-content: center;
  width: 100%;
  top: 10vw;
}

.PredictionCard:hover {
  transform: scale(1.02);
}
.PredictionCard:active {
  transform: scale(0.97);
}
/*农业专家 */
.AICard {
  position: absolute;
  top: 23.6vw;
  width: 600.5px;
  height: 315px;
  left: 62.6vw;
  background-image: url(../img/Round-robin_diagram/AI_Card.jpg);
  background-size: 100%;
  border-radius: 2vw;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  transition: transform 0.4s;
}

.AI_surveillance_Card {
  position: absolute;
  top: 4vw;
  width: 600.5px;
  height: 315px;
  left: 62.6vw;
  background-image: url(../img/Round-robin_diagram/AI_surveillance_Card.jpg);
  background-size: 100%;
  border-radius: 2vw;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  transition: transform 0.4s;
}
/* 农场卡片样式 */


.PredictionCard {
  position: absolute;
  top: 4vw;
  width: 315px;
  height: 652.5px;
  left: 1.5vw;
  background-image: url('../img/pic2.jpg');
  background-size: 100%;
  border-radius: 2vw;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  transition: transform 0.4s;
}

.FarmCard {
  position: relative; /* 相对定位，便于内部元素绝对定位 */
  top: 4vw; /* 距离顶部4vw，vw表示视口宽度的百分比 */
  width: 315px; /* 卡片宽度固定为315像素 */
  height: 652.5px; /* 卡片高度固定为652.5像素 */
  left: 21.7vw;
  background-image: url(../img/Round-robin_diagram/Farm_Card.jpg); /* 设置背景图片路径 */
  background-size: 100%; /* 背景图片大小为容器100%，即完全覆盖容器 */
  border-radius: 2vw; /* 圆角大小为2vw，vw表示视口宽度的百分比 */
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19); /* 设置阴影效果，增强立体感 */
  transition: transform 0.4s; /* 设置transform属性的过渡效果，持续0.4秒 */
}

.pastureCard {
  position: absolute;
  top: 4vw;
  width: 315px;
  height: 652.5px;
  left: 42.0vw;
  background-image: url(../img/Round-robin_diagram/Crop_storage_Card.jpg);
  background-size: 100%;
  border-radius: 2vw;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  transition: transform 0.4s;
}
</style>
