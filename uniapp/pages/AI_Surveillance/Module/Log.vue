<template>
	<view>
		<view class="box">
			<view class="title">AI监控日志</view>
			<img ref="image" src="../../../img/index/unfold.png" alt="" class="unfold_img" @click="rotateImage" />
			<hr class="line" ref="Line" v-if="rotation !== 0" />
			<view class="echarts" ref="echarts_box">
				<view class="Log">
					<view v-if="MonitorLogs">
						<view class="message">
							<view v-for="item in MonitorLogs" :key="item.id">
								<!-- 将 key 设为每个列表项的唯一标识 -->
								<view v-if="item.Level1Class == 1" style="color: #ffb82a">
									{{ item.Datetime }}
									{{ item.Log }}
								</view>
								<view v-if="item.Level1Class == 2" style="color: #6dd51d">
									{{ item.Datetime }}
									{{ item.Log }}
								</view>
								<view v-if="item.Level1Class == 3" style="color: #ff0000">
									{{ item.Datetime }}
									{{ item.Log }}
								</view>
								<view class="br-ui"></view>
							</view>
						</view>
					</view>
					<view v-else>
						<view class="message">Loading...</view>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import mqtt from 'mqtt';

	export default {
		data() {
			return {
				rotation: 0,
				MonitorLogs: null
			};
		},
		mounted() {
			uni.request({
				url: 'http://127.0.0.1:20201/IoTA/ai/GetMonitorLogs',
				method: 'GET',
				data: {
					TimePeriod: 50
				},
				header: {
					Authorization: `Bearer ${uni.getStorageSync('token')}`
				},
				success: (res) => {
					if (res.statusCode === 200) {
						this.MonitorLogs = res.data.data.MonitorLogs;
					} else {
						console.error('Error fetching data:', res.data);
					}
				},
				fail: (error) => {
					console.error('Error fetching data:', error);
				}
			});
			this.initMqttConnection(); // 监听 MQTT 实时数据
		},
		methods: {
			initMqttConnection() {
				// 使用 MQTT 协议连接到服务器
				console.log('Connecting to MQTT broker...');

				// 配置 MQTT 连接参数
				const options = {
					username: 'qiyao', // 设置用户名
					password: 'qiyao123456', // 设置密码
					qos: 2 // 设置 QoS
				};

				this.client = mqtt.connect('ws://1.94.221.79:8083/mqtt', options); // 使用配置参数连接

				// 订阅主题
				this.client.subscribe('Go_IoT_Logs');

				// 监听消息
				this.client.on('message', (topic, message) => {
					const mqttData = JSON.parse(message.toString());
					// 将 MQTT 数据添加到 MonitorLogs 数组中
					const newLog = {
						Datetime: mqttData.Datetime,
						Log: mqttData.Log,
						Level1Class: mqttData.Level1Class // 添加 Level1Class 字段
					};
					this.MonitorLogs.push(newLog);

					// 添加$nextTick确保DOM更新
					this.$nextTick(() => {
						// 滚动到底部
						this.scrollToBottom();
					});
				});
			},
			scrollToBottom() {
				// 获取消息容器
				const container = this.$el.querySelector('.message');
				// 滚动到底部
				container.scrollTop = container.scrollHeight;
			},
			rotateImage() {
				this.rotation += 90;
				const echartsBox = this.$refs.echarts_box.$el;
				if (this.rotation === 180) {
					this.rotation = 0;
					echartsBox.style.display = 'none';
				} else {
					echartsBox.style.display = 'block';
				}
				const img = this.$refs.image;
				img.style.transform = `rotate(${this.rotation}deg)`;
				
				// 添加$nextTick确保DOM更新
				this.$nextTick(() => {
					// 滚动到底部
					this.scrollToBottom();
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
		height: 310rpx;
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

	.AI_Prevention p {
		position: absolute;
		font-weight: bold;
		top: -0.5vw;
		left: 9.52vw;
		color: #ffffff;
	}

	.AI_Prevention hr {
		position: absolute;
		top: 2vw;
		height: 0.2px;
		width: 23.4vw;
		background: #ffffff;
	}

	.message {
		/* margin-top: 59px; */
		margin-left: 30rpx;
		margin-right: 20rpx;
		max-height: 300rpx;
		overflow: auto;
		color: #ffffff;
		font-weight: bold;
		width: 700rpx;
		font-size: 35rpx;

		/* 针对Webkit浏览器隐藏滚动条 */
		scrollbar-width: thin;
		scrollbar-color: transparent transparent;
	}

	/* 针对Webkit浏览器隐藏滚动条 */
	.message::-webkit-scrollbar {
		width: 0;
		height: 0;
	}

	/* 针对Webkit浏览器隐藏滚动条 */
	.message::-webkit-scrollbar-thumb {
		background-color: transparent;
	}

	.br-ui {
		height: 10rpx;
	}
</style>