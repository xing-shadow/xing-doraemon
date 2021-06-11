import httpClient from '../utils/request';

export function LoginUser(data) {
    return httpClient({
        method:'post',
        url:"/api/v1/user/login",
        data:data
    })
}