import httpClient from "../utils/request";

export function GetRuleList(param) {
    return httpClient({
        method:'get',
        url:"/api/v1/rules",
        params:param,
    })
}

export function GetRule(param) {
    return httpClient({
        method:'get',
        url:"/api/v1/rule",
        params:param,
    })
}

export function AddRule(param) {
    return httpClient({
        method:'post',
        url:"/api/v1/rule/add",
        data:param,
    })
}

export function EditRule(param) {
    return httpClient({
        method:'post',
        url:"/api/v1/rule/update",
        data:param,
    })
}

export function DeleteRule(param) {
    return httpClient({
        method:'post',
        url:"/api/v1/rule/delete",
        data:param,
    })
}
