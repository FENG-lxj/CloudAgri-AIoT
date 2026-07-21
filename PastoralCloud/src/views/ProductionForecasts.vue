<template>
  <div class="app-container">
    <Navbar />
    <div class="content-container">
      <div v-if="!showResult" class="input-view">
        <div class="header-section">
          <h1>农作物产量预测系统</h1>
          <p class="subtitle">根据土壤条件和气候环境，科学预测作物产量</p>
        </div>

        <div class="input-section">
          <div class="input-header">
            <h2>产量预测参数输入</h2>
            <p>请选择作物类型并填写相关参数，我们将为您预测作物产量</p>
          </div>

          <div class="form-group">
            <label>作物类型</label>
            <select v-model="selectedCrop" class="crop-select">
              <option v-for="crop in CROPS" :key="crop" :value="crop">{{ crop }}</option>
            </select>
          </div>

          <div class="input-table">
            <!-- 第一行 -->
            <div class="input-row">
              <div class="input-cell">
                <label>氮含量(N)(kg/亩)</label>
                <div class="number-input">
                  <input type="number" v-model.number="soilData.N" min="0" max="100" step="0.1">
                </div>
                <span class="hint">土壤中的氮含量，一般在0-100kg/亩之间</span>
              </div>
              <div class="input-cell">
                <label>磷含量(P)(kg/亩)</label>
                <div class="number-input">
                  <input type="number" v-model.number="soilData.P" min="0" max="100" step="0.1">
                </div>
                <span class="hint">土壤中的磷含量，一般在0-100kg/亩之间</span>
              </div>
            </div>

            <!-- 第二行 -->
            <div class="input-row">
              <div class="input-cell">
                <label>钾含量(K)(kg/亩)</label>
                <div class="number-input">
                  <input type="number" v-model.number="soilData.K" min="0" max="100" step="0.1">
                </div>
                <span class="hint">土壤中的钾含量，一般在0-100kg/亩之间</span>
              </div>
              <div class="input-cell">
                <label>年均温度(℃)</label>
                <div class="number-input">
                  <input type="number" v-model.number="soilData.temperature" min="-10" max="50" step="0.1">
                </div>
                <span class="hint">地区年平均温度，一般在-10-50℃之间</span>
              </div>
            </div>

            <!-- 第三行 -->
            <div class="input-row">
              <div class="input-cell">
                <label>湿度(%)</label>
                <div class="number-input">
                  <input type="number" v-model.number="soilData.humidity" min="0" max="100" step="0.1">
                </div>
                <span class="hint">空气湿度，范围在0-100%之间</span>
              </div>
              <div class="input-cell">
                <label>pH值</label>
                <div class="number-input">
                  <input type="number" v-model.number="soilData.ph" min="3" max="10" step="0.1">
                </div>
                <span class="hint">土壤pH值，范围在3-10之间</span>
              </div>
            </div>

            <!-- 第四行 -->
            <div class="input-row">
              <div class="input-cell">
                <label>年降雨量(mm)</label>
                <div class="number-input">
                  <input type="number" v-model.number="soilData.rainfall" min="200" max="4000" step="1">
                </div>
                <span class="hint">地区年降雨量，一般在200-4000mm之间</span>
              </div>
              <div class="input-cell">
                <label>农药用量(kg/ha)</label>
                <div class="number-input">
                  <input type="number" v-model.number="soilData.pesticide" min="0" max="100" step="0.1">
                </div>
                <span class="hint">每公顷农药使用量，一般在0-100kg/ha之间</span>
              </div>
            </div>
          </div>

          <div class="button-group">
            <button class="btn secondary" @click="useExampleData">使用示例数据</button>
            <button class="btn secondary" @click="resetData">重置数据</button>
            <button class="btn primary" @click="predictYield" :disabled="isPredicting">
              {{ isPredicting ? '预测中...' : '开始预测' }}
            </button>
          </div>
        </div>
      </div>

      <!-- 结果视图 -->
      <div v-else class="result-view">
        <div class="header-section">
          <h1>农作物产量预测系统</h1>
          <p class="subtitle">根据土壤条件和气候环境，科学预测作物产量</p>
        </div>

        <div class="result-section">
          <div class="result-header">
            <h2>产量预测结果</h2>
            <button class="btn back-btn" @click="showResult = false">返回修改</button>
          </div>

          <div class="yield-display">
            <div class="crop-icon">🌾</div>
            <h3>{{ predictionResult.crop }}</h3>
            <div class="yield-value">
              {{ predictionResult.yield }} <span class="unit">公斤/公顷</span>
            </div>
            <div class="yield-rating" :class="getRatingClass(predictionResult.rating)">
              {{ predictionResult.rating }}
            </div>
            <p class="yield-comment">{{ predictionResult.comment }}</p>
          </div>

          <div class="input-summary">
            <h3>您输入的参数</h3>
            <div class="summary-table">
              <div class="summary-row">
                <div class="summary-cell">氮含量(N)</div>
                <div class="summary-cell">磷含量(P)</div>
                <div class="summary-cell">钾含量(K)</div>
                <div class="summary-cell">年均温度</div>
              </div>
              <div class="summary-row">
                <div class="summary-cell">{{ soilData.N }} kg/亩</div>
                <div class="summary-cell">{{ soilData.P }} kg/亩</div>
                <div class="summary-cell">{{ soilData.K }} kg/亩</div>
                <div class="summary-cell">{{ soilData.temperature }}℃</div>
              </div>
              <div class="summary-row">
                <div class="summary-cell">湿度</div>
                <div class="summary-cell">pH值</div>
                <div class="summary-cell">年降雨量</div>
                <div class="summary-cell">农药用量</div>
              </div>
              <div class="summary-row">
                <div class="summary-cell">{{ soilData.humidity }}%</div>
                <div class="summary-cell">{{ soilData.ph }}</div>
                <div class="summary-cell">{{ soilData.rainfall }}mm</div>
                <div class="summary-cell">{{ soilData.pesticide }}kg/ha</div>
              </div>
            </div>
          </div>

          <div class="suggestions-section">
            <h3>提高产量建议</h3>
            <ul>
              <li v-for="(suggestion, index) in predictionResult.suggestions" :key="index">
                {{ suggestion }}
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios from 'axios';
import Navbar from "../components/Navbar.vue";

const showResult = ref(false);
const isPredicting = ref(false);
const selectedCrop = ref("木豆");

const CROPS = ref([
  "苹果", "香蕉", "大麦", "蓝莓", "荞麦", "卷心菜",
  "胡萝卜和萝卜", "樱桃", "黄瓜和小黄瓜",
  "茄子", "葡萄", "青玉米", "青蒜",
  "花生", "猕猴桃", "韭菜和其他葱属蔬菜", "生菜和菊苣",
  "玉米", "芒果、番石榴和山竹", "小米", "燕麦",
  "橙子", "梨",  "柿子", "菠萝", "柚子和葡萄柚", "马铃薯", "南瓜、西葫芦和葫芦",
  "芝麻", "高粱", "大豆", "菠菜", "草莓", "甜菜",
  "甘蔗", "甘薯", "番茄", "黑小麦", "西瓜", "小麦"
]);

const soilData = ref({
  N: 20,
  P: 10,
  K: 30,
  temperature: 24,
  humidity: 75,
  ph: 6.8,
  rainfall: 1000,
  pesticide: 2.5
});

const predictionResult = ref({
  crop: "",
  yield: 0,
  rating: "",
  comment: "",
  suggestions: []
});

// 作物中英文映射表
const CROP_MAPPING = {
  "苹果": "Apples",
  "香蕉": "Bananas",
  "大麦": "Barley",
  "蓝莓": "Blueberries",
  "荞麦": "Buckwheat",
  "卷心菜": "Cabbages",
  "胡萝卜和萝卜": "Carrots and turnips",
  "樱桃": "Cherries",
  "黄瓜和小黄瓜": "Cucumbers and gherkins",
  "茄子": "Eggplants",
  "葡萄": "Grapes",
  "青玉米": "Green corn",
  "青蒜": "Green garlic",
  "花生": "Groundnuts",
  "猕猴桃": "Kiwi fruit",
  "韭菜和其他葱属蔬菜": "Leeks and other alliaceous vegetables",
  "生菜和菊苣": "Lettuce and chicory",
  "玉米": "Maize",
  "芒果、番石榴和山竹": "Mangoes, guavas and mangosteens",
  "小米": "Millet",
  "燕麦": "Oats",
  "橙子": "Oranges",
  "梨": "Pears",
  "柿子": "Persimmons",
  "菠萝": "Pineapples",
  "柚子和葡萄柚": "Pomelos and grapefruits",
  "马铃薯": "Potatoes",
  "南瓜、西葫芦和葫芦": "Pumpkins, squash and gourds",
  "芝麻": "Sesame seed",
  "高粱": "Sorghum",
  "大豆": "Soya beans",
  "菠菜": "Spinach",
  "草莓": "Strawberries",
  "甜菜": "Sugar beet",
  "甘蔗": "Sugar cane",
  "甘薯": "Sweet potatoes",
  "番茄": "Tomatoes",
  "黑小麦": "Triticale",
  "西瓜": "Watermelons",
  "小麦": "Wheat"
};

// 使用示例数据
const useExampleData = () => {
  soilData.value = {
    N: 45,
    P: 25,
    K: 55,
    temperature: 24,
    humidity: 55,
    ph: 6.8,
    rainfall: 300,
    pesticide: 2.5
  };
  selectedCrop.value = "木豆";
};

// 重置数据
const resetData = () => {
  soilData.value = {
    N: 0,
    P: 0,
    K: 0,
    temperature: 0,
    humidity: 0,
    ph: 0,
    rainfall: 0,
    pesticide: 0
  };
  selectedCrop.value = "木豆";
};

// 获取评级样式
const getRatingClass = (rating) => {
  switch(rating) {
    case "优秀": return "excellent";
    case "良好": return "good";
    case "一般": return "average";
    case "较差": return "poor";
    case "很差": return "very-poor";
    default: return "";
  }
};

// 预测产量
const predictYield = async () => {
  try {
    isPredicting.value = true;

    // 准备发送到后端的数据
    const requestData = {
      crop: CROP_MAPPING[selectedCrop.value] || selectedCrop.value,
      features: {
        N: Number(soilData.value.N) || 0,
        P: Number(soilData.value.P) || 0,
        K: Number(soilData.value.K) || 0,
        temperature: Number(soilData.value.temperature) || 0,
        humidity: Number(soilData.value.humidity) || 0,
        ph: Number(soilData.value.ph) || 0,
        rainfall: Number(soilData.value.rainfall) || 0,
        pesticide: Number(soilData.value.pesticide) || 0
      }
    };

    console.log("发送请求数据:", JSON.stringify(requestData, null, 2));

    // 发送请求到后端API
    const response = await axios.post('http://127.0.0.1:8901/predict',
        requestData,

    );

    console.log("收到响应:", response.data);

    // 验证响应数据
    if (!response.data || typeof response.data.yield !== 'number') {
      throw new Error('无效的响应数据');
    }

    // 更新预测结果
    predictionResult.value = {
      crop: selectedCrop.value,
      yield: response.data.yield,
      rating: response.data.rating || "未知",
      comment: response.data.comment || "无评价",
      suggestions: response.data.suggestions || ["暂无建议"]
    };

    showResult.value = true;
  } catch (error) {
    console.error('预测失败:', error);

    // 更友好的错误提示
    let errorMessage = '预测失败';
    if (error.response) {
      errorMessage = error.response.data?.error ||
          `服务器错误: ${error.response.status}`;
    } else if (error.request) {
      errorMessage = '无法连接到服务器，请检查:';
      errorMessage += '\n1. 后端服务是否运行';
      errorMessage += '\n2. 网络连接是否正常';
      errorMessage += '\n3. 端口是否被占用';
    } else {
      errorMessage = error.message;
    }

    alert(`错误: ${errorMessage}`);
  } finally {
    isPredicting.value = false;
  }
};
</script>
<style scoped>
/* 基础布局样式 - 使用清新自然的农业风格 */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html, body, #app {
  height: 100%;
  width: 100%;
}

/* 主应用容器 - 浅绿色背景 */
.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  background-color: #f5faf5; /* 浅绿色背景 */
  color: #2c3e2c; /* 深绿色文字 */
  font-family: 'Arial', sans-serif;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  overflow-y: auto;
}

/* 内容容器 - 白色卡片效果 */
.content-container {
  flex: 1;
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

/* 头部区域 - 农业主题绿色 */
.header-section {
  position: relative;
  text-align: center;
  margin-bottom: 20px;
  padding: 40px;
  color: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  overflow: hidden;
}

/* 背景图片层 */
.header-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-image: url('../img/Round-robin_diagram/true.jpg');
  background-size: cover;
  background-position: center;
  opacity: 0.8; /* 80%透明度 */
  z-index: -1;
}

/* 颜色叠加层 */
.header-section::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #4CAF50; /* 农业绿 */
  opacity: 0.7; /* 70%透明度 */
  z-index: -1;
}

.header-section h1 {
  position: relative;
  font-size: 28px;
  margin-bottom: 10px;
  text-shadow: 0 2px 4px rgba(0,0,0,0.3);
}

.subtitle {
  position: relative;
  font-size: 16px;
  color: rgba(255,255,255,0.9);
}

/* 输入区域 - 白色卡片 */
.input-section, .result-section {
  background-color: white;
  border-radius: 8px;
  padding: 30px;
  margin-top: 20px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.05);
  border: 1px solid #e0e0e0;
}

.input-header {
  margin-bottom: 20px;
}

.input-header h2 {
  font-size: 22px;
  color: #2c3e2c; /* 深绿色 */
  margin-bottom: 8px;
}

.input-header p {
  font-size: 14px;
  color: #666;
}

/* 表单组样式 */
.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: #444;
}

.crop-select {
  width: 100%;
  padding: 12px;
  border: 1px solid #a5d6a7;
  border-radius: 6px;
  font-size: 16px;
  background-color: #f8fff8;
  color: #2c3e2c;
}

.crop-select:focus {
  outline: none;
  border-color: #4CAF50;
  box-shadow: 0 0 0 2px rgba(76, 175, 80, 0.2);
}

/* 输入表格 - 两列布局 */
.input-table {
  width: 100%;
  margin: 20px 0;
}

.input-row {
  display: flex;
  margin-bottom: 15px;
}

.input-cell {
  flex: 1;
  padding: 10px;
  margin-right: 20px;
}

.input-cell:last-child {
  margin-right: 0;
}

.input-cell label {
  display: block;
  font-size: 16px;
  margin-bottom: 8px;
  color: #2c3e2c; /* 深绿色 */
  font-weight: 500;
}

/* 数字输入框样式 - 居中显示 */
.number-input {
  display: flex;
  justify-content: center; /* 水平居中 */
  width: 100%; /* 确保容器宽度100% */
}

.number-input input {
  width: 80%; /* 保持80%宽度 */
  padding: 12px 15px;
  border: 1px solid #a5d6a7; /* 浅绿色边框 */
  border-radius: 6px;
  font-size: 16px;
  color: #2c3e2c;
  background-color: #f8fff8; /* 浅绿色背景 */
  text-align: center; /* 输入文字居中 */
  display: block; /* 块级元素 */
  margin: 0 auto; /* 水平居中 */
}

.number-input input:focus {
  outline: none;
  border-color: #4CAF50;
  box-shadow: 0 0 0 2px rgba(76, 175, 80, 0.2);
}

.hint {
  display: block;
  font-size: 13px;
  color: #888;
  margin-top: 8px;
}

/* 按钮组 - 农业主题色 */
.button-group {
  display: flex;
  justify-content: flex-end;
  gap: 15px;
  margin-top: 25px;
}

.btn {
  padding: 12px 24px;
  border-radius: 6px;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  border: none;
}

.btn.primary {
  background-color: #4CAF50; /* 农业绿 */
  color: white;
}

.btn.secondary {
  background-color: #81c784; /* 浅绿色 */
  color: white;
}

.btn.back-btn {
  background-color: #2196F3; /* 蓝色 */
  color: white;
}

.btn:hover {
  opacity: 0.9;
  transform: translateY(-2px);
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

/* 结果区域 - 保持清新风格 */
.result-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 25px;
}

.result-header h2 {
  font-size: 22px;
  color: #2c3e2c;
}

.yield-display {
  text-align: center;
  padding: 20px 0;
  border-bottom: 1px solid #e0e0e0;
  margin-bottom: 25px;
}

.crop-icon {
  font-size: 60px;
  margin-bottom: 15px;
  color: #4CAF50;
}

.yield-display h3 {
  font-size: 24px;
  color: #4CAF50;
  margin-bottom: 8px;
}

.yield-value {
  font-size: 36px;
  font-weight: bold;
  color: #4CAF50;
  margin-bottom: 10px;
}

.unit {
  font-size: 20px;
  color: #666;
}

.yield-rating {
  display: inline-block;
  padding: 8px 16px;
  border-radius: 20px;
  font-weight: 500;
  margin-bottom: 15px;
}

.yield-rating.excellent {
  background-color: #4CAF50;
  color: white;
}

.yield-rating.good {
  background-color: #8BC34A;
  color: white;
}

.yield-rating.average {
  background-color: #FFC107;
  color: white;
}

.yield-rating.poor {
  background-color: #FF9800;
  color: white;
}

.yield-rating.very-poor {
  background-color: #F44336;
  color: white;
}

.yield-comment {
  font-size: 16px;
  color: #666;
}

/* 参数汇总表格 */
.input-summary {
  margin-bottom: 25px;
}

.input-summary h3 {
  font-size: 20px;
  color: #2c3e2c;
  margin-bottom: 15px;
}

.summary-table {
  width: 100%;
  border-collapse: collapse;
}

.summary-row {
  display: flex;
}

.summary-cell {
  flex: 1;
  padding: 12px;
  text-align: center;
  border: 1px solid #e0e0e0;
  font-size: 16px;
}

.summary-row:nth-child(odd) {
  background-color: #f8fff8; /* 非常浅的绿色背景 */
}

/* 建议区域 */
.suggestions-section {
  background-color: #f9f9f9;
  border-radius: 8px;
  padding: 20px;
}

.suggestions-section h3 {
  font-size: 20px;
  color: #2c3e2c;
  margin-bottom: 15px;
}

.suggestions-section ul {
  padding-left: 20px;
}

.suggestions-section li {
  margin-bottom: 10px;
  color: #666;
  line-height: 1.6;
}

/* 响应式设计 - 移动端适配 */
@media (max-width: 768px) {
  .content-container {
    padding: 15px;
  }

  .header-section {
    padding: 20px;
  }

  .header-section h1 {
    font-size: 22px;
  }

  .subtitle {
    font-size: 14px;
  }

  .input-row {
    flex-direction: column;
  }

  .input-cell {
    margin-right: 0;
    margin-bottom: 15px;
  }

  .button-group {
    flex-direction: column;
  }

  .btn {
    width: 100%;
  }

  .result-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 15px;
  }

  .summary-row {
    flex-wrap: wrap;
  }

  .summary-cell {
    flex: 0 0 50%;
  }
}
</style>