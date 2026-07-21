<template>
  <div>
    <div class="Nitrogen">
      <p>土壤氮：{{this.$store.state.Go_IoTAE_N_DataArray[7].toFixed(2)}}mg/kg</p>
      <hr class="line" />
      <div class="Nitrogen_Echart" id="Nitrogen_Echart"></div>
    </div>
  </div>
</template>

<script>
import * as echarts from "echarts/core";
import { GridComponent } from "echarts/components";
import { BarChart } from "echarts/charts";
import { CanvasRenderer } from "echarts/renderers";

echarts.use([GridComponent, BarChart, CanvasRenderer]);
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
    this.update_Nitrogen_echarts();
    //监听窗口改变，重新渲染
    this.$watch("$store.state.Go_IoTAE_N_DataArray", () => {
      this.update_Nitrogen_echarts();
    });
  },
  methods: {
    initChart() {
      this.myChart = echarts.init(document.getElementById("Nitrogen_Echart"));
    },
    update_Nitrogen_echarts() {
      const option = {
        tooltip: {
          trigger: "axis",
        },
        grid: {
          left: "5%",
          containLabel: true,
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
            data: this.$store.state.Go_IoTAE_N_DataArray,
            type: "bar",
            lineStyle: {
              color: "green",
            },
            itemStyle: {
              color: "green",
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
.Nitrogen {
  position: relative;
  height: 16vw;
  width: 12vw;
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
.Nitrogen_Echart {
  position: relative;
  top: -1vw;
  left: -0.8vw;
  height: 21vw;
  width: 13vw;
}
.Nitrogen p {
  position: absolute;
  top: -0.5vw;
  left: 0.3vw;
  font-weight: bold;
  color: aliceblue;
}
.line {
  position: absolute;
  width: 12vw;
  top: 1.7vw;
}
</style>