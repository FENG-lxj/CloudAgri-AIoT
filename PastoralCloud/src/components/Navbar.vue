<template>
  <div class="navbar">
    <div class="logo">农业智能系统</div>
    <div class="nav-links">
      <router-link to="/YuceHome" active-class="active" exact-active-class="active">首页</router-link>
      <router-link to="/YuceHome/CropRecommendations" active-class="active" exact-active-class="active">作物推荐</router-link>
      <router-link to="/YuceHome/ProductionForecasts" active-class="active" exact-active-class="active">产量预测</router-link>
      <router-link to="/YuceHome/AboutUs" active-class="active" exact-active-class="active">关于我们</router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Navbar',
};
</script>

<style scoped>
.navbar {
  display: flex;
  align-items: center;
  background-color: #ffffff;
  padding: 15px 30px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
}

.logo {
  font-size: 20px;
  font-weight: bold;
  color: #4CAF50;
  margin-right: 40px;
}

.nav-links {
  display: flex;
  justify-content: center;
  flex-grow: 1;
  gap: 20px;
}

.navbar a {
  color: #555;
  text-decoration: none;
  padding: 8px 15px;
  border-radius: 20px;
  transition: all 0.3s ease;
  font-size: 20px;
}

.navbar a:hover {
  background-color: #f0f8f0;
  color: #2E7D32;
}

.navbar a.active {
  color: #2E7D32;
  font-weight: bold;
  background-color: #e8f5e9;
}
</style>