export default [
    {
        path: '/departments',
        component: resolve => require(['../components/page/sys-user/department/departments.vue'], resolve),
        meta: {title: '部门管理'}
    },
    {
        path: '/new-user',
        component: resolve => require(['../components/page/sys-user/user/userList.vue'], resolve),
        meta: {title: '内部用户'}
    },
    {
        path: '/role-list',
        component: resolve => require(['../components/page/sys-user/role/role-list.vue'], resolve),
        meta: {title: '角色管理'}
    },
    {
        path: '/role-edit',
        component: resolve => require(['../components/page/sys-user/role/role-edit.vue'], resolve),
        meta: {title: '角色编辑'}
    },
    {
        path: '/permission-list',
        component: resolve => require(['../components/page/sys-user/permission/permission-list.vue'], resolve),
        meta: {title: '资源管理'}
    },
    {
        path: '/page-resource-list',
        component: resolve => require(['../components/page/sys-user/permission/page-resource-list.vue'], resolve),
        meta: {title: '页面元素管理'}
    }
]


