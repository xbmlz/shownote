import { createRouter, createWebHistory, RouteRecordRaw, NavigationGuard } from "vue-router";
import Layout from "@/layouts/index.vue";

// modules
import dashboard from "./modules/dashboard";

export const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "Home",
        component: Layout,
        redirect: "/dashboard"
    },
    dashboard
];

const router = createRouter({
    history: createWebHistory('/'),
    routes
});


export default router;