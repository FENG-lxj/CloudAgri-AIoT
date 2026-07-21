<template>
	<!--  外部view  -->
	<view class="tab-bar">
		<view class="semicircle">

		</view>
		<!--  循环每个 item 菜单 -->
		<view v-for="(item, index) in tabBarList" :key="index"
			:class="{'tab-bar-item': true, 'currentTar': selected == item.id}" @click="switchTab(item, index)">

			<img class="tab_img" :src="selected == index ? item.selectedIconPath : item.iconPath">

		</view>
	</view>
</template>

<script>
	export default {
		props: {
			selected: {
				type: Number,
				default: 0
			}
		},
		// 在组件创建之前执行
		beforeCreate() {
			// 定义变量
			this.color = '#000'
			this.selectedColor = '#ffb2b2'
			this.tabBarList = [{
					"id": 0,
					"pagePath": "/pages/index/index",
					"iconPath": "https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/%E7%94%9F%E6%80%81%E7%9B%91%E6%B5%8B.png",
					"selectedIconPath": "https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/%E7%94%9F%E6%80%81%E7%9B%91%E6%B5%8B1.png",
				},
				{
					"id": 1,
					"pagePath": "/pages/Monitoring_Points/Monitoring_Points",
					"iconPath": "https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/%E4%BC%81%E4%B8%9A%E7%9B%91%E6%B5%8B%E7%82%B9%E4%BF%A1%E6%81%AF.png",
					"selectedIconPath": "https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/%E4%BC%81%E4%B8%9A%E7%9B%91%E6%B5%8B%E7%82%B9%E4%BF%A1%E6%81%AF1.png",
				},
				{
					"id": 2,
					"pagePath": "/pages/GPT/GPT",
					"iconPath": "https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/ChatGPT.png",
					"selectedIconPath": "https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/ChatGPT.png",
				},
				{
					"id": 3,
					"pagePath": "/pages/AI_Surveillance/AI_Surveillance",
					"iconPath": "https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/AI%E7%9B%91%E6%8E%A7.png",
					"selectedIconPath": "https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/AI%E7%9B%91%E6%8E%A71.png",
				},
				{
					"id": 4,
					"pagePath": "/pages/Settings/Settings",
					"iconPath": "https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/%E8%AE%BE%E7%BD%AE.png",
					"selectedIconPath": "https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/%E8%AE%BE%E7%BD%AE1.png",
				}
			]
		},
		// 定义方法
		methods: {
			// 跳转tabBar菜单栏
			switchTab(item) {
				let url = item.pagePath
				uni.switchTab({
					url
				})
			}
		}
	}
	// 隐藏原生TabBar
	uni.hideTabBar();
</script>

<style lang="less" scoped>
	.semicircle {
		width: 150rpx;
		height: 75rpx;
		background-color: #ffffff;
		border-radius: 0 0 100px 100px;
		/* 只将下左角和下右角设为半圆 */
		overflow: hidden;
		/* 确保超出部分被隐藏 */
		position: absolute;
		margin-top: -66rpx;
	}

	// 外部装修
	.tab-bar {
		z-index: 9999;
		position: fixed;
		bottom: 0rpx;
		left: 0rpx;
		right: 0rpx;
		height: 100rpx;
		background: white;
		padding: 20rpx;
		// border-radius: 30rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: #262F3E;
		// box-shadow: 0 4px 15px rgba(165, 168, 171, 0.83) !important;

		// 当前点击的
		.currentTar {
			border-radius: 15rpx;
			// box-shadow: 0 0 15px #D7D7D7FF !important;
			transition: all 0.5s ease-in-out;
		}

		// 给每个 item 设置样式
		.tab-bar-item {
			//flex: 0.5;
			text-align: center;
			display: flex;
			justify-content: center;
			align-items: center;
			flex-direction: column;
			width: 150rpx;
			padding: 15rpx;
			background-color: transparent;
			transition: all 0.5s ease-in-out;
			margin: auto;

			// 限制每个icon的大小
			.tab_img {
				width: 80rpx;
				height: 80rpx;
			}

			// 限制文字大小

		}
	}

	.tab-bar-item:nth-child(4) {
		height: 80rpx;
		width: 80rpx;
		transform: translateY(-80rpx);
		/* 向上移动20px */
		border-radius: 300rpx;
		background-color: #495AD4;
		/* 添加背景颜色，将 "yourColor" 替换为你想要的颜色值 */
		box-shadow: 1rpx 1rpx 15px #939393ff;
	}
</style>