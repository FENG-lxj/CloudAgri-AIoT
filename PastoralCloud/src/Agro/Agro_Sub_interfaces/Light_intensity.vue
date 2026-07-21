<template>
  <div>
    <div class="Light_intensity">
      <p>光照强度：{{ this.$store.state.Go_IoTAE_Light_intensity_DataArray[7].toFixed(2) }}lx</p>
      <hr class="line" />
      <div class="Light_intensity_Echart" id="Light_intensity_Echart"></div>
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
  mounted() {
    //初始化echarts
    this.initChart();
    //初始化图表渲染数据
    this.update_Light_intensity_echarts();
    //监听窗口改变，重新渲染
    this.$watch("$store.state.Go_IoTAE_Light_intensity_DataArray", () => {
      this.update_Light_intensity_echarts();
    });
  },
  methods: {
    initChart() {
      this.myChart = echarts.init(
        document.getElementById("Light_intensity_Echart")
      );
    },
    update_Light_intensity_echarts() {
      const option = {
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
        grid: {
          left: "2%",
          containLabel: true,
        },

        series: [
          {
            data: this.$store.state.Go_IoTAE_Light_intensity_DataArray,
            type: "bar",
            showBackground: true,
            backgroundStyle: {
              color: "rgba(180, 180, 180, 0.2)",
            },
          },
        ],
        itemStyle: {
          color: "#008000", // 设置柱状图颜色为绿色
        },
      };

      this.myChart.setOption(option, true);
    },
  },
};
</script>



<style scoped>
.Light_intensity {
  position: relative;
  height: 10vw;
  width: 17.5vw;
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
.Light_intensity_Echart {
  position: relative;
  top: -1vw;
  left: 0.3vw;
  height: 15.2vw;
  width: 18.5vw;
}
.Light_intensity p {
  position: absolute;
  font-weight: bold;
  top: -0.5vw;
  left: 4.5vw;
  color: aliceblue;
}
.line {
  position: absolute;
  width: 17.5vw;
  top: 1.7vw;
}
</style>