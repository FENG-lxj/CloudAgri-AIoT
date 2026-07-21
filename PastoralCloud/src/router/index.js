import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import about from '../views/AboutView'
import login from '../views/login.vue'
import Agro_ecological from '../Agro/Agro_ecological'
import Storage_Main from '../Storage/Storage_Main'
import AI_Vision_Main from '../AI_Vision/AI_Vision_Main'
import SuperAdmin from '../SuperAdmin/SuperAdmin'
import Device_Management_Page from '../Device_Management_Page/Device_Management'
import Agricultural_Equipment from '../Device_Management_Page/Page/Agricultural_Equipment'
import Storage_Devices from '../Device_Management_Page/Page/Storage_Devices'
import AI_Devices from '../Device_Management_Page/Page/AI_Devices'
import error from '../router/error'
import AboutUs from "@/views/AboutUs.vue";
import CropRecommendations from "@/views/CropRecommendations.vue";
import ProductionForecasts from "@/views/ProductionForecasts.vue";
import YuceHome from "@/views/YuceHome.vue";
Vue.use(VueRouter)

const routes = [
  {
    path: '/home',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/',
    redirect: '/login', // 添加重定向规则，将根路径重定向到 /login
  },

  {
    path: '/login',
    name: 'login',
    component: login
  },
  {
    path: '/SuperAdmin',
    name: 'SuperAdmin',
    component: SuperAdmin
  },
  {
    path: '/Device_Management_Page',
    name: 'Device_Management_Page',
    component: Device_Management_Page,
    children: [
      {
        path: '/',
        name: '/',
        component: Agricultural_Equipment,
      },
      {
        path: 'Storage_Devices',
        name: 'Storage_Devices',
        component: Storage_Devices,
      },
      {
        path: 'AI_Devices',
        name: 'AI_Devices',
        component: AI_Devices
      }
    ],
  },
  {
    path: '/about',
    name: 'about',
    component: about
  },
  {
    path: '/Agro_ecological',
    name: 'Agro_ecological',
    component: Agro_ecological
  },
  {
    path: '/Storage_Main',
    name: 'Storage_Main',
    component: Storage_Main
  },
  {
    path: '/AI_Vision_Main',
    name: 'AI_Vision_Main',
    component: AI_Vision_Main
  },
  {
    path: '/error',
    name: 'error',
    component: error
  },

  {
    path: '/YuceHome',
    name: 'YuceHome',
    component: YuceHome,
    children: [
      {
        path: 'ProductionForecasts',
        name: 'ProductionForecasts',
        component: ProductionForecasts
      },
      {
        path: 'CropRecommendations',
        name: 'CropRecommendations',
        component: CropRecommendations
      },
      {
        path: 'AboutUs',
        name: 'AboutUs',
        component: AboutUs
      }
    ]
  }

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

