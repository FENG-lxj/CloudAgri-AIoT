<template>
	<view>
		<view class="box">
			<view class="title">病虫害防治小建议</view>
			<img ref="image" src="../../../img/index/refresh_img.png" alt="" class="unfold_img" @click="rotateImage" />
			<hr class="line" ref="Line" />
			<view class="echarts">
				<view class="Air_Temperature_Echart">
					<view class="AIProposal" v-if="AIProposal != ''">
						{{ AIProposal }}
					</view>
					<view class="AIProposal" v-if="AIProposal == ''">正在获取病虫害防治小建议...</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				AIProposal: '',
			};
		},
		mounted() {
			this.GET_AI();
		},
		methods: {
			rotateImage() {
				this.$refs.image.classList.add('rotate'); // 添加旋转类
				setTimeout(() => {
					this.$refs.image.classList.remove('rotate'); // 一段时间后移除旋转类
				}, 1000); // 调整这个时间以控制旋转时间
				this.GET_AI();
			},
			GET_AI() {
				console.log('GET_AI');
				this.AIProposal = '';
				uni.request({
					url: 'http://127.0.0.1:20201/IoTA/ai/PestProposal',
					method: 'GET',
					data: {
						Pest: '稻瘟病'
					},
					header: {
						'Authorization': `Bearer ${uni.getStorageSync('token')}`,
					},
					success: (res) => {
						if (res.statusCode === 200) {
							// 请求成功
							// 更新 pestProposal 数据
							this.AIProposal = res.data.data.Proposal;
						} else {
							console.error('Error fetching data:', res.data);
						}
					},
					fail: (error) => {
						// 请求失败时的处理
						console.error('Error fetching data:', error);
					}
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
	}

	.line {
		margin-top: 10rpx;
		width: 700rpx;
		border: 1rpx solid #bdbebe;
	}

	.echarts {
		/* display: none; */
	}

	.Air_Temperature_Echart {
		margin-top: 10rpx;
		height: 450rpx;
		width: 650rpx;
		margin-left: 25rpx;
		margin-bottom: -90rpx;
	}

	.AIProposal {
		height: 350rpx;
		width: 650rpx;
		white-space: pre-line;
		overflow-y: auto;
		overflow-x: hidden;
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

	.rotate {
		transform: rotate(-360deg); // 旋转360度
		transition: transform 1s; // 使用过渡效果
	}
</style>