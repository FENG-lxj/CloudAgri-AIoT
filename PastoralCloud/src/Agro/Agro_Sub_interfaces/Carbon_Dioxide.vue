<template>
  <div>
    <div class="Carbon_Dioxide">
      <p>二氧化碳浓度：{{ this.$store.state.Go_IoTAE_CO2_DataArray[7].toFixed(2) }}ppm</p>
      <hr class="line" />
      <div class="Carbon_Dioxide_Echart" id="Carbon_Dioxide_Echart"></div>
    </div>
  </div>
</template>


<script>
import * as echarts from "echarts/core";
import { GridComponent } from "echarts/components";
import { LineChart } from "echarts/charts";
import { UniversalTransition } from "echarts/features";
import { CanvasRenderer } from "echarts/renderers";
import store from "@/store";

echarts.use([GridComponent, LineChart, CanvasRenderer, UniversalTransition]);

export default {
  data() {
    return {
      myChart: null,
    };
  },
  mounted() {
    //初始化echarts
    this.initChart();
    //初始化图表渲染数据
    this.update_Carbon_Dioxide_echarts();
    //监听窗口改变，重新渲染
    this.$watch("$store.state.Go_IoTAE_CO2_DataArray", () => {
      this.update_Carbon_Dioxide_echarts();
    });
  },
  methods: {
    initChart() {
      this.myChart = echarts.init(
        document.getElementById("Carbon_Dioxide_Echart")
      );
    },
    update_Carbon_Dioxide_echarts() {
      const option = {
        tooltip: {
          trigger: "axis",
        },
        xAxis: {
          type: "category",
          data: this.$store.state.Time,
          // data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
          axisLabel: {
            textStyle: {
              color: "white", // 设置 x 轴文字颜色为白色
            },
          },
        },
        yAxis: {
          type: "value",
          axisLabel: {
            textStyle: {
              color: "white", // 设置 x 轴文字颜色为白色
            },
          },
        },
        series: [
          {
            data: this.$store.state.Go_IoTAE_CO2_DataArray,
            type: "line",
            smooth: true,
            lineStyle: {
              color: "red",
            },
            itemStyle: {
              color: "red",
            },
          },
        ],
      };

      this.myChart.setOption(option, true);
    },
  },
};
</script>

<style scoped>
.Carbon_Dioxide {
  position: relative;
  height: 12.8vw;
  width: 23vw;
  border-radius: 1vw;
  /* background: #7474FA;
    box-shadow: 0px 3px 6px #7474fabc; */
  background-color: hsla(
    180,
    100%,
    50%,
    0.2
  ); /* 设置背景颜色为带有50%透明度的蓝色 */
  box-shadow: 5px 4px 8px 0 rgb(7, 226, 233); /* 添加阴影效果 */
}
.Carbon_Dioxide_Echart {
  position: relative;
  top: -1vw;
  height: 17vw;
  widows: 23vw;
  left: 1vw;
}
.Carbon_Dioxide p {
  position: absolute;
  top: -0.5vw;
  left: 7vw;
  font-weight: bold;
  color: aliceblue;
}
.line {
  position: absolute;
  width: 23vw;
  top: 1.7vw;
}
</style>