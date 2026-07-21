<template>
  <div>
    <div class="Air_quality">
      <p class="title">
        空气质量：{{ this.$store.state.Go_IoTAE_Air_quality_DataArray[7].toFixed(2) }}
      </p>
      <hr class="line" />
      <div class="Air_quality_Echart" id="Air_quality_Echart"></div>
    </div>
  </div>
</template>

<script>
import * as echarts from "echarts/core";
import { GridComponent } from "echarts/components";
import { LineChart } from "echarts/charts";
import { UniversalTransition } from "echarts/features";
import { CanvasRenderer } from "echarts/renderers";

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
    this.update_Air_quality_echarts();
    //监听窗口改变，重新渲染
    this.$watch("$store.state.Go_IoTAE_Air_quality_DataArray", () => {
      this.update_Air_quality_echarts();
    });
  },
  methods: {
    initChart() {
      this.myChart = echarts.init(
        document.getElementById("Air_quality_Echart")
      );
    },
    update_Air_quality_echarts() {
      const option = {
        tooltip: {
          trigger: "axis",
        },
        xAxis: {
          type: "category",
          // data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
          data: this.$store.state.Time,
          axisLabel: {
            textStyle: {
              color: "white", // 设置 x 轴文字颜色为白色
            },
          },
        },
        tooltip: {
          trigger: "axis",
          axisPointer: {
            type: "cross",
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
            data: this.$store.state.Go_IoTAE_Air_quality_DataArray,
            type: "line",
          },
        ],
      };

      this.myChart.setOption(option, true);
    },
  },
};
</script>

<style scoped>
.Air_quality {
  position: relative;
  /* background: #7474FA; */
  height: 10.2vw;
  width: 22.3vw;
  border-radius: 1vw;
  background-color: hsla(
    180,
    100%,
    50%,
    0.2
  ); /* 设置背景颜色为带有50%透明度的蓝色 */
  box-shadow: 5px 4px 8px 0 rgb(7, 226, 233); /* 添加阴影效果 */

  /* box-shadow: 0px 3px 6px #7474fabc; */
}
.Air_quality_Echart {
  position: relative;
  top: -1vw;
  left: -0.3vw;
  height: 14.5vw;
  width: 25vw;
}
.title {
  position: absolute;
  left: 7.5vw;
  top: -0.5vw;
  font-weight: bold;
  color: aliceblue;
}
.line {
  position: absolute;
  width: 22.3vw;
  top: 1.7vw;
}
</style>