<template>
  <div>
    <div id="container"></div>
    <div class="right_box">
      <div
        v-for="(device, index) in DevInfo_data"
        :key="index"
        class="Equipment_One"
        @click="openModal(index)"
      >
        <div class="ID">设备ID:{{ device.DevID }}</div>
        <div class="State" v-if="device.Status != '0'">状态:在线</div>
        <img
          v-if="device.Status != '0'"
          src="../../img/Device_Management/Green_Dots.png"
          alt=""
          class="Red_Dots"
        />
        <div class="State" v-if="device.Status == '0'">状态:离线</div>
        <img
          v-if="device.Status == '0'"
          src="../../img/Device_Management/Red_Dots.png"
          alt=""
          class="Red_Dots"
        />
        <div class="Longitude">经度:{{ device.Longitude }}</div>
        <div class="Dimension">维度:{{ device.Dimension }}</div>
        <img
          src="../../img/Device_Management/Instrument_Control.png"
          alt=""
          class="Equipment_Img"
        />
        <div class="Cropper">区域种植作物:{{ device.Cropper }}</div>
      </div>
    </div>
    <!-- 修改设备弹窗 -->
    <div v-if="isModalOpen" class="modal">
      <div class="modal-content">
        <!-- <span @click="closeModal" class="close">&times;</span> -->
        <img
          src="../../img/Device_Management/close_img.png"
          alt=""
          class="close"
          @click="closeModal"
        />
        <p>设备ID：{{ Modal_DevID }}</p>
        <p v-if="this.Modal_Status == 1">状态：正常</p>
        <p v-if="this.Modal_Status == 0">状态：异常</p>
        <p>
          经度：{{ Modal_Longitude }}
          <img
            src="../../img/Device_Management/revise.png"
            alt=""
            class="revise"
            @click="Modal_revise('Modal_Longitude')"
          />
        </p>
        <p>
          维度：{{ Modal_Dimension }}
          <img
            src="../../img/Device_Management/revise.png"
            alt=""
            class="revise"
            @click="Modal_revise('Modal_Dimension')"
          />
        </p>
        <p>
          区域种植作物：{{ Modal_Cropper }}
          <img
            src="../../img/Device_Management/revise.png"
            alt=""
            class="revise"
            @click="Modal_revise('Modal_Cropper')"
          />
        </p>
        <img
          src="../../img/Device_Management/Red_Dots.png"
          alt=""
          class="state_Img"
          v-if="this.Modal_Status == 0"
        />
        <img
          src="../../img/Device_Management/Green_Dots.png"
          alt=""
          class="state_Img"
          v-if="this.Modal_Status == 1"
        />
        <img
          src="../../img/Device_Management/Instrument_Control.png"
          alt=""
          class="modal_Equipment_Img"
        />
        <button class="Save_Configuration" @click="Save_Configuration">
          保存配置
        </button>
        <button class="Delete_Device" @click="Delete_Device">删除设备</button>
      </div>
    </div>
    <el-dialog
      title="提示"
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="handleClose"
    >
      <span>确认删除</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="dialogVisible = false"
          >确 定</el-button
        >
      </span>
    </el-dialog>
  </div>
</template>
  <script>
import AMapLoader from "@amap/amap-jsapi-loader";
import axios from "axios";
export default {
  name: "map-view",
  data() {
    return {
      isModalOpen: false,
      dialogVisible: false,
      //储存AE数据
      DevInfo_data: {},
      map: null,
      positions: [],
      markers: [], //标记点文本
      //修改设备数据
      Modal_DevID: "",
      Modal_Status: "",
      Modal_Longitude: "",
      Modal_Dimension: "",
      Modal_Cropper: "",
    };
  },
  created() {
    this.get_data_Agricultural_Equipment();
  },
  mounted() {
    console.log("0", this.$store.state.phoneCaptcha);
  },
  unmounted() {
    this.map?.destroy();
  },
  methods: {
    get_data_Agricultural_Equipment() {
      axios
        .get("http://127.0.0.1:20201/IoTA/bigdata/GetDevInfo", {
          params: {
            Types: "AE",
          },
          headers: {
            Authorization: `Bearer ${this.$store.state.phoneCaptcha}`,
          },
        })
        .then((res) => {
          // 更新 DevInfo_data 数据
          this.DevInfo_data = res.data.data.DevInfo;
          console.log("获取设备管理列表信息", res.data.data.DevInfo);
          console.log("获取设备管理列表信息", this.DevInfo_data);
          this.positions = this.DevInfo_data.map((item) => [
            item.Longitude,
            item.Dimension,
          ]);
          this.initAMap();
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
        });
    },
    openModal(index) {
      (this.Modal_DevID = this.DevInfo_data[index].DevID),
        (this.Modal_Status = this.DevInfo_data[index].Status),
        (this.Modal_Longitude = this.DevInfo_data[index].Longitude),
        (this.Modal_Dimension = this.DevInfo_data[index].Dimension),
        (this.Modal_Cropper = this.DevInfo_data[index].Cropper),
        (this.isModalOpen = true);
    },
    closeModal() {
      this.isModalOpen = false;
    },
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((_) => {
          done();
        })
        .catch((_) => {});
    },
    initAMapMarker(position, text) {
      const marker = new AMap.Marker({
        icon: new AMap.Icon({
          size: new AMap.Size(60, 46),
          image: require("@/img/Device_Management/location_Two.png"),
          imageSize: new AMap.Size(60, 46),
        }),
        position: position,
        offset: new AMap.Pixel(-13, -30),
      });

      // 创建文字标注
      const textMarker = new AMap.Text({
        text: text, // 文字内容
        position: position, // 标注位置
        offset: new AMap.Pixel(0, -25), // 文字标注偏移量
        style: {
          position: "absolute",
          fontSize: 12,
          fontWeight: "bold",
          color: "#000",
          background: "rgba(255,255,255,0)",
          border: "0px",
          top: "-30px",
          left: "-16px",
        },
      });

      this.map.add(marker);
      this.map.add(textMarker);

      // 存储标点及其文字信息，以便后续操作
      this.markers.push({ marker, textMarker });
    },
    initAMapMarker_Error(position, text) {
      const marker = new AMap.Marker({
        icon: new AMap.Icon({
          size: new AMap.Size(60, 46),
          image: require("@/img/Device_Management/location_One.png"),
          imageSize: new AMap.Size(60, 46),
        }),
        position: position,
        offset: new AMap.Pixel(-13, -30),
      });

      // 创建文字标注
      const textMarker = new AMap.Text({
        text: text, // 文字内容
        position: position, // 标注位置
        offset: new AMap.Pixel(0, -25), // 文字标注偏移量
        style: {
          position: "absolute",
          fontSize: 12,
          fontWeight: "bold",
          color: "#000",
          background: "rgba(255,255,255,0)",
          border: "0px",
          top: "-30px",
          left: "-16px",
        },
      });

      this.map.add(marker);
      this.map.add(textMarker);

      // 存储标点及其文字信息，以便后续操作
      this.markers.push({ marker, textMarker });
    },
    initAMap() {
      AMapLoader.load({
        key: "e479c3ada256dd8af6beff5da9a1df31",
        version: "2.0",
        plugins: [],
      })
        .then((AMap) => {
          this.map = new AMap.Map("container", {
            viewMode: "3D",
            zoom: 11,
            center:
              this.DevInfo_data.length > 0
                ? [
                    this.DevInfo_data[0].Longitude,
                    this.DevInfo_data[0].Dimension,
                  ]
                : [0, 0],
          });

          this.DevInfo_data.forEach((item) => {
            const position = [item.Longitude, item.Dimension];
            const text = `设备ID：${item.DevID}`;
            if(item.Status == 0)
            {
              this.initAMapMarker(position, text);
            }
            else
            {
              this.initAMapMarker_Error(position, text);
            }
          });
        })
        .catch((e) => {
          console.log(e);
        });
    },
    //修改设备
    Modal_revise(revise_text) {
      this.$prompt("请输入修改内容", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        // inputErrorMessage: '邮箱格式不正确'
      })
        .then(({ value }) => {
          // console.log(value);
          // console.log(revise_text);
          if (revise_text == "Modal_Longitude") {
            this.Modal_Longitude = value;
          }
          if (revise_text == "Modal_Dimension") {
            this.Modal_Dimension = value;
          }
          if (revise_text == "Modal_Cropper") {
            this.Modal_Cropper = value;
          }
        })
        .catch(() => {});
    },
    Save_Configuration() {
      axios
        .post(
          "http://127.0.0.1:20201/IoTA/SuperAdmin/UpDev",
          {
            Types: "AE",
            DevID: this.Modal_DevID,
            Status: this.Modal_Status,
            Longitude: parseInt(this.Modal_Longitude),
            Dimension: parseInt(this.Modal_Dimension),
            Cropper: this.Modal_Cropper,
          },
          {
            headers: {
              Authorization: `Bearer ${this.$store.state.phoneCaptcha}`,
            },
          }
        )
        .then((res) => {
          console.log("修改设备成功", res.data);
          this.get_data_Agricultural_Equipment();
          this.isModalOpen = false;
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
          console.log("修改设备失败");
        });
    },
    Delete_Device() {
      axios
        .post(
          "http://127.0.0.1:20201/IoTA/SuperAdmin/RemoveDev",
          {
            Types: "IC",
            DevID: this.Modal_DevID,
          },
          {
            headers: {
              Authorization: `Bearer ${this.$store.state.phoneCaptcha}`,
            },
          }
        )
        .then((res) => {
          console.log("删除设备成功", res.data);
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
          console.log("删除设备失败");
        });
    },
  },
};
</script>
  <style scoped>
#container {
  position: absolute;
  width: 700px;
  height: 713px;
}
.right_box {
  position: absolute;
  left: 700px;
  height: 713px;
  width: 700px;
}
.Equipment_One {
  position: relative;
  height: 154px;
  width: 700px;
  background: #ebebeb;
  margin-top: 1vw;
  left: 1vw;
  border-radius: 1vw;
  box-shadow: 2px 2px 5px #888888;
  transition: transform 0.3s;
}
.Equipment_One:hover {
  transform: scale(1.01);
}
.Red_Dots {
  position: absolute;
  height: 18px;
  left: 100px;
  top: 45px;
  border-radius: 1000px;
}
.ID {
  position: absolute;
  top: 8px;
  left: 20px;
  font-size: 18px;
  font-weight: 600;
}
.State {
  position: absolute;
  top: 41px;
  left: 20px;
  font-size: 18px;
  font-weight: 600;
}
.Longitude {
  position: absolute;
  top: 71px;
  left: 20px;
  font-size: 18px;
  font-weight: 600;
}
.Dimension {
  position: absolute;
  top: 100px;
  left: 20px;
  font-size: 18px;
  font-weight: 600;
}
.Cropper {
  position: absolute;
  top: 126px;
  left: 20px;
  font-size: 18px;
  font-weight: 600;
}
.Equipment_Img {
  position: absolute;
  height: 88px;
  width: 88px;
  left: 590px;
  top: 30px;
}
/* 弹窗 */
.modal {
  position: fixed;
  top: 50%;
  left: 50%;
  height: 200px;
  width: 450px;
  transform: translate(-50%, -50%);
  background-color: #fefefe;
  border-radius: 10px;
  z-index: 9999;
  box-shadow: 2px 2px 3px #9e9e9e;
  text-align: left;
}

.modal-content {
  position: relative;
  top: 10px;
}
p {
  margin-left: 35px;
  margin-bottom: -15px;
  font-weight: bold;
}
.close {
  position: absolute;
  top: -23px;
  right: 5px;
  color: #aaa;
  font-size: 28px;
  font-weight: bold;
  cursor: pointer;
  height: 35px;
}
.state_Img {
  position: absolute;
  height: 16px;
  margin-left: 118px;
  margin-top: -69px;
}
.modal_Equipment_Img {
  position: absolute;
  height: 80px;
  margin-left: 324px;
  margin-top: -67px;
}
.revise {
  position: relative;
  height: 15px;
  top: 1px;
  cursor: pointer;
}

.Save_Configuration {
  height: 37px;
  width: 130px;
  margin-left: 64px;
  margin-top: 26px;
  font-weight: bold;
  background: #1ed806;
  border: 0px;
  border-radius: 14px;
  box-shadow: 2px 2px 2px #9e9e9e;
  cursor: pointer;
}
.Delete_Device {
  position: absolute;
  height: 37px;
  width: 130px;
  margin-left: 50px;
  margin-top: 26px;
  font-weight: bold;
  background: #d81e06;
  border: 0px;
  border-radius: 14px;
  box-shadow: 2px 2px 2px #9e9e9e;
  cursor: pointer;
}
</style>
  