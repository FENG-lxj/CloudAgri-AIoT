<template>
  <div>
    <div class="AI_Camera_Two">
      <img
        src="../../img/Agro/上一个.png"
        alt=""
        class="AI_Camera_Two_left_Img"
        @click="Add"
      />
      <img :src="Images[index]" alt="" class="AI_Camera_Two_img" />
      <!-- <img src="../../img/病虫害.png" alt="" class="AI_Camera_Two_img" /> -->
      <div class="text-container">
        <p class="Detect_Diseases">
          发现入侵：{{ this.DetectionInfo[this.index].InvadeClass }}
          <!-- 发现病虫害：稻纹枯病 -->
        </p>
        <p class="Time_Discovery">
          发现时间：{{ this.DetectionInfo[this.index].Datetime }}
          <!-- 发现时间：2024-04-8 20:27:06 -->
        </p>
        <p class="ID">
          识别ID：{{ this.DetectionInfo[this.index].IdentifyID }}
        </p>
      </div>
      <img
        src="../../img/Agro/下一个.png"
        alt=""
        class="AI_Camera_Two_right_Img"
        @click="Decrease"
      />
    </div>
  </div>
</template>
<script>
import axios from "axios";

export default {
  data() {
    return {
      DetectionInfo: {},
      AI_Camera_Two_ICID: 1,
      index: 0,
      id: [],
      Images: [] // 存储图片的数组
    };
  },
  created() {
    this.index = 0;
    this.$store.commit("set_AI_Camera_One_ICID", 1);
    this.$watch(
      () => this.$store.state.AI_Camera_Two_ICID,
      (newValue, oldValue) => {
        console.log("AI_Camera_One_ICID changed:", newValue, oldValue);
        this.getdata();
      }
    );
  },
  mounted() {
    this.getdata();
  },
  methods: {
    Add() {
      if (this.index > 0) {
        this.index--;
        console.log("++");
      }
    },
    Decrease() {
      if (this.index < this.DetectionInfo.length - 1) {
        this.index++;
        console.log("--");
      }
    },
    getdata() {
      axios
        .get("http://127.0.0.1:20201/IoTA/ai/DetectionInfo", {
          params: {
            ICID: this.$store.state.AI_Camera_One_ICID,
            Level1Class: 2,
            StartNum: 0,
            Num: 10,
          },
          headers: {
            Authorization: `Bearer ${this.$store.state.phoneCaptcha}`,
          },
        })
        .then((res) => {
          this.DetectionInfo = res.data.data.DetectionInfo;
          console.log("获取动物监测结果", res.data.data.DetectionInfo);

          const detectionInfo_data = res.data.data.DetectionInfo;
          const extracted_id = [];

          detectionInfo_data.forEach((item) => {
            extracted_id.push(item.id);
            this.GET_Img(item.id);
          });

          console.log("32", extracted_id);

          this.id = extracted_id;
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
        });
    },
    GET_Img(id) {
      axios
        .get("http://127.0.0.1:20201/IoTA/ai/DetectionImage", {
          params: {
            Id: id,
            Level1Class: 2,
          },
          headers: {
            Authorization: `Bearer ${this.$store.state.phoneCaptcha}`,
          },
          responseType: "blob", // 声明响应类型为 blob
        })
        .then((res) => {
          const imageUrl = URL.createObjectURL(res.data);
          this.Images.push(imageUrl); // 将图片URL添加到Images数组中
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
        });
    },
  },
};
</script>
<style scoped>
.AI_Camera_Two {
  position: relative;
  height: 17.5vw;
  width: 47vw;
  border-radius: 0.8vw;
  background-color: hsla(180, 100%, 50%, 0.2);
  box-shadow: 3px 3px 4px 0 rgb(7, 226, 233);
  /* 使用flex布局使内部内容居中对齐 */
  display: flex;
  justify-content: center; /* 水平居中对齐 */
  align-items: center; /* 垂直居中对齐 */
}

.text-container {
  display: flex;
  flex-direction: column; /* 竖直排列 */
  align-items: center; /* 居中对齐 */
  margin-top: 193px;
}

.AI_Camera_Two_left_Img {
  position: absolute;
  height: 10vw;
  width: 8vw;
  top: 3.6vw;
  left: 0px;
  cursor: pointer;
}
.AI_Camera_Two_right_Img {
  position: absolute;
  left: 39vw;
  height: 10vw;
  width: 8vw;
  top: 3.6vw;
  cursor: pointer;
}
.AI_Camera_Two_img {
  position: absolute;
  height: 147px;
  width: 258px;
  top: 19px;
}
.Detect_Diseases {
  margin-top: -15px;
  font-size: 18px;
  color: #fff;
}
.Time_Discovery {
  margin-top: -10px;
  font-size: 18px;
  color: #fff;
}
.ID {
  margin-top: -10px;
  font-size: 18px;
  color: #fff;
}
</style>
