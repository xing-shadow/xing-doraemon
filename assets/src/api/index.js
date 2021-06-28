import httpClient from '../utils/request';

export function LoginUser(data) {
    return httpClient({
        method:'post',
        url:"/api/v1/user/login",
        data:data
    })
}

export function ListUser(data) {
    return httpClient({
        method:'get',
        url:"/api/v1/user/list",
        params:data
    })
}

export function CreateUser(data) {
    return httpClient({
        method:'post',
        url:"/api/v1/user/create",
        params:data
    })
}


export function UpdateUser(data) {
    return httpClient({
        method:'post',
        url:"/api/v1/user/update",
        params:data
    })
}

export function DeleteUser(data) {
    return httpClient({
        method:'post',
        url:"/api/v1/user/delete",
        params:data
    })
}
