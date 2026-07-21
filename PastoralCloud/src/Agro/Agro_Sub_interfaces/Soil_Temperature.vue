<template>
  <div>
    <div class="Soil_Temperature">
      <p>土壤温度：{{this.$store.state.Go_IoTAE_ST_DataArray[7].toFixed(2)}}℃</p>
      <hr class="line" />
      <div class="Soil_Temperature_Echart" id="Soil_Temperature_Echart"></div>
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
  VisualMapComponent,
  MarkAreaComponent,
} from "echarts/components";
import { LineChart } from "echarts/charts";
import { UniversalTransition } from "echarts/features";
import { CanvasRenderer } from "echarts/renderers";

echarts.use([
  TitleComponent,
  ToolboxComponent,
  TooltipComponent,
  GridComponent,
  VisualMapComponent,
  MarkAreaComponent,
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
    this.update_Soil_Temperature_echarts();
    //监听窗口改变，重新渲染
    this.$watch("$store.state.Go_IoTAE_ST_DataArray", () => {
      this.update_Soil_Temperature_echarts();
    });
  },
  methods: {
    initChart() {
      this.myChart = echarts.init(
        document.getElementById("Soil_Temperature_Echart")
      );
    },
    update_Soil_Temperature_echarts() {
      const option = {
        title: {
          // text: "Distribution of Electricity",
          // subtext: "Fake Data",
        },
        tooltip: {
          trigger: "axis",
          axisPointer: {
            type: "cross",
          },
        },
        toolbox: {
          show: false,
          feature: {
            saveAsImage: {},
          },
        },
        xAxis: {
          type: "category",
          boundaryGap: false,
          // prettier-ignore
          // data: ['00:00', '01:15', '02:30', '03:45', '05:00', '06:15', '07:30', '08:45', '10:00', '11:15', '12:30', '13:45', '15:00', '16:15', '17:30', '18:45', '20:00', '21:15', '22:30', '23:45'],
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
            formatter: "{value} W",
          },
          axisPointer: {
            snap: true,
          },
          axisLabel: {
            textStyle: {
              color: "white", // 设置 x 轴文字颜色为白色
            },
          },
        },
        visualMap: {
          show: false,
          dimension: 0,
          pieces: [
            {
              lte: 3,
              color: "#FF9300",
            },
            {
              gt: 3,
              lte: 5,
              color: "#FF9300",
            },
            {
              gt: 5,
              lte: 14,
              color: "#FF9300",
            },
            {
              gt: 14,
              lte: 17,
              color: "#FF9300",
            },
            {
              gt: 17,
              color: "#FF9300",
            },
          ],
        },
        series: [
          {
            name: "Electricity",
            type: "line",
            smooth: true,
            // prettier-ignore
            data: this.$store.state.Go_IoTAE_ST_DataArray,
          },
        ],
      };

      this.myChart.setOption(option, true);
    },
  },
};
</script>

<style scoped>
.Soil_Temperature {
  position: relative;
  height: 13vw;
  width: 19.5vw;
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
.Soil_Temperature_Echart {
  position: relative;
  top: 0vw;
  left: 0vw;
  height: 16vw;
  width: 21vw;
}
.Soil_Temperature p {
  position: absolute;
  top: -0.5vw;
  left: 5.5vw;
  font-weight: bold;
  color: aliceblue;
}
.line {
  position: absolute;
  width: 19.5vw;
  top: 1.7vw;
}
</style>