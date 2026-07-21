<template>
	<view>
		<view class="box">
			<view class="title">
				发现病虫害&nbsp;&nbsp;
				{{ DetectionInfo[index].Datetime }}
			</view>
			<img ref="image" src="../../../img/index/unfold.png" alt="" class="unfold_img" @click="rotateImage" />
			<hr class="line" ref="Line" v-if="rotation !== 0" />
			<view class="echarts" ref="echarts_box">
				<p class="Detect_Diseases">{{ DetectionInfo[index].Pest }}</p>
				<view class="AI_Camera_One">
					<img src="@/img/index/Last.png" alt="" class="AI_Camera_One_left_Img" @click="Decrease" />
					<img :src="Images[index]" alt="" class="AI_Camera_One_img" />
					<p class="ID">
						识别ID：{{ DetectionInfo[index].IdentifyID }}
					</p>
					<img src="@/img/index/next.png" alt="" class="AI_Camera_One_right_Img" @click="Add" />
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				rotation: 0,
				DetectionInfo: [],
				index: 0,
				Images: [], // 存储图片的数组
			};
		},
		created() {
			this.index = 0;
			this.$store.commit('set_AI_Camera_One_ICID', 1);
			this.$watch(
				() => this.$store.state.AI_Camera_One_ICID,
				(newValue, oldValue) => {
					console.log('AI_Camera_One_ICID changed:', newValue, oldValue);
					this.forceRefresh();
				}
			);
		},
		mounted() {
			this.getdata();
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
			Decrease() {
				if (this.index > 0) {
					this.index--;
					console.log(this.index, '++');
				}
			},
			Add() {
				if (this.index < this.DetectionInfo.length - 1) {
					this.index++;
					console.log(this.index, '--');
				}
			},
			async getdata() {
				try {
					const res = await uni.request({
						url: 'http://127.0.0.1:20201/IoTA/ai/DetectionInfo',
						method: 'GET',
						data: {
							ICID: this.$store.state.AI_Camera_One_ICID,
							Level1Class: 1,
							StartNum: 0,
							Num: 10
						},
						header: {
							'Authorization': `Bearer ${uni.getStorageSync('token')}`,
						},
					});
					if (res[1].data.code === 200) {
						// 按时间降序排序
						const sortedData = res[1].data.data.DetectionInfo.sort(
							(a, b) => new Date(b.Datetime) - new Date(a.Datetime)
						);

						this.DetectionInfo = sortedData;
						this.Images = new Array(sortedData.length).fill(null); // 初始化图片数组

						// 并行获取图片并保持顺序
						const imageRequests = sortedData.map((item, index) =>
							this.GET_Img(item.id, index)
						);
						await Promise.all(imageRequests);
					} else {
						console.error('请求病虫害检测信息组出错:', res);
					}
				} catch (error) {
					console.error('请求病虫害检测信息组出错-2:', error);
				}
			},
			async GET_Img(id, index) {
				try {
					const res = await uni.request({
						url: 'http://127.0.0.1:20201/IoTA/ai/DetectionImage',
						method: 'GET',
						data: {
							Id: id,
							Level1Class: 1
						},
						header: {
							Authorization: `Bearer ${uni.getStorageSync('token')}`
						},
						responseType: 'arraybuffer',
					});

					if (res[1].statusCode === 200) {
						const imageUrl = 'data:image/jpeg;base64,' + uni.arrayBufferToBase64(res[1].data);
						this.$set(this.Images, index, imageUrl); // 确保响应式更新
					} else {
						console.error('病虫害请求图片出错:', res);
						this.$set(this.Images, index, 'placeholder.png'); // 可设置占位图
					}
				} catch (error) {
					console.error('病虫害请求图片出错-2:', error);
					this.$set(this.Images, index, 'placeholder.png'); // 可设置占位图
				}
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
		height: 350rpx;
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

	//////////////////1
	.AI_Camera_One {
		position: relative;
		height: 17.5vw;
		width: 47vw;
		border-radius: 0.8vw;
		/* 使用flex布局使内部内容居中对齐 */
		display: flex;
		/* justify-content: center;
	align-items: center */
		;
	}

	.text-container {
		/* display: flex; */
		flex-direction: column;
		/* 竖直排列 */
		align-items: center;
		/* 居中对齐 */
		margin-top: 193px;
	}

	.AI_Camera_One_left_Img {
		position: absolute;
		height: 150rpx;
		width: 150rpx;
		top: 95rpx;
		left: -30rpx;
		cursor: pointer;
	}

	.AI_Camera_One_right_Img {
		position: absolute;
		left: 580rpx;
		height: 150rpx;
		width: 150rpx;
		top: 95rpx;
		cursor: pointer;
	}

	.AI_Camera_One_img {
		position: absolute;
		height: 230rpx;
		width: 475rpx;
		top: 55rpx;
		left: 115rpx;
	}

	.Detect_Diseases {
		position: absolute;
		font-size: 36rpx;
		margin: 5rpx 0 0 0;
		font-family: DingTalk_JinBuTi;
		left: 50%;						/* 向右移动 50% */
		transform: translateX(-50%);	/* 向左移动自身宽度的 50% */
	}

	.Time_Discovery {
		margin-top: -10px;
		font-size: 18px;
	}

	.ID {
		position: absolute;
		width: 600rpx;
		margin: 298rpx 0 0 0;
		font-size: 32rpx;
		font-family: DingTalk_JinBuTi;
		left: 50%;						/* 向右移动 50% */
		transform: translateX(-16%);	/* 向左移动自身宽度的 50% */
	}
</style>