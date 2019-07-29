import axios from 'axios'
// import Qs from 'qs'
import {Message} from 'element-ui'
import {fmtToken} from "./utils";

/**
 const postData=JSON.stringify(this.formCustomer);
 'Content-Type':'application/json'}

 const postData=Qs.stringify(this.formCustomer);//过滤成？&=格式
 'Content-Type':'application/xxxx-form'}
 */

var axiosInstance = axios.create({
    baseURL: 'http://localhost:8099',
    // config里面有这个transformRquest，这个选项会在发送参数前进行处理。
    // 这时候我们通过Qs.stringify转换为表单查询参数
    transformRequest: [function (data) {
        // data = Qs.stringify(data)
        data = JSON.stringify(data); // 转换json
        return data
    }],
    // 设置Content-Type
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
        // 'Authorization': 'bearer ' + JSON.parse(window.sessionStorage.getItem('user')).token
    }
});

// 请求拦截
axiosInstance.interceptors.request.use(config => {
    config.headers['Authorization'] = fmtToken();
    // config.headers.Authorization = 'bearer ' + token
    return config
}, err => {
    Message.error({message: '请求超时!', showClose: true});
    return Promise.resolve(err)
});
// response拦截
axiosInstance.interceptors.response.use(data => {
    // if (data.status && data.status === 200 && data.data.code != 200) {
    //   console.log(data)
    //   Message.error({message: data.data.msg})
    //   return
    // }
    return data
}, err => {
    if (err.response) {
        var status = err.response.status;
        var msg = err.response.data.msg;
        switch (status) {
            case 400:
                Message.error({message: (msg != null && msg !== '') ? msg : '操作失败,参数不对!', showClose: true});
                break;
            case 401:
                Message.error({message: (msg != null && msg !== '') ? msg : '未认证或认证过期,请重新登录!', showClose: true});
                break;
            case 404:
                Message.error({message: '请求的资源不存在!', showClose: true});
                break;
            case 403:
                Message.error({message: '权限不足,请联系管理员!', showClose: true});
                break;
            case 500:
                Message.error({message: '服务器被吃了 ⊙﹏⊙...', showClose: true});
                break;
            case 504:
                Message.error({message: '网关连接超时!', showClose: true});
                break;
            default:
                Message.error({message: '未知错误!', showClose: true})
        }
    }
    return Promise.resolve(err)
});

// ---------- rest请求封装 -----------
let base = 'http://localhost:8088'
// let base = '/api'
// let base = ''
export const ftmURL = function url(url) {
    return `${base}${url}`
    // return base + url
};

export const postRequest = function (url, params) {
    return axiosInstance({
        method: 'post',
        url: ftmURL(url),
        data: params,
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    })
};

export const uploadFileRequest = function (url, params) {
    return axiosInstance({
        method: 'post',
        url: ftmURL(url),
        data: params,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    })
};

export const putRequest = function (url, params) {
    return axiosInstance({
        method: 'put',
        url: ftmURL(url),
        data: params
    })
};

export const deleteRequest = function (url) {
    return axiosInstance({
        method: 'delete',
        url: ftmURL(url)
    })
};

export const getRequest = function (url) {
    return axiosInstance({
        method: 'get',
        url: ftmURL(url)
    })
};
