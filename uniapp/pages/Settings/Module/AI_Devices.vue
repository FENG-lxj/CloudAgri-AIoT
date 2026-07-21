<template>
	<view>
		<view class="page-body">
			<view class="page-section page-section-gap">
				<map id="map-container" class="mapS" :latitude="latitude" :longitude="longitude" :markers="covers"></map>
			</view>
		</view>
		<view class="content">
			<view v-for="(device, index) in DevInfo_data" :key="index" class="box" @click="openModifyModal(device)">
				<view class="ID">设备ID: {{ device.DevID }}</view>
				<view class="Status" v-if="device.Status != '0'">状态: 在线</view>
				<img v-if="device.Status != '0'" src="../../../img/Setting/Green_Dots.png" alt="" class="Dots_img" />
				<view class="Status" v-if="device.Status == '0'">状态: 离线</view>
				<img v-if="device.Status == '0'" src="../../../img/Setting/Red_Dots.png" alt="" class="Dots_img" />
				<view class="latitude">经度: {{ device.Longitude }}</view>
				<view class="longitude">维度: {{ device.Dimension }}</view>
				<view class="zuowu">区域种植作物: {{ device.Cropper }}</view>
				<img src="../../../img/Setting/Instrument_Control.png" alt="" class="shebei_img" />
			</view>
		</view>

		<!-- //修改设备信息窗口 -->
		<view class="xiugai-confirm-modal" v-if="isOpen">
			<view class="modal-content">
				<view class="modal-item">
					<text>设备ID：</text>
					<!-- 这里显示当前设备的ID -->
					<text>{{ selectedDevice.DevID }}</text>
				</view>
				<view class="modal-item">
					<text>经度：</text>
					<!-- 这里添加经度的输入框 -->
					<input type="text" v-model="selectedDevice.Longitude" />
				</view>
				<view class="modal-item">
					<text>维度：</text>
					<!-- 这里添加维度的输入框 -->
					<input type="text" v-model="selectedDevice.Dimension" />
				</view>
				<view class="modal-item">
					<text>区域种植作物：</text>
					<!-- 这里添加区域种植农作物的输入框或选择框 -->
					<input type="text" v-model="selectedDevice.Cropper" />
				</view>
				<view class="modal-item">
					<!-- 这里添加确认修改按钮 -->
					<button @click="confirmModification">确认修改</button>
					<button @click="deletemodal">删除设备</button>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
import AMapLoader from '@amap/amap-jsapi-loader';

export default {
	data() {
		return {
			covers: [],
			DevInfo_data: [],
			longitude: 0,
			latitude: 0,
			mapInitialized: false,
			// isOpen: true,
			selectedDevice: {}, // 新增选中设备对象
			isOpen: false
		};
	},
	mounted() {
		this.initAMap();
	},
	methods: {
		openModifyModal(device) {
			// 打开修改设备信息窗口
			this.isOpen = true;
			// 将选中设备的信息传递给修改窗口
			this.selectedDevice = device;
		},
		confirmModification() {
			console.log(this.selectedDevice);
			uni.request({
				url: 'http://127.0.0.1:20201/IoTA/SuperAdmin/UpDev',
				method: 'POST',
				header: {
					Authorization: `Bearer ${uni.getStorageSync('token')}`
				},
				data: {
					Types: 'IC',
					DevID: this.Modal_DevID,
					Status: this.Modal_Status,
					Longitude: parseInt(this.Modal_Longitude),
					Dimension: parseInt(this.Modal_Dimension),
					Cropper: this.Modal_Cropper
				},
				success: (res) => {
					console.log('修改设备成功', res.data);
					this.get_data_Agricultural_Equipment();
					// this.isModalOpen = false;
					this.isOpen = false;
				},
				fail: (error) => {
					console.error('Error fetching data:', error);
					console.log('修改设备失败');
				}
			});
		},
		deletemodal() {
			uni.request({
				url: 'http://127.0.0.1:20201/IoTA/SuperAdmin/RemoveDev',
				method: 'POST',
				header: {
					Authorization: `Bearer ${uni.getStorageSync('token')}`
				},
				data: {
					Types: 'IC',
					DevID: this.Modal_DevID
				},
				success: (res) => {
					console.log('删除设备成功', res.data);
					this.isOpen = false;
				},
				fail: (error) => {
					console.error('Error fetching data:', error);
					console.log('删除设备失败');
				}
			});
		},
		async initAMap() {
			try {
				const AMap = await AMapLoader.load({
					key: 'e479c3ada256dd8af6beff5da9a1df31',
					version: '2.0',
					plugins: []
				});

				const map = new AMap.Map('map-container', {
					// 使用地图容器的 id
					viewMode: '3D',
					zoom: 11,
					center: [0, 0] // 默认中心点，待数据加载后更新
				});

				this.mapInitialized = true;
				this.get_data_Agricultural_Equipment();
			} catch (error) {
				console.error('Error initializing AMap:', error);
			}
		},
		get_data_Agricultural_Equipment() {
			uni.request({
				url: 'http://127.0.0.1:20201/IoTA/bigdata/GetDevInfo',
				method: 'GET',
				data: {
					Types: 'IC'
				},
				header: {
					Authorization: `Bearer ${uni.getStorageSync('token')}`
				},
				success: (res) => {
					console.log(res);
					if (res.statusCode === 200) {
						this.DevInfo_data = res.data.data.DevInfo;
						this.renderMarkers();
					} else {
						console.error('Error fetching data:', res.data);
					}
				},
				fail: (error) => {
					console.error('Error fetching data:', error);
				}
			});
		},
		renderMarkers() {
			if (!this.mapInitialized) return; // 等待地图初始化完成
			const map = new AMap.Map('map-container');
			map.setCenter([this.DevInfo_data[0].Longitude, this.DevInfo_data[0].Dimension]);
			this.covers = this.DevInfo_data.map((device) => ({
				latitude: device.Dimension,
				longitude: device.Longitude,
				iconPath: device.Status === '0' ? 'https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/location_Two.png' : 'https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/location_One.png',
				width: '40',
				height: '35'
			}));

			this.covers.forEach((cover) => {
				const marker = new AMap.Marker({
					icon: new AMap.Icon({
						size: new AMap.Size(60, 46),
						image: cover.iconPath,
						imageSize: new AMap.Size(60, 46)
					}),
					position: [cover.longitude, cover.latitude],
					offset: new AMap.Pixel(-13, -30),
					map: map
				});
			});
		}
	}
};
</script>

<style scoped>
@font-face {
	font-family: DingTalk_JinBuTi;
	src: url('../../../img/font/DingTalk_JinBuTi.ttf');
}
.content {
	margin-top: 30rpx;
	width: 690rpx;
	margin-left: 30rpx;
	height: 275rpx;
	position: relative;
	background-color: #ebebeb;
	border-radius: 25rpx;
	box-shadow: 6rpx 6rpx 10rpx #8c8d8d;
}

/* 删除确认弹窗样式 */
.xiugai-confirm-modal {
	position: fixed;
	top: 0;
	left: 0;
	right: 0;
	bottom: 0;
	background-color: rgba(0, 0, 0, 0.5); /* 半透明黑色背景 */
	display: flex;
	align-items: center;
	justify-content: center;
	z-index: 9999; /* 确保弹窗在最顶层 */
}
/* .modal-content {
	height: 300rpx;
	width: 90%;
	background-color: rgba(237, 238, 238);
	border-radius: 30rpx;
} */
.modal-content {
	height: auto;
	width: 90%;
	background-color: rgba(237, 238, 238);
	border-radius: 30rpx;
	padding: 20rpx;
}

.modal-item {
	margin-bottom: 20rpx;
	font-family: DingTalk_JinBuTi;
	font-size: 35rpx;
}

.modal-label {
	font-size: 32rpx;
	color: #333333;
}

.modal-input {
	width: 100%;
	height: 60rpx;
	border: 1rpx solid #cccccc;
	border-radius: 10rpx;
	padding: 10rpx;
	font-family: DingTalk_JinBuTi;
	font-size: 35rpx;
	margin-top: 50rpxs;
}

.modal-text {
	font-family: DingTalk_JinBuTi;
	font-size: 35rpx;
	color: #666666;
}

.modal-button {
	width: 100%;
	height: 60rpx;
	background-color: #409eff;
	border: none;
	border-radius: 10rpx;
	color: #ffffff;
	font-family: DingTalk_JinBuTi;
	font-size: 35rpx;
	text-align: center;
	line-height: 60rpx;
	cursor: pointer;
}
.xiugai_img {
	position: relative;
	height: 35rpx;
	margin-left: 20rpx;
}
/* 123 */
.box {
	position: absolute;
	margin-top: 10rpx;
}
.longitude,
.latitude,
.ID,
.Status,
.zuowu {
	margin-left: 20rpx;
	margin-top: 7rpx;
	font-family: DingTalk_JinBuTi;
	font-size: 35rpx;
}
.shebei_img {
	position: absolute;
	height: 160rpx;
	margin-top: -190rpx;
	margin-left: 490rpx;
}
.Dots_img {
	position: absolute;
	height: 35rpx;
	margin-top: -40rpx;
	margin-left: 180rpx;
}
.modal {
	position: fixed;
	height: 300rpx;
	width: 80%;
	z-index: 999;
	margin-top: -600rpx;
	margin-left: 10%;
	background-color: aqua;
}
.mapS{
	width: 90%; 
	height: 500rpx; 
	margin-left: 5%;
	box-shadow: 0px 0px 7px #202547;
	border: 1rpx solid #8c8d8d00;
}
</style>
