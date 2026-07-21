// Go_IoTAS.js
import mqtt from 'mqtt';
import {
	mapMutations
} from 'vuex';

export default {
	data() {
		return {
			client: null,
			Time: [],
		};
	},
	methods: {
		AE_init() {
			console.log('init');
			uni.request({
				url: 'http://127.0.0.1:20201/IoTA/bigdata/GetASdata',
				data: {
					ASID: this.$store.state.Go_IoTAS_ASID,
					StartTime: '-24h'
				},
				header: {
					'Authorization': `Bearer ${uni.getStorageSync('token')}`,
				},
				success: (res) => {
					console.log("获取农业数据", res.data.data.ASDatas.slice(-8));
					// 使用循环将数据逐个存入 Air_quality_DataArray
					res.data.data.ASDatas.slice(-8).forEach((data) => {
						// this.$store.dispatch('update_Air_quality_DataArray', data);
						console.log('data', data);
						//更新ASID
						this.$store.dispatch('update_Go_IoTAS_ASID_DataArray', data.ASID);
						// console.log(`Received MQTT message from ${topic} topic: ASID`, mqttData.ASID);
						//更新空气温度数据
						this.$store.dispatch('update_Go_IoTAS_T_DataArray', data.T);
						// console.log(`Received MQTT message from ${topic} topic: T`, mqttData.T);
						// 更新空气湿度数据
						this.$store.dispatch('update_Go_IoTAS_H_DataArray', data.H);
						// console.log(`Received MQTT message from ${topic} topic: H`, mqttData.H);
						//更新二氧化碳浓度数据
						this.$store.dispatch('update_Go_IoTAS_CO2_DataArray', data.CO2);
						//更新氧气浓度数据
						this.$store.dispatch('update_Go_IoTAS_O2_DataArray', data.O2);
					});
				},
				fail: (error) => {
					console.error("Error fetching data:", error);
				}
			});
		},

		initMqttConnection() {
			// 使用 MQTT 协议连接到服务器
			console.log('Connecting to MQTT broker...');
			this.client = mqtt.connect('ws://1.94.221.79:8083/mqtt');

			this.client.subscribe('Go_IoTAS');

			this.client.on('message', (topic, message) => {
				const mqttData = JSON.parse(message.toString());

				//更新ASID
				this.$store.dispatch('update_Go_IoTAS_ASID_DataArray', mqttData.ASID);
				// console.log(`Received MQTT message from ${topic} topic: ASID`, mqttData.ASID);
				//更新空气温度数据
				this.$store.dispatch('update_Go_IoTAS_T_DataArray', mqttData.T);
				// console.log(`Received MQTT message from ${topic} topic: T`, mqttData.T);
				// 更新空气湿度数据
				this.$store.dispatch('update_Go_IoTAS_H_DataArray', mqttData.H);
				// console.log(`Received MQTT message from ${topic} topic: H`, mqttData.H);
				//更新二氧化碳浓度数据
				this.$store.dispatch('update_Go_IoTAS_CO2_DataArray', mqttData.CO2);
				console.log(`Received MQTT message from ${topic} topic: CO2`, this.$store.state.Go_IoTAS_CO2);
				//更新氧气浓度数据
				this.$store.dispatch('update_Go_IoTAS_O2_DataArray', mqttData.O2);
				// console.log(`Received MQTT message from ${topic} topic: O2`, mqttData.O2);
				this.updateDate();
			});
		},
		updateDate() {
			const now = new Date();
			this.currentMonth = (now.getMonth() + 1).toString();
			this.currentDay = now.getDate().toString();

			// 保存过去七个小时的时间
			for (let i = 6; i >= 0; i--) {
				const pastDate = new Date(now);
				pastDate.setHours(now.getHours() - i);

				// 格式化为 "某时" 格式
				const formattedTime = `${pastDate.getHours()}:00`;

				this.Time.push(formattedTime);
			}

			this.$store.commit('set_Time', this.Time);
			console.log('time', this.$store.state.Time);
			this.Time = [];
		},
	},
};