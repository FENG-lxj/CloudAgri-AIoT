<template>
  <div class="app-container">
    <Navbar />
    <div class="content-container">
      <div v-if="!showResult" class="input-view">
        <div class="header-section">
          <h1>作物推荐系统</h1>
          <p class="subtitle">根据您的土壤条件和气候环境，科学推荐最适合种植的作物</p>
        </div>

        <div class="input-section">
          <div class="input-header">
            <h2>输入土壤和气候条件</h2>
            <p>请填写以下参数，我们将根据这些数据为您推荐最适合的作物种类</p>
          </div>

          <div class="input-table">
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
                  <input type="number" v-model.number="soilData.ph" min="0" max="14" step="0.1">
                </div>
                <span class="hint">土壤pH值，范围在0-14之间</span>
              </div>
            </div>

            <div class="input-row">
              <div class="input-cell full-width">
                <label>年降雨量(mm)</label>
                <div class="number-input">
                  <input type="number" v-model.number="soilData.rainfall" min="0" max="5000" step="0.1">
                </div>
                <span class="hint">地区年降雨量，一般在0-5000mm之间</span>
              </div>
            </div>
          </div>

          <div class="button-group">
            <button class="btn secondary" @click="useExampleData">使用示例数据</button>
            <button class="btn secondary" @click="resetData">重置数据</button>
            <button class="btn primary" @click="getRecommendation">开始推荐</button>
          </div>
        </div>
      </div>

      <div v-else class="result-view">
        <div class="header-section">
          <h1>作物推荐系统</h1>
          <p class="subtitle">根据您的土壤条件和气候环境，科学推荐最适合种植的作物</p>
        </div>

        <div class="result-section">
          <div class="result-header">
            <h2>推荐结果</h2>
            <button class="btn back-btn" @click="showResult = false">返回修改</button>
          </div>

          <div class="crop-display" v-if="recommendedCrop">
            <div class="crop-icon">{{ cropEmoji }}</div>
            <h3>{{ recommendedCrop.chinese || '未知作物' }} ({{ recommendedCrop.english || 'Unknown' }})</h3>
            <div class="yield-info" v-if="recommendedCrop.similarity !== undefined">
              <span class="yield-value">{{ (recommendedCrop.similarity * 100).toFixed(1) }}</span>
              <span class="yield-unit">% 匹配度</span>
            </div>
            <p class="yield-comment">这是最适合您当前土壤和气候条件的作物</p>
          </div>

          <div class="top-crops" v-if="topCrops && topCrops.length > 0">
            <h3>其他推荐作物</h3>
            <div class="crop-grid">
              <div class="crop-item" v-for="(crop, index) in topCrops" :key="index">
                <div class="crop-emoji">{{ getCropEmoji(crop.english) }}</div>
                <div class="crop-name">{{ crop.chinese || '未知作物' }}</div>
                <div class="crop-yield" v-if="crop.similarity !== undefined">
                  {{ (crop.similarity * 100).toFixed(1) }}% 匹配度
                </div>
              </div>
            </div>
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
                <div class="summary-cell"></div>
              </div>
              <div class="summary-row">
                <div class="summary-cell">{{ soilData.humidity }}%</div>
                <div class="summary-cell">{{ soilData.ph }}</div>
                <div class="summary-cell">{{ soilData.rainfall }}mm</div>
                <div class="summary-cell"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import axios from 'axios';
import Navbar from "../components/Navbar.vue";

const showResult = ref(false);
const recommendedCrop = ref({
  english: '',
  chinese: '',
  similarity: 0
});
const topCrops = ref([]);

const soilData = ref({
  N: 45,
  P: 25,
  K: 55,
  temperature: 25,
  humidity: 50,
  ph: 6.5,
  rainfall: 300
});

// 作物表情符号映射
const cropEmojis = {
  'Apples': '🍎',
  'Bananas': '🍌',
  'Barley': '🌾',
  'Blueberries': '🫐',
  'Buckwheat': '🌾',
  'Cabbages': '🥬',
  'Carrots_and_turnips': '🥕',
  'Cherries': '🍒',
  'Cucumbers_and_gherkins': '🥒',
  'Eggplants': '🍆',
  'Ginger_raw': '🫚',
  'Grapes': '🍇',
  'Green_corn': '🌽',
  'Green_garlic': '🧄',
  'Groundnuts': '🥜',
  'Kiwi_fruit': '🥝',
  'Leeks_and_other_alliaceous_vegetables': '🧅',
  'Lettuce_and_chicory': '🥬',
  'Maize': '🌽',
  'Millet': '🌾',
  'Oats': '🌾',
  'Oranges': '🍊',
  'Pears': '🍐',
  'Peas_green': '🫛',
  'Persimmons': '🍅',
  'Pineapples': '🍍',
  'Pomelos_and_grapefruits': '🍈',
  'Potatoes': '🥔',
  'Sesame_seed': '🌰',
  'Sorghum': '🌾',
  'Soya_beans': '🫘',
  'Spinach': '🍃',
  'Strawberries': '🍓',
  'Sugar_beet': '🍠',
  'Sugar_cane': '🎋',
  'Sweet_potatoes': '🍠',
  'Tomatoes': '🍅',
  'Triticale': '🌾',
  'Watermelons': '🍉',
  'Wheat': '🌾'
};

// 计算当前推荐作物的表情符号
const cropEmoji = computed(() => {
  return cropEmojis[recommendedCrop.value.english] || '🌱';
});

// 获取作物表情符号
const getCropEmoji = (cropName) => {
  return cropEmojis[cropName] || '🌱';
};

// 使用示例数据
const useExampleData = () => {
  soilData.value = {
    N: 45,
    P: 25,
    K: 55,
    temperature: 25,
    humidity: 50,
    ph: 6.5,
    rainfall: 300
  };
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
    rainfall: 0
  };
};

// 获取推荐
const getRecommendation = async () => {
  try {
    const response = await axios.post('http://localhost:8900/recommend_crop', {
      features: {
        N: soilData.value.N,
        P: soilData.value.P,
        K: soilData.value.K,
        temperature: soilData.value.temperature,
        humidity: soilData.value.humidity,
        ph: soilData.value.ph,
        rainfall: soilData.value.rainfall
      }
    });

    if (response.data?.recommended_crop) {
      recommendedCrop.value = {
        english: response.data.recommended_crop_english,
        chinese: response.data.recommended_crop,
        similarity: response.data.similarity_score
      };

      topCrops.value = response.data.top_crops.map(crop => ({
        english: crop.english_name,
        chinese: crop.name,
        similarity: crop.similarity_score
      }));
    }

    showResult.value = true;
  } catch (error) {
    console.error('获取推荐失败:', error);
    alert('获取推荐失败，请检查后端服务是否运行');
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
  background-image: url('../img/Round-robin_diagram/true.jpg');
  opacity: 0.8; /* 80%透明度 */
  z-index: -1; /* 置于内容下方 */
  text-align: center;
  margin-bottom: 0px;
  padding: 40px;
  background-color: #4CAF50; /* 农业绿 */
  color: green;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.header-section h1 {
  font-size: 28px;
  margin-bottom: 10px;
}

.subtitle {
  font-size: 20px;
  color: green;
}

/* 输入区域 - 白色卡片 */
.input-section, .result-section {
  background-color: white;
  border-radius: 8px;
  padding: 25px;
  margin-top: 20px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
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

.input-cell.full-width {
  flex: 2;
}

.input-cell label {
  display: block;
  font-size: 16px;
  margin-bottom: 8px;
  color: #2c3e2c; /* 深绿色 */
  font-weight: 500;
}

/* 数字输入框 - 农业风格 */
.number-input {
  display: flex;
  align-items: center;
}

/* 数字输入框样式 - 居中显示 */
.number-input {
  display: flex;
  justify-content: center; /* 水平居中 */
  width: 100%; /* 确保容器宽度100% */
}

.number-input input {
  width: 80%; /* 保持70%宽度 */
  padding: 12px 15px;
  border: 1px solid #a5d6a7; /* 浅绿色边框 */
  border-radius: 6px;
  font-size: 18px;
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

/* 特别标注磷含量输入框 */
.input-row:nth-child(1) .input-cell:nth-child(2) input {
  border-color: #4CAF50;
}

/* 特别标注pH值输入框 */
.input-row:nth-child(3) .input-cell:nth-child(2) input {
  border-color: #4CAF50; /* 绿色边框 */
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
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
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

.crop-display {
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

.crop-display h3 {
  font-size: 24px;
  color: #4CAF50;
  margin-bottom: 8px;
}

.crop-display p {
  font-size: 16px;
  color: #666;
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

/* 响应式设计 - 移动端适配 */
@media (max-width: 768px) {
  .content-container {
    padding: 15px;
  }

  .header-section {
    padding: 15px;
  }

  .header-section h1 {
    font-size: 22px;
  }

  .subtitle {
    font-size: 14px;
  }

  .input-grid {
    grid-template-columns: 1fr;
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

  .parameter-grid {
    grid-template-columns: 1fr;
  }

  .crop-icon {
    font-size: 48px;
  }

  .crop-display h3 {
    font-size: 20px;
  }
}

/* 按钮样式 - 三种不同颜色 */
.button-group {
  display: flex;
  justify-content: flex-end;
  gap: 15px;
  margin-top: 30px;
}

.btn {
  padding: 12px 24px;
  border-radius: 6px;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  border: none;
  color: white; /* 统一白色文字 */
}

/* 1. 使用示例数据按钮 - 中性色 */
.btn.secondary:nth-child(1) {
  background-color: #607D8B; /* 蓝灰色 */
}

/* 2. 重置数据按钮 - 警示色 */
.btn.secondary:nth-child(2) {
  background-color: #F44336; /* 红色 */
}

/* 3. 开始推荐按钮 - 主题绿色 */
.btn.primary {
  background-color: #4CAF50; /* 农业绿 */
}

/* 悬停效果 */
.btn:hover {
  opacity: 0.9;
  transform: translateY(-2px);
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

/* 新增样式 */
.yield-info {
  margin: 15px 0;
  display: flex;
  align-items: baseline;
  justify-content: center;
}

.yield-value {
  font-size: 36px;
  font-weight: bold;
  color: #4CAF50;
  margin-right: 8px;
}

.yield-unit {
  font-size: 18px;
  color: #666;
}

.yield-comment {
  font-size: 16px;
  color: #666;
  margin-bottom: 20px;
}

.top-crops {
  margin: 30px 0;
}

.top-crops h3 {
  font-size: 20px;
  color: #2c3e2c;
  margin-bottom: 15px;
  text-align: center;
}

.crop-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 20px;
  margin-top: 15px;
}

.crop-item {
  background-color: #f8fff8;
  border-radius: 8px;
  padding: 15px;
  text-align: center;
  transition: all 0.3s ease;
}

.crop-item:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.crop-emoji {
  font-size: 32px;
  margin-bottom: 8px;
}

.crop-name {
  font-weight: 500;
  margin-bottom: 5px;
}

.crop-yield {
  font-size: 14px;
  color: #666;
}
</style>