import {createRouter, createWebHashHistory, RouteRecordRaw} from "vue-router";

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
    history: createWebHashHistory(),
    routes,
});

export default router;
