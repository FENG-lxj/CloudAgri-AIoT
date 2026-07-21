<template>
  <div>
    <div v-if="$store.state.isModalVisible" class="modal" @click="closeModal">
      <div class="modal-content" @click.stop>
        <!-- 关闭按钮 -->
        <span class="close" @click="closeModal">
          <img class="close_img" src="../img/Device_Management/close_img.png" alt="Close" />
        </span>

        <!-- 标题 + 喇叭 -->
        <div class="title">
          <p>AI农业专家</p>
          <i
              v-if="conversation.some(m => m.role === 'ai')"
              class="el-icon-microphone title-speaker"
              :title="speaking ? '停止朗读' : '朗读内容'"
              @click="toggleSpeak"
          ></i>
        </div>

        <!-- 聊天区 -->
        <div class="Chat_area" ref="chatArea">
          <div v-for="(message, index) in conversation" :key="index">
            <div v-if="message.role === 'user'" class="message-container">
              <img src="../img/SuperAdmin/AI.jpg" class="avatar" />
              <div class="message-text">{{ message.text }}</div>
              <hr class="message-hr" />
            </div>
            <div v-else class="message-container">
              <img src="../img/GPT.png" class="avatar" />
              <div class="message-text">{{ message.text }}</div>
              <hr class="message-hr" />
            </div>
          </div>
        </div>

        <!-- 输入框（仅文本） -->
        <input
            type="text"
            class="Enter_text"
            placeholder="输入问题按回车发送"
            v-model="inputText"
            @keyup.enter="sendMessage"
            @keyup.esc="closeModal"
        />

        <!-- 发送按钮 -->
        <button class="Send" @click="sendMessage">发送</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "AgricultureChat",
  data() {
    return {
      conversation: [],
      speaking: false,
      inputText: "",
    };
  },
  methods: {
    closeModal() {
      this.$store.commit("set_isModalVisible", false);
      document.body.style.overflow = "auto";
      speechSynthesis.cancel();
    },

    sendMessage() {
      const content = this.inputText.trim();
      if (!content) return;

      this.conversation.push({ role: "user", text: content });
      this.inputText = "";

      axios
          .get("http://127.0.0.1:4999/IoTA/ai/AgricultureExpertAI", {
            params: { Issue: content },
            headers: { Authorization: `Bearer ${this.$store.state.phoneCaptcha}` },
            timeout: 500000,
          })
          .then((res) => {
            const aiReply = res.data.data?.AIReply || "暂无回复";
            this.conversation.push({ role: "ai", text: aiReply });
            this.$nextTick(this.scrollToBottom);
          })
          .catch(() => {
            this.conversation.push({ role: "ai", text: "网络异常，请稍后重试" });
            this.$nextTick(this.scrollToBottom);
          });
    },

    scrollToBottom() {
      const area = this.$refs.chatArea;
      area && (area.scrollTop = area.scrollHeight);
    },

    /* 📢 朗读最新一条 AI 回答 */
    toggleSpeak() {
      if (this.speaking) {
        speechSynthesis.cancel();
        this.speaking = false;
        return;
      }

      // 只取最后一条 AI 消息
      const aiMessages = this.conversation.filter((m) => m.role === "ai");
      const latestAiMessage = aiMessages.length > 0 ? aiMessages[aiMessages.length - 1].text : "";

      if (!latestAiMessage) return;

      const utterance = new SpeechSynthesisUtterance(latestAiMessage);
      utterance.lang = "zh-CN";
      utterance.onstart = () => (this.speaking = true);
      utterance.onend = () => (this.speaking = false);
      speechSynthesis.speak(utterance);
    },
  },
};
</script>

<style scoped>
/* 保持原样式不变，仅移除语音按钮相关样式 */
.modal {
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.7);
  z-index: 1000;
}

/* 内容区域样式 */
.modal-content {
  background: #dce9f9;
  padding: 20px;
  border-radius: 20px;
  width: 50vw;
  height: 33vw;
  position: relative;
  overflow: hidden;
}

/* 关闭按钮样式 */
.close {
  position: absolute;
  top: 0.4vw;
  right: 0.4vw;
  font-size: 20px;
  cursor: pointer;
  z-index: 1;
}
.close_img {
  width: 3vw;
}
.title {
  position: absolute;
  text-align: left;
  margin-bottom: 20px;
  top: 0;
  left: 0;
  width: 55vw;
  height: 3.8vw;
  background: #b8d3f0;
  border-top-left-radius: 20px;
  border-top-right-radius: 20px;
  border-bottom-left-radius: 0;
  border-bottom-right-radius: 0;
}
.title p {
  font-size: 2vw;
  color: #346e81;
  margin-left: 1vw;
  margin-top: 0.6vw;
}
.Chat_area {
  position: absolute;
  height: 27vw;
  width: 52.2vw;
  /* border: 1px solid black; */
  top: 3.8vw;
  left: 0;
  overflow-y: scroll; /* 允许 Chat_area 区域滚动 */
}

/* 隐藏滚动条 */
.Chat_area::-webkit-scrollbar {
  width: 0.8em;
}

.Chat_area::-webkit-scrollbar-thumb {
  background-color: transparent;
}

.Enter_text {
  position: absolute;
  top: 31.4vw;
  left: 15px;
  height: 3vw;
  width: 44vw;
  font-size: 1.6vw;
  border-radius: 20px;
  background: #b7d2f0;
  text-indent: 0.6vw;
}
.Send {
  position: absolute;
  top: 31.5vw;
  left: 46.2vw;
  height: 3.1vw;
  width: 80px;
  border-radius: 20px;
  background: #b7d2f0;
  border: none;
  cursor: pointer;
  font-size: 21px;
  color: #334e6b;
  font-weight: 600;
  box-shadow: 2px 3px 2px #334e6b61;
}

.message-container {
  display: flex;
  margin-bottom: 10px;
  margin-left: 20px;
  margin-top: 0.5vw;
  position: relative;
  align-items: flex-start;
}

.avatar {
  width: 45px;
  height: 45px;
  margin-right: 10px;
  margin-top: 7px;
  border-radius: 100px;
}

.message-text {
  margin-top: 15px;
  font-size: 21px;
  color: #3c6592;
  text-align: left;
  white-space: pre-line;
  margin-bottom: 15px;
}
.message-hr {
  position: absolute;
  width: 100%;
  bottom: -17px;
}
.title-speaker {
  font-size: 2vw; /* 调整大小，避免过大 */
  color: #346e81;
  cursor: pointer;
  position: absolute;
  top: 0.6vw;
  right: 6vw; /* 放在 X 按钮左侧一点点（X 位于 0.4vw） */
}
</style>