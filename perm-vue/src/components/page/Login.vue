<template>
    <div class="login-wrap">
        <div class="ms-login">
            <div class="ms-title">后台管理系统</div>
            <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="0px" class="ms-content">
                <el-form-item prop="username">
                    <el-input v-model="ruleForm.username" placeholder="请输入登录账号">
                        <el-button slot="prepend" icon="el-icon-lx-people"></el-button>
                    </el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input type="password" placeholder="请输入密码" v-model="ruleForm.password"
                              @keyup.enter.native="submitForm('ruleForm')">
                        <el-button slot="prepend" icon="el-icon-lx-lock"></el-button>
                    </el-input>
                </el-form-item>
                <div class="login-btn">
                    <el-button type="primary" @click="submitForm('ruleForm')">登录</el-button>
                </div>
            </el-form>
        </div>
    </div>
</template>

<script>
    export default {
        data: function () {
            const isvalidateMobile = (rule, value, callback) => {
                var XLphoneReg = /^1(3[0-9]|4[4-9]|5[012356789]|66|7[0-8]|8[0-9]|9[89])[0-9]{8}$/; //手机号正则 
                if (value != null && value != "") {
                    if (!XLphoneReg.test(value)) {
                        callback(new Error("手机号格式错误"));
                    } else {
                        callback();
                    }
                } else {
                    callback();
                }
            };
            const isvalidateRegexn = (rule, value, callback) => {
                var XLphoneReg = /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,15}$/; //密码正则   由数字和字母组成，并且要同时含有数字和字母，且长度要在6-15位之间   
                if (value != null && value != "") {
                    if (!regexn(value)) {
                        callback(new Error('含有非法字符(只能输入字母、汉字)!'))
                    } else {
                        callback()
                    }
                } else {
                    callback();
                }
            };
            return {
                ruleForm: {
                    username: '',
                    password: ''
                },
                rules: {
                    username: [
                        {required: true, message: '请输入用登录账号', trigger: 'blur'},
                        // {validator: isvalidateMobile, trigger: "blur"}
                    ],
                    password: [
                        {required: true, message: '请输入密码', trigger: 'blur'},
                        // {validator: isvalidateRegexn, trigger: "blur"}
                    ]
                }
            }
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        let param = {
                            userAccount: this.ruleForm.username,
                            // UserPassword: this.$md5(this.ruleForm.password)
                            userPassword: this.ruleForm.password
                        };
                        this.$axios.defaults.headers['urlencoded'] = '-1';
                        this.$axios.post("/perm/user/login", param, {}
                        ).then(resp => {
                            let data = resp.data;
                            if (data && data.code == 200) {
                                localStorage.setItem('ms_username', data.data.userName);
                                window.sessionStorage.setItem('user', JSON.stringify(data.data));
                                this.$router.push('/');
                            } else {
                                this.$message.error(data.errorMessage);
                            }
                        });
                    }
                });
            }
        }
    }
</script>

<style scoped>
    .login-wrap {
        position: relative;
        width: 100%;
        height: 100%;
        background-image: url(../../assets/login-bg.jpg);
        background-size: 100%;
    }

    .ms-title {
        width: 100%;
        line-height: 50px;
        text-align: center;
        font-size: 20px;
        color: #fff;
        border-bottom: 1px solid #ddd;
    }

    .ms-login {
        position: absolute;
        left: 50%;
        top: 50%;
        width: 350px;
        margin: -190px 0 0 -175px;
        border-radius: 5px;
        background: rgba(255, 255, 255, 0.3);
        overflow: hidden;
    }

    .ms-content {
        padding: 30px 30px;
    }

    .login-btn {
        text-align: center;
    }

    .login-btn button {
        width: 100%;
        height: 36px;
        margin-bottom: 10px;
    }

    .login-tips {
        font-size: 12px;
        line-height: 30px;
        color: #fff;
    }

    .el-input-group {
        width: 100% !important;
    }
</style>