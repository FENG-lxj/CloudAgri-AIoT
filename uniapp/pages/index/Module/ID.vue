<template>
	<view class="box">
		<view>
			<img src="../../../img/index/Last.png" alt="" class="Last"/>
			<view class="title">
				设备ID：1
			</view>
			<img src="../../../img/index/next.png" alt="" class="Next"/>
		</view>
	</view>
</template>
<style scoped>
	@font-face {
		font-family: DingTalk_JinBuTi;
		src: url('../../../img/font/DingTalk_JinBuTi.ttf');
	}
	.box {
		position: relative;
		height: 60rpx;
		width: 700rpx;
		margin: 25rpx;
		border-radius: 20rpx;
		box-shadow: 0px 0px 7px #202547;
		border: 1rpx solid #8c8d8d00;
	}
	.Last{
		position: absolute;
		margin-top: -10rpx;
		height: 60rpx;
	}
	.title{
	font-size: 35rpx;
	font-family: DingTalk_JinBuTi;
	margin-top: 10rpx;
	margin-left: 260rpx;
	margin-bottom: 10rpx;
	}
	.Next{
		position: absolute;
		margin-top: -63rpx;
		margin-left: 640rpx;
		height: 60rpx;
	}
</style>