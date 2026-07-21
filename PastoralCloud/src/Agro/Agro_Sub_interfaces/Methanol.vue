<template>
  <div>
    <div class="Methanol">
      <p>甲醇浓度：{{ this.$store.state.Go_IoTAE_Methanol_DataArray[7].toFixed(2) }}</p>
      <hr class="line" />
      <div class="Methanol_Echart" id="Methanol_Echart"></div>
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
  MarkLineComponent,
  MarkPointComponent,
} from "echarts/components";
import { BarChart } from "echarts/charts";
import { CanvasRenderer } from "echarts/renderers";

echarts.use([
  TitleComponent,
  ToolboxComponent,
  TooltipComponent,
  GridComponent,
  LegendComponent,
  MarkLineComponent,
  MarkPointComponent,
  BarChart,
  CanvasRenderer,
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
    this.update_Methanol_echarts();
    //监听窗口改变，重新渲染
    this.$watch("$store.state.Go_IoTAE_Methanol_DataArray", () => {
      this.update_Methanol_echarts();
      console.log(this.$store.state.Go_IoTAE_Methanol_DataArray);
    });
  },
  methods: {
    initChart() {
      this.myChart = echarts.init(document.getElementById("Methanol_Echart"));
    },
    update_Methanol_echarts() {
      const option = {
        title: {
          // text: "Rainfall vs Evaporation",
          // subtext: "Fake Data",
        },
        tooltip: {
          trigger: "axis",
        },
        legend: {
          show: false, // 将show设置为false关闭图例
          data: ["Rainfall", "Evaporation"],
        },
        toolbox: {
          show: false,
          feature: {
            dataView: { show: false, readOnly: false },
            magicType: { show: false, type: ["line", "bar"] },
            restore: { show: false },
            saveAsImage: { show: false },
          },
        },
        calculable: false,
        xAxis: [
          {
            type: "category",
            // prettier-ignore
            data: this.$store.state.Time,
            // data: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
            axisLabel: {
              textStyle: {
                color: "white", // 设置 x 轴文字颜色为白色
              },
            },
          },
        ],
        yAxis: [
          {
            type: "value",
            axisLabel: {
              textStyle: {
                color: "white", // 设置 x 轴文字颜色为白色
              },
            },
          },
        ],
        grid: {
          left: "5%",
          containLabel: true,
        },
        series: [
          {
            name: "Evaporation",
            type: "bar",
            data: this.$store.state.Go_IoTAE_Methanol_DataArray,
            itemStyle: {
              color: "#7474FA", // 设置柱状图颜色为绿色
            },
            markPoint: {
              data: [
                { name: "Max", value: 182.2, xAxis: 7, yAxis: 183 },
                { name: "Min", value: 2.3, xAxis: 11, yAxis: 3 },
              ],
            },

            // markLine: {
            //   data: [{ type: "average", name: "Avg" }],
            // },
          },
        ],
      };

      this.myChart.setOption(option, true);
    },
  },
};
</script>

<style scoped>
.Methanol {
  position: relative;
  height: 12.8vw;
  width: 16.3vw;
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
.Methanol_Echart {
  position: relative;
  top: 0vw;
  left: -1vw;
  height: 17vw;
  width: 18vw;
}
.Methanol p {
  position: absolute;
  font-weight: bold;
  top: -0.5vw;
  left: 4vw;
  color: aliceblue;
}
.line {
  position: absolute;
  width: 16.3vw;
  top: 1.7vw;
}
</style>