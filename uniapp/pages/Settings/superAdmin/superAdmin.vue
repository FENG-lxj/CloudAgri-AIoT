<template>
	<view class="box">
		<u-toast ref="uToast" />
		<view class="user">
			<img src="@/img/Setting/userImg.png" alt="" class="UserImg" />
			<view class="Phone">{{this.phone}}</view>
			<view class="text">超级管理员</view>
			<img src="../../../img/index/next.png" alt="" class="usernext" @tap="goBack()" />
		</view>
		
		<view class="kon"></view>

		<!-- 二级管理员列表 -->
		<view class="Super">
			<view class="title">二级管理员</view>
			<img ref="image" src="../../../img/Setting/add_img.png" alt="" class="unfold_img"
				@click="showAddPanel = true" />
			<hr class="line" />
			<view class="admin">
				<view class="content">
					<!-- 使用 v-for 遍历渲染管理员列表 -->
					<view v-for="(user, index) in list" :key="index" class="content_admin">
						<view class="ID">{{ index + 1 }}</view>
						<view class="admin_Phone">{{ user.phone }}</view>
						<view class="IP">{{ user.last_login_ip }}</view>
						<img src="../../../img/Setting/trashcan_img.png" alt="" class="Delete_img"
							@click="showDeleteConfirm(user.phone)" />
						<hr class="admin_line" />
					</view>
				</view>
			</view>
		</view>

		<view class="API_KEY">
			<view class="API_title">AI API Key</view>
			<view class="API_bt">
				<view class="API_bt_text" @click="Replace_Key_btn()">更换</view>
			</view>
			<hr class="API_KEY_line" />
			<input type="text" placeholder=" New API Key" class="GPT_Input" />
		</view>


		<!-- 内存环形进度条 -->
		<view class="Memory">
			<p class="baifenbi">{{this.MemPer}}%</p>
			<p class="Memory_title">服务器内存使用情况</p>
			<view class="Memory_title_progress">
				<canvas canvas-id="memoryCanvas"
					:style="{ width: circleSize + 'px', height: circleSize + 'px' }"></canvas>

				<view class="Memory_message">
					<view class="Memory_message_one_two">{{ memory_percentage/100*2 }}GB/2GB</view>
				</view>
			</view>
		</view>
		<!-- CPU环形进度条 -->
		<!-- CPU环形进度条 -->
		<view class="Memory_CPU">
			<p class="baifenbi">{{ this.CPUper }}%</p>
			<p class="Memory_title">服务器CPU使用情况</p>
			<view class="Memory_title_progress">
				<canvas canvas-id="cpuCanvas" :style="{ width: circleSize + 'px', height: circleSize + 'px' }"></canvas>

				<view class="Memory_message_cpu">
					2核心
				</view>
			</view>
		</view>

		<!-- 添加管理员的输入面板 -->
		<view v-if="showAddPanel" class="add-panel">
			<view class="Addpanelmodel">
				<input v-model="form.Addphone" placeholder="请输入手机号">
				<view class="option_true" @click="AddManagement_btn('二级管理员')">添加</view>
				<view class="option" @click="showAddPanel = false">放弃</view>
			</view>
		</view>
		<!-- 删除确认弹窗 -->
		<view v-if="showDeleteConfirmModal" class="delete-confirm-modal">
			<view class="modal-content">
				<view class="modal-text">确定要删除该二级管理员吗？</view>
				<view class="modal-buttons">
					<view class="modal-button" @click="confirmDelete">确定</view>
					<view class="modal-button" @click="cancelDelete">取消</view>
				</view>
			</view>
		</view>
	</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				list: [],
				// 添加管理员
				dialogFormVisible: false,
				form: {
					Addphone: "",
				},
				formLabelWidth: "120px",
				//删除管理员
				dialogVisible: false,
				//更换GPT-Key
				GPT_key: "",
				// 初始百分比
				memory_percentage: 0,
				CPU_percentage: 0,
				MemPer: 0,
				CPUper: 0,
				numberFontSize: 30,
				circleSize: 120,
				phone: 0,
				// 控制输入面板显示与隐藏的变量
				showAddPanel: false,
				// 表单数据
				form: {
					Addphone: ""
				},
				// 其他数据省略...
				showDeleteConfirmModal: false, // 控制删除确认弹窗的显示与隐藏
				deleteUserId: null, // 要删除的二级管理员的ID
			};
		},
		mounted() {
			this.phone = uni.getStorageSync('phone');
			this.getServerData();
			this.getList();
		},

		methods: {
			success_submit() {
				this.$refs.uToast.show({
					title: '成功',
					type: 'success'
				});
			},
			error_submit() {
				this.$refs.uToast.show({
					title: '失败',
					type: 'error'
				});
			},
			goBack() {
				uni.navigateBack({
					//关闭当前页面，返回上一页面或多级页面。
					delta: 1
				});
			},
			Replace_Key_btn() {
				uni.request({
					url: 'https://qpi.qiyao.ink/IoTA/SuperAdmin/UpChatGPTKey',
					method: 'POST',
					header: {
						Authorization: `Bearer ${uni.getStorageSync('token')}`
					},
					data: {
						APIKEY: this.GPT_key,
					},
					success: (res) => {
						console.log('success:', res);
						this.Key_success();
						if (res.data.code == 200) {
							this.success_submit();
						} else {
							this.error_submit();
						}
					},
					fail: (error) => {
						console.error('Error fetching data:', error);
						this.Key_error();
					}
				});
			},

			// 显示删除确认弹窗，并传递要删除的管理员的手机号
			showDeleteConfirm(phone) {
				this.deleteUserId = phone;
				this.showDeleteConfirmModal = true;
			},

			// 取消删除
			cancelDelete() {
				this.showDeleteConfirmModal = false;
				this.deleteUserId = null;
			},

			// 确认删除
			confirmDelete() {
				if (this.deleteUserId) {
					// 调用删除方法
					this.deleteUser(this.deleteUserId);
				}
				this.cancelDelete(); // 关闭确认弹窗
			},

			// 发送请求删除二级管理员
			deleteUser(phone) {
				uni.request({
					url: "http://127.0.0.1:20201/IoTA/SuperAdmin/RemoveUser",
					method: 'POST',
					header: {
						Authorization: `Bearer ${uni.getStorageSync('token')}`
					},
					data: {
						RemovePhone: phone,
					},
					success: (res) => {
						// 请求成功的处理
						console.log("删除成功:", res);
						this.getList(); // 刷新管理员列表
						if (res.data.code == 200) {
							this.success_submit();
						} else {
							this.error_submit();
						}
					},
					fail: (error) => {
						// 请求失败的处理
						console.error("删除失败:", error);
					}
				});
			},
			// 添加管理员的方法
			AddManagement_btn(role) {
				// 验证手机号是否为空
				if (!this.form.Addphone) {
					uni.showToast({
						title: '请输入手机号',
						icon: 'none'
					});
					return;
				}

				// 验证手机号格式
				const reg = /^1[0-9]{10}$/;
				if (!reg.test(this.form.Addphone)) {
					uni.showToast({
						title: '请输入正确的手机号',
						icon: 'none'
					});
					return;
				}

				// 添加管理员逻辑...
				console.log("手机号码：", this.form.Addphone);
				console.log("选择的角色：", role);

				// 发起添加管理员请求
				uni.request({
					url: 'http://127.0.0.1:20201/IoTA/SuperAdmin/SuperAdminAddUser',
					method: 'POST',
					header: {
						Authorization: `Bearer ${uni.getStorageSync('token')}`
					},
					data: {
						phone: this.form.Addphone
					},
					success: (res) => {
						console.log('success:', res);
						this.getList(); // 刷新页面
						if (res.data.code == 200) {
							this.success_submit();
						} else {
							this.error_submit();
						}
					},
					fail: (error) => {
						console.error('Error fetching data:', error);
					}
				});

				// 关闭输入面板
				this.showAddPanel = false;
			},

			//获取管理员列表
			getList() {
				uni.request({
					url: 'http://127.0.0.1:20201/IoTA/SuperAdmin/GetUserList',
					method: 'GET',
					header: {
						Authorization: `Bearer ${uni.getStorageSync('token')}`
					},
					success: (res) => {
						this.list = res.data.data.Users;
						console.log("success:", res.data.data.Users);
					},
					fail: (error) => {
						console.error("Error fetching data:", error);
					}
				});
			},

			drawMemoryProgress() {
				const ctx = uni.createCanvasContext('memoryCanvas', this);
				const radius = this.circleSize / 2 - 4;
				const centerX = this.circleSize / 2;
				const centerY = this.circleSize / 2;

				// 绘制背景条
				ctx.beginPath();
				ctx.arc(centerX, centerY, radius, 0, 2 * Math.PI);
				ctx.strokeStyle = '#e0e0e0';
				ctx.lineWidth = 3;
				ctx.stroke();

				// 绘制前景条
				const startAngle = -0.5 * Math.PI;
				const endAngle = 2 * Math.PI * this.MemPer / 100 - 0.5 * Math.PI;
				ctx.beginPath();
				ctx.arc(centerX, centerY, radius, startAngle, endAngle);
				ctx.strokeStyle = '#00ffff';
				ctx.lineWidth = 3;
				ctx.stroke();

				ctx.draw();


				// 绘制CPU环形进度条
				const cpuCtx = uni.createCanvasContext('cpuCanvas', this);
				const cpuradius = this.circleSize / 2 - 4;
				const CPUcenterX = this.circleSize / 2;
				const CPUcenterY = this.circleSize / 2;

				// 绘制背景条
				cpuCtx.beginPath();
				cpuCtx.arc(CPUcenterX, CPUcenterY, radius, 0, 2 * Math.PI);
				cpuCtx.strokeStyle = '#e0e0e0';
				cpuCtx.lineWidth = 3;
				cpuCtx.stroke();

				// 绘制前景条
				const cpustartAngle = -0.5 * Math.PI;
				const cpuendAngle = 2 * Math.PI * this.CPUper / 100 - 0.5 * Math.PI;
				cpuCtx.beginPath();
				cpuCtx.arc(CPUcenterX, CPUcenterY, cpuradius, cpustartAngle, cpuendAngle);
				cpuCtx.strokeStyle = '#00ffff';
				cpuCtx.lineWidth = 3;
				cpuCtx.stroke();

				cpuCtx.draw();
			},

			// 获取服务器数据
			getServerData() {
				uni.request({
					url: 'http://127.0.0.1:20201/IoTA/bigdata/GetMemCPU',
					method: 'GET',
					header: {
						Authorization: `Bearer ${uni.getStorageSync('token')}`
					},
					success: (res) => {
						if (res.statusCode === 200) {
							// 请求成功
							this.MemPer = parseFloat(res.data.data.MemPer).toFixed(2);
							this.CPUper = parseFloat(res.data.data.CPUPer).toFixed(2);
							// 更新内存和CPU的百分比，使用 Math.round 进行四舍五入
							this.memory_percentage = res.data.data.MemPer;
							this.CPU_percentage = res.data.data.CPUper;
							console.log('服务器success:', this.CPU_percentage);
							this.drawMemoryProgress();
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
		position: fixed;
		height: 100%;
		background-color: #ffffff;
	}

	/* 删除确认弹窗样式 */
	.delete-confirm-modal {
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background-color: rgba(0, 0, 0, 0.5);
		/* 半透明黑色背景 */
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: 9999;
		/* 确保弹窗在最顶层 */
	}

	.modal-content {
		background-color: #fff;
		padding: 20px;
		border-radius: 10px;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
	}

	.modal-text {
		font-size: 16px;
		margin-bottom: 20px;
		text-align: center;
	}

	.modal-buttons {
		display: flex;
		justify-content: center;
	}

	.modal-button {
		padding: 10px 20px;
		margin: 0 10px;
		border: 1px solid #ccc;
		border-radius: 5px;
		cursor: pointer;
	}

	.modal-button:hover {
		background-color: #f0f0f0;
	}

	.add-panel {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background-color: rgba(0, 0, 0, 0.5);
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.Addpanelmodel {
		background-color: #edeeee;
		width: 80%;
		border-radius: 20rpx;
	}

	.add-panel input {
		width: 80%;
		padding: 30rpx;
		margin-top: 20rpx;
		margin-left: 30rpx;
		height: 0rpx;
		background-color: #ffffff;
	}

	.option {
		padding: 10px 20px;
		width: 20%;
		margin-top: -105rpx;
		margin-left: 320rpx;
		background-color: #fff;
		border-radius: 5px;
		font-size: 30rpx;
		margin-bottom: 10px;
		cursor: pointer;
		text-align: center;
	}

	.option_true {
		padding: 10px 20px;
		width: 20%;
		font-size: 30rpx;
		margin-top: 10rpx;
		margin-left: 50rpx;
		background-color: #fff;
		border-radius: 5px;
		margin-bottom: 10px;
		cursor: pointer;
		text-align: center;
	}

	.option:hover {
		background-color: #f0f0f0;
	}

	.option_true:hover {
		background-color: #f0f0f0;
	}

	.user {
		position: relative;
		height: 150rpx;
		width: 700rpx;
		margin: 25rpx;
		background-color: #f2f2f2;
		border-radius: 15rpx;
		box-shadow: 0px 0px 7px #202547;
		border: 1rpx solid #8c8d8d00;
	}

	.UserImg {
		position: absolute;
		height: 110rpx;
		margin-top: 20rpx;
		margin-left: 25rpx;
	}

	.Phone {
		position: absolute;
		margin-left: 150rpx;
		margin-top: 40rpx;
		font-size: 42rpx;
		font-weight: bold;
	}

	.text {
		position: absolute;
		margin-left: 153rpx;
		margin-top: 85rpx;
		font-size: 30rpx;
		font-weight: bold;
		font-family: DingTalk_JinBuTi;
	}

	.usernext {
		position: absolute;
		height: 100rpx;
		margin-left: 600rpx;
		margin-top: 30rpx;
		transform: rotate(90deg);
		cursor: pointer;
	}

	@font-face {
		font-family: DingTalk_JinBuTi;
		src: url('../../../img/font/DingTalk_JinBuTi.ttf');
	}
	
	.kon{
		height: 0.5rpx;
	}

	.Super {
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

	.admin {
		margin-top: -90rpx;
		height: 590rpx;
		width: 700rpx;
		margin-bottom: -90rpx;
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

	.content {
		margin-top: 100rpx;
		height: 480rpx;
		width: 700rpx;
		/* background-color: aqua; */
	}

	.content_admin {
		height: 80rpx;
		width: 700rpx;
		/* background-color: pink; */
	}

	.ID {
		position: absolute;
		font-size: 32rpx;
		font-family: DingTalk_JinBuTi;
		margin-top: 10rpx;
		margin-left: 20rpx;
	}

	.admin_Phone {
		position: absolute;
		font-size: 32rpx;
		font-family: DingTalk_JinBuTi;
		margin-top: 10rpx;
		margin-left: 80rpx;
	}

	.IP {
		position: absolute;
		font-size: 32rpx;
		font-family: DingTalk_JinBuTi;
		margin-top: 10rpx;
		margin-left: 330rpx;
	}

	.Delete_img {
		position: absolute;
		height: 50rpx;
		margin-left: 640rpx;
		margin-top: 5rpx;
	}

	.admin_line {
		position: absolute;
		margin-top: 65rpx;
		width: 700rpx;
		border: 1rpx solid #bdbebe;
	}

	.API_KEY {
		position: absolute;
		height: 170rpx;
		width: 700rpx;
		margin: 25rpx;
		border-radius: 30rpx;
		box-shadow: 0px 0px 7px #202547;
		border: 1rpx solid #8c8d8d00;
	}

	.API_title {
		position: absolute;
		font-size: 35rpx;
		font-family: DingTalk_JinBuTi;
		margin-top: 10rpx;
		margin-left: 20rpx;
	}

	.API_bt {
		position: absolute;
		margin-top: 8rpx;
		margin-left: 578rpx;
		height: 45rpx;
		width: 90rpx;
		box-shadow: 0px 0px 4px #202547;
		border: 1rpx solid #8c8d8d00;
		border-radius: 15rpx;
	}

	.API_bt_text {
		margin-top: 5rpx;
		margin-left: 15rpx;
		font-size: 30rpx;
		font-family: DingTalk_JinBuTi;
	}

	.API_KEY_line {
		position: absolute;
		margin-top: 65rpx;
		width: 700rpx;
		border: 1rpx solid #bdbebe;
	}

	.GPT_Input {
		position: absolute;
		top: 90rpx;
		left: 20rpx;
		height: 55rpx;
		width: 650rpx;
		border-radius: 15rpx;
		border: 2px solid #9e9e9e;
	}

	.Memory {
		position: absolute;
		margin-top: 245rpx;
		height: 430rpx;
		width: 320rpx;
		margin-left: 28rpx;
		border-radius: 27rpx;
		box-shadow: 0px 0px 7px #202547;
		border: 1rpx solid #8c8d8d00;
		background-color: #edeeee;
	}

	.Memory_title {
		position: absolute;
		font-size: 30rpx;
		font-family: DingTalk_JinBuTi;
		margin-top: 30rpx;
		margin-left: 28rpx;
	}

	.Memory_title_progress canvas {
		position: absolute;
		margin-top: 100rpx;
		margin-left: 40rpx;
		width: 100%;
		height: 100%;
	}

	.Memory_message {
		position: absolute;
		font-size: 45rpx;
		font-family: DingTalk_JinBuTi;
		margin-top: 360rpx;
		margin-left: 35rpx;

	}

	.baifenbi {
		position: absolute;
		font-size: 40rpx;
		font-family: DingTalk_JinBuTi;
		margin-top: 195rpx;
		margin-left: 90rpx;
	}

	.Memory_CPU {
		position: absolute;
		margin-top: 245rpx;
		height: 430rpx;
		width: 320rpx;
		margin-left: 400rpx;
		border-radius: 27rpx;
		box-shadow: 0px 0px 7px #202547;
		border: 1rpx solid #8c8d8d00;
		background-color: #edeeee;
	}

	.Memory_message_cpu {
		position: absolute;
		font-size: 45rpx;
		font-family: DingTalk_JinBuTi;
		margin-top: 360rpx;
		margin-left: 105rpx;
	}
</style>