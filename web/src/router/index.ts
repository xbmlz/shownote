import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        redirect:'/notes'
    },
    {
        path: '/notes',
        name: 'Notes',
        component: () => import('@/views/notes/index.vue')
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/login/index.vue')
    }
]

const router = createRouter({
    // history: createWebHistory(process.env.BASE_URL),
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

export default router;