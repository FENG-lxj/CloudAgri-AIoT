<template>
	<div>
		<!-- 自定义导航栏 -->
		<view class="custom-nav-bar">
			<!-- 返回箭头按钮 -->
			<button class="back-button" @click="goBackHome">
				<uni-icons type="arrowleft" size="20" color="#fff"></uni-icons>
			</button>
			<!-- 页面标题 -->
			<view class="nav-bar-title">AI农业专家</view>
		</view>

		<div class="chat-container">
			<!-- 聊天消息区域 -->
			<scroll-view class="chat-messages" scroll-y="true" :scroll-top="scrollTop">
				<div v-for="(message, index) in conversation" :key="index" class="message-container"
					:class="{'user-message-container': message.role === 'user', 'bot-message-container': message.role === 'gpt'}">
					<img :src="message.role === 'user' ? userAvatar : gptAvatar" class="avatar"
						:class="{'user-avatar': message.role === 'user', 'bot-avatar': message.role === 'gpt'}" />
					<div class="message-bubble"
						:class="{'user-bubble': message.role === 'user', 'bot-bubble': message.role === 'gpt'}">
						{{ message.text }}
					</div>
				</div>
			</scroll-view>

			<!-- 输入区域 -->
			<div class="input-container">
				<input class="message-input" v-model="inputText" placeholder="输入咨询内容..." @keyup.enter="sendMessage"
					@keyup.esc="closeModal" />
				<button class="send-button" @click="sendMessage">
					<uni-icons type="paperplane" size="25" color="#ffffff" class="icon-ui"></uni-icons>
				</button>
			</div>
		</div>
	</div>
</template>

<script>
	export default {
		data() {
			return {
				inputText: '',
				conversation: [],
				userAvatar: 'https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/userImg.png',
				gptAvatar: 'https://pastoral-cloud.obs.cn-north-4.myhuaweicloud.com/uniapp/ChatGPT.png',
				scrollTop: 0
			};
		},
		methods: {
			// 返回主页的方法
			goBackHome() {
				// 使用 switchTab 跳转到主页（假设主页路径为 pages/index/index）
				uni.switchTab({
					url: '/pages/index/index'
				});
			},
			sendMessage() {
				const inputText = this.inputText.trim();
				if (inputText !== '') {
					// 将用户的消息添加到对话数组中
					this.conversation.push({
						role: 'user',
						text: inputText
					});
					// 清空输入框
					this.inputText = '';
					// 在下一个 DOM 更新周期后滚动到底部
					this.$nextTick(() => {
						this.scrollToBottom();
					});
					// 发送消息给 GPT
					uni.request({
						url: 'http://127.0.0.1:4999/IoTA/ai/AgricultureExpertAI',
						method: 'GET',
						data: {
							Issue: inputText
						},
						header: {
							Authorization: `Bearer ${uni.getStorageSync('token')}`
						},
						timeout: 500000,
						success: (res) => {
							// 将 AI 的回复文本中的 '**' 替换为空字符串，即过滤掉 '**'
							const filteredReply = res.data.data.AIReply.replace(/\*\*/g, '');
							// 将 GPT 的回复添加到对话数组中
							this.conversation.push({
								role: 'gpt',
								text: filteredReply
							});

							// 在下一个 DOM 更新周期后滚动到底部
							this.$nextTick(() => {
								this.scrollToBottom();
							});

							// 清空输入框
							this.inputText = '';
						},
						fail: (error) => {
							console.error('获取数据时出错：', error);
						}
					});
				}
			},
			// 发送消息后滚动到底部
			scrollToBottom() {
				this.$nextTick(() => {
					this.scrollTop = Math.random() * 1000000 // 强制触发滚动
				})
			}
		}
	}
</script>

<style scoped>
	@font-face {
		font-family: DingTalk_JinBuTi;
		src: url('../../img/font/DingTalk_JinBuTi.ttf');
	}

	/* 自定义导航栏样式 */
	.custom-nav-bar {
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		height: 44px;
		/* 导航栏高度 */
		background-color: #262F3E;
		/* 与配置中 navigationBarBackgroundColor 一致 */
		display: flex;
		align-items: center;
		padding: 0 15px;
		z-index: 100;
	}

	.back-button {
		background: none;
		border: none;
		color: white;
		cursor: pointer;
		margin-left: -35rpx;
	}

	.nav-bar-title {
		color: white;
		font-size: 16px;
		font-weight: bold;
		flex: 1;
		text-align: center;
		font-family: DingTalk_JinBuTi;
		margin-left: -40rpx;
	}

	.chat-container {
		position: fixed;
		top: 44px;
		/* 导航栏高度 */
		bottom: 0;
		left: 0;
		right: 0;
		display: flex;
		flex-direction: column;
		background: #f5f5f5;
	}

	.chat-messages {
		flex: 1;
		padding: 15px;
		overflow-y: auto;
		margin-left: -35rpx;
	}

	.message-container {
		display: flex;
		margin-bottom: 20px;
	}

	.user-message-container {
		flex-direction: row-reverse;
	}

	.bot-message-container {
		flex-direction: row;
	}

	.avatar {
		width: 40px;
		height: 40px;
		border-radius: 50%;
		margin: 0 10px;
	}

	.message-bubble {
	  max-width: 70%;
	  padding: 12px 18px;
	  border-radius: 18px;
	  font-size: 16px;
	  line-height: 1.4;
	  word-break: break-word;
	  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
	  white-space: pre-wrap; /* 添加换行属性 */
	}

	.user-bubble {
		background: #1890ff;
		color: white;
		border-bottom-right-radius: 4px;
	}

	.bot-bubble {
		background: white;
		color: #333;
		border-bottom-left-radius: 4px;
	}

	.input-container {
		display: flex;
		align-items: center;
		padding: 12px 15px;
		background: white;
		border-top: 1px solid #eee;
	}

	.message-input {
		flex: 1;
		padding: 10px 15px;
		border: 1px solid #ddd;
		border-radius: 25px;
		margin-right: 10px;
		font-size: 16px;
		transition: all 0.3s;
	}

	.message-input:focus {
		outline: none;
		border-color: #1890ff;
		box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
	}

	.send-button {
		width: 40px;
		height: 40px;
		border: none;
		border-radius: 50%;
		background: #1890ff;
		display: flex;
		align-items: center;
		justify-content: center;
		transition: all 0.3s;
	}

	.send-button:active {
		transform: scale(0.95);
		background: #096dd9;
	}

	/* 保持原有导航栏样式不变 */
	.custom-nav-bar {
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		height: 44px;
		background-color: #262F3E;
		display: flex;
		align-items: center;
		padding: 0 15px;
		z-index: 100;
	}

	.nav-bar-title {
		color: white;
		font-size: 16px;
		font-weight: bold;
		flex: 1;
		text-align: center;
		font-family: DingTalk_JinBuTi;
	}
	.icon-ui {
		margin: 5rpx 5rpx 0 0;
	}
</style>