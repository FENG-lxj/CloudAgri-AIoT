<template>
  <div>
    <div class="AI_log">
      <p>病虫害防治小建议</p>
      <hr />
      <div v-if="pestProposal != ''">
        <!-- 显示获取到的建议 -->
        <p class="message">{{ pestProposal }}</p>
      </div>
      <div v-if="pestProposal == ''">
        <!-- 如果数据还在加载中，可以显示加载中的提示 -->
        <p class="message">Loading...</p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      pestProposal: null, // 存储病虫害防治建议的数据
    };
  },
  mounted() {
    this.init();
    setInterval(() => {
      this.init();
      }, 10000); // 每1000毫秒（1秒）执行一次
  },
  methods: {
    init() {
      axios
        .get("http://127.0.0.1:20201/IoTA/ai/PestProposal", {
          params: {
            Pest: "疫病",
          },
          headers: {
            Authorization: `Bearer ${this.$store.state.phoneCaptcha}`,
          },
        })
        .then((res) => {
          // 更新 pestProposal 数据
          this.pestProposal = res.data.data.Proposal;
          console.log("获取log", res.data);
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
        });
    },
      
  },
};
</script>

<style scoped>
.AI_log {
  position: relative;
  height: 17.3vw;
  width: 23.5vw;
  margin-top: 1vw;
  border-radius: 0.8vw;
  background-color: hsla(180, 100%, 50%, 0.2);
  box-shadow: 3px 3px 4px 0 rgb(7, 226, 233);
}
.AI_log p {
  position: absolute;
  font-weight: bold;
  top: -0.5vw;
  left: 7.5vw;
  color: #ffffff;
}
.AI_log hr {
  position: absolute;
  top: 2vw;
  height: 0.2px;
  width: 23.4vw;
  background: #ffffff;
}
.message {
  margin-top: 67px;
  margin-left: -95px;
  margin-right: 10px;
  max-height: 189px;
  overflow: auto;

  /* 针对Webkit浏览器隐藏滚动条 */
  scrollbar-width: thin;
  scrollbar-color: transparent transparent;
}

/* 针对Webkit浏览器隐藏滚动条 */
.message::-webkit-scrollbar {
  width: 0;
  height: 0;
}

/* 针对Webkit浏览器隐藏滚动条 */
.message::-webkit-scrollbar-thumb {
  background-color: transparent;
}
</style>