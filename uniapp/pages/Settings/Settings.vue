<template>
	<view class="box">
		<view class="user">
			<img src="@/img/Setting/userImg.png" alt="" class="UserImg"/>
			<view class="Phone">
				{{phone}}
			</view>
			<view class="text" v-if="this.phone==18607229637">
				超级管理员
			</view>
			<view class="text" v-else>
				管理员
			</view>
			<img src="../../img/index/next.png" alt="" class="usernext" @click="nextsuper"/>
		</view>
		
		<view class="Equipment">
			<img src="../../img/index/Last.png" alt="" class="EquipmentLast" @click="previousEquipment"/>
			<view class="title">
				{{ equipmentTypes[currentEquipmentIndex] }}
			</view>
			<img src="../../img/index/next.png" alt="" class="Equipmentnext" @click="nextEquipment"/>
		</view>
		<!-- 根据当前设备类型的索引显示对应的组件 -->
		<component :is="currentComponent"></component>
		<tabber :selected="4"></tabber>
	</view>
</template>

<script>
import tabber from '../../tabbar/tabbar.vue'
import Agricultural from './Module/Agricultural_Equipment.vue'
import AIDevices from './Module/AI_Devices.vue'
import StorageDevices from './Module/Storage_Devices.vue'

export default {
	components: {
		tabber,
		Agricultural,
		StorageDevices,
		AIDevices,
		
	},
	data() {
		return {
			currentEquipmentIndex: 0,
			equipmentTypes: ["农田环境采集设备", "粮仓环境采集设备", "农田AI监控设备"],
			phone:0,
		}
	},
	computed: {
		// 根据当前设备类型索引返回对应的组件名称
		currentComponent() {
			switch (this.currentEquipmentIndex) {
				case 0:
					return 'Agricultural';
				case 1:
					return 'StorageDevices';
				case 2:
					return 'AIDevices';
				default:
					return null;
			}
		}
	},
	mounted() {
		this.phone = uni.getStorageSync('phone');
	},
	methods: {
		nextsuper(){
			uni.navigateTo({
			  url: '/pages/Settings/superAdmin/superAdmin', 
			  success: (res) => {
			    console.log('跳转成功');
			  },
			});
		},
		previousEquipment() {
			if (this.currentEquipmentIndex > 0) {
				this.currentEquipmentIndex--;
			}
		},
		nextEquipment() {
			if (this.currentEquipmentIndex < this.equipmentTypes.length - 1) {
				this.currentEquipmentIndex++;
			}
		}
	}
}
</script>

<style scoped>
	@font-face {
		font-family: DingTalk_JinBuTi;
		src: url('../../img/font/DingTalk_JinBuTi.ttf');
	}
	.box{
		position: fixed;
		height: 100%;
		background-color: #ffffff;
	}

	.user {
		position: relative;
		height: 150rpx;
		width: 700rpx;
		margin: 25rpx;
		background-color: #F2F2F2;
		border-radius: 15rpx;
		box-shadow: 0px 0px 7px #202547;
		border: 1rpx solid #8c8d8d00;
	}
	.UserImg{
		position: absolute;
		height: 110rpx;
		margin-top: 20rpx;
		margin-left: 25rpx;
	}
	.Phone{
		position: absolute;
		margin-left: 150rpx;
		margin-top: 40rpx;
		font-size: 42rpx;
		font-weight: bold;
	}
	.text{
		position: absolute;
		margin-left: 153rpx;
		margin-top: 85rpx;
		font-size: 30rpx;
		font-weight: bold;
		font-family: DingTalk_JinBuTi;
	}
	.usernext{
		position: absolute;
		height: 100rpx;
		margin-left: 600rpx;
		margin-top: 30rpx;
		cursor: pointer;
	}
	.Equipment{
		position: relative;
		height: 70rpx;
		width: 700rpx;
		margin: 25rpx;
		font-family: DingTalk_JinBuTi;
		font-size: 36rpx;
		background-color: #F2F2F2;
		border-radius: 30rpx;
		box-shadow: 0px 0px 7px #202547;
		border: 1rpx solid #8c8d8d00;
	}
	.title{
		margin-left: 54%;						/* 向右移动 50% */
		transform: translateX(-50%);	/* 向左移动自身宽度的 50% */
		margin-top: 15rpx;
	}
	.EquipmentLast{
		position: absolute;
		height: 70rpx;
		cursor: pointer;
	}
	.Equipmentnext{
		position: absolute;
		height: 70rpx;
		margin-left: 620rpx;
		margin-top: -55rpx;
		cursor: pointer;
	}
</style>
