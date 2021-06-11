import httpClient from "../utils/request";

export function GetAlertList(param) {
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