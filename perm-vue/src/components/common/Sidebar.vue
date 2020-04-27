<template>
    <div class="sidebar">
        <el-menu class="sidebar-el-menu" :default-active="onRoutes" :collapse="collapse" background-color="#324157"
                 text-color="#bfcbd9" active-text-color="#20a0ff" unique-opened router>
            <template v-for="item in items">
                <template v-if="item.subs">
                    <el-submenu :index="item.index" :key="item.index">
                        <template slot="title">
                            <i :class="item.icon"></i><span slot="title">{{ item.title }}</span>
                        </template>
                        <template v-for="subItem in item.subs">
                            <el-submenu v-if="subItem.subs" :index="subItem.index" :key="subItem.index">
                                <template slot="title">{{ subItem.title }}</template>
                                <el-menu-item v-for="(threeItem,i) in subItem.subs" :key="i" :index="threeItem.index">
                                    {{ threeItem.title }}
                                </el-menu-item>
                            </el-submenu>
                            <el-menu-item v-else :index="subItem.index" :key="subItem.index">
                                {{ subItem.title }}
                            </el-menu-item>
                        </template>
                    </el-submenu>
                </template>
                <template v-else>
                    <el-menu-item :index="item.index" :key="item.index">
                        <i :class="item.icon"></i><span slot="title">{{ item.title }}</span>
                    </el-menu-item>
                </template>
            </template>
        </el-menu>
    </div>
</template>

<script>
    import bus from '../common/bus';
    import {getUserId} from 'permissionManage/perm-server/utils/utils'

    export default {
        data() {
            return {
                onRoutes: '',
                menuList: [],
                collapse: true,
                items: []
            }
        },
        methods: {
            getMenuList() {
                let param = {
                    userId: getUserId()
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/resoure/findResourceTree", param, {}
                ).then(resp => {
                    let data = resp.data;
                    if (data && data.code == 200 && data.data != '') {
                        this.items = data.data;
                        this.onRoutes = this.items[0].index;
                    } else {
                        this.Common.GetSessionStorage(this, data.msg);
                    }
                });
            }
        },
        mounted() {
            this.getMenuList();
        },
        created() {
            // this.items = this.navMenus;
            // 通过 Event Bus 进行组件间通信，来折叠侧边栏
            bus.$on('collapse', msg => {
                this.collapse = msg;
            });
            // this.getMenuList();
        }
    }
</script>

<style scoped>
    .sidebar {
        display: block;
        position: absolute;
        left: 0;
        top: 70px;
        bottom: 0;
        overflow-y: scroll;
    }

    .sidebar::-webkit-scrollbar {
        width: 0;
    }

    .sidebar-el-menu:not(.el-menu--collapse) {
        width: 220px;
    }

    .sidebar > ul {
        height: 100%;
    }
</style>
