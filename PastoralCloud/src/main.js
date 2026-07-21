import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import axios from 'axios';

Vue.config.productionTip = false
Vue.use(ElementUI);

import { Buffer } from 'buffer';
window.Buffer = Buffer;

import process from 'process';
window.process = process;



const whiteList = ['/login', 'error'];

router.beforeEach((to, from, next) => {
  // 检查当前路由是否在白名单中
  if (whiteList.includes(to.path)) {
    // 如果在白名单中，直接允许访问
    next();
  } else {
    const token = store.state.phoneCaptcha;
    // 检查 token 是否为空
    if (token) {
      // 如果 token 不为空，则允许路由跳转
      next();
    } else {
      // 如果 token 为空，禁止跳转，重定向到登录页面
      console.error('token为空，禁止跳转');
      next('/login');
    }
  }
});




new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
