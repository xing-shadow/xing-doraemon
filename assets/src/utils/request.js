import axios from 'axios';

const httpClient = axios.create({
    // baseURL:window.location.origin,
    baseURL: "http://localhost:8080",
});


httpClient.interceptors.response.use(
    function (resp) {
        if (resp.status === 200) {
            if (resp.data.code === 302) {
                window.location.href = resp.data.msg;
            } else {
                return resp.data;
            }
        } else {
            Promise.reject("get http code:" + resp.status);
        }
    },
    function (error) {
        Promise.reject(error);
    }
)

export default httpClient; 