<template>
	<view>
		<view class="box">
			<view class="title">甲醇浓度&nbsp;&nbsp;{{ this.$store.state.Go_IoTAE_Methanol_DataArray[7] }}&nbsp;mg/kg</view>
			<img ref="image" src="../../../img/index/unfold.png" alt="" class="unfold_img" @click="rotateImage" />
			<hr class="line" ref="Line" v-if="rotation !== 0" />
			<view class="echarts" ref="echarts_box">
				<view class="Methanol_Echart" id="Methanol_Echart"></view>
			</view>
		</view>
	</view>
</template>

<script>
import * as echarts from 'echarts/core';
import { TitleComponent, ToolboxComponent, TooltipComponent, GridComponent, LegendComponent, MarkLineComponent, MarkPointComponent } from 'echarts/components';
import { BarChart } from 'echarts/charts';
import { CanvasRenderer } from 'echarts/renderers';

echarts.use([TitleComponent, ToolboxComponent, TooltipComponent, GridComponent, LegendComponent, MarkLineComponent, MarkPointComponent, BarChart, CanvasRenderer]);

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
		this.update_Methanol_echarts();
		//监听窗口改变，重新渲染
		this.$watch('$store.state.Go_IoTAE_Methanol_DataArray', () => {
			this.update_Methanol_echarts();
			console.log(this.$store.state.Go_IoTAE_Methanol_DataArray);
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
			this.myChart = echarts.init(document.getElementById('Methanol_Echart'));
		},
		update_Methanol_echarts() {
			const option = {
				title: {
					// text: "Rainfall vs Evaporation",
					// subtext: "Fake Data",
				},
				tooltip: {
					trigger: 'axis'
				},
				legend: {
					show: false, // 将show设置为false关闭图例
					data: ['Rainfall', 'Evaporation']
				},
				toolbox: {
					show: false,
					feature: {
						dataView: { show: false, readOnly: false },
						magicType: { show: false, type: ['line', 'bar'] },
						restore: { show: false },
						saveAsImage: { show: false }
					}
				},
				calculable: false,
				xAxis: [
					{
						type: 'category',
						// prettier-ignore
						// data: this.$store.state.Time,
						data: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
					}
				],
				yAxis: [
					{
						type: 'value'
					}
				],
				grid: {
					left: '5%',
					containLabel: true
				},
				series: [
					{
						name: 'Evaporation',
						type: 'bar',
						data: this.$store.state.Go_IoTAE_Methanol_DataArray,
						itemStyle: {
							color: '#7474FA' // 设置柱状图颜色为绿色
						},
						markPoint: {
							data: [
								{ name: 'Max', value: 182.2, xAxis: 7, yAxis: 183 },
								{ name: 'Min', value: 2.3, xAxis: 11, yAxis: 3 }
							]
						}

						// markLine: {
						//   data: [{ type: "average", name: "Avg" }],
						// },
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
.Methanol_Echart {
	margin-top: -90rpx;
	height: 500rpx;
	width: 700rpx;
	margin-left: 0rpx;
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
