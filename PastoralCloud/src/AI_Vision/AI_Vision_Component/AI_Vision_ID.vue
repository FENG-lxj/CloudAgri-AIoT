<template>
  <div>
    <div class="ID">
      <img
        src="../../img/Agro/上一个.png"
        alt=""
        class="ID_left_Img"
        @click="Decrease"
      />
      <p>设备ID：{{ this.$store.state.AI_Camera_One_ICID }}</p>
      <!-- <button class="EquipmentID_Right">></button> -->
      <img
        src="../../img/Agro/下一个.png"
        alt=""
        class="ID_right_Img"
        @click="Add"
      />
    </div>
  </div>
</template>
<script>
import AI_Camera_One from "./AI_Camera_One.vue";
export default {
  components: {
    AI_Camera_One,
  },
  data() {
    return {};
  },
  mounted() {},
  methods: {
    Add() {
      if (this.$store.state.AI_Camera_One_ICID < 2) {
        const num = this.$store.state.AI_Camera_One_ICID + 1;
        this.$store.commit("set_AI_Camera_One_ICID", num);
      }
    },
    Decrease() {
      if(this.$store.state.AI_Camera_One_ICID > 1){
        const num1 = this.$store.state.AI_Camera_One_ICID - 1;
      this.$store.commit("set_AI_Camera_One_ICID", num1);
      }
    },
  },
};
</script>
<style scoped>
.ID {
  position: relative;
  height: 6vw;
  width: 47vw;
  border-radius: 0.8vw;
  background-color: hsla(180, 100%, 50%, 0.2);
  box-shadow: 3px 3px 4px 0 rgb(7, 226, 233);
  display: flex;
  margin-top: 1vw;
  justify-content: center;
  align-items: center;
  /* box-shadow: 3px 3px 5px 1px #a5a9a9; */
}
.ID p {
  position: absolute;
  font-size: 2.5vw;
  /* font-weight: bold; */
  color: white;
}
.ID_left_Img {
  position: absolute;
  left: 0.5vw;
  height: 3.8vw;
  width: 3vw;
  cursor: pointer;
}
.ID_right_Img {
  position: absolute;
  left: 43.6vw;
  height: 3.8vw;
  width: 3vw;
  cursor: pointer;
}
</style>