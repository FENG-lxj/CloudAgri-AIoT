<template>
  <div>
    <div class="Soil_PH">
      <p>土壤PH：{{ this.$store.state.Go_IoTAE_PH_DataArray[7].toFixed(2) }}</p>
      <hr class="line" />
      <div class="Soil_PH_Echart" id="Soil_PH_Echart"></div>
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
    this.update_Soil_PH_echarts();
    //监听窗口改变，重新渲染
    this.$watch("$store.state.Go_IoTAE_PH_DataArray", () => {
      this.update_Soil_PH_echarts();
    });
  },
  methods: {
    initChart() {
      this.myChart = echarts.init(document.getElementById("Soil_PH_Echart"));
    },
    update_Soil_PH_echarts() {
      var chartDom = document.getElementById("Soil_PH_Echart");
      var myChart = echarts.init(chartDom);
      var option;

      option = {
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
            data: this.$store.state.Go_IoTAE_PH_DataArray,
            type: "line",
            lineStyle: {
              color: "#00ff95",
            },
            itemStyle: {
              color: "#00ff95",
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
.Soil_PH {
  position: relative;
  height: 10vw;
  width: 21.5vw;
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
.Soil_PH_Echart {
  position: relative;
  top: -1vw;
  left: 0vw;
  height: 14.5vw;
  width: 23vw;
}
.Soil_PH p {
  position: absolute;
  top: -0.5vw;
  left: 7vw;
  font-weight: bold;
  color: aliceblue;
}
.line {
  position: absolute;
  width: 21.5vw;
  top: 1.7vw;
}
</style>