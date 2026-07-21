<template>
  <div class="box">
    <div class="top">
      <img
        src="../img/Device_Management/logo.jpg"
        alt=""
        class="logo"
        @click="home"
      />
      <div class="logo_text" @click="home">螭耀科技</div>
      <div class="vertical_line"></div>
      <div class="logo_text_two" @click="home">云端田园</div>
      <span class="Vertical">|</span>
      <img
        src="../img/SuperAdmin/AI.jpg"
        alt=""
        class="UserImg"
      />
      <span class="phone">{{this.$store.state.Phone}}</span>
      <span class="Super"> 超级管理员 </span>
    </div>
    <div class="left_box">
      <p class="left_box_title">二级管理员</p>
      <div class="addAdmin">
        <img
          src="../img/SuperAdmin/add.png"
          alt=""
          class="addAdmin_Img"
          @click="dialogFormVisible = true"
        />
      </div>
      <hr class="line" />
      <div class="list">
        <div v-for="(item, index) in list" :key="index" class="Admin_One">
          <div class="ID">{{ item.headImg }}</div>
          <div class="Phone">{{ item.phone }}</div>
          <div class="Address">{{ item.last_login_ip }}</div>
          <img
            src="../img/SuperAdmin/Trash_Can.png"
            alt=""
            class="Trash_Can_Img"
            @click="Trash_Can_btn(item.phone)"
          />
        </div>
      </div>
    </div>
    <div class="GPT">
      <div class="title_txt">AI API Token</div>
      <hr class="GPT_line" />
      <input
        type="text"
        placeholder="New API Token"
        class="GPT_Input"
        v-model="GPT_key"
      />
      <button class="Replace_Key" @click="Replace_Key_btn">更换</button>
    </div>
    <!-- 内存环形进度条 -->
    <div class="Memory">
      <p class="Memory_title">服务器内存使用情况</p>
      <div class="Memory_title_progress">
        <svg
          :width="circleSize"
          :height="circleSize"
          xmlns="http://www.w3.org/2000/svg"
        >
          <circle
            class="progress-bar-background"
            :cx="circleSize / 2"
            :cy="circleSize / 2"
            :r="circleRadius"
          ></circle>
          <circle
            class="progress-bar"
            :cx="circleSize / 2"
            :cy="circleSize / 2"
            :r="circleRadius"
            :style="memory_progressStyle"
          ></circle>
          <text x="50%" y="50%" text-anchor="middle" dy="8">
            <tspan :font-size="numberFontSize">{{ MemPer }}</tspan>
            <tspan font-size="12">%</tspan>
          </text>
        </svg>
        <div class="Memory_message">
          <div class="Memory_message_one">{{ memory_percentage }}GB/2GB</div>
          <!-- <button class="Memory_message_btn">设置告警</button> -->
          <div class="Memory_message_one_two">已使用</div>
        </div>
      </div>
    </div>

    <!-- CPU环形进度条 -->
    <div class="CPU">
      <p class="CPU_title">服务器CPU使用情况</p>
      <div class="CPU_title_progress">
        <svg
          :width="circleSize"
          :height="circleSize"
          xmlns="http://www.w3.org/2000/svg"
        >
          <circle
            class="progress-bar-background"
            :cx="circleSize / 2"
            :cy="circleSize / 2"
            :r="circleRadius"
          ></circle>
          <circle
            class="progress-bar"
            :cx="circleSize / 2"
            :cy="circleSize / 2"
            :r="circleRadius"
            :style="CPU_progressStyle"
          ></circle>
          <text x="50%" y="50%" text-anchor="middle" dy="8">
            <tspan :font-size="numberFontSize">{{ CPUper }}</tspan>
            <tspan font-size="12">%</tspan>
          </text>
        </svg>
        <div class="CPU_message">
          <div class="CPU_message_one">2核心</div>
          <!-- <button class="CPU_message_btn">设置告警</button> -->
          <div class="CPU_message_one_two">已使用</div>
        </div>
      </div>
    </div>
    <img
      src="../img/SuperAdmin/EMQX.png"
      alt=""
      class="EMQX_Img"
      @click="EMQX"
    />
    <img src="../img/SuperAdmin/influxdb.png" alt="" class="influxdb_Img" @click="influxdb"/>
    <!-- 添加用户 -->
    <el-dialog
      title="添加管理用户"
      :visible.sync="dialogFormVisible"
      class="AddManagement"
      :append-to-body="true"
    >
      <el-form :model="form">
        <el-form-item label="用户手机号" :label-width="formLabelWidth">
          <el-input v-model="form.Addphone" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="AddManagement_btn">确定</el-button>
      </div>
    </el-dialog>
    <!-- 删除管理员 -->
    <el-dialog
      title="提示"
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="handleClose"
      :append-to-body="true"
    >
      <span>是否删除该管理员</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="Remove_Administrator"
          >确 定</el-button
        >
      </span>
    </el-dialog>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      list: [],
      // 添加管理员
      dialogFormVisible: false,
      form: {
        Addphone: "",
      },
      formLabelWidth: "120px",
      //删除管理员
      dialogVisible: false,
      //更换GPT-Key
      GPT_key: "",
      // 初始百分比
      memory_percentage: 0,
      CPU_percentage: 0,
      MemPer: 0,
      CPUper: 0,
      numberFontSize: 30,
      circleSize: 120,
    };
  },
  mounted() {
    this.getList();
    this.getServerData();
  },
  computed: {
    circleRadius() {
      return this.circleSize / 2 - 4; // 减去一些边距
    },
    memory_progressStyle() {
      const offset =
        2 * Math.PI * this.circleRadius * (1 - this.MemPer / 100);
      return {
        strokeDasharray: `${2 * Math.PI * this.circleRadius} ${
          2 * Math.PI * this.circleRadius
        }`,
        strokeDashoffset: `${offset}`,
        transform: "rotate(-90deg)", // 顺时针旋转90度
        transformOrigin: "center", // 以中心为旋转点
      };
    },
    CPU_progressStyle() {
      const offset =
        2 * Math.PI * this.circleRadius * (1 - this.CPUper / 100);
      return {
        strokeDasharray: `${2 * Math.PI * this.circleRadius} ${
          2 * Math.PI * this.circleRadius
        }`,
        strokeDashoffset: `${offset}`,
        transform: "rotate(-90deg)", // 顺时针旋转90度
        transformOrigin: "center", // 以中心为旋转点
      };
    },
  },
  watch: {},
  methods: {
    // 获取服务器数据
    getServerData() {
      axios
        .get("http://127.0.0.1:20201/IoTA/bigdata/GetMemCPU", {
          headers: {
            Authorization: `Bearer ${this.$store.state.phoneCaptcha}`,
          },
        })
        .then((res) => {
          this.MemPer = res.data.data.MemPer;
          this.CPUper = res.data.data.CPUPer;
          // 更新内存和CPU的百分比，使用 Math.round 进行四舍五入
          this.memory_percentage = parseFloat((res.data.data.MemPer / 50).toFixed(2));
          this.CPU_percentage = res.data.data.CPUper / 300;
          console.log("服务器success:", res.data.data.CPUPer);
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
        });
    },
    //获取管理员列表
    getList() {
      axios
        .get("http://127.0.0.1:20201/IoTA/SuperAdmin/GetUserList", {
          headers: {
            Authorization: `Bearer ${this.$store.state.phoneCaptcha}`,
          },
        })
        .then((res) => {
          this.list = res.data.data.Users;
          console.log("success:", res.data.code);
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
        });
    },
    //添加管理圆通知
    notify_AddManagement_success() {
      const h = this.$createElement;

      this.$notify({
        title: "添加成功",
        type: "success",
      });
    },
    notify_AddManagement_error() {
      const h = this.$createElement;

      this.$notify({
        title: "添加失败，请重试",
        type: "error",
      });
    },
    //添加管理
    AddManagement_btn() {
      this.dialogFormVisible = false;
      axios
        .post(
          "http://127.0.0.1:20201/IoTA/SuperAdmin/SuperAdminAddUser",
          {
            phone: this.form.Addphone,
          },
          {
            headers: {
              Authorization: `Bearer ${this.$store.state.phoneCaptcha}`,
            },
          }
        )
        .then((res) => {
          console.log("success:", res);
          this.notify_AddManagement_success();
          this.getList();
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
          this.notify_AddManagement_error();
        });
    },
    //删除管理员
    Trash_Can_btn(phone) {
      console.log(phone);
      this.dialogVisible = true;
      this.$store.commit("setRemovePhone", phone);
      console.log(this.$store.state.RemovePhone);
    },
    //删除管理员
    Remove_Administrator() {
      this.dialogVisible = false;
      axios
        .post(
          "http://127.0.0.1:20201/IoTA/SuperAdmin/RemoveUser",
          {
            RemovePhone: this.$store.state.RemovePhone,
          },
          {
            headers: {
              Authorization: `Bearer ${this.$store.state.phoneCaptcha}`,
            },
          }
        )
        .then((res) => {
          console.log("success:", res);
          this.notify_Remove_success();
          this.getList();
          this.$store.commit("RemovePhone", 0);
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
          this.notify_Remove_error();
        });
    },
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((_) => {
          done();
        })
        .catch((_) => {});
    },
    //删除管理员通知
    notify_Remove_success() {
      const h = this.$createElement;

      this.$notify({
        title: "删除成功",
        type: "success",
      });
    },
    notify_Remove_error() {
      const h = this.$createElement;

      this.$notify({
        title: "删除失败，请重试",
        type: "error",
      });
    },
        //更换KEY
        Key_success() {
      const h = this.$createElement;

      this.$notify({
        title: "更换成功",
        type: "success",
      });
    },
    Key_error() {
      const h = this.$createElement;

      this.$notify({
        title: "更换失败，请重试",
        type: "error",
      });
    },
    //更换GPT-Key
    Replace_Key_btn() {
      axios
        .post(
          "http://127.0.0.1:20201/IoTA/SuperAdmin/UpChatGPTKey",
          {
            APIKEY: this.GPT_key,
          },
          {
            headers: {
              Authorization: `Bearer ${this.$store.state.phoneCaptcha}`,
            },
          }
        )
        .then((res) => {
          console.log("success:", res);
          this.Key_success()
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
          this.Key_error()
        });
    },
    home() {
      this.$router.replace("home");
    },
    EMQX() {
      window.open("http://127.0.0.1:18083/#/login?to=/dashboard/overview");
    },
    influxdb() {
      window.open("http://127.0.0.1:8086/signin");
    },
  },
};
</script>

<style scoped>
.box {
  position: fixed;
  height: 100%;
  width: 100%;
  top: 0;
  left: 0;
}
.top {
  position: fixed;
  height: 4vw;
  width: 100%;
  background: #000;
}
.logo {
  position: absolute;
  border-radius: 15vw;
  height: 3vw;
  width: 3vw;
  top: 0.5vw;
  left: 0.8vw;
  cursor: pointer;
}
.logo_text {
  position: absolute;
  left: 4.5vw;
  top: 1.2vw;
  color: white;
  font-size: 1.2vw;
  cursor: pointer;
}
.logo_text_two {
  position: absolute;
  left: 10.2vw;
  top: 1.2vw;
  color: white;
  font-size: 1.2vw;
  cursor: pointer;
}
.vertical_line {
  position: absolute;
  border-left: 1px solid rgb(177, 177, 177);
  height: 2vw;
  left: 9.7vw;
  top: 1.1vw;
}
.UserImg {
  position: absolute;
  width: 3vw;
  height: 3vw;
  left: 87vw;
  top: 0.5vw;
}
.Vertical {
  position: relative;
  left: 86vw;
  top: 0.1vw;
  font-size: 5vh;
  font-weight: lighter;
  color: #ffffff;
}
.phone {
  color: #ffffff;
  position: absolute;
  width: 10vw;
  height: 3vw;
  left: 90.8vw;
  top: 0.5vw;
  font-size: 1.3vw;
}
.Super {
  color: #ffffff;
  position: absolute;
  width: 10vw;
  height: 3vw;
  left: 91vw;
  top: 2.5vw;
  font-size: 0.8vw;
}
.left_box {
  position: fixed;
  height: 69vh;
  top: 6.3vw;
  left: 2vw;
  width: 56vw;
  border-radius: 1vw;
  /* border-style: solid; */
  border-color: #9e9e9e;
  /* box-shadow: 2px 2px 3px #9e9e9e; */
  box-shadow: 5px 4px 8px 0 rgb(85, 98, 98); /* 添加阴影效果 */
  transition: transform 0.4s;
}
.left_box:hover {
  transform: scale(1.02);
}
.left_box_title {
  position: relative;
  font-size: 1.5vw;
  font-weight: 800;
  top: -0.7vw;
  left: 1vw;
}
.addAdmin {
  position: absolute;
  top: 1vw;
  left: 53vw;
  height: 2vw;
  width: 2vw;
}
.addAdmin_Img {
  height: 2vw;
  width: 2vw;
}
.line {
  position: absolute;
  top: 3vw;
  width: 55.9vw;
}
.list {
  position: relative;
  left: 1vw;
  width: 54vw;
}
.Admin_One {
  display: flex;
  border-bottom: 2px solid #9e9e9e;
  margin-bottom: 1vw;
}
.ID {
  position: relative;
  left: 2vw;
  font-size: 1.2vw;
  font-weight: 600;
}
.Phone {
  position: relative;
  left: 9vw;
  font-size: 1.2vw;
  font-weight: 600;
}
.Address {
  position: relative;
  left: 20vw;
  font-size: 1.2vw;
  font-weight: 600;
}
.Trash_Can_Img {
  position: absolute;
  left: 52vw;
  margin-top: -0.5vw;
  height: 2vw;
}
.ID .Phone .Address .Trash_Can_Img {
  font-size: 1.2vw;
  font-weight: 600;
}
.GPT {
  position: absolute;
  top: 84vh;
  left: 2vw;
  height: 13vh;
  width: 56vw;
  border-radius: 1vw;
  /* box-shadow: 2px 2px 3px #9e9e9e; */
  /* border: 2px solid #9e9e9e; */
  box-shadow: 5px 4px 8px 0 rgb(85, 98, 98); /* 添加阴影效果 */
  transition: transform 0.4s;
}
.GPT:hover {
  transform: scale(1.02);
}
.title_txt {
  position: absolute;
  left: 2vw;
  top: 0.5vh;
  font-size: 1.3vw;
  font-weight: 600;
}
.GPT_line {
  position: absolute;
  top: 4vh;
  left: 1vw;
  width: 53.8vw;
}
.GPT_Input {
  position: absolute;
  top: 7vh;
  left: 1vw;
  height: 2vw;
  width: 49vw;
  border-radius: 1vw;
  border: 2px solid #9e9e9e;
  font-size: 16px;
  padding-left: 10px;
}
.Replace_Key {
  position: absolute;
  top: 7vh;
  left: 51.5vw;
  height: 2.2vw;
  width: 3.5vw;
  font-weight: 600;
  border-radius: 1vw;
  background: #ffffff;
  border: 0px solid #9e9e9e;
  box-shadow: 2px 2px 3px #9e9e9e;
  cursor: pointer;
}
.Memory {
  position: absolute;
  top: 12vh;
  left: 60vw;
  height: 48vh;
  width: 18vw;
  background: #ffffff;
  border-radius: 1vw;
  /* border: 2px solid #9e9e9e; */
  /* box-shadow: 2px 2px 3px #9e9e9e; */
  box-shadow: 5px 4px 8px 0 rgb(85, 98, 98); /* 添加阴影效果 */
  transition: transform 0.4s;
}
.Memory:hover {
  transform: scale(1.02);
}
.Memory_title,
.CPU_title {
  position: absolute;
  left: 1vw;
  top: -0.5vh;
  font-size: 1.3vw;
  font-weight: 600;
}
.Memory_title_progress,
.CPU_title_progress {
  position: absolute;
  top: 150px;
  left: 77px;
}
.Memory_message,
.CPU_message {
  position: absolute;
}
.Memory_message_btn,
.CPU_message_btn {
  position: absolute;
  top: -16vh;
  left: -5.2vw;
  width: 6vw;
  background: none;
  border: none;
  padding: 0;
  font: inherit;
  cursor: pointer;
  outline: inherit;
  font-size: 13px;
  color: #006eff;
}
.Memory_message_one {
  position: absolute;
  top: -20vh;
  left: 0.5vw;
  font-size: 18px;
  width: 5vw;
}
.CPU_message_one{
  position: absolute;
  top: -20vh;
  left: 2.5vw;
  font-size: 18px;
  width: 5vw;
}
.Memory_message_one_two,
.CPU_message_one_two {
  position: absolute;
  top: 1vh;
  left: 2.7vw;
  font-size: 15px;
  width: 5vw;
  color: #9e9e9e;
}
.CPU {
  position: absolute;
  top: 12vh;
  left: 80vw;
  height: 48vh;
  width: 18vw;
  background: #ffffff;
  border-radius: 1vw;
  /* border: 2px solid #9e9e9e; */
  /* box-shadow: 2px 2px 3px #9e9e9e; */
  box-shadow: 5px 4px 8px 0 rgb(85, 98, 98); /* 添加阴影效果 */
  transition: transform 0.4s;
}
.CPU:hover {
  transform: scale(1.02);
}
.EMQX_Img {
  position: absolute;
  height: 18vw;
  left: 60vw;
  top: 63.3vh;
  border-radius: 1vw;
  cursor: pointer;
  box-shadow: 5px 4px 8px 0 rgb(85, 98, 98); /* 添加阴影效果 */
  transition: transform 0.4s;
}
.EMQX_Img:hover {
  transform: scale(1.02);
}
.influxdb_Img {
  position: absolute;
  height: 18vw;
  left: 80vw;
  top: 63.3vh;
  border-radius: 1vw;
  cursor: pointer;
  box-shadow: 5px 4px 8px 0 rgb(85, 98, 98); /* 添加阴影效果 */
  transition: transform 0.4s;
}
.influxdb_Img:hover {
  transform: scale(1.02);
}
.progress-container {
  text-align: center;
}

.progress-bar-background {
  fill: none;
  stroke: #ddd;
  stroke-width: 2;
}

.progress-bar {
  fill: none;
  stroke: #006eff;
  stroke-width: 2;
  transition: stroke-dashoffset 0.5s ease-out;
}
</style>