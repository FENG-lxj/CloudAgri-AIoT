<template>
  <div>
    <div class="Air_Humidity">
      <p>空气湿度：{{ this.$store.state.Go_IoTAE_AH_DataArray[7].toFixed(2) }}%</p>
      <hr class="line" />
      <div class="Air_Humidity_Echart" id="Air_Humidity_Echart"></div>
    </div>
  </div>
</template>

<script>
import * as echarts from "echarts/core";
import {
  TitleComponent,
  ToolboxComponent,
  TooltipComponent,
  GridComponent,
  LegendComponent,
} from "echarts/components";
import { LineChart } from "echarts/charts";
import { UniversalTransition } from "echarts/features";
import { CanvasRenderer } from "echarts/renderers";

echarts.use([
  TitleComponent,
  ToolboxComponent,
  TooltipComponent,
  GridComponent,
  LegendComponent,
  LineChart,
  CanvasRenderer,
  UniversalTransition,
]);

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
    this.update_Air_Humidity_echarts();
    //监听窗口改变，重新渲染
    this.$watch("$store.state.Go_IoTAE_AH_DataArray", () => {
      this.update_Air_Humidity_echarts();
    });
  },
  methods: {
    initChart() {
      this.myChart = echarts.init(
        document.getElementById("Air_Humidity_Echart")
      );
    },
    update_Air_Humidity_echarts() {
      const option = {
        title: {
          // text: "Step Line",
        },
        tooltip: {
          trigger: "axis",
          axisPointer: {
            type: "cross",
          },
        },
        legend: {
          show: false, // 将show设置为false关闭图例
          data: ["Step Start", "Step Middle", "Step End"],
        },
        grid: {
          left: "3%",
          right: "4%",
          bottom: "3%",
          containLabel: true,
        },

        xAxis: {
          type: "category",
          data: this.$store.state.Time,
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
            name: "Step Start",
            type: "line",
            step: "start",
            data: this.$store.state.Go_IoTAE_AH_DataArray,
            lineStyle: {
              color: "#FAC858",
            },
            itemStyle: {
              color: "#FAC858",
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
.Air_Humidity {
  position: relative;
  height: 10vw;
  width: 21.5vw;
  border-radius: 1vw;
  box-shadow: 5px 4px 8px 0 rgb(7, 226, 233); /* 添加阴影效果 */

  background-color: hsla(
    180,
    100%,
    50%,
    0.2
  ); /* 设置背景颜色为带有50%透明度的蓝色 */
}
.Air_Humidity_Echart {
  position: relative;
  top: -1vw;
  height: 11.5vw;
  width: 21.5vw;
}
.Air_Humidity p {
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