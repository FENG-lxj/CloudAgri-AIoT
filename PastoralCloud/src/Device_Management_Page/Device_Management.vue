<template>
  <div class="box">
    <div class="top">
      <img src="../img/Device_Management/logo.jpg" alt="" class="logo" @click="home"/>
      <div class="logo_text" @click="home">螭耀科技</div>
      <div class="vertical_line"></div>
      <div class="logo_text_two" @click="home">云端田园</div>
      <div class="vertical_line_two"></div>
      <div class="logo_text_three">设备管理</div>
    </div>
    <div class="Navigation_Bar">
      <div
        class="Navigation_Bar_One"
        id="Navigation_Bar_One"
        @click="Navigation_Bar_One_btn"
        :class="{ Bar_clicked: activeNavigation === 'one' }"
      >
        农业采集设备
      </div>
      <div
        class="Navigation_Bar_Two"
        id="Navigation_Bar_Two"
        @click="Navigation_Bar_Two_btn"
        :class="{ Bar_clicked: activeNavigation === 'two' }"
      >
        存储采集设备
      </div>
      <div
        class="Navigation_Bar_Three"
        id="Navigation_Bar_Three"
        @click="Navigation_Bar_Three_btn"
        :class="{ Bar_clicked: activeNavigation === 'three' }"
      >
        AI采集设备
      </div>
    </div>
    <div class="main">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Agricultural_Equipment from "./Page/Agricultural_Equipment.vue";
import Storage_Devices from "./Page/Storage_Devices.vue";
import AI_Devices from "./Page/AI_Devices.vue";
export default {
  components: {
    Agricultural_Equipment,
    Storage_Devices,
    AI_Devices,
  },
  //设置子路由
  
  data() {
    return {
      activeNavigation: "one",
    };
  },
  methods: {
    home(){
        this.$router.replace({ path: '/home' });
    },
    Navigation_Bar_One_btn() {
      this.activeNavigation = "one";
      if(this.$route.name!='/'){
        this.$router.push({ name: '/' });
      }
    },
    Navigation_Bar_Two_btn() {
      this.activeNavigation = "two";
      if(this.$route.name!='Storage_Devices'){
        this.$router.push({ name: 'Storage_Devices' });
      }
    },
    Navigation_Bar_Three_btn() {
      this.activeNavigation = "three";
      if(this.$route.name!='AI_Devices'){
        this.$router.push({ name: 'AI_Devices' });
      }
    },
  },
};
</script>

<style scoped>
.box {
  position: fixed;
  left: 0;
  top: 0;
  height: 100%;
  width: 100%;
}
.top {
  position: fixed;
  background: #000;
  height: 4vw;
  width: 100%;
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
  top: 1.2vw;
  color: white;
  font-size: 1.2vw;
  cursor: pointer;
}
.logo_text_two {
  position: absolute;
  left: 10.2vw;
  top: 1.2vw;
  color: white;
  font-size: 1.2vw;
  cursor: pointer;
}
.logo_text_three {
  position: absolute;
  left: 94vw;
  top: 1.2vw;
  color: white;
  font-size: 1.2vw;
}
.vertical_line {
  position: absolute;
  border-left: 1px solid rgb(177, 177, 177);
  height: 2vw;
  left: 9.7vw;
  top: 1.1vw;
}
.vertical_line_two {
  position: absolute;
  border-left: 1px solid rgb(177, 177, 177);
  height: 2vw;
  left: 93vw;
  top: 1.1vw;
}
.Navigation_Bar {
  position: absolute;
  left: 0;
  width: 4vw;
  height: 96vw;
  top: 4vw;
  background: #000;
}
.Navigation_Bar_Two {
  position: relative;
  color: white;
  height: 16vw;
  writing-mode: vertical-rl; /* vertical right-to-left */
  text-orientation: upright;
}
.Navigation_Bar_Three {
  position: relative;
  color: white;
  height: 16vw;
  writing-mode: vertical-rl; /* vertical right-to-left */
  text-orientation: upright;
}
.Navigation_Bar_One,
.Navigation_Bar_Two,
.Navigation_Bar_Three {
  display: flex;
  width: 4vw;
  align-items: center; /* 竖直方向居中 */
  justify-content: center; /* 水平方向居中 */
  writing-mode: vertical-rl; /* 竖排，从右到左 */
  border-bottom: 1px solid white; /* 添加横线，可以根据需要调整颜色和宽度 */
  cursor: pointer;
  background: #000;
}
.Navigation_Bar_One {
  position: relative;
  color: white;
  height: 16vw;
  writing-mode: vertical-rl; /* vertical right-to-left */
  text-orientation: upright;
}
.Navigation_Bar_One:hover,
.Navigation_Bar_Two:hover,
.Navigation_Bar_Three:hover {
  background: #525864;
}

.Bar_clicked {
  background: #525864;
}
.main {
  position: fixed;
  top: 4vw;
  left: 4vw;
  width: 96vw;
  height: 96vw;
  /* background: #b40dfc; */
}
</style>