import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate';
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
    //空气质量
    Go_IoTAE_Air_quality_DataArray: [10, 11, 13, 11, 12, 12, 9],
    //光照强度
    Go_IoTAE_Light_intensity_DataArray: [120, 200, 150, 80, 70, 110, 130],
    //二氧化碳浓度
    Go_IoTAE_CO2_DataArray: [820, 932, 901, 934, 1290, 1330, 1320],
    //甲醇浓度
    Go_IoTAE_Methanol_DataArray: [2.6, 5.9, 9.0, 26.4, 28.7, 70.7, 175.6],
    //空气湿度
    Go_IoTAE_AH_DataArray: [120, 200, 150, 80, 70, 110, 130],
    //空气温度
    Go_IoTAE_Air_temperature_DataArray: [10, 11, 13, 11, 12, 12, 9],
    //土壤PH值
    Go_IoTAE_PH_DataArray: [150, 230, 224, 218, 135, 147, 260],
    //氮元素浓度
    Go_IoTAE_N_DataArray: [120, 200, 150, 80, 70, 110, 130],
    //磷元素浓度
    Go_IoTAE_P_DataArray: [120, 200, 150, 80, 70, 110, 130],
    //钾元素浓度
    Go_IoTAE_K_DataArray: [120, 200, 150, 80, 70, 110, 130],
    //土壤温度
    Go_IoTAE_ST_DataArray: [500, 600, 750, 800, 700, 600, 400],
    //土壤湿度
    Go_IoTAE_SH_DataArray: [120, 200, 150, 80, 70, 110, 130],
    //控制窗口显示
    isModalVisible: false,
    //AI监控数据-Go_IoTAS
    Go_IoTAS_ASID: 1,
    //环境温度
    Go_IoTAS_T_DataArray: [120, 200, 150, 80, 70, 110, 130],
    //环境湿度
    Go_IoTAS_H_DataArray: [120, 200, 150, 80, 70, 110, 130],
    // 环境二氧化碳浓度
    Go_IoTAS_CO2_DataArray: [120, 200, 150, 80, 70, 110, 130],
    //环境氧气浓度
    Go_IoTAS_O2_DataArray: [120, 200, 150, 80, 70, 110, 130],
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

      if (state.Go_IoTAE_Air_quality_DataArray.length > 8) {
        state.Go_IoTAE_Air_quality_DataArray.shift();
      }
    },
    //更新光照强度数据
    update_Light_intensity_DataArray(state, newData) {
      state.Go_IoTAE_Light_intensity_DataArray.push(newData);

      if (state.Go_IoTAE_Light_intensity_DataArray.length > 8) {
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
    createPersistedState(),
  ],
})
export default store