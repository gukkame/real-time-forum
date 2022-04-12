import { createRouter, createWebHistory } from 'vue-router'

import SignUp from '../views/SignUp.vue'
import HomeView from '../views/HomeView.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/home',
      name: 'home',
      component: HomeView
    },
    {
      path: '/',
      name: 'default',
      component: HomeView
    },
    {
      path: '/signup',
      name: 'signup',
      component: SignUp
    },
    {
      path: '/profile',
      name: 'profile',
      component: () => import('../views/profile.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LogIn.vue')
    },
    {
      path: '/logout',
      name: 'logout',
      component: () => import('../views/Logout.vue')
    },
    {
      path: '/home/nft',
      name: 'nft',
      component: () => import('../views/nft.vue')
    },
    {
      path: '/home/nftgames',
      name: 'nftgames',
      component: () => import('../views/nftgames.vue')
    },
    {
      path: '/home/crypto',
      name: 'crypto',
      component: () => import('../views/crypto.vue')
    },
    {
      path: '/home/blockchain',
      name: 'blockchain',
      component: () => import('../views/blockchain.vue')
    },
    {
      path: '/home/web3',
      name: 'web3',
      component: () => import('../views/web3.vue')
    }
  ]
})

export default router
