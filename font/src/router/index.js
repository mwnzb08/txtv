import { createRouter, createWebHistory } from 'vue-router'
import RouterMapping from './RouterMapping'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../components/Home')
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'redirect',
    redirect: '404'
  }
]
for (const map in RouterMapping) {
  const value = RouterMapping[map]
  const routerConfig = {
    path: '/' + map,
    name: value,
    component: () => import('../components/' + value)
  }
  routes.push(routerConfig)
}

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
