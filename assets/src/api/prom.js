import httpClient from "../utils/request";

export function GetPromList(params) {
    return httpClient({
        method:'get',
        url:"/api/v1/proms",
        params:params
    })
}

export function GetPromAllName() {
    return httpClient({
        method:'get',
        url:"/api/v1/prom/allName",
    })
}

export function GetProm(params) {
    return httpClient({
        method:'get',
        url:"/api/v1/prom",
        params:params
    })
}

export function AddProm(params) {
    return httpClient({
        method:'post',
        url:"/api/v1/prom/add",
        data:params
    })
}

export function UpdateProm(params) {
    return httpClient({
        method:'post',
        url:"/api/v1/prom/update",
        data:params
    })
}

export function DeleteProm(params) {
    return httpClient({
        method:'post',
        url:"/api/v1/prom/delete",
        data:params
    })
}