import httpClient from '../utils/request';

// /api/v1/user
export function LoginUser(data) {
    return httpClient({
        method:'post',
        url:"/api/v1/user/login",
        data:data
    })
}

// /api/v1/plans

export function GetPormList(params) {
    return httpClient({
        method:'get',
        url:"/api/v1/prom",
        params:params
    })
}

export function DeleProm(params) {
    return httpClient({
        method:'delete',
        url:"/api/v1/prom",
        data:params
    })
}

export function GetProm(params) {
    return httpClient({
        method:'get',
        url:"/api/v1/promId",
        params:params
    })
}

export function UpdataProm(params) {
    return httpClient({
        method:'put',
        url:"/api/v1/prom",
        data:params
    })
}

export function AddProm(params) {
    return httpClient({
        method:'post',
        url:"/api/v1/prom",
        data:params
    })
}

export function AddPlan(params) {
    return httpClient({
        method:'post',
        url:"/api/v1/plan",
        data:params
    })
}

export function GetPlan(params) {
    return httpClient({
        method:'get',
        url:"/api/v1/plan",
        params:params
    })
}

export function DeletePlan(params) {
    return httpClient({
        method:'delete',
        url:"/api/v1/plan",
        data:params
    })
}

export function GetPromAllName() {
    return httpClient({
        method:'get',
        url:"/api/v1/prom/allName",
    })
}

export function GetPlanAllName() {
    return httpClient({
        method:'get',
        url:"/api/v1/plan/allName",
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
        url:"/api/v1/rule",
        data:param,
    })
}

export function GetRuleId(param) {
    return httpClient({
        method:'get',
        url:"/api/v1/ruleId",
        params:param,
    })
}

export function EditRule(param) {
    return httpClient({
        method:'put',
        url:"/api/v1/rule",
        data:param,
    })
}

export function DeleteRule(param) {
    return httpClient({
        method:'delete',
        url:"/api/v1/rule",
        data:param,
    })
}

export function GetAlerts(param) {
    return httpClient({
        method:'get',
        url:"/api/v1/alerts",
        params:param,
    })
}

export function ConfirmAlert(param) {
    return httpClient({
        method:'post',
        url:"/api/v1/alerts/confirm",
        data:param,
    })
}