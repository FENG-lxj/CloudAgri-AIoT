import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    Phone: 0,
    AI_Camera_One_ICID: 1,
    Time:[],
    count: 0,
    phoneCaptcha: 0,
    RemovePhone: 0,
    //AEID
    AEID: 1,
    // 空气质量
	// "AEID":1,"LI":31863.06,"PH":5.5,"ST":21.55,"SM":34.22,"N":309.75,"P":14.65,
	// "K":651.76,"AT":9.29,"AH":20.79,"AQ":85.6,"CO2":0.028,"CH4O":0.003
    Go_IoTAE_Air_quality_DataArray: [86.2, 87.7, 88.3, 85.8, 84.3, 86.5, 89.4, 86.6],
    // 光照强度
    Go_IoTAE_Light_intensity_DataArray: [150, 180, 160, 210, 230, 190, 170],
    // 二氧化碳浓度
    Go_IoTAE_CO2_DataArray: [520, 532, 541, 564, 690, 730, 720],
    // 甲醇浓度
    Go_IoTAE_Methanol_DataArray: [1.8, 4.2, 6.5, 18.4, 21.7, 50.7, 135.6],
    // 空气湿度
    Go_IoTAE_AH_DataArray: [45, 50, 55, 60, 65, 58, 52],
    // 空气温度
    Go_IoTAE_Air_temperature_DataArray: [15, 16, 18, 17, 19, 18, 16],
    // 土壤 PH 值
    Go_IoTAE_PH_DataArray: [6.2, 6.5, 6.3, 6.4, 6.7, 6.6, 6.8],
    // 氮元素浓度
    Go_IoTAE_N_DataArray: [50, 70, 60, 80, 90, 75, 65],
    // 磷元素浓度
    Go_IoTAE_P_DataArray: [30, 40, 35, 45, 50, 42, 38],
    // 钾元素浓度
    Go_IoTAE_K_DataArray: [60, 80, 70, 90, 100, 85, 75],
    // 土壤温度
    Go_IoTAE_ST_DataArray: [25, 26, 27, 28, 29, 28, 27],
    // 土壤湿度
    Go_IoTAE_SH_DataArray: [30, 35, 40, 45, 50, 42, 38],
    //控制窗口显示
    isModalVisible: false,
    //AI监控数据-Go_IoTAS
    Go_IoTAS_ASID: 1,
    // 环境温度
    Go_IoTAS_T_DataArray: [20, 22, 24, 23, 25, 24, 22],
    // 环境湿度
    Go_IoTAS_H_DataArray: [55, 60, 65, 70, 75, 68, 62],
    // 环境二氧化碳浓度
    Go_IoTAS_CO2_DataArray: [420, 432, 441, 464, 590, 630, 620],
    // 环境氧气浓度
    Go_IoTAS_O2_DataArray: [21, 21.2, 21.5, 21.3, 21.4, 21.3, 21.2]
  },
  mutations: {
    //更新手机号
    setPhone(state, newPhone) {
      state.Phone = newPhone
    },
    //更新日期
    set_Time(state, newData) {
      state.Time = newData;
    },    
    set_AI_Camera_One_ICID(state, newData) {
      state.AI_Camera_One_ICID = newData;
    },
    //控制窗口
    set_isModalVisible(state, newData) {
      state.isModalVisible = newData;
    },
    setphoneCaptcha(state, newphoneCaptcha) {
      state.phoneCaptcha = newphoneCaptcha
    },
    setRemovePhone(state, newRemovePhone) {
      state.RemovePhone = newRemovePhone
    },
    //MQTT协议修改数据
    setMqttData(state, newData) {
      state.mqttData = newData;
    },
    //更新AEID
    update_AEID_DataArray(state, newData) {
      state.AEID = newData;
    },
    //更新空气质量数据
    update_Air_quality_DataArray(state, newData) {
      state.Go_IoTAE_Air_quality_DataArray.push(newData);

      if (state.Go_IoTAE_Air_quality_DataArray.length > 7) {
        state.Go_IoTAE_Air_quality_DataArray.shift();
      }
    },
    //更新光照强度数据
    update_Light_intensity_DataArray(state, newData) {
      state.Go_IoTAE_Light_intensity_DataArray.push(newData);

      if (state.Go_IoTAE_Light_intensity_DataArray.length > 9) {
        state.Go_IoTAE_Light_intensity_DataArray.shift();
      }
    },
    //更新二氧化碳浓度数据
    update_CO2_DataArray(state, newData) {
      state.Go_IoTAE_CO2_DataArray.push(newData);

      if (state.Go_IoTAE_CO2_DataArray.length > 8) {
        state.Go_IoTAE_CO2_DataArray.shift();
      }
    },
    //更新甲醇浓度数据
    update_Methanol_DataArray(state, newData) {
      state.Go_IoTAE_Methanol_DataArray.push(newData);

      if (state.Go_IoTAE_Methanol_DataArray.length > 8) {
        state.Go_IoTAE_Methanol_DataArray.shift();
      }
    },
    //更新空气湿度
    update_AH_DataArray(state, newData) {
      state.Go_IoTAE_AH_DataArray.push(newData);

      if (state.Go_IoTAE_AH_DataArray.length > 8) {
        state.Go_IoTAE_AH_DataArray.shift();
      }
    },
    //更新空气温度
    update_Air_temperature_DataArray(state, newData) {
      state.Go_IoTAE_Air_temperature_DataArray.push(newData);

      if (state.Go_IoTAE_Air_temperature_DataArray.length > 8) {
        state.Go_IoTAE_Air_temperature_DataArray.shift();
      }
    },
    //更新土壤PH值数据
    update_PH_DataArray(state, newData) {
      state.Go_IoTAE_PH_DataArray.push(newData);

      if (state.Go_IoTAE_PH_DataArray.length > 8) {
        state.Go_IoTAE_PH_DataArray.shift();
      }
    },
    //更新氮元素浓度数据
    update_N_DataArray(state, newData) {
      state.Go_IoTAE_N_DataArray.push(newData);

      if (state.Go_IoTAE_N_DataArray.length > 8) {
        state.Go_IoTAE_N_DataArray.shift();
      }
    },
    //更新磷元素浓度数据
    update_P_DataArray(state, newData) {
      state.Go_IoTAE_P_DataArray.push(newData);

      if (state.Go_IoTAE_P_DataArray.length > 8) {
        state.Go_IoTAE_P_DataArray.shift();
      }
    },
    //更新钾元素浓度数据
    update_K_DataArray(state, newData) {
      state.Go_IoTAE_K_DataArray.push(newData);

      if (state.Go_IoTAE_K_DataArray.length > 8) {
        state.Go_IoTAE_K_DataArray.shift();
      }
    },
    //更新土壤温度数据
    update_ST_DataArray(state, newData) {
      state.Go_IoTAE_ST_DataArray.push(newData);

      if (state.Go_IoTAE_ST_DataArray.length > 8) {
        state.Go_IoTAE_ST_DataArray.shift();
      }
    },
    //更新土壤湿度数据
    update_SH_DataArray(state, newData) {
      state.Go_IoTAE_SH_DataArray.push(newData);

      if (state.Go_IoTAE_SH_DataArray.length > 8) {
        state.Go_IoTAE_SH_DataArray.shift();
      }
    },
    //更新AI监控数据-Go_IoTAS
    update_Go_IoTAS_ASID_DataArray(state, newData) {
      state.Go_IoTAS_ASID = newData;

    },
    //更新环境温度数据
    update_Go_IoTAS_T_DataArray(state, newData) {
      state.Go_IoTAS_T_DataArray.push(newData);

      if (state.Go_IoTAS_T_DataArray.length > 8) {
        state.Go_IoTAS_T_DataArray.shift();
      }
    },
    //更新环境湿度数据
    update_Go_IoTAS_H_DataArray(state, newData) {
      state.Go_IoTAS_H_DataArray.push(newData);

      if (state.Go_IoTAS_H_DataArray.length > 8) {
        state.Go_IoTAS_H_DataArray.shift();
      }
    },
    //更新环境二氧化碳浓度数据
    update_Go_IoTAS_CO2_DataArray(state, newData) {
      state.Go_IoTAS_CO2_DataArray.push(newData);

      if (state.Go_IoTAS_CO2_DataArray.length > 8) {
        state.Go_IoTAS_CO2_DataArray.shift();
      }
    },
    //更新环境氧气浓度数据
    update_Go_IoTAS_O2_DataArray(state, newData) {
      state.Go_IoTAS_O2_DataArray.push(newData);

      if (state.Go_IoTAS_O2_DataArray.length > 8) {
        state.Go_IoTAS_O2_DataArray.shift();
      }
    },
  },
  actions: {
    updateMqttData({ commit }, newData) {
      commit('setMqttData', newData);
    },
    //更新AEID
    update_AEID_DataArray({ commit }, newData) {
      commit('update_AEID_DataArray', newData);
    },
    //更新空气质量
    update_Air_quality_DataArray({ commit }, newData) {
      commit('update_Air_quality_DataArray', newData);
    },
    //更新光照质量
    update_Light_intensity_DataArray({ commit }, newData) {
      commit('update_Light_intensity_DataArray', newData);
    },
    //更新二氧化碳浓度
    update_CO2_DataArray({ commit }, newData) {
      commit('update_CO2_DataArray', newData);
    },
    //更新甲醇浓度
    update_Methanol_DataArray({ commit }, newData) {
      commit('update_Methanol_DataArray', newData);
    },
    //更新空气湿度
    update_AH_DataArray({ commit }, newData) {
      commit('update_AH_DataArray', newData);
    },
    //更新空气温度
    update_Air_temperature_DataArray({ commit }, newData) {
      commit('update_Air_temperature_DataArray', newData);
    },
    //更新土壤PH值
    update_PH_DataArray({ commit }, newData) {
      commit('update_PH_DataArray', newData);
    },
    //更新氮元素浓度
    update_N_DataArray({ commit }, newData) {
      commit('update_N_DataArray', newData);
    },
    //更新磷元素浓度
    update_P_DataArray({ commit }, newData) {
      commit('update_P_DataArray', newData);
    },
    //更新钾元素浓度
    update_K_DataArray({ commit }, newData) {
      commit('update_K_DataArray', newData);
    },
    //更新土壤温度
    update_ST_DataArray({ commit }, newData) {
      commit('update_ST_DataArray', newData);
    },
    //更新土壤湿度
    update_SH_DataArray({ commit }, newData) {
      commit('update_SH_DataArray', newData);
    },
    //更新AI监控数据-Go_IoTAS
    update_Go_IoTAS_ASID_DataArray({ commit }, newData) {
      commit('update_Go_IoTAS_ASID_DataArray', newData);
    },
    //更新环境温度
    update_Go_IoTAS_T_DataArray({ commit }, newData) {
      commit('update_Go_IoTAS_T_DataArray', newData);
    },
    //更新环境湿度
    update_Go_IoTAS_H_DataArray({ commit }, newData) {
      commit('update_Go_IoTAS_H_DataArray', newData);
    },
    //更新环境二氧化碳浓度
    update_Go_IoTAS_CO2_DataArray({ commit }, newData) {
      commit('update_Go_IoTAS_CO2_DataArray', newData);
    },
    //更新环境氧气浓度
    update_Go_IoTAS_O2_DataArray({ commit }, newData) {
      commit('update_Go_IoTAS_O2_DataArray', newData);
    },
  },
  plugins: [
    // createPersistedState(),
  ],
})
export default store