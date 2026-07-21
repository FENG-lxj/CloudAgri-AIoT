<template>
  <div>
    <div class="Storage_CarbonDioxide">

      <p>二氧化碳浓度:{{ this.$store.state.Go_IoTAS_CO2_DataArray[7].toFixed(2) }}pmm</p>
      <hr />
      <div
        class="Storage_CarbonDioxide_Echart"
        id="Storage_CarbonDioxide_Echart"
      ></div>
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
    this.update_Storage_CarbonDioxide_echarts();
    //监听窗口改变，重新渲染
    this.$watch("$store.state.Go_IoTAS_CO2_DataArray", () => {
      this.update_Storage_CarbonDioxide_echarts();
    });
  },
  methods: {
    initChart() {
      this.myChart = echarts.init(
        document.getElementById("Storage_CarbonDioxide_Echart")
      );
    },
    update_Storage_CarbonDioxide_echarts() {
      const option = {
        tooltip: {
          trigger: "axis",
        },
        xAxis: {
          type: "category",
          data:this.$store.state.Time,
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
            data: this.$store.state.Go_IoTAS_CO2_DataArray,
            type: "line",
            lineStyle: {
              color: "#ff0000",
            },
            itemStyle: {
              color: "#ff0000",
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
.Storage_CarbonDioxide {
  position: relative;
  height: 20.5vw;
  width: 25vw;
  background-color: hsla(180, 100%, 50%, 0.2);
  border-radius: 0.7vw;
  box-shadow: 3px 3px 4px 0 rgb(7, 226, 233);
}
.Storage_CarbonDioxide p {
  position: absolute;
  font-size: 1.2vw;
  font-weight: bold;
  top: -0.8vw;
  color: white;
  left: 5.7vw;
}

.Storage_CarbonDioxide_Echart {
  position: absolute;
  height: 23.5vw;
  width: 26.5vw;
}
.Storage_CarbonDioxide hr {
  position: absolute;
  border: 0; /* 移除默认的边框 */
  height: 0.8px; /* 设置横线的高度 */
  width: 25vw;
  top: 1.7vw;
  background: #ffffff;
  box-shadow: 0 0 10px rgba(255, 255, 255, 0.5); /* 设置模糊阴影效果 */
}
</style>