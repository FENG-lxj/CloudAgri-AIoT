<template>
  <div class="about-page">
    <Navbar />
    <div class="about-container">
      <div class="page-header">
        <h1>关于我们</h1>
        <p class="subtitle">了解农业智能预测平台背后的理念与技术</p>
      </div>

      <div class="content-section">
        <div class="section project-intro">
          <h2>项目简介</h2>
          <p>农业智能预测平台是一款基于人工智能和大数据技术的现代农业辅助工具，旨在帮助农民和农业工作者做出更科学、更精准的种植决策。</p>
          <p>我们的系统融合了先进的机器学习算法、多年积累的农业数据和专业的农学知识，通过分析土壤条件、气候环境和作物特性，为用户提供个性化的作物推荐和产量预测服务。</p>
        </div>

        <div class="section vision">
          <h2>我们的愿景</h2>
          <p>让每一位农业从业者都能够享受科技带来的便利，通过数据驱动决策，提高作物产量，降低资源浪费，实现农业生产的可持续发展。</p>
        </div>

        <div class="section core-values">
          <h2>核心价值</h2>
          <ul class="value-list">
            <li>
              <span class="value-icon">●</span>
              <span class="value-text">科学决策-基于数据和模型，而非经验和猜测</span>
            </li>
            <li>
              <span class="value-icon">●</span>
              <span class="value-text">精准农业-优化资源投入，提高产出效率</span>
            </li>
            <li>
              <span class="value-icon">●</span>
              <span class="value-text">可持续发展-减少资源浪费，保护环境生态</span>
            </li>
            <li>
              <span class="value-icon">●</span>
              <span class="value-text">知识共享-将农业科技成果转化为实用工具</span>
            </li>
          </ul>
        </div>

        <div class="section tech-support">
          <h2>技术支持</h2>
          <p>我们的系统采用多种机器学习算法，包括随机森林、梯度提升树和深度神经网络，对农业数据进行分析和预测，为用户提供精准的决策建议。</p>
          <div class="tech-tags">
            <span class="tech-tag">机器学习</span>
            <span class="tech-tag">数据分析</span>
            <span class="tech-tag">云计算</span>
            <span class="tech-tag">实时监控</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import Navbar from "../components/Navbar.vue";
</script>

<style scoped>
.about-page {
  font-family: 'Arial', sans-serif;
  background-color: #f5f5f5;
  min-height: 100vh;
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

.about-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.page-header {
  text-align: center;
  margin-bottom: 40px;
  padding-top: 20px;
}

.page-header h1 {
  font-size: 36px;
  color: #4CAF50;
  margin-bottom: 10px;
}

.subtitle {
  font-size: 18px;
  color: #666;
}

.content-section {
  background-color: white;
  border-radius: 8px;
  padding: 30px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
}

.section {
  margin-bottom: 40px;
}

.section h2 {
  font-size: 24px;
  color: #4CAF50;
  margin-bottom: 20px;
  padding-bottom: 10px;
  border-bottom: 2px solid #e0e0e0;
}

.section p {
  font-size: 16px;
  line-height: 1.8;
  color: #333;
  margin-bottom: 15px;
}

.value-list {
  list-style: none;
  padding: 0;
}

.value-list li {
  margin-bottom: 15px;
  display: flex;
  align-items: center;
}

.value-icon {
  color: #4CAF50;
  margin-right: 10px;
  font-size: 20px;
}

.value-text {
  color: #333;
}

.tech-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 20px;
}

.tech-tag {
  background-color: #e8f5e9;
  color: #2E7D32;
  padding: 5px 15px;
  border-radius: 20px;
  font-size: 14px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .about-container {
    padding: 10px;
  }

  .page-header h1 {
    font-size: 28px;
  }

  .subtitle {
    font-size: 16px;
  }

  .content-section {
    padding: 20px;
  }

  .section h2 {
    font-size: 20px;
  }

  .tech-tags {
    justify-content: center;
  }
}
</style>