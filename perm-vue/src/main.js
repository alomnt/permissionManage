import Vue from 'vue';
import App from './App';
import router from './router';
import axios from 'axios';
import ElementUI from 'element-ui';
import './assets/reset.css';
import 'element-ui/lib/theme-chalk/index.css'; // 默认主题
// import '../static/css/theme-green/index.css';       // 浅绿色主题
import 'jquery';
import '../static/css/icon.css';
import "babel-polyfill";
import md5 from 'js-md5';
import Common from './assets/util';
import common from './assets/common.js';

import {fmtToken} from './utils/utils'

Vue.use(ElementUI, {size: 'small'});
Vue.prototype.$axios = axios;
Vue.config.productionTip = false;
Vue.prototype.$md5 = md5;
Vue.prototype.Common = Common;
Vue.prototype.$Common = common;

axios.interceptors.request.use(request => {
    request.headers.Authorization = fmtToken();
    let type = axios.defaults.headers['urlencoded'];
    let requestData = [];
    //上传
    if (type > -1) {
        request.headers['Content-Type'] = 'application/x-www-form-urlencoded; charset=UTF-8';
        let _data = '';
        if (request.data) {
            for (let key in request.data) {
                requestData.push(request.data[key]);
                _data += `${key}=${request.data[key]}`;
                Object.keys(request.data).pop() !== key && (_data += '&');
            }
            request.data = _data;
        }
        return request;
    } else {
        request.headers['Content-Type'] = 'application/json;charset=UTF-8';
        if (request.method != 'get') {
        }
        return request;
    }
});

//使用钩子函数对路由进行权限跳转
router.beforeEach((to, from, next) => {
    const role = localStorage.getItem('ms_username');
    if (!role && to.path !== '/login') {
        next('/login');
    } else if (to.meta.permission) {
        // 如果是管理员权限则可进入，这里只是简单的模拟管理员权限而已
        role === 'admin' ? next() : next('/403');
    } else {
        // 简单的判断IE10及以下不进入富文本编辑器，该组件不兼容
        if (navigator.userAgent.indexOf('MSIE') > -1 && to.path === '/editor') {
            Vue.prototype.$alert('vue-quill-editor组件不兼容IE10及以下浏览器，请使用更高版本的浏览器查看', '浏览器不兼容通知', {
                confirmButtonText: '确定'
            });
        } else {
            next();
        }
    }
});

new Vue({
    router,
    render: h => h(App)
}).$mount('#app');