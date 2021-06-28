import httpClient from "../utils/request";

export function GetPlanList(params) {
    return httpClient({
        method:'get',
        url:"/api/v1/plans",
        params:params
    })
}

export function GetPlan(params) {
    return httpClient({
        method:'get',
        url:"/api/v1/plan",
        params:params
    })
}

export function GetPlanAllName() {
    return httpClient({
        method:'get',
        url:"/api/v1/plan/allName",
    })
}

export function AddPlan(params) {
    return httpClient({
        method:'post',
        url:"/api/v1/plan/add",
        data:params
    })
}

export function UpdatePlan(params) {
    return httpClient({
        method:'post',
        url:"/api/v1/plan/update",
        data:params
    })
}


export function DeletePlan(params) {
    return httpClient({
        method:'delete',
        url:"/api/v1/plan/delete",
        data:params
    })
}




