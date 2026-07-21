<template>
  <div class="box">
    <!-- <div class="top">
            <p>云端田园-农作物存储环境监测平台</p>
        </div> -->
    <div class="table_text">农作物存储环境监测平台</div>
    <img src="../img/daohlan1.png" alt="" class="top" />
    <div class="content">
      <div class="left">
        <div class="ID">
          <img
            src="../img/Agro/上一个.png"
            alt=""
            class="ID_left_Img"
            @click="Agro_Last"
          />
          <p>设备ID：{{ this.$store.state.Go_IoTAS_ASID }}</p>
          <!-- <button class="EquipmentID_Right">></button> -->
          <img
            src="../img/Agro/下一个.png"
            alt=""
            class="ID_right_Img"
            @click="Agro_next"
          />
        </div>
        <Storage_AI class="AI"> </Storage_AI>
      </div>
      <div class="right_box">
        <Storage_Temperature class="Storage_Temperature"></Storage_Temperature>
        <Storage_Humidity class="Storage_Humidity"></Storage_Humidity>
      </div>
      <div class="right_bottm_box">
        <Storage_Oxygen class="Storage_Oxygen"></Storage_Oxygen>
        <Storage_CarbonDioxide class="Storage_CarbonDioxide"></Storage_CarbonDioxide>
      </div>
    </div>
  </div>
</template>

<script>
// import Storage_AI from "./Storage_Monitoring/Storage_AI";
import Storage_AI from "./Storage_Monitoring/Storage_AI.vue";
import Storage_Temperature from "./Storage_Monitoring/Storage_Temperature.vue";
import Storage_Humidity from "./Storage_Monitoring/Storage_Humidity.vue";
import Storage_Oxygen from "./Storage_Monitoring/Storage_Oxygen.vue";
import Storage_CarbonDioxide from "./Storage_Monitoring/Storage_CarbonDioxide.vue";
import Go_IoTAS from "./Go_IoTAS";
export default {
  components: {
    Storage_AI,
    Storage_Temperature,
    Storage_Humidity,
    Storage_Oxygen,
    Storage_CarbonDioxide,
  },
  mixins: [Go_IoTAS],
  mounted() {
    this.init(); // 初始化
    this.initMqttConnection(); // 初始化 MQTT 连接
  },
  methods: {
    Agro_next() {
      const num = this.$store.state.Go_IoTAS_ASID + 1;
      this.$store.commit("update_Go_IoTAS_ASID_DataArray", num);
      this.init();
      console.log(this.$store.state.Go_IoTAS_ASID);
    },
    Agro_Last() {
      const num = this.$store.state.Go_IoTAS_ASID - 1;
      if (num < 1) {
        return;
      }
      this.$store.commit("update_Go_IoTAS_ASID_DataArray", num);
      this.init();
      console.log(this.$store.state.Go_IoTAS_ASID);
    },
  },
};
</script>

<style scoped>
.box {
  position: fixed;
  height: 100%;
  width: 100%;
  bottom: 0;
  right: 0;
  top: 0;
  background-image: url(../img/bg.png);
  background-size: 100% 100%;
}
.top {
  /* height: 4vw; */
  width: 100%;
  height: 5vw;
  margin: 0 0 1vw 0;
}

.table_text {
  color: white;
  font-size: 1.8rem; /* 字号加大一号 */
  position: absolute;
  margin: 1.4vw 0 0 38.5vw; /* 向右移动一点 */
  font-family: "思源宋体 SemiBold", sans-serif;
}
.content {
  position: fixed;
  height: 45vw;
  width: 96vw;
  top: 6vw;
  left: 2vw;
  background-color: hsla(
    180,
    100%,
    50%,
    0.2
  ); /* 设置背景颜色为带有50%透明度的蓝色 */
  border-radius: 0.7vw;
  /* box-shadow: 0px 0px 6px 3px #bdb7ab; */
}
.left {
  position: relative;
  height: 45vw;
  width: 48vw;
  border-radius: 0.7vw;
  /* background: #2cdcc2; */
}
.ID {
  position: absolute;
  height: 4vw;
  width: 40vw;
  box-shadow: 3px 3px 4px 0 rgb(7, 226, 233);
  background-color: hsla(180, 100%, 50%, 0.2);
  top: 1.5vw;
  left: 1.5vw;
  border-radius: 0.7vw;
  /* box-shadow:inset 12px 12px 7px -10px #898585;  */
  display: flex;
  justify-content: center;
  align-items: center;
}
.ID_left_Img{
    position: absolute;
    left: 0.3vw;
    height: 3.5vw;
    width: 3vw;
    cursor: pointer;
}
.ID_right_Img{
    position: absolute;
    left: 36.7vw;
    height: 3.5vw;
    width: 3vw;
    cursor: pointer;
}
.ID p {
  position: absolute;
  font-weight: bold;
  font-size: 2vw;
  color: white;
}
.AI {
  position: absolute;
  top: 7vw;
  left: 1.5vw;
}
.right_box {
  position: relative;
  top: -44.5vw;
  left: 43vw;
  height: 21.5vw;
  width: 51.5vw;
  display: flex;
}
.Storage_Temperature {
  position: relative;
  top: 1vw;
}
.Storage_Humidity {
  position: relative;
  top: 1vw;
  left: 1.5vw;
}
.right_bottm_box {
  position: relative;
  top: -44.5vw;
  left: 43vw;
  height: 21.5vw;
  width: 51.5vw;
  display: flex;
}
.Storage_Oxygen {
  position: relative;
  top: 1.5vw;
}
.Storage_CarbonDioxide {
  position: relative;
  top: 1.5vw;
  left: 1.5vw;
}
</style>