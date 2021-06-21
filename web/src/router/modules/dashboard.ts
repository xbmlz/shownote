import {RouteRecordRaw} from "vue-router";
import Layout from "@/layouts/index.vue";

export const dashboard: RouteRecordRaw = {
    path: "/dashboard",
    name: "Dashboard",
    meta: {title: 'Dashboard', icon: ''},
    component: Layout,
    redirect: "/dashboard/map",
    children: [
        {
            path: "map",
            name: "Map",
            component: () => import("@/views/dashboard/index.vue"),
            meta: {title: "Map", icon: ""}
        }
    ]
}

export default dashboard;