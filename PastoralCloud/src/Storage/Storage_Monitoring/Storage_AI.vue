<template>
  <div>
    <div class="AI">
      <p v-if="time !== ''">最新AI建议时间：{{ time }}</p>
      <p v-else>最新AI建议时间：</p>
      <hr />

      <!-- 获取建议按钮 -->
      <button class="Advice_btn" @click="GET_AI">获取建议</button>

      <!-- 喇叭图标 -->
      <i
          v-if="AIProposal"
          class="el-icon-microphone AI-speaker"
          @click="toggleSpeak"
      ></i>

      <!-- 建议内容 -->
      <div class="AIProposal_text" v-if="AIProposal !== ''">
        {{ AIProposal }}
      </div>
      <div class="AIProposal_text" v-else>正在获取最新AI建议...</div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "AI",
  data() {
    return {
      AIProposal: "",
      time: "",
    };
  },
  mounted() {
    this.GET_AI();
  },
  methods: {
    GET_AI() {
      this.AIProposal = "";
      this.time = "";
      axios
          .get("http://127.0.0.1:20201/IoTA/ai/ProposalAI_AS", {
            params: { ASID: 1 },
            headers: { Authorization: `Bearer ${this.$store.state.phoneCaptcha}` },
            timeout: 500000,
          })
          .then((res) => {
            this.AIProposal = res.data.data.AIProposal;
            this.time = res.data.data.time;
            console.log("AI建议", res.data);
          })
          .catch((error) => {
            console.error("Error fetching data:", error);
          });
    },

    /* 原有提示方法保留 */
    GET_AIProposal_success() {
      this.$notify({ title: "获取最新AI建议中~" });
    },
    AIProposal_success() {
      this.$notify({ title: "最新AI建议请求成功", type: "success" });
    },
    AIProposal_error() {
      this.$notify.error({
        title: "最新AI建议请求失败",
        message: "请检查网络连接",
      });
    },

    /* 新增：开始 / 停止 语音 */
    toggleSpeak() {
      if (!this.AIProposal) return;

      if (speechSynthesis.speaking) {
        speechSynthesis.cancel(); // 立即停止
      } else {
        const utterance = new SpeechSynthesisUtterance(this.AIProposal);
        utterance.lang = "zh-CN";
        speechSynthesis.speak(utterance);
      }
    },
  },
};
</script>

<style scoped>
.AI {
  position: relative;
  width: 40vw;
  height: 37vw;
  border-radius: 1vw;
  background-color: hsla(180, 100%, 50%, 0.2);
  box-shadow: 3px 3px 4px 0 rgb(7, 226, 233);
}
.AI p {
  position: absolute;
  font-weight: bold;
  color: aliceblue;
  top: -0.8vw;
  left: 0.6vw;
  font-size: 1.3vw;
}
.AI hr {
  position: absolute;
  height: 1px;
  top: 2.3vw;
  background: #ffffff;
  border: none;
  width: 40vw;
}

/* 获取建议按钮 */
.Advice_btn {
  position: absolute;
  top: 0.4vw;
  left: 34.5vw;
  height: 2vw;
  width: 5vw;
  border-radius: 1vw;
  font-weight: bold;
  font-size: 1vw;
  color: #fff;
  background-color: hsla(180, 100%, 50%, 0.453);
  box-shadow: 1px 1px 3px 0 rgb(7, 226, 233);
  cursor: pointer;
  border: none;
}
.Advice_btn:hover {
  background-color: hsla(180, 100%, 50%, 0.382);
  color: rgba(255, 255, 255, 0.751);
}

/* 喇叭图标：紧贴按钮左侧 */
.AI-speaker {
  position: absolute;
  top: 0.4vw;
  right: 5.8vw; /* 34.5vw - 5vw - 0.3vw = 29.2vw ≈ 5.8vw from right */
  font-size: 1.8vw;
  color: #fff;
  cursor: pointer;
  z-index: 10;
}
.AI-speaker:hover {
  color: #07e2e9;
}

/* 建议文本区 */
.AIProposal_text {
  color: #ffffff;
  margin-top: 55px;
  position: absolute;
  margin-left: 10px;
  margin-right: 22px;
  max-height: 32.3vw;
  overflow: auto;
  white-space: pre-line;
  width: 38.8vw;
}
.AIProposal_text::-webkit-scrollbar {
  width: 0px;
}
.AIProposal_text::-webkit-scrollbar-thumb {
  background-color: transparent;
}
</style>