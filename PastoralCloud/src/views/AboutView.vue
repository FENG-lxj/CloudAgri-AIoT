<template>
  <div class="about">
    <!-- 文本输入框，用于用户输入消息 -->
    <input type="text" v-model="message" />
    <!-- 提交按钮，点击时触发 subim 方法 -->
    <button @click="subim">提交</button>
  </div>
</template>

<script>
// 引入 mqtt 库
import mqtt from "mqtt";

export default {
  name: "AboutView",
  data() {
    return {
      // mqtt 客户端对象
      client: "",
      // 用户输入的消息内容
      message: "",
    };
  },
  mounted() {
    // 创建 mqtt 连接，并将其赋值给 client 属性
    this.client = mqtt.connect("ws://127.0.0.1:8083/mqtt");

    //订阅主题
    this.client.subscribe("Go_IoTAE");

    //监听消息
    this.client.on("message",(topic,message)=>{
      //收到消息时的处理逻辑
      console.log(`收到来自${topic}主题的消息：${message.toString()}`);
    });
  },
  methods: {
    subim() {
      // 发布用户输入的消息到 "test" 主题
      this.client.publish("test", this.message);
    },
  },
};
</script>
