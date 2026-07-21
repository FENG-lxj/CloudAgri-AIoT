<template>
	<view class="top">
		<image src="@/img/index/agro.png" class="swiper"></image>
		
		<ID></ID>
		<Light_intensity></Light_intensity>
		<Air_quality></Air_quality>
		<Air_Humidity></Air_Humidity>
		<Air_Temperature></Air_Temperature>
		<Carbon_Dioxide></Carbon_Dioxide>
		<Methanol></Methanol>
		<Nitrogen></Nitrogen>
		<Kalium></Kalium>
		<Phosphorus></Phosphorus>
		<Soil_Temperature></Soil_Temperature>
		<Soil_Moisture></Soil_Moisture>
		<Soil_PH></Soil_PH>
		<AI></AI>
		<view class="bottom-ui"></view>
		<tabber :selected="0"></tabber>
	</view>
</template>

<script>
import Go_IoTAE from './Go_IoTAE';
import ID from './Module/ID.vue';
import Air_quality from './Module/Air_quality.vue';
import Air_Humidity from './Module/Air_Humidity.vue'
import Air_Temperature from './Module/Air_Temperature.vue'
import Carbon_Dioxide from './Module/Carbon_Dioxide.vue'
import Kalium from './Module/Kalium.vue'
import Light_intensity from './Module/Light_intensity.vue'
import Methanol from './Module/Methanol.vue'
import Nitrogen from './Module/Nitrogen.vue'
import Phosphorus from './Module/Phosphorus.vue'
import Soil_Moisture from './Module/Soil_Moisture.vue'
import Soil_PH from './Module/Soil_PH.vue'
import Soil_Temperature from './Module/Soil_Temperature.vue'
import AI from './AI.vue'
import tabber from '../../tabbar/tabbar.vue'
export default {
	components: {
		ID,
		Air_quality,
		Air_Humidity,
		Air_Temperature,
		Carbon_Dioxide,
		Kalium,
		Light_intensity,
		Methanol,
		Nitrogen,
		Phosphorus,
		Soil_Moisture,
		Soil_PH,
		Soil_Temperature,
		AI,
		tabber
	},
	data() {
		return {
			swipers: [require('@/img/index/agro.png'), 
						require('@/img/index/AI.png'), 
						require('@/img/index/AI_surveillance.png'), 
						require('@/img/index/Crop_storage.png')]
		};
	},
	mixins: [Go_IoTAE],
	mounted() {
		this.AE_init(); // 初始化
		this.initMqttConnection();
			const now = new Date();
			this.currentMonth = (now.getMonth() + 1).toString();
			this.currentDay = now.getDate().toString();
		
			// 保存过去七个小时的时间
			for (let i = 6; i >= 0; i--) {
				const pastDate = new Date(now);
				pastDate.setHours(now.getHours() - i);
		
				// 格式化为 "某时" 格式
				const formattedTime = `${pastDate.getHours()}：00`;
		
				this.Time.push(formattedTime);
			}
		
			this.$store.commit('set_Time', this.Time);
			console.log('time', this.$store.state.Time);
			this.Time = [];
	},
	methods: {
		// int() {
		// 	var echarts = require('echarts');

		// 	var chartDom = document.getElementById('main');
		// 	var myChart = echarts.init(chartDom);
		// 	var option;

		// 	option = {
		// 		xAxis: {
		// 			type: 'category',
		// 			data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
		// 		},
		// 		yAxis: {
		// 			type: 'value'
		// 		},
		// 		series: [
		// 			{
		// 				data: [150, 230, 224, 218, 135, 147, 260],
		// 				type: 'line'
		// 			}
		// 		]
		// 	};

		// 	option && myChart.setOption(option);
		// }
	}
};
</script>

<style scoped>
.swiper {
	height: 230rpx;
	width: 700rpx;
	margin-top: 25rpx;
	margin-left: 25rpx;
	border-radius: 30rpx;
	box-shadow: 6rpx 6rpx 5rpx #8c8d8d;
}
.swiper-item {
	height: 300rpx;
	width: 700rpx;
	border-radius: 30rpx;
}
.swiper-item-img {
	height: 100%;
	width: 100%;
}
.main {
	height: 300rpx;
	width: 300rpx;
}
.bottom-ui{
	height: 200rpx;
}
</style>
