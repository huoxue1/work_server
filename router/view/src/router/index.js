import { createRouter, createWebHashHistory } from 'vue-router'


const routes = [
  {
    path: '/upload',
    name: 'Upload',
    component: ()=> import('../views/Upload')
  },{
    path: '/token_manager',
    name: 'token_Manager',
    component: ()=> import('../views/TokenManager')
  },
  {
    path: '/work_manager',
    name: 'work_Manager',
    component: ()=> import('../views/WorkerManager')
  },
  // {
  //   path: '/about',
  //   name: 'About',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
