<template>
	<view>
		<view class="box">
			<view class="title">空气温度&nbsp;&nbsp;{{ this.$store.state.Go_IoTAE_Air_temperature_DataArray[7] }}&nbsp;°C</view>
			<img ref="image" src="../../../img/index/unfold.png" alt="" class="unfold_img" @click="rotateImage" />
			<hr class="line" ref="Line" v-if="rotation !== 0" />
			<view class="echarts" ref="echarts_box">
				<view class="Air_Temperature_Echart" id="Air_Temperature_Echart"></view>
			</view>
		</view>
	</view>
</template>

<script>
import * as echarts from 'echarts/core';
import { TitleComponent, ToolboxComponent, TooltipComponent, GridComponent, LegendComponent, MarkLineComponent, MarkPointComponent } from 'echarts/components';
import { LineChart } from 'echarts/charts';
import { UniversalTransition } from 'echarts/features';
import { CanvasRenderer } from 'echarts/renderers';

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
	UniversalTransition
]);
export default {
	data() {
		return {
			myChart: null,
			rotation: 0
		};
	},
	mounted() {
		//初始化echarts
		this.initChart();
		//初始化图表渲染数据
		this.update_Air_Temperature_echarts();
		//监听窗口改变，重新渲染
		this.$watch('$store.state.Go_IoTAE_Air_temperature_DataArray', () => {
			this.update_Air_Temperature_echarts();
		});
	},
	methods: {
		rotateImage() {
			this.rotation += 90;
			const echartsBox = this.$refs.echarts_box.$el; // 获取echarts_box的DOM元素
			if (this.rotation === 180) {
				this.rotation = 0;
				echartsBox.style.display = 'none';
			} else {
				echartsBox.style.display = 'block';
			}
			const img = this.$refs.image;
			img.style.transform = `rotate(${this.rotation}deg)`;
		},

		initChart() {
			this.myChart = echarts.init(document.getElementById('Air_Temperature_Echart'));
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
					type: 'category',
					boundaryGap: false,
					// data: this.$store.state.Time,
					data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
					axisLabel: {
						textStyle: {
							// color: 'white' // 设置 x 轴文字颜色为白色
						}
					}
				},
				tooltip: {
					trigger: 'axis',
					axisPointer: {
						type: 'cross'
					}
				},
				yAxis: {
					type: 'value',
					boundaryGap: true,
					axisLabel: {
						formatter: '{value}',
						textStyle: {
							// color: 'white'
						}
					}
				},

				series: [
					{
						name: 'Highest',
						type: 'line',
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
							color: '#ff9300'
						},
						itemStyle: {
							color: '#ff9300'
						}
					}
				]
			};

			this.myChart.setOption(option, true);
		}
	}
};
</script>

<style scoped>
@font-face {
	font-family: DingTalk_JinBuTi;
	src: url('../../../img/font/DingTalk_JinBuTi.ttf');
}
.box {
	position: relative;
	/* height: 400rpx; */
	width: 700rpx;
	margin: 25rpx;
	border-radius: 30rpx;
	box-shadow: 0px 0px 7px #202547;
	border: 1rpx solid #8c8d8d00;
}
.title {
	font-size: 35rpx;
	font-family: DingTalk_JinBuTi;
	margin-top: 10rpx;
	margin-left: 20rpx;
	margin-bottom: 10rpx;
}
.line {
	margin-top: 10rpx;
	width: 700rpx;
	border: 1rpx solid #bdbebe;
}
.echarts {
	display: none;
}
.Air_Temperature_Echart {
	margin-top: -90rpx;
	height: 590rpx;
	width: 700rpx;
	margin-bottom: -90rpx;
}
.unfold_img {
	position: absolute;
	top: 10rpx;
	right: 10rpx;
	height: 50rpx;
	width: 50rpx;
	cursor: pointer;
	z-index: 999;
}
</style>
