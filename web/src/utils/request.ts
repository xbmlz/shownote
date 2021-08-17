import axios, {AxiosInstance, AxiosRequestConfig, AxiosResponse} from "axios";
import {ElMessage} from 'element-plus';
import Router from "@/router"

export interface ResponseData {
    code: number;
    data?: any;
    message: string;
}


// console.log('import.meta.env: ', import.meta.env);

// 创建 axios 实例
let service: AxiosInstance | any;
localStorage.baseURL = import.meta.env.MODE === "development" ? 'http://127.0.0.1:8000/' : window.location.origin + '/api/'
service = axios.create({
    baseURL: localStorage.baseURL, // api 的 base_url
    timeout: 50000 // 请求超时时间
});

// request 拦截器 axios 的一些配置
service.interceptors.request.use(
    (config: AxiosRequestConfig) => {
        if (localStorage.token) {
            config.headers['Authorization'] = localStorage.token
        }
        return config;
    },
    (error: any) => {
        // Do something with request error
        console.error("error:", error); // for debug
        Promise.reject(error);
    }
);

// respone 拦截器 axios 的一些配置
service.interceptors.response.use(
    (res: AxiosResponse) => {
        // Some example codes here:
        // code == 0: success
        if (res.status === 200) {
            const data: ResponseData = res.data
            if (data.code === 0) {
                return data;
            } else if (data.code === 4010 || data.code === 4011 || data.code === 4012) {
                ElMessage.error("Token失效，请重新登录");
                localStorage.removeItem('token')
                Router.push('/login')
            } else {
                ElMessage.error(data.message);
                return Promise.reject(new Error(res.data.message || "Error"));
            }
        } else {
            ElMessage.error("网络错误");
            return Promise.reject(new Error(res.data.message || "Error"));
        }
    },
    (error: any) => Promise.reject(error)
);

export default service;
