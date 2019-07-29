import Vue from 'vue';
import Router from 'vue-router';
import sys_user_router from './sys-user-router'

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: '/',
            redirect: '/dashboard'
        },
        {
            path: '/',
            component: resolve => require(['../components/common/Home.vue'], resolve),
            children: [
                {
                    path: '/dashboard',
                    component: resolve => require(['../components/page/demo/Dashboard.vue'], resolve),
                    meta: {title: '系统首页'}
                },
                ...sys_user_router
            ]
        },
        {
            path: '/login',
            component: resolve => require(['../components/page/Login.vue'], resolve)
        },
        {
            path: '*',
            redirect: '/404'
        }
    ]
})
