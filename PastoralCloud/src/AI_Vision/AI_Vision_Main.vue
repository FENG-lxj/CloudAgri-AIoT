<template>
  <div class="box">
    <!-- <div class="top">
        <p>云端田园-AI智能视觉监控</p>
    </div> -->
    <div class="table_text">AI智能视觉监控平台</div>
    <img src="../img/daohlan1.png" alt="" class="top" />

    <div class="content">
        <div class="left_box">
        <AI_Vision_ID class="AI_Vision_ID"></AI_Vision_ID>
        <AI_Camera_One class="AI_Camera_One"></AI_Camera_One>
        <AI_Camera_Two class="AI_Camera_Two"></AI_Camera_Two>
    </div>
    <video ref="videoElement" controls class="AI_Video"></video>
    <AI_Prevention class="AI_Prevention2"></AI_Prevention>
    <AI_log class="AI_log2"></AI_log>
    </div>
  </div>
</template>

<script>
import AI_Vision_ID from './AI_Vision_Component/AI_Vision_ID.vue'
import AI_Camera_One from './AI_Vision_Component/AI_Camera_One.vue'
import AI_Camera_Two from './AI_Vision_Component/AI_Camera_Two.vue'
import AI_Video from './AI_Vision_Component/AI_Video.vue'
import AI_Prevention from './AI_Vision_Component/AI_Prevention.vue'
import AI_log from './AI_Vision_Component/AI_log.vue'
import flvjs from 'flv.js';

export default{
    components:{
        AI_Vision_ID,
        AI_Camera_One,
        AI_Camera_Two,
        AI_Video,
        AI_Prevention,
        AI_log,
    },
    mounted() {
    if (flvjs.isSupported()) {
      this.initFlvPlayer();
      console.error('成功显示');

    } else {
      console.error('FLV not supported');
    }
  },
    methods: {
      initFlvPlayer() {
        const videoElement = this.$refs.videoElement;
        const flvPlayer = flvjs.createPlayer({
          type: 'flv',
          // url: 'http://39.102.208.105:8080/PastoralCloud/monitor_1.flv',
          url: 'http://127.0.0.1:8080/live/livestream.flv',
        });
        flvPlayer.attachMediaElement(videoElement);
        flvPlayer.load();
      },
    },
}
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
.content{
  display: flex;
}
.top {
  /* height: 4vw; */
  width: 100%;
  height: 5vw;
  margin: 0 0 1vw 0;
}
.table_text{
  color: white;
  font-size: 1.7rem;
  position: absolute;
  margin: 1.4vw 0 0 36.5vw;
  font-family: "思源宋体 SemiBold", sans-serif;
}
.left_box{
  position: relative;
  height: 50vw;
  width: 50vw;
  top: -1vw;
}
.AI_Vision_ID{
  position: relative;
  top: 0vw;
  left: 1.5vw;
}
.AI_Camera_One{
  position: relative;
  top: 1.5vw;
  left: 1.5vw;
}
.AI_Camera_Two{
  position: relative;
  top: 3vw;
  left: 1.5vw;
}
.AI_Video{
  position: relative;
  top: 0vw;
}
.AI_Prevention2{
  position: absolute;
  top: 32vw;
  left: 50vw;
}
.AI_log2{
  position: absolute;
  top: 32vw;
  left: 75vw;
}
</style>