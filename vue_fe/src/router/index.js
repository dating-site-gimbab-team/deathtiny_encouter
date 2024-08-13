import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Profile from '../views/Profile.vue'
import Matches from '../views/Matches.vue'
import Chat from '../views/Chat.vue'
import NotFound from '../views/NotFound.vue'

const routes = [
    { path: '/', component: Home },
    { path: '/profile', component: Profile },
    { path: '/matches', component: Matches },
    { path: '/chat', component: Chat },
    { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
