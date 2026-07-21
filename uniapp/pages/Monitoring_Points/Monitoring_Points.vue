<template>
	<view class="top">
		<image src="@/img/index/Crop_storage.png" class="swiper"></image>
		<ID></ID>
		<Air_Oxygen></Air_Oxygen>
		<Air_Humidity></Air_Humidity>
		<Air_Temperature></Air_Temperature>
		<Carbon_Dioxide></Carbon_Dioxide>
		<AI></AI>
		<view class="bottom-ui"></view>
		<tabber :selected="1"></tabber>
	</view>
</template>

<script>
	import Go_IoTAS from './Go_IoTAS';
import ID from './Module/ID.vue';
import Air_Oxygen from './Module/Air_Oxygen.vue';
import Air_Humidity from './Module/Air_Humidity.vue';
import Air_Temperature from './Module/Air_Temperature.vue';
import Carbon_Dioxide from './Module/Carbon_Dioxide.vue';
import AI from './AI.vue'
import tabber from '../../tabbar/tabbar.vue';
export default {
	components: {
		ID,
		Air_Oxygen,
		Air_Humidity,
		Air_Temperature,
		Carbon_Dioxide,
		AI,
		tabber
	},
	data() {
		return {};
	},
	mixins: [Go_IoTAS],
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
	methods: {}
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
