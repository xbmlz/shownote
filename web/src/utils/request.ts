import {useRouter} from "vue-router";
import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import { ElMessage  } from 'element-plus';

const router = useRouter();
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
            } else if(data.code === 401) {
                router.push('/login')
            } else{
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
