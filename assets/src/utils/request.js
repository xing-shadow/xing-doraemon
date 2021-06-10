import axios from 'axios';

const httpClient = axios.create({
    // baseURL:window.location.origin,
    baseURL: "http://localhost:8080",
});


httpClient.interceptors.response.use(
    function (resp) {
        if (resp.status === 200) {
            return resp.data
        }
    },
    function (error) {
        Promise.reject(error)
    }
)

export default httpClient; 