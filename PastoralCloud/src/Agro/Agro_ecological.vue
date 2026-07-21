<template>
  <div class="box">
    <!-- <div class="top">
            <p>云端田园农业生态大数据监测平台</p>
        </div> -->
    <div class="table_text">农业生态环境大数据监测平台</div>
    <img src="../img/daohlan1.png" alt="" class="top" />
    <div class="content">
      <EquipmentID class="EquipmentID_One"></EquipmentID>
      <Air_quality class="Air_quality_Two"></Air_quality>
      <Light_intensity class="Light_intensity_Three"></Light_intensity>
      <Carbon_Dioxide class="Carbon_Dioxide_Four"></Carbon_Dioxide>
      <Methanol class="Methanol_Five"></Methanol>
      <AI class="AI_Six"></AI>
      <Air_Temperature class="Air_Temperature_Seven"></Air_Temperature>
      <Air_Humidity class="Air_Humidity_Eight"></Air_Humidity>
      <Soil_PH class="Soil_PHv_Nine"></Soil_PH>
    </div>
    <div class="content_One">
      <Nitrogen class="Nitrogen_Ten"></Nitrogen>
      <Phosphorus class="Phosphorus_Eleven"></Phosphorus>
      <Kalium class="Kalium_Twelve"></Kalium>
      <Soil_Temperature class="Soil_Temperature_Thirteen"></Soil_Temperature>
      <Soil_Moisture class="Soil_Moisture_Fourteen"></Soil_Moisture>
    </div>
  </div>
</template>

<script>
import EquipmentID from "@/Agro/Agro_Sub_interfaces/EquipmentID.vue";
import Air_quality from "@/Agro/Agro_Sub_interfaces/Air_quality.vue";
import Light_intensity from "@/Agro/Agro_Sub_interfaces/Light_intensity.vue";
import Carbon_Dioxide from "@/Agro/Agro_Sub_interfaces/Carbon_Dioxide.vue";
import Methanol from "@/Agro/Agro_Sub_interfaces/Methanol.vue";
import AI from "@/Agro/Agro_Sub_interfaces/AI.vue";
import Air_Temperature from "@/Agro/Agro_Sub_interfaces/Air_Temperature.vue";
import Air_Humidity from "@/Agro/Agro_Sub_interfaces/Air_Humidity.vue";
import Soil_PH from "@/Agro/Agro_Sub_interfaces/Soil_PH.vue";
import Nitrogen from "@/Agro/Agro_Sub_interfaces/Nitrogen.vue";
import Phosphorus from "@/Agro/Agro_Sub_interfaces/Phosphorus.vue";
import Kalium from "@/Agro/Agro_Sub_interfaces/Kalium.vue";
import Soil_Temperature from "@/Agro/Agro_Sub_interfaces/Soil_Temperature.vue";
import Soil_Moisture from "@/Agro/Agro_Sub_interfaces/Soil_Moisture.vue";
import Go_IoTAE from "@/Agro/Go_IoTAE";
import { MapState } from "vuex";
import { mapMutations } from "vuex";
export default {
  components: {
    EquipmentID,
    Air_quality,
    Light_intensity,
    Carbon_Dioxide,
    Methanol,
    AI,
    Air_Temperature,
    Air_Humidity,
    Soil_PH,
    Nitrogen,
    Phosphorus,
    Kalium,
    Soil_Temperature,
    Soil_Moisture,
  },
  mixins: [Go_IoTAE],
  mounted() {
    this.init(); // 初始化
    this.initMqttConnection(); // 初始化 MQTT 连接
  },
};
</script>

<style scoped>
.EquipmentID_One {
  position: relative;
}
.Air_quality_Two {
  position: relative;
  left: 1.8vw;
  top: 1vw;
}
.Light_intensity_Three {
  position: relative;
  top: 1.3vw;
  left: 3.3vw;
}
.Carbon_Dioxide_Four {
  position: relative;
  top: 1.3vw;
  left: 4.8vw;
}
.Methanol_Five {
  position: relative;
  top: 1vw;
  left: 6.2vw;
}
.AI_Six {
  position: relative;
  left: 1vw;
}
.Air_Temperature_Seven {
  position: relative;
  left: 2.2vw;
}
.Air_Humidity_Eight {
  position: relative;
  top: 11.2vw;
  left: -19.2vw;
}
.Soil_PHv_Nine {
  position: relative;
  top: 22.5vw;
  left: -40.6vw;
}
.Nitrogen_Ten {
  position: relative;
  top: -1vw;
}
.Phosphorus_Eleven {
  position: relative;
  top: -1vw;
  left: 2vw;
}
.Kalium_Twelve {
  position: relative;
  top: -1vw;
  left: 4vw;
}
.Soil_Temperature_Thirteen {
  position: relative;
  top: 0vw;
}
.Soil_Moisture_Fourteen {
  position: relative;
  top: 0vw;
  left: 1vw;
}
.box {
  position: fixed;
  height: 100%;
  width: 100%;
  bottom: 0;
  right: 0;
  background-image: url(../img/bg.png);
  background-size: 100% 100%;
}
.content {
  display: flex;
  flex-wrap: wrap;
  /* align-items: flex-start; 或者 baseline，根据需要选择 */
  position: relative;
  top: -1vw;
}
.content_One {
  position: relative;
  display: flex;
  flex-wrap: wrap;
  top: -30vw;
  left: 58.3vw;
  width: 50vw;
}
.top {
  /* height: 4vw; */
  width: 100%;
  height: 5vw;
  margin: 0 0 1vw 0;
}
.table_text {
  color: white;
  font-size: 1.7rem;
  position: absolute;
  margin: 1.4vw 0 0 36.5vw;
  font-family: "思源宋体 SemiBold", sans-serif;
}
</style>