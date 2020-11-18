import axios from 'axios';
import {GetToken,DeleteToken} from '@/utils/Token';
const httpClient = axios.create({
    baseURL:window.location.origin,
});


httpClient.interceptors.request.use(
    function(config) {
        config.headers.Authorization = GetToken();
        return config;
    },
    function(error) {
        Promise.reject(error);
    }
)

httpClient.interceptors.response.use(
    function(resp) {
        if (resp.status === 302) {
            DeleteToken();
        }
        if (resp.data) {
            return resp.data
        }else{
            Promise.reject("请求失败")
        }
    },
    function(error) {
       Promise.reject(error)
    }
)

export default httpClient; 