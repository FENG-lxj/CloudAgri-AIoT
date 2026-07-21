<template>
	<view>
		<view class="top_bg">
			<u-toast ref="uToast" />
			<img src="../../img/login/logo.jpg" alt="" class="logo" />
			<view class="title">螭耀科技&nbsp;|&nbsp;云端田园</view>
			<img src="../../img/login/left_img.png" alt="" class="top_img" />
		</view>

		<view class="right_box">
			<view class="login_text">登录</view>
			<view class="form">
				<view class="input">
					<input class="input__inner" type="text" placeholder="手机号" v-model="Phone"
						@input="limitLength($event, 11)" />
				</view>
				<view class="input">
					<input class="input__inner" type="number" placeholder="验证码" v-model="Captcha"
						@input="limitLength($event, 6)" />
				</view>
				<span class="Vertical">|</span>
				<button class="Send_Captcha" @click="sendCaptcha">
					<p v-if="countdown == 0">发送验证码</p>
					<p v-else>{{ countdown }}秒后重试</p>
				</button>
				<button class="btn" @click="submit">登录</button>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				Phone: '17548750710',
				Captcha: null,
				countdown: 0,
				timer: null
			};
		},
		methods: {
			limitLength(event, maxLength) {
				if (event.target.value.length > maxLength) {
					event.target.value = event.target.value.slice(0, maxLength);
				}
			},
			startCountdown() {
				this.timer = setInterval(() => {
					if (this.countdown > 0) {
						this.countdown--;
					}
				}, 1000);
			},
			sendCaptcha() {
				this.countdown = 60;
				this.startCountdown();
				uni.request({
					url: 'http://127.0.0.1:20201/IoTA/users/captcha',
					method: 'POST',
					data: {
						phone: this.Phone
					},
					success: (res) => {
						this.Captcha = res.data.data.code;
						console.log('success:', res.data.code);
						console.log('Captcha:', res.data.data.code);
					}
				});
			},
			success_submit() {
				this.$refs.uToast.show({
					title: '登录成功',
					type: 'success'
				});
			},
			error_submit() {
				this.$refs.uToast.show({
					title: 'asd',
					type: 'error'
				});
			},

			submit() {
				this.startCountdown();
				uni.request({
					url: 'http://127.0.0.1:20201/IoTA/users/login',
					method: 'POST',
					data: {
						phone: this.Phone,
						phoneCaptcha: this.Captcha
					},
					success: (res) => {
						if (res.data.code == 200) {
							this.success_submit();
							// 登录成功处理
							this.countdown = 0;
							console.log(res.data.data.token);
							//数据缓存
							uni.setStorage({
								key: 'phone',
								data: this.Phone,
								success: function() {
									console.log('success');
								}
							});
							//数据缓存
							uni.setStorage({
								key: 'token',
								data: res.data.data.token,
								success: function() {
									console.log('success');
								}
							});
							this.go();
						} else {
							// 登录失败处理
							this.error_submit();
						}
					}
				});
			},
			go() {
				uni.reLaunch({
					url: '/pages/index/index', // 要跳转到的页面路径
					success: (res) => {
						console.log('跳转成功');
					},
					fail: (err) => {
						console.log('跳转失败', err);
					}
				});
			}
		},
		beforeDestroy() {
			clearInterval(this.timer); // 组件销毁前清除定时器
		}
	};
</script>

<style scoped lang="scss">
	/* 注意要写在第一行，同时给style标签加入lang="scss"属性 */
	@import 'uview-ui/index.scss';

	@font-face {
		font-family: DingTalk_JinBuTi;
		src: url('../../img/font/DingTalk_JinBuTi.ttf');
	}

	.top_bg {
		position: fixed;
		width: 100%;
		/* width: 750rpx; */
		height: 1050rpx;
		background: #495ad4;
	}

	.logo {
		position: fixed;
		height: 90rpx;
		width: 90rpx;
		border-radius: 100rpx;
		margin-top: 40rpx;
		margin-left: 40rpx;
	}

	.title {
		font-size: 40rpx;
		margin-top: 65rpx;
		margin-left: 145rpx;
		color: aliceblue;
		font-family: DingTalk_JinBuTi;
	}

	.top_img {
		height: 700rpx;
		/* width: 650rpx; */
		margin-left: 30rpx;
	}

	/* .login {
	position: fixed;
	height: 500rpx;
	width: 600rpx;
	background-color: #ffffff;
	top: 900rpx;
	margin-left: 75rpx;
	border-radius: 45rpx;
	box-shadow: 7rpx 7rpx 10rpx #97a19f;
} */
	/* .login_title {
	font-size: 60rpx;
	margin-top: 50rpx;
	margin-left: 75rpx;
	color: #495ad4;
	font-family: DingTalk_JinBuTi;
}
.input_Phone {
	height: 75rpx;
	width: 450rpx;
	margin-left: 75rpx;
	margin-top: 20rpx;
	border-radius: 100rpx;
	padding-left: 20rpx;
} */

	.right_box {
		position: absolute;
		height: 550rpx;
		width: 600rpx;
		margin-left: 75rpx;
		bottom: 250rpx;
		border-radius: 45rpx;
		box-shadow: 0px 6px 20px #202547;
		background: #fafafa;
	}

	.right_box p {
		position: absolute;
		color: #495ad4;
		left: 30rpx;
		font-size: 30rpx;
	}

	.login_text {
		font-family: DingTalk_JinBuTi;
		font-size: 60rpx;
		color: #495ad4;
		text-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2);
		/* 添加阴影效果 */
		margin: 70rpx 0 50rpx 60rpx;
	}

	.form {
		/* border: 5px solid #55ff00; */
		width: 150rpx;
		height: 400rpx;
		margin: 25px auto 0px auto;
	}

	.input {
		display: flex;
		height: 75rpx;
		width: 430rpx;
		border-radius: 20rpx;
		padding: 0 15px;
		background-color: #f9f9f9;
		box-shadow: inset 1px 1px 7px rgba(95, 95, 95, 0.3);
		transition: 300ms ease-in-out;
		margin: 30rpx 0 0 -170rpx;
	}

	.input:focus {
		background-color: white;
		transform: scale(1.05);
		box-shadow: 13px 13px 100px #c2c2c2, -13px -13px 100px #ffffff;
	}

	.input__inner {
		flex: 1;
		border: none;
		outline: none;
		width: 100%;
		background-color: transparent;
		font-size: 50rpx;
		line-height: 30px;
		color: #333;
		-moz-appearance: textfield;
		/* Firefox */
		appearance: textfield;
		/* Chrome */
	}

	/* 隐藏输入框右边的上下调整箭头 */
	.input__inner::-webkit-inner-spin-button,
	.input__inner::-webkit-outer-spin-button {
		-webkit-appearance: none;
		margin: 0;
	}

	.input__inner::-webkit-outer-spin-button,
	.input__inner::-webkit-inner-spin-button {
		display: none;
	}

	.btn {
		margin: 50rpx auto 0rpx auto;
		left: -125rpx;
		height: 90rpx;
		width: 400rpx;
		color: #fff;
		background-color: #495ad4;
		border-radius: 20px;
		text-align: center;
		/* 水平居中 */
		display: flex;
		justify-content: center;
		/* 水平居中 */
		align-items: center;
		/* 垂直居中 */
		padding: 10px 20px;
		/* 按钮的内边距 */
		font-size: 50rpx;
		box-shadow: 2px 3px 3px #65656582;
		border: none;
		/* 去掉边框 */
		font-family: DingTalk_JinBuTi;
	}

	.btn:active {
		background-color: #7284d4;
		border-color: #495ad4;
	}

	.Vertical {
		position: absolute;
		left: 325rpx;
		top: 275rpx;
		font-size: 50rpx;
		font-weight: 100;
		color: #c7bdbd;
		text-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2);
		/* 添加阴影效果 */
	}

	.Send_Captcha {
		position: absolute;
		left: 320rpx;
		top: 280rpx;
		height: 50rpx;
		width: 300rpx;
		font: inherit;
		color: inherit;
		background-color: transparent;
		cursor: pointer;
		text-shadow: 0px 2px 4px rgba(0, 0, 0, 0.2);
		z-index: 999;
	}

	button::after {
		border: none;
	}

	.Send_Captcha p {
		color: #495ad4;
		font-size: 35rpx;
	}

	.Send_Captcha p:active {
		color: #4959d4bd;
	}
</style>