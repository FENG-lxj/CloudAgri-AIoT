<template>
  <div>
    <div class="AI_Prevention">
      <p>AI监控日志</p>
      <hr />
      <div v-if="MonitorLogs">
        <!-- 显示获取到的建议 -->
        <div class="message">
          <div v-for="item in MonitorLogs" :key="item.Datetime">
            <div :style="{ color: getColor(item.Level1Class) }">
              {{ item.Datetime }} - {{ item.Log }}
            </div>
            <br />
          </div>
        </div>
      </div>
      <div v-else>
        <!-- 如果数据还在加载中，可以显示加载中的提示 -->
        <div class="message">Loading...</div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import mqtt from "mqtt";

export default {
  data() {
    return {
      MonitorLogs: [], // 存储病虫害防治建议的数据
    };
  },
  mounted() {
    this.fetchHistoricalLogs(); // 获取历史日志
    this.initMqttConnection(); // 监听 MQTT 实时数据
  },
  methods: {
    fetchHistoricalLogs() {
      axios
        .get("https:///IoTA/ai/GetMonitorLogs", {
          params: {
            TimePeriod: 500,
          },
          headers: {
            Authorization: `Bearer ${this.$store.state.phoneCaptcha}`,
          },
        })
        .then((res) => {
          // 更新 MonitorLogs 数据
          this.MonitorLogs = res.data.data.MonitorLogs.map((log) => ({
            Datetime: log.Datetime,
            Log: log.Log,
            Level1Class: log.Level1Class, // 添加 Level1Class 字段
          }));
          console.log("获取历史日志", this.MonitorLogs);
          // 等待DOM更新后滚动到底部
          this.$nextTick(() => {
            this.scrollToBottom();
          });
        })
        .catch((error) => {
          console.error("Error fetching historical data:", error);
        });
    },
    initMqttConnection() {
      // 使用 MQTT 协议连接到服务器
      console.log("Connecting to MQTT broker...");

      // 配置 MQTT 连接参数
      const options = {
        username: "qiyao", // 设置用户名
        password: "qiyao123456", // 设置密码
        qos: 2, // 设置 QoS
      };

      this.client = mqtt.connect("ws://127.0.0.1:8083/mqtt", options); // 使用配置参数连接

      // 订阅主题
      this.client.subscribe("Go_IoT_Logs");

      // 监听消息
      this.client.on("message", (topic, message) => {
        const mqttData = JSON.parse(message.toString());
        console.log("Received message:", mqttData);

        // 将 MQTT 数据添加到 MonitorLogs 数组中
        const newLog = {
          Datetime: mqttData.Datetime,
          Log: mqttData.Log,
          Level1Class: mqttData.Level1Class, // 添加 Level1Class 字段
        };
        this.MonitorLogs.push(newLog);

        // 滚动到底部
        this.scrollToBottom();
      });
    },
    scrollToBottom() {
      // 获取消息容器
      const container = this.$el.querySelector(".message");
      // 滚动到底部
      container.scrollTop = container.scrollHeight;
    },
    getColor(level) {
      if (level === 1) return "#fffb00"; // yellow
      if (level === 2) return "#00ff2a"; // green
      if (level === 3) return "#ff0000"; // red
      return "#ffffff"; // default white
    },
  },
};
</script>

<style scoped>
.AI_Prevention {
  position: absolute;
  height: 17.3vw;
  width: 23.5vw;
  margin-top: 1vw;
  border-radius: 0.8vw;
  background-color: hsla(180, 100%, 50%, 0.2);
  box-shadow: 3px 3px 4px 0 rgb(7, 226, 233);
}
.AI_Prevention p {
  position: absolute;
  font-weight: bold;
  top: -0.5vw;
  left: 9.52vw;
  color: #ffffff;
}
.AI_Prevention hr {
  position: absolute;
  top: 2vw;
  height: 0.2px;
  width: 23.4vw;
  background: #ffffff;
}
.message {
  margin-top: 59px;
  margin-left: 15px;
  margin-right: 10px;
  max-height: 189px;
  overflow: auto;
  color: #ffffff;
  font-weight: bold;
  width: 355px;

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
