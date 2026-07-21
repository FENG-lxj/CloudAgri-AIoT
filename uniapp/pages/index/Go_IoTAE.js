// Go_IoTAE.js
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
				url: 'http://127.0.0.1:20201/IoTA/bigdata/GetAEdata',
				data: {
					AEID: this.$store.state.AEID,
					StartTime: '-24h'
				},
				header: {
				    'Authorization': `Bearer ${uni.getStorageSync('token')}`,
				},
				success: (res) => {
					// 使用循环将数据逐个存入 Air_quality_DataArray
					res.data.data.AEDatas.slice(-7).forEach((data) => {
						// this.$store.dispatch('update_Air_quality_DataArray', data);
						console.log('data', data);
						//更新空气质量数据
						this.$store.dispatch('update_Air_quality_DataArray', data.AQ);
						//更新光照质量数据
						this.$store.dispatch('update_Light_intensity_DataArray', data.LI);
						//更新二氧化碳浓度数据
						this.$store.dispatch('update_CO2_DataArray', data.CO2);
						//更新甲醇浓度数据
						this.$store.dispatch('update_Methanol_DataArray', data.CH4O);
						// 更新空气湿度数据
						this.$store.dispatch('update_AH_DataArray', data.AH);
						//更新空气温度数据
						this.$store.dispatch('update_Air_temperature_DataArray', data.AT);
						//更新土壤PH值数据
						this.$store.dispatch('update_PH_DataArray', data.PH);
						//更新氮元素浓度数据
						this.$store.dispatch('update_N_DataArray', data.N);
						//更新磷元素浓度数据
						this.$store.dispatch('update_P_DataArray', data.P);
						//更新钾元素浓度数据
						this.$store.dispatch('update_K_DataArray', data.K);
						//更新土壤温度数据
						this.$store.dispatch('update_ST_DataArray', data.ST);
						//更新土壤湿度数据
						this.$store.dispatch('update_SH_DataArray', data.SM);
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
			this.client = mqtt.connect('ws://39.102.208.105:8083/mqtt');

			this.client.subscribe('Go_IoTAE');

			this.client.on('message', (topic, message) => {
				const mqttData = JSON.parse(message.toString());

				console.log(`Received MQTT message from ${topic} topic: AH`, mqttData.AH);
				// console.log(`Received MQTT message from ${topic} topic: O2`, mqttData);

				// 更新 Vuex store 中的数据
				this.$store.dispatch('updateMqttData', mqttData);
				// this.$store.dispatch('updateMqttDataArray', mqttData.CO2);
				//更新AEID
				this.$store.dispatch('update_AEID_DataArray', mqttData.AEID);
				console.log(`Received MQTT message from ${topic} topic: AEID`, mqttData.AEID);
				//更新空气质量数据
				this.$store.dispatch('update_Air_quality_DataArray', mqttData.AQ);
				//更新光照质量数据
				this.$store.dispatch('update_Light_intensity_DataArray', mqttData.LI);
				//更新二氧化碳浓度数据
				this.$store.dispatch('update_CO2_DataArray', mqttData.CO2);
				//更新甲醇浓度数据
				this.$store.dispatch('update_Methanol_DataArray', mqttData.CH4O);
				// 更新空气湿度数据
				this.$store.dispatch('update_AH_DataArray', mqttData.AH);
				//更新空气温度数据
				this.$store.dispatch('update_Air_temperature_DataArray', mqttData.AT);
				//更新土壤PH值数据
				this.$store.dispatch('update_PH_DataArray', mqttData.PH);
				//更新氮元素浓度数据
				this.$store.dispatch('update_N_DataArray', mqttData.N);
				//更新磷元素浓度数据
				this.$store.dispatch('update_P_DataArray', mqttData.P);
				//更新钾元素浓度数据
				this.$store.dispatch('update_K_DataArray', mqttData.K);
				//更新土壤温度数据
				this.$store.dispatch('update_ST_DataArray', mqttData.ST);
				//更新土壤湿度数据
				this.$store.dispatch('update_SH_DataArray', mqttData.SM);
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