import { createRouter,createWebHistory } from "vue-router"

// import $store from '@/store'

const routes = [
  {
      path: '/:catchAll(.*)',
      name: '404',
      component: ()=>import('../views/404.vue')
  },
  {
    path: '/',
    name: 'home',
    component: ()=>import('../views/Home.vue')
  },
  {
    path: '/login',
    name: 'login',
    component: ()=>import('../views/Login.vue')
  },
  {
    path: '/register',
    name: 'register',
    component: ()=>import('../views/Register.vue')
  },
  {
    path: '/repass',
    name: 'repass',
    component: ()=>import('../views/Repass.vue')
  },
  {
    path: '/profile',
    name: 'profile',
    component: ()=>import('../views/Profile.vue')
  },
]

const router = createRouter({
    history: createWebHistory("/"),
    routes
})
// router.beforeEach((to,from,next) => {
//   let token = localStorage.getItem("token")
//   // 有token时访问
//   if(token){
//     // 不应该访问登录页面和注册页面
//     if(to.path === "/login"){
//       // router.push("/")
//       // $store.dispatch("getToken")
      
      
//       // if($store.state.tokenInfo){
//       //   router.push("/")
//       // }
//       // 应该验证token是否
//       // 如果登录了跳转到首页
//     }
//     next()
//   }else{
//     // 主页和登录界面不需要token
//     if(to.path === "/" || to.path === "/login"){
//       next()
//     }else{
//       // 其他页面如果没有token则跳转到登录页面
//       next('/login')
//     }
//   }
// })

export default router