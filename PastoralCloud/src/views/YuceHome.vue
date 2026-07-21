<template>
  <div class="home-page">
    <Navbar />

    <!-- 顶部背景和标题 -->
    <div class="hero-section">
      <div class="hero-content">
        <h1 class="hero-title">智能农业预测平台</h1>
        <p class="hero-subtitle">基于人工智能算法，为您提供精准的作物推荐和产量预测</p>
      </div>
    </div>

    <!-- 核心功能 -->
    <div class="features-section">
      <h2 class="section-title">核心功能</h2>
      <div class="features-container">
        <div class="feature-card" @click="$router.push('/YuceHome/CropRecommendations')">
          <div class="feature-icon">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#38a169">
              <path d="M12 2L4 5v6.09c0 5.05 3.41 9.76 8 10.91 4.59-1.15 8-5.86 8-10.91V5l-8-3zm-1.06 13.54L7.4 12l1.41-1.41 2.12 2.12 4.24-4.24 1.41 1.41-5.64 5.66z"/>
            </svg>
          </div>
          <h3>作物推荐</h3>
          <p>根据您的土壤和气候条件，智能推荐最适合种植的作物品种，提高种植成功率。</p>
          <div class="feature-link">了解更多 →</div>
        </div>

        <div class="feature-card" @click="$router.push('/YuceHome/ProductionForecasts')">
          <div class="feature-icon">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#38a169">
              <path d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zM9 17H7v-7h2v7zm4 0h-2V7h2v10zm4 0h-2v-4h2v4z"/>
            </svg>
          </div>
          <h3>产量预测</h3>
          <p>基于历史数据和当前环境参数，预测作物产量，帮助您制定更好的生产计划。</p>
          <div class="feature-link">了解更多 →</div>
        </div>
      </div>
    </div>

    <!-- 使用步骤 -->
    <div class="steps-section">
      <h2 class="section-title">使用步骤</h2>
      <div class="steps-container">
        <div class="step">
          <div class="step-number">1</div>
          <h3>输入参数</h3>
          <p>输入您的土壤、气候和环境参数，如氮磷钾含量、温度、湿度、pH值等。</p>
        </div>

        <div class="step">
          <div class="step-number">2</div>
          <h3>AI分析</h3>
          <p>我们的AI模型会分析您的数据，与历史数据进行比对和学习。</p>
        </div>

        <div class="step">
          <div class="step-number">3</div>
          <h3>获取结果</h3>
          <p>系统会生成详细的分析报告，包括推荐作物或产量预测及科学建议。</p>
        </div>
      </div>
    </div>

    <!-- 底部引导 -->
    <div class="cta-section">
      <h2>开始您的智能农业之旅</h2>
      <p>使用我们的智能推荐和预测服务，提高您的农业生产效率</p>
      <button class="cta-button" @click="$router.push('/croprec')">开始体验</button>
    </div>

    <!-- 子路由显示区域 -->
    <router-view></router-view>
  </div>
</template>

<script>
import Navbar from '../components/Navbar.vue'

export default {
  name: 'YuceHome',
  components: {
    Navbar
  }
}
</script>

<style scoped>
.home-page {
  font-family: 'Arial', sans-serif;
  color: #2d3748;
  background-color: #f7fafc;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  overflow-y: auto;
  margin: 0;
  padding: 0;
  z-index: 0;
}

/* 顶部背景和标题 */
.hero-section {
  background-image: url('../img/Round-robin_diagram/true.jpg');
  background-size: cover;
  background-position: center;
  height: 400px;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  color: white;
  position: relative;
}

.hero-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
}

.hero-content {
  position: relative;
  z-index: 1;
  max-width: 800px;
  padding: 0 20px;
}
.hero-title {
  font-size: 3rem;
  margin-bottom: 1rem;
}
.hero-subtitle {
  font-size: 1.5rem;
  margin-bottom: 2rem;
}

.hero-content h1 {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.hero-content .subtitle {
  font-size: 1.5rem;
  margin-bottom: 2rem;
}

/* 核心功能 */
.section-title {
  text-align: center;
  font-size: 2rem;
  margin-bottom: 50px;
  color: #2f855a;
}

.features-section {
  max-width: 1200px;
  margin: 0 auto;
}

.features-container {
  display: flex;
  justify-content: center;
  gap: 90px;
}

.feature-card {
  flex: 1;
  max-width: 350px;
  padding: 30px;
  border-radius: 12px;
  background-color: white;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 1px solid #e2e8f0;
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(56, 161, 105, 0.1);
  border-color: #c6f6d5;
}

.feature-icon {
  margin-bottom: 1.5rem;
  height: 48px;
}

.feature-icon svg {
  width: 48px;
  height: 48px;
}

.feature-card h3 {
  font-size: 1.5rem;
  margin-bottom: 1rem;
  color: #2f855a;
}

.feature-card p {
  color: #4a5568;
  line-height: 1.6;
  margin-bottom: 1.5rem;
}

.feature-link {
  color: #38a169;
  font-weight: 500;
  transition: all 0.2s;
}

.feature-card:hover .feature-link {
  transform: translateX(5px);
}

/* 使用步骤 */
.steps-section {
  background-color: white;
  padding: 60px 20px;
  border-top: 1px solid #edf2f7;
  border-bottom: 1px solid #edf2f7;
}

.steps-container {
  display: flex;
  justify-content: center;
  gap: 30px;
  max-width: 1200px;
  margin: 0 auto;
}

.step {
  flex: 1;
  max-width: 350px;
}

.step-number {
  width: 60px;
  height: 60px;
  background-color: #38a169;
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.8rem;
  font-weight: bold;
  margin: 0 auto 20px;
}

.step h3 {
  font-size: 1.3rem;
  margin-bottom: 1rem;
  color: #2d3748;
}

.step p {
  color: #4a5568;
  line-height: 1.6;
}

/* 底部引导 */
.cta-section {
  background-color: #38a169;
  color: white;
  text-align: center;
  padding: 80px 20px;
}

.cta-section h2 {
  font-size: 2.2rem;
  margin-bottom: 1rem;
}

.cta-section p {
  font-size: 1.2rem;
  margin-bottom: 2rem;
  opacity: 0.9;
}

.cta-button {
  background-color: white;
  color: #38a169;
  border: none;
  padding: 15px 40px;
  font-size: 1.2rem;
  border-radius: 30px;
  cursor: pointer;
  transition: all 0.3s;
  font-weight: bold;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.cta-button:hover {
  transform: scale(1.05);
  box-shadow: 0 8px 15px rgba(0, 0, 0, 0.2);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .features-container, .steps-container {
    flex-direction: column;
    align-items: center;
  }

  .feature-card, .step {
    max-width: 100%;
    margin-bottom: 30px;
  }

  .hero-content h1 {
    font-size: 2rem;
  }

  .hero-content .subtitle {
    font-size: 1.2rem;
  }

  .hero-section {
    height: 400px;
  }
}
</style>