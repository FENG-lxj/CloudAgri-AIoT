<template>
	<view>
		<image src="@/img/index/AI_surveillance.png" class="swiper"></image>
		<ID></ID>
		<view>
			<video class="AI_Video" :src="videoSrc" controls></video>
		</view>
		<Detect_Diseases></Detect_Diseases>
		<Farm_Invasion></Farm_Invasion>
		<Log></Log>
		<AI></AI>
		<view class="bottom-ui"></view>
		<tabber :selected="3"></tabber>
	</view>
</template>

<script>
import tabber from '../../tabbar/tabbar.vue';
import ID from './Module/ID.vue';
import Detect_Diseases  from './Module/Detect_Diseases.vue'
import Farm_Invasion from './Module/Farm_Invasion.vue'
import Log from './Module/Log.vue'
import AI from './Module/AI.vue'

export default {
	components: {
		ID,
		Detect_Diseases,
		Farm_Invasion,
		Log,
		AI,
		tabber,
	},
	data() {
		return {
			videoSrc: "http://39.102.208.105:8080/PastoralCloud/monitor_1.flv",
		};
	},
	mounted() {
		
	},
	methods: {
		
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
.AI_Video{
	height: 325rpx;
	width: 685rpx;
	margin-left: 30rpx;
	border-radius: 20rpx;
	box-shadow: 6rpx 6rpx 5rpx #8c8d8d;
}
.bottom-ui{
	height: 200rpx;
}
</style>
