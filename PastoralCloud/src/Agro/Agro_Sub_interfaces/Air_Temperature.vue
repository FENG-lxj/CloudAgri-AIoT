<template>
  <div>
    <div class="Air_Temperature">
      <p>
        空气温度：{{
          this.$store.state.Go_IoTAE_Air_temperature_DataArray[7].toFixed(2)
        }}℃
      </p>
      <hr class="line" />
      <div class="Air_Temperature_Echart" id="Air_Temperature_Echart"></div>
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
import { LineChart } from "echarts/charts";
import { UniversalTransition } from "echarts/features";
import { CanvasRenderer } from "echarts/renderers";

echarts.use([
  TitleComponent,
  ToolboxComponent,
  TooltipComponent,
  GridComponent,
  LegendComponent,
  MarkLineComponent,
  MarkPointComponent,
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
    this.update_Air_Temperature_echarts();
    //监听窗口改变，重新渲染
    this.$watch("$store.state.Go_IoTAE_Air_temperature_DataArray", () => {
      this.update_Air_Temperature_echarts();
    });
  },
  methods: {
    initChart() {
      this.myChart = echarts.init(
        document.getElementById("Air_Temperature_Echart")
      );
    },
    update_Air_Temperature_echarts() {
      const option = {
        title: {
          // text: "Temperature Change in the Coming Week",
        },
        // tooltip: {
        //   trigger: "axis",
        // },
        // legend: {},
        // toolbox: {
        //   show: true,
        //   feature: {
        //     dataZoom: {
        //       yAxisIndex: "none",
        //     },
        //     dataView: { readOnly: false },
        //     magicType: { type: ["line", "bar"] },
        //     restore: {},
        //     saveAsImage: {},
        //   },
        // },
        xAxis: {
          type: "category",
          // boundaryGap: false,
          data: this.$store.state.Time,
          // data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
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
          boundaryGap: true, 
          axisLabel: {
            formatter: "{value}",
            textStyle: {
              color: "white",
            },
          },
        },

        series: [
          {
            name: "Highest",
            type: "line",
            data: this.$store.state.Go_IoTAE_AH_DataArray,
            // markPoint: {
            //   data: [
            //     { type: "max", name: "Max" },
            //     { type: "min", name: "Min" },
            //   ],
            // },
            // markLine: {
            //   data: [{ type: "average", name: "Avg" }],
            // },
            lineStyle: {
              color: "#ff9300",
            },
            itemStyle: {
              color: "#ff9300",
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
.Air_Temperature {
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
.Air_Temperature_Echart {
  position: relative;
  top: -1vw;
  left: 0vw;
  height: 14.5vw;
  width: 23vw;
}
.Air_Temperature p {
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