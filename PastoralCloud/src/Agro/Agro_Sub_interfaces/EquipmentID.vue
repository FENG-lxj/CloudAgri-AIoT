<template>
  <div>
    <div class="EquipmentID">
      <!-- <button class="EquipmentID_Left">&lt;</button> -->
      <img
        src="../../img/Agro/上一个.png"
        alt=""
        class="EquipmentID_left_Img"
        @click="Agro_Last"
      />
      <!-- <p>设备ID：{{ AEID }}</p> -->
      <p class="IDtext">设备ID：{{ this.$store.state.AEID }}</p>
      <!-- <button class="EquipmentID_Right">></button> -->
      <img
        src="../../img/Agro/下一个.png"
        alt=""
        class="EquipmentID_right_Img"
        @click="Agro_next"
      />
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import { mapMutations } from "vuex";
import Go_IoTAE from "@/Agro/Go_IoTAE";
export default {
  mixins: [Go_IoTAE],
  computed: {
    ...mapState(["AEID"]),
  },
  created() {
    console.log(this.$store.state.AEID);
  },
  methods: {
    Agro_next() {
      const num = this.$store.state.AEID + 1;
      this.$store.commit("update_AEID_DataArray", num);
      this.init();
      console.log(this.$store.state.AEID);
    },
    Agro_Last() {
      const num = this.$store.state.AEID - 1;
      if (num < 1) {
        return;
      }
      this.$store.commit("update_AEID_DataArray", num);
      this.init();
      console.log(this.$store.state.AEID);
    },
  },
};
</script>

<style scoped>
.EquipmentID {
  display: flex;
  position: relative;
  height: 10.2vw;
  width: 12.6vw;
  margin-left: 1vw;
  margin-top: 1vw;
  /* background: #7474FA; */
  border-radius: 1vw;
  justify-content: center;
  align-items: center;
  background-color: hsla(
    180,
    100%,
    50%,
    0.2
  ); /* 设置背景颜色为带有50%透明度的蓝色 */
  box-shadow: 5px 4px 8px 0 rgb(7, 226, 233); /* 添加阴影效果 */
  /* box-shadow: 0px 3px 6px #7474fabc; */
}
.IDtext {
  font-size: 23px;
}
.EquipmentID_Left {
  position: relative;
  left: -1vw;
}
.EquipmentID_left_Img {
  position: absolute;
  left: 0vw;
  height: 5vw;
  width: 3vw;
  cursor: pointer;
}
.EquipmentID_Right {
  position: relative;
  right: -1vw;
}
.EquipmentID_right_Img {
  position: absolute;
  left: 9.5vw;
  height: 5vw;
  width: 3vw;
  cursor: pointer;
}
.EquipmentID p {
  position: relative;
  font-weight: bold;
  color: aliceblue;
}
</style>