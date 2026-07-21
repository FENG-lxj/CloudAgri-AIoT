<template>
  <div>
    <div v-if="$store.state.isModalVisible" class="modal" @click="closeModal">
      <div class="modal-content" @click.stop>
        <!-- 关闭按钮 -->
        <span class="close" @click="closeModal">
          <img class="close_img" src="../img/Device_Management/close_img.png" alt="Close" />
        </span>

        <!-- 标题 + 喇叭（替换话筒为图片喇叭） -->
        <div class="title">
          <p>AI农业专家</p>
          <img
              v-if="conversation.some(m => m.role === 'ai')"
              class="title-speaker"
              :title="speaking ? '停止朗读' : '朗读内容'"
              @click="toggleSpeak"
              src="../img/laba.png"
          />
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

        <!-- 输入框 + 话筒图标 -->
        <div class="input-wrapper">
          <input
              type="text"
              class="Enter_text"
              placeholder="输入问题按回车发送"
              v-model="inputText"
              @keyup.enter="sendMessage"
              @keyup.esc="closeModal"
          />
          <img
              class="mic-icon"
              src="../img/huatong.png"
              alt="话筒"
              @click="onMicClick"
          />
        </div>

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

      // 语音识别相关变量
      ws: null,
      audioContext: null,
      mediaStream: null,
      processor: null,
      sampleBuffer: new Uint8Array(0),
      sendTimer: null,
      firstFrameSent: false,
      recognizing: false, // 当前是否在录音识别中
      statusText: "未连接", // 状态显示（可自行绑定）
    };
  },
  methods: {
    closeModal() {
      this.$store.commit("set_isModalVisible", false);
      document.body.style.overflow = "auto";
      speechSynthesis.cancel();
      this.stopRecognition();
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
      if (area) area.scrollTop = area.scrollHeight;
    },

    toggleSpeak() {
      if (this.speaking) {
        speechSynthesis.cancel();
        this.speaking = false;
        return;
      }

      const aiMessages = this.conversation.filter((m) => m.role === "ai");
      const latestAiMessage = aiMessages.length > 0 ? aiMessages[aiMessages.length - 1].text : "";

      if (!latestAiMessage) return;

      const utterance = new SpeechSynthesisUtterance(latestAiMessage);
      utterance.lang = "zh-CN";
      utterance.onstart = () => (this.speaking = true);
      utterance.onend = () => (this.speaking = false);
      speechSynthesis.speak(utterance);
    },

    // 语音识别相关方法

    async onMicClick() {
      if (this.recognizing) {
        await this.stopRecognition();
      } else {
        await this.startRecognition();
      }
    },

    setStatus(text) {
      this.statusText = text;
      console.log("状态：" + text);
    },

    uint8ToBase64(u8arr) {
      const CHUNK = 0x8000;
      let binary = "";
      for (let i = 0; i < u8arr.length; i += CHUNK) {
        binary += String.fromCharCode.apply(null, u8arr.subarray(i, i + CHUNK));
      }
      return btoa(binary);
    },

    arrayBufferToBase64(ab) {
      return this.uint8ToBase64(new Uint8Array(ab));
    },

    textToBase64UTF8(str) {
      return this.arrayBufferToBase64(new TextEncoder().encode(str).buffer);
    },

    async hmacSha256Base64(message, keyStr) {
      const keyData = new TextEncoder().encode(keyStr);
      const key = await crypto.subtle.importKey("raw", keyData, { name: "HMAC", hash: "SHA-256" }, false, ["sign"]);
      const sig = await crypto.subtle.sign("HMAC", key, new TextEncoder().encode(message));
      return this.arrayBufferToBase64(sig);
    },

    async makeAuthUrl() {
      const WS_BASE = "wss://iat-api.xfyun.cn/v2/iat";
      const WS_HOST = "iat-api.xfyun.cn";
      const WS_PATH = "/v2/iat";

      const date = new Date().toUTCString();
      const signature_origin = `host: ${WS_HOST}\ndate: ${date}\nGET ${WS_PATH} HTTP/1.1`;
      const signature = await this.hmacSha256Base64(signature_origin, "ZGQxNTFkY2MwMjkwYTYxZDlkZWNjZGYz");
      const authorization_origin = `api_key="d155d36e5660db9beaa7cdd5a711a9b8", algorithm="hmac-sha256", headers="host date request-line", signature="${signature}"`;
      const authorization = this.textToBase64UTF8(authorization_origin);
      return `${WS_BASE}?authorization=${encodeURIComponent(authorization)}&date=${encodeURIComponent(date)}&host=${WS_HOST}`;
    },

    downsampleBuffer(buffer, srcRate, dstRate) {
      if (dstRate === srcRate) return buffer;
      const ratio = srcRate / dstRate;
      const newLength = Math.round(buffer.length / ratio);
      const result = new Float32Array(newLength);
      let offsetResult = 0;
      let offsetBuffer = 0;
      while (offsetResult < result.length) {
        const nextOffsetBuffer = Math.round((offsetResult + 1) * ratio);
        let accum = 0,
            count = 0;
        for (let i = offsetBuffer; i < nextOffsetBuffer && i < buffer.length; i++) {
          accum += buffer[i];
          count++;
        }
        result[offsetResult] = count ? accum / count : 0;
        offsetResult++;
        offsetBuffer = nextOffsetBuffer;
      }
      return result;
    },

    floatTo16BitPCM(float32Array) {
      const l = float32Array.length;
      const buffer = new ArrayBuffer(l * 2);
      const view = new DataView(buffer);
      let offset = 0;
      for (let i = 0; i < l; i++, offset += 2) {
        let s = Math.max(-1, Math.min(1, float32Array[i]));
        view.setInt16(offset, s < 0 ? s * 0x8000 : s * 0x7fff, true);
      }
      return new Uint8Array(buffer);
    },

    concatUint8(a, b) {
      const c = new Uint8Array(a.length + b.length);
      c.set(a, 0);
      c.set(b, a.length);
      return c;
    },

    startSendLoop() {
      if (this.sendTimer) return;
      const FRAME_BYTES = 1280;
      const SEND_INTERVAL_MS = 40;

      this.sendTimer = setInterval(() => {
        if (!this.ws || this.ws.readyState !== WebSocket.OPEN) return;
        while (this.sampleBuffer.length >= FRAME_BYTES) {
          const chunk = this.sampleBuffer.slice(0, FRAME_BYTES);
          this.sampleBuffer = this.sampleBuffer.slice(FRAME_BYTES);
          const b64 = this.uint8ToBase64(chunk);
          if (!this.firstFrameSent) {
            const msg = {
              common: { app_id: "c8c5c148" },
              business: { language: "zh_cn", domain: "iat", accent: "mandarin", vad_eos: 3000 },
              data: { status: 0, format: "audio/L16;rate=16000", encoding: "raw", audio: b64 },
            };
            this.ws.send(JSON.stringify(msg));
            this.firstFrameSent = true;
          } else {
            const msg = {
              data: { status: 1, format: "audio/L16;rate=16000", encoding: "raw", audio: b64 },
            };
            this.ws.send(JSON.stringify(msg));
          }
        }
      }, SEND_INTERVAL_MS);
    },

    stopSendLoop() {
      if (this.sendTimer) {
        clearInterval(this.sendTimer);
        this.sendTimer = null;
      }
    },

    async startRecognition() {
      try {
        this.inputText = "";
        this.setStatus("正在生成鉴权 URL...");
        const url = await this.makeAuthUrl();
        this.setStatus("正在连接 WebSocket...");
        this.ws = new WebSocket(url);

        this.ws.onopen = () => {
          this.setStatus("WebSocket 已连接，准备开始录音...");
          this.initRecording();
          this.startSendLoop();
          this.recognizing = true;
        };

        this.ws.onmessage = (evt) => {
          try {
            const msg = JSON.parse(evt.data);
            if (msg.data && msg.data.result) {
              let text = "";
              if (msg.data.result.ws) {
                msg.data.result.ws.forEach((w) => {
                  if (w.cw) {
                    w.cw.forEach((cw) => {
                      text += cw.w;
                    });
                  }
                });
              }
              this.inputText = text;
            }
          } catch (e) {}
        };

        this.ws.onclose = () => {
          this.setStatus("WebSocket 已关闭");
          this.stopSendLoop();
          this.cleanupAudio();
          this.recognizing = false;
        };

        this.ws.onerror = () => {
          this.setStatus("WebSocket 错误");
          this.recognizing = false;
          this.stopRecognition();
        };
      } catch (err) {
        this.setStatus("启动失败");
        this.recognizing = false;
      }
    },

    async stopRecognition() {
      this.setStatus("正在停止录音...");
      this.stopSendLoop();

      if (this.ws && this.ws.readyState === WebSocket.OPEN) {
        if (this.sampleBuffer.length > 0) {
          const b64 = this.uint8ToBase64(this.sampleBuffer);
          this.ws.send(
              JSON.stringify({
                data: { status: 1, format: "audio/L16;rate=16000", encoding: "raw", audio: b64 },
              })
          );
          this.sampleBuffer = new Uint8Array(0);
        }
        this.ws.send(JSON.stringify({ data: { status: 2 } }));
      }

      this.cleanupAudio();

      setTimeout(() => {
        if (this.ws && this.ws.readyState === WebSocket.OPEN) {
          try {
            this.ws.close();
          } catch (e) {}
        } else {
          this.setStatus("已停止");
          this.recognizing = false;
        }
      }, 300);
    },

    async initRecording() {
      try {
        const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
        this.mediaStream = stream;
        this.audioContext = new (window.AudioContext || window.webkitAudioContext)();
        const src = this.audioContext.createMediaStreamSource(stream);
        const bufferSize = 4096;
        this.processor = this.audioContext.createScriptProcessor(bufferSize, 1, 1);
        this.processor.onaudioprocess = (e) => {
          const input = e.inputBuffer.getChannelData(0);
          const down = this.downsampleBuffer(input, this.audioContext.sampleRate, 16000);
          const pcm16 = this.floatTo16BitPCM(down);
          this.sampleBuffer = this.concatUint8(this.sampleBuffer, pcm16);
        };
        src.connect(this.processor);
        this.processor.connect(this.audioContext.destination);
      } catch (e) {
        this.setStatus("麦克风权限或设备不可用");
      }
    },

    cleanupAudio() {
      if (this.processor) {
        try {
          this.processor.onaudioprocess = null;
          this.processor.disconnect();
        } catch (e) {}
        this.processor = null;
      }
      if (this.mediaStream) {
        try {
          this.mediaStream.getTracks().forEach((t) => t.stop());
        } catch (e) {}
        this.mediaStream = null;
      }
      if (this.audioContext) {
        try {
          this.audioContext.close();
        } catch (e) {}
        this.audioContext = null;
      }
      this.sampleBuffer = new Uint8Array(0);
      this.firstFrameSent = false;
    },
  },
};
</script>

<style scoped>
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

.modal-content {
  background: #dce9f9;
  padding: 20px;
  border-radius: 20px;
  width: 50vw;
  height: 33vw;
  position: relative;
  overflow: hidden;
}

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
  top: 3.8vw;
  left: 0;
  overflow-y: scroll;
}
.Chat_area::-webkit-scrollbar {
  width: 0.8em;
}
.Chat_area::-webkit-scrollbar-thumb {
  background-color: transparent;
}

.input-wrapper {
  position: absolute;
  top: 31.4vw;
  left: 15px;
  width: 44vw;
  height: 3vw;
}

.Enter_text {
  width: 100%;
  height: 100%;
  padding-right: 3.5vw; /* 给右侧图标留空间 */
  font-size: 1.6vw;
  border-radius: 20px;
  background: #b7d2f0;
  text-indent: 0.6vw;
  box-sizing: border-box;
  border: none;
  outline: none;
}

.mic-icon {
  position: absolute;
  top: 50%;
  right: 0.8vw;
  transform: translateY(-50%);
  width: 2.4vw;
  height: 2.4vw;
  cursor: pointer;
  user-select: none;
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
  width: 3vw;
  height: 3vw;
  cursor: pointer;
  position: absolute;
  top: 0.6vw;
  right: 6.5vw;
  object-fit: contain;
}
</style>
