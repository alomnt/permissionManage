<template>
    <div class="header">
        <!-- 折叠按钮 -->
        <div class="collapse-btn" @click="collapseChage">
            <i class="el-icon-menu"></i>
        </div>
        <div class="logo">后台管理系统</div>
        <div class="header-right">
            <div class="header-user-con">
                <!-- 全屏显示 -->
                <div class="btn-fullscreen" @click="handleFullScreen">
                    <el-tooltip effect="dark" :content="fullscreen?`取消全屏`:`全屏`" placement="bottom">
                        <i class="el-icon-rank"></i>
                    </el-tooltip>
                </div>
                <!-- 消息中心 -->
                <div class="btn-bell">
                    <el-tooltip effect="dark" :content="message?`有${message}条未读消息`:`消息中心`" placement="bottom">
                        <router-link to="/tabs">
                            <i class="el-icon-bell"></i>
                        </router-link>
                    </el-tooltip>
                    <span class="btn-bell-badge" v-if="message"></span>
                </div>
                <!-- 用户头像 -->
                <div class="user-avator"><img src="static/img/img.jpg"></div>
                <!-- 用户名下拉菜单 -->
                <el-dropdown class="user-name" trigger="click" @command="handleCommand">
                    <span class="el-dropdown-link">
                        {{username}} <i class="el-icon-caret-bottom"></i>
                    </span>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item divided command="updatePassWord">修改密码</el-dropdown-item>
                        <el-dropdown-item divided command="loginout">退出登录</el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
            </div>
        </div>
        <el-dialog title="修改密码" :visible.sync='display' width="40%" :show-close="false">
            <el-form :model="permissionForm" ref="permissionForm" label-width="100px">
                <el-row>
                    <el-col :span="20">
                        <el-form-item label="原密码:" prop="icon">
                            <el-input v-model="permissionForm.oldPassWord" placeholder="原密码" type="password"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="20">
                        <el-form-item label="新密码:" prop="icon">
                            <el-input v-model="permissionForm.newPassWord" placeholder="新密码" type="password" @blur.native.capture="oldKONew"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="20">
                        <el-form-item label="再次输入密码:" prop="icon">
                            <el-input v-model="permissionForm.new2PassWord" placeholder="再次输入新密码" type="password" @blur.native.capture="newKONew"></el-input>
                        </el-form-item>
                    </el-col>
                </el-row>
                <div style="color:red;">* 密码至少是6位包含数字和英文</div>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="goBanck('permissionForm')">取 消</el-button>
                <el-button type="primary" @click="saveEdit()">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>
<script>
    import bus from '../common/bus';

    export default {
        data() {
            return {
                collapse: true,
                fullscreen: false,
                name: 'linxin',
                message: 2,
                display: false,
                permissionForm: {
                    oldPassWord: '',
                    new2PassWord: '',
                    newPassWord: '',
                },
            }
        },
        computed: {
            username() {
                let username = localStorage.getItem('ms_username');
                return username ? username : this.name;
            }
        },
        methods: {
            goBanck(PermissionForm) {
                this.$refs[PermissionForm].resetFields();
                this.display = false;
            },
            saveEdit() {
                this.$axios.defaults.headers['urlencoded'] = '1';
                let param = {
                    'password': this.permissionForm.oldPassWord,
                    'newPass': this.$md5(this.permissionForm.newPassWord),
                };
                this.$axios.post("/perm/user/modifyPass", param, {headers: {'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'}}
                ).then(resp => {
                    let data = resp.data;
                });
            },
            //新老比较
            oldKONew() {
                if (this.permissionForm.oldPassWord == this.permissionForm.newPassWord) {
                    this.$message({
                        showClose: true,
                        message: '新密码不能和原密码一样'
                    });
                }
            },
            //新新比较
            newKONew() {
                if (this.permissionForm.new2PassWord != this.permissionForm.newPassWord) {
                    this.$message({
                        showClose: true,
                        message: '两次密码输入不一致'
                    });
                }
            },
            // 用户名下拉菜单选择事件
            handleCommand(command) {
                if (command == 'loginout') {
                    localStorage.removeItem('ms_username');
                    this.$router.push('/login');
                    this.$axios.defaults.headers['urlencoded'] = '1';
                    this.$axios.post("/perm/user/logout", {}, {}
                    ).then(resp => {
                        let data = resp.data;
                    });
                } else if (command == 'updatePassWord') {
                    this.display = true;
                }
            },
            // 侧边栏折叠
            collapseChage() {
                this.collapse = !this.collapse;
                bus.$emit('collapse', this.collapse);
            },
            // 全屏事件
            handleFullScreen() {
                let element = document.documentElement;
                if (this.fullscreen) {
                    if (document.exitFullscreen) {
                        document.exitFullscreen();
                    } else if (document.webkitCancelFullScreen) {
                        document.webkitCancelFullScreen();
                    } else if (document.mozCancelFullScreen) {
                        document.mozCancelFullScreen();
                    } else if (document.msExitFullscreen) {
                        document.msExitFullscreen();
                    }
                } else {
                    if (element.requestFullscreen) {
                        element.requestFullscreen();
                    } else if (element.webkitRequestFullScreen) {
                        element.webkitRequestFullScreen();
                    } else if (element.mozRequestFullScreen) {
                        element.mozRequestFullScreen();
                    } else if (element.msRequestFullscreen) {
                        // IE11
                        element.msRequestFullscreen();
                    }
                }
                this.fullscreen = !this.fullscreen;
            }
        },
        mounted() {
            if (document.body.clientWidth < 1500) {
                this.collapseChage();
            }
        }
    }
</script>
<style scoped>
    .header {
        position: relative;
        box-sizing: border-box;
        width: 100%;
        height: 70px;
        font-size: 22px;
        color: #fff;
    }

    .collapse-btn {
        float: left;
        padding: 0 21px;
        cursor: pointer;
        line-height: 70px;
    }

    .header .logo {
        float: left;
        width: 250px;
        line-height: 70px;
    }

    .header-right {
        float: right;
        padding-right: 50px;
    }

    .header-user-con {
        display: flex;
        height: 70px;
        align-items: center;
    }

    .btn-fullscreen {
        transform: rotate(45deg);
        margin-right: 5px;
        font-size: 24px;
    }

    .btn-bell, .btn-fullscreen {
        position: relative;
        width: 30px;
        height: 30px;
        text-align: center;
        border-radius: 15px;
        cursor: pointer;
    }

    .btn-bell-badge {
        position: absolute;
        right: 0;
        top: -2px;
        width: 8px;
        height: 8px;
        border-radius: 4px;
        background: #f56c6c;
        color: #fff;
    }

    .btn-bell .el-icon-bell {
        color: #fff;
    }

    .user-name {
        margin-left: 10px;
    }

    .user-avator {
        margin-left: 20px;
    }

    .user-avator img {
        display: block;
        width: 40px;
        height: 40px;
        border-radius: 50%;
    }

    .el-dropdown-link {
        color: #fff;
        cursor: pointer;
    }

    .el-dropdown-menu__item {
        text-align: center;
    }
</style>
