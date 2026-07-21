<template>
	<view>
		<view class="box">
			<view class="title">氧气浓度：{{ this.$store.state.Go_IoTAS_O2_DataArray[7] }}%</view>
			<img ref="image" src="../../../img/index/unfold.png" alt="" class="unfold_img" @click="rotateImage" />
			<hr class="line" ref="Line" v-if="rotation !== 0" />
			<view class="echarts" ref="echarts_box">
				<view class="Air_quality_Echart" id="Air_quality_Echart"></view>
			</view>
		</view>
	</view>
</template>

<script>
import * as echarts from 'echarts/core';
import { GridComponent } from 'echarts/components';
import { LineChart } from 'echarts/charts';
import { UniversalTransition } from 'echarts/features';
import { CanvasRenderer } from 'echarts/renderers';
import { TooltipComponent } from 'echarts/components';

echarts.use([GridComponent, LineChart, CanvasRenderer, UniversalTransition, TooltipComponent]);

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
		this.update_Air_quality_echarts();
		//监听窗口改变，重新渲染
		this.$watch('$store.state.Go_IoTAS_O2_DataArray', () => {
			this.update_Air_quality_echarts();
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
			this.myChart = echarts.init(document.getElementById('Air_quality_Echart'));
		},
		update_Air_quality_echarts() {
			const option = {
				tooltip: {
					trigger: 'axis'
				},
				xAxis: {
					type: 'category',
					// data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
					data: this.$store.state.Time
				},
				tooltip: {
					trigger: 'axis',
					axisPointer: {
						type: 'cross'
					}
				},
				yAxis: {
					type: 'value',
					axisLabel: {
						textStyle: {
							// color: 'white' // 设置 x 轴文字颜色为白色
						}
					}
				},
				series: [
					{
						data: this.$store.state.Go_IoTAS_O2_DataArray,
						type: 'line'
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
	margin-bottom: 10rpx;
	margin-left: 20rpx;
}
.line {
	margin-top: 10rpx;
	width: 700rpx;
	border: 1rpx solid #bdbebe;
}
.echarts {
	display: none;
}
.Air_quality_Echart {
	margin-top: -90rpx;
	height: 500rpx;
	width: 700rpx;
	margin-bottom: -90rpx;
	margin-left: 10rpx;
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
