<template>
    <div class="table">
        <div class="container">
            <div class="handle-box">
                <el-row>
                    <el-col :span="8">
                        <div class="demo-input-suffix">编号:</div>
                        <el-input style="width:55%;" placeholder="员工编号" size="small"
                                  v-model="search.userNumber"></el-input>
                    </el-col>
                    <el-col :span="8">
                        <div class="demo-input-suffix">姓名:</div>
                        <el-input style="width:55%;" placeholder="请输入姓名" size="small"
                                  v-model="search.userName"></el-input>
                    </el-col>
                    <el-col :span="8">
                        <div class="demo-input-suffix">所属部门:</div>
                        <el-select style="width:55%;" placeholder="请选择所属部门" v-model="search.departmentId">
                            <el-option v-for="item in departmentIdList" :key="item.id" :label="item.depName"
                                       :value="item.id"></el-option>
                        </el-select>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="8">
                        <div class="demo-input-suffix">电话:</div>
                        <el-input style="width:55%;" placeholder="请输入内容" size="small"
                                  v-model="search.mobile"></el-input>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="4" :offset="10">
                        <el-button type="primary" icon="search" @click="searchBtn">搜 索</el-button>
                        <el-button icon="search" @click="resetBtn">重 置</el-button>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="24">
                        <el-button type="primary" icon="search" @click="newUsersBtn">创建账户</el-button>
                    </el-col>
                </el-row>
            </div>
            <el-table :data="tableData" border class="table" ref="multipleTable"
                      @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="40"></el-table-column>
                <el-table-column prop="userNumber" label="编号" width="150"></el-table-column>
                <el-table-column prop="userName" label="姓名" width="120"></el-table-column>
                <el-table-column prop="sex" label="性别" width="120"></el-table-column>
                <el-table-column prop="userAccount" label="账号" width="120"></el-table-column>
                <el-table-column prop="departmentName" label="所属部门"></el-table-column>
                <el-table-column prop="mobile" label="电话"></el-table-column>
                <el-table-column prop="email" label="邮箱"></el-table-column>
                <el-table-column fixed="right" label="操作" width="110" align="center">
                    <template slot-scope="scope">
                        <el-dropdown @command="handleCommand">
                            <el-button type="primary">
                                操作<i class="el-icon-arrow-down el-icon--right"></i>
                            </el-button>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item :command="{name:'handlePermissions',index:scope.$index}">查看权限
                                </el-dropdown-item>
                                <el-dropdown-item :command="{name:'handleDAllot',index:scope.$index}">分配角色
                                </el-dropdown-item>
                                <el-dropdown-item :command="{name:'handleReset',index:scope.$index}">重置密码
                                </el-dropdown-item>
                                <el-dropdown-item :command="{name:'editUser',index:scope.$index}">用户编辑
                                </el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination background @current-change="handleCurrentChange" layout="prev, pager, next"
                               :total="Total"></el-pagination>
            </div>
        </div>
        <!-- 新增弹出框 -->
        <el-dialog :title="setTitle" :visible.sync="newVisible" width="37%" :show-close="false"
                   :close-on-press-escape="false" :close-on-click-modal="false" @close="handleDialogClose('userForm')">
            <el-form :model="userForm" :rules="rules" ref="userForm" label-width="100px">
                <el-form-item label="所属部门" prop="departmentId">
                    <el-select placeholder="请选择所属部门" v-model="userForm.departmentId">
                        <el-option v-for="item in departmentIdList" :key="item.id" :label="item.depName"
                                   :value="item.id"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="姓名" prop="userName">
                    <el-input v-model="userForm.userName" placeholder="请输入姓名"></el-input>
                </el-form-item>
                <el-form-item label="账号" prop="userAccount">
                    <el-input v-model="userForm.userAccount" placeholder="请输入账号"
                              onkeyup="this.value=this.value.replace(/[^\w_]/g,'');"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="userPassword">
                    <el-input type="password" maxlength="20" v-model="userForm.userPassword"
                              placeholder="请输入密码"></el-input>
                </el-form-item>
                <el-form-item label="确认密码" prop="reUserPassword">
                    <el-input type="password" maxlength="20" v-model="userForm.reUserPassword"
                              placeholder="请输入确认密码"></el-input>
                </el-form-item>
                <el-form-item label="邮箱" prop="email">
                    <el-input v-model="userForm.email" placeholder="请输入邮箱"></el-input>
                </el-form-item>
                <el-form-item label="手机号" prop="mobile">
                    <el-input v-model="userForm.mobile" placeholder="请输入手机号" maxlength="11"></el-input>
                </el-form-item>
                <el-form-item label="性别" prop="sex">
                    <el-radio-group v-model="userForm.sex">
                        <el-radio label="1">男</el-radio>
                        <el-radio label="0">女</el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="newVisibleBtn('userForm')">取 消</el-button>
                <el-button type="primary" @click="saveNew('userForm')">确 定</el-button>
            </span>
        </el-dialog>
        <!-- 查看权限-->
        <el-dialog :title="setTitle" :visible.sync="powerVisible" width="37%" :close-on-press-escape="false"
                   :close-on-click-modal="false">
            <el-table :data="powerData" border class="table" ref="multipleTable">
                <el-table-column prop="label" label="角色名称"></el-table-column>
                <el-table-column prop="description" label="角色描述"></el-table-column>
            </el-table>
        </el-dialog>
        <!-- 分配角色 -->
        <el-dialog :title="setTitle" :visible.sync="roleVisible" width="45%" :close-on-press-escape="false"
                   :close-on-click-modal="false">
            <template>
                <el-transfer
                        v-model="value3"
                        filterable
                        :left-default-checked="[]"
                        :right-default-checked="value3"
                        :render-content="renderFunc"
                        :button-texts="['撤销', '授权']"
                        :titles="['未授权','已授权']"
                        :format="{
                        noChecked: '${total}',
                        hasChecked: '${checked}/${total}'
                         }"
                        @change="handleChange"
                        :data="rolesAll">
                    <!--<el-button class="transfer-footer" slot="left-footer" size="small">操作</el-button>
                    <el-button class="transfer-footer" slot="right-footer" size="small">操作</el-button>-->
                </el-transfer>
            </template>
        </el-dialog>
        <!-- 重置密码 -->
        <el-dialog :title="setTitle" :visible.sync="resetVisible" width="37%" :close-on-press-escape="false"
                   :close-on-click-modal="false">
            <el-form :model="resetPasswordForm" status-icon :rules="rules" ref="resetPasswordForm" label-width="100px"
                     class="demo-ruleForm">
                <el-form-item label="新密码" prop="password">
                    <el-input type="password" v-model="resetPasswordForm.password"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="submitForm">提交</el-button>
                    <el-button @click="resetForm">重置</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
        <!-- 编辑 -->
        <el-dialog :title="setTitle" :visible.sync="editUserVisible" width="37%" :close-on-press-escape="false"
                   :close-on-click-modal="false">
            <el-form :model="editUserForm" status-icon ref="editUserForm" label-width="100px" class="demo-ruleForm">
                <el-form-item label="所属部门" prop="departmentId">
                    <el-select placeholder="请选择所属部门" v-model="departmentId">
                        <el-option v-for="item in departmentIdList" :key="item.id" :label="item.depName"
                                   :value="item.id"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="姓名:" prop="username">
                    <el-input v-model="editUserForm.userName"></el-input>
                </el-form-item>
                <el-form-item label="账号:" prop="nickName">
                    <el-input v-model="editUserForm.userAccount" readonly="true"></el-input>
                </el-form-item>
                <el-form-item label="手机号" prop="mobile">
                    <el-input v-model="editUserForm.mobile" maxlength="11"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="submitEditUser">提交</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>

<script>
    export default {
        data() {
            var validatePass2 = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请再次输入密码'))
                } else if (value !== this.userForm.password) {
                    callback(new Error('两次输入密码不一致!'))
                } else {
                    callback()
                }
            };
            return {
                rolesAll: [],//所有角色
                value3: [],//选中角色
                renderFunc(h, option) {
                    return
                <
                    span > {option.label} < /span>
                },


                setTitle: '新增',
                Total: 0,    //总页数
                cur_page: 1,   //当前页
                multipleSelection: [],   //全选
                tableData: [],  //表格数据
                powerData: [], //用户权限数据
                departmentIdList: [],//部门数据
                rolesList: [],  //角色数据
                userTypeList: [],  //员工类型
                departmentId: '',
                userId: "",  //用户id
                roleKeyList: [],//用户已有角色数据
                batchRoleUserId: [],//用户id
                imageUrl: '',
                search: {//搜索
                    userNumber: '',   //员工编号
                    userName: '',   //姓名
                    mobile: '',    //电话
                    departmentId: '',     //所属部门
                    current: 1,
                    rowCount: 10,
                },
                rules: {//校验规则
                    username: [
                        {required: true, message: '请输入用户名', trigger: 'blur'},
                    ],
                    nickName: [
                        {required: true, message: '请输入真实姓名', trigger: 'blur'},
                    ],
                    password: [
                        {required: true, message: '请输入用密码', trigger: 'blur'},
                    ],
                    repassword: [
                        {required: true, validator: validatePass2, trigger: 'blur'}
                    ],
                    email: [
                        {required: true, message: '邮箱格式不正确', trigger: 'blur'},
                        {type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change']}
                    ],
                    mobile: [
                        {required: true, message: '手机号格式不正确', trigger: 'blur'},
                        {
                            pattern: /^((\+?[0-9]{1,4})|(\(\+86\)))?(13[0-9]|14[57]|15[012356789]|17[0678]|18[0-9])\d{8}$/,
                            message: '手机号格式不正确',
                            trigger: 'blur'
                        }
                    ],
                    sex: [
                        {required: true, message: '请选择性别', trigger: 'change'},
                    ],
                    status: [
                        {required: true, message: '请选择状态', trigger: 'change'},
                    ],
                    departmentId: [
                        {required: true, message: '请选择所属部门', trigger: 'change'},
                    ],
                    headPhoto: [
                        {required: true, message: '请上传头像', trigger: 'blur'},
                    ],
                    type: [
                        {required: true, message: '请选择用户类型', trigger: 'change'},
                    ],
                    roleArr: [
                        {required: true, message: '请选择用户角色', trigger: 'change'},
                    ],
                    userType: [
                        {required: true, message: '请选择员工类型', trigger: 'change'},
                    ],
                },
                userForm: {//新增from表单
                    userName: '',  //用户名
                    userAccount: '', //真实姓名
                    userPassword: '',   //密码
                    reUserPassword: '', //确认密码
                    email: '123@qq.com',    //邮箱
                    mobile: '15918605210',     //手机号
                    sex: '',       //性别
                    departmentId: '',   //部门
                    roleArr: []      //角色
                },
                editUserForm: {},
                resetPasswordForm: {//重置重置密码信息
                    id: "",
                    password: ""
                },
                newVisible: false,   //新增弹出框
                batchRolesVisible: false,//批量分配弹出框
                powerVisible: false, //查看权限弹出框
                roleVisible: false, //分配角色弹出框
                resetVisible: false,//重置密码弹出框
                editUserVisible: false,//用户编辑弹出框
                formData: new FormData(),
            }
        },
        methods: {
            handleSelectionChange(val) {    //全选
                this.multipleSelection = val;
                this.batchRoleUserId = [];
                for (let i = 0; i < val.length; i++) {//批量分配获取用户id
                    this.batchRoleUserId.push(val[i].id);
                }
            },
            handleCurrentChange(val) {// 分页导航
                this.cur_page = val;
                this.getData(val, 10);
            },
            searchBtn() {//搜索
                this.getData(1, 10);
            },
            resetBtn() {//重置搜索
                this.search.employeeNum = '';
                this.search.nickName = '';
                this.search.mobile = '';
                this.search.roleid = '';
                this.search.departmentId = '';
                this.getData(1, 10);
            },
            newUsersBtn() { //新增用户按钮
                this.rolesList = [];
                this.setTitle = '新增内部用户';
                this.newVisible = true;
                //this.getRole();
            },
            saveNew(userForm) { //新增保存
                this.$refs[userForm].validate((valid) => {
                    if (valid) {
                        const loading = this.$loading({
                            lock: true,
                            text: 'Loading',
                            spinner: 'el-icon-loading',
                            background: 'rgba(0, 0, 0, 0.7)'
                        });
                        let param = {
                            userName: this.userForm.userName,
                            userAccount: this.userForm.userAccount,
                            userPassword: this.userForm.userPassword,
                            email: this.userForm.email,
                            mobile: this.userForm.mobile,
                            sex: this.userForm.sex,
                            departmentId: this.userForm.departmentId
                        };
                        this.$axios.defaults.headers['urlencoded'] = '-1';
                        this.$axios.post("/perm/user/addUser", param, {}
                        ).then(resp => {
                            let data = resp.data;
                            loading.close();
                            if (data && data.code == 200) {
                                this.$refs[userForm].resetFields();
                                this.newVisible = false;
                                this.$message({
                                    message: '保存成功！',
                                    type: 'success'
                                });
                                this.cleanFormData();
                            } else {
                                this.$refs[userForm].resetFields();
                                this.newVisible = false;
                                this.$message.error('保存失败');
                                this.cleanFormData();
                            }
                            this.getData(1, 10);
                            this.deleteFormData();
                        });
                    }
                })
            },
            cleanFormData() {
                //清空from表单
                this.userForm.username = '',
                    this.userForm.password = '',
                    this.userForm.repassword = '',
                    this.userForm.email = '',
                    this.userForm.mobile = '',
                    this.userForm.nickName = '',
                    this.userForm.sex = '',
                    this.userForm.status = '',
                    this.userForm.type = '',
                    this.userForm.headPhoto = '',
                    this.userForm.departmentId = '',
                    this.userForm.roleArr = [],
                    this.userForm.description = '',
                    this.userForm.userType = ''
            },
            deleteFormData() {
                //清空fromDate表单
                this.formData.delete('username');
                this.formData.delete('nickName');
                this.formData.delete('password');
                this.formData.delete('email');
                this.formData.delete('mobile');
                this.formData.delete('nickName');
                this.formData.delete('sex');
                this.formData.delete('status');
                this.formData.delete('typeSta');
                this.formData.delete('userType');
                this.formData.delete('departmentId');
                this.formData.delete('roleArr');
                this.formData.delete('headPhotoFile');
                this.formData.delete('description');
            },
            handleDialogClose(userForm) {  //新增弹窗关闭
                this.$refs[userForm].resetFields();
            },
            newVisibleBtn(userForm) {  //新增取消
                this.$refs[userForm].resetFields();
                this.newVisible = false;
            },
            searchRoleBtn() {//批量分配搜索
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/role/findRoleList", {}, {}).then(resp => {
                    let data = resp.data;
                    if (data && data.code == 200) {
                        this.rolesList = [];
                        this.rolesAll = [];
                        let Data = data.data;
                        for (let i = 0; i < Data.length; i++) {
                            let role = {};
                            role.key = Data[i].id;
                            role.label = Data[i].roleName;
                            // role.disabled = Data[i].roleCode;
                            this.rolesList.push(role);
                            this.rolesAll.push(role);
                        }
                        console.log(this.rolesAll)
                    } else {
                        console.log(resp);
                    }
                });
            },
            saveRoles() {
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let params = {
                    id: this.batchRoleUserId.join(","),
                    roleArr: this.roleKeyList
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/user/batchRoles", params, {}
                ).then(resp => {
                    let data = resp.data;
                    loading.close();
                    if (data.success) {
                        this.batchRolesVisible = false;
                        this.$message({
                            message: '角色分配成功！',
                            type: 'success'
                        });
                    } else {
                        this.batchRolesVisible = false;
                        this.$message.error('角色分配失败');
                    }
                });
            },
            handleCommand(command) {//操作
                if (command.name == 'handlePermissions') {
                    this.handlePermissions(command.index);
                } else if (command.name == 'handleDAllot') {
                    this.handleDAllot(command.index);
                } else if (command.name == 'handleReset') {
                    this.handleReset(command.index);
                } else if (command.name == 'editUser') {
                    this.editUser(command.index);
                }
            },
            handlePermissions(index) {  //查看权限
                this.powerData = [];
                let item = this.tableData[index];
                this.setTitle = item.userName + '的权限';
                this.powerVisible = true;
                let params = {
                    id: item.id
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/user/findUserRoleSByUserId/", params, {}
                ).then(resp => {
                    let data = resp.data;
                    if (data && data.code == 200) {
                        let Data = data.data;
                        for (let i = 0; i < Data.length; i++) {
                            let role = {};
                            role.key = Data[i].id;
                            role.label = Data[i].roleName;
                            role.description = Data[i].roleCode;
                            this.powerData.push(role);
                        }
                        //this.powerDataToObj(this.powerData, Data.data);
                    } else {
                        console.log(resp)
                    }
                });
            },
            handleDAllot(index) {  //分配角色
                this.value3 = [];
                this.setTitle = "分配角色";
                this.roleVisible = true;
                let item = this.tableData[index];
                this.userId = item.id;
                this.searchRoleBtn();
                let params = {
                    id: item.id
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/user/findUserRoleSByUserId/", params, {}
                ).then(resp => {
                    let data = resp.data;
                    if (data && data.code == 200 && data.data != null) {
                        let Data = data.data;
                        for (let i = 0; i < Data.length; i++) {
                            this.value3.push(Data[i].id);
                        }
                    } else {
                        console.log(resp)
                    }
                });
            },
            handleChange(value, direction, movedKeys) {//分配角色点击授权撤权
                let param = {
                    userId: this.userId,
                    roleIds: movedKeys,
                    checkType: "right" == direction ? "1" : "0",
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/user/handlePermission", param, {}
                ).then(resp => {
                    let data = resp.data;
                    if (data && data.code == 200) {
                        this.$message({
                            message: '角色分配成功！',
                            type: 'success'
                        });
                    } else {
                        this.$message.error('角色分配失败');
                        this.roleVisible = false;
                    }
                });
            },
            handleReset(index) {   //重置密码
                this.setTitle = '重置密码';
                this.resetVisible = true;
                let item = this.tableData[index];
                this.resetPasswordForm.id = item.id
            },
            editUser(index) {//编辑用户
                this.setTitle = '编辑用户';
                this.editUserVisible = true;
                this.departmentIdList = [];
                this.departmentId = [];
                this.getDepartment();
                let item = this.tableData[index];
                let param = {
                    id: item.id
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/user/findUserById", param, {}
                ).then(resp => {
                    let data = resp.data;
                    if (data && data.code == 200) {
                        this.editUserForm = data.data;
                        this.departmentId = data.data.departmentId;
                    } else {
                        console.log(resp)
                    }
                });
            },
            submitEditUser() {//编辑用户提交
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let param = {
                    id: this.editUserForm.id,
                    departmentId: this.departmentId,
                    mobile: this.editUserForm.mobile,
                    userName: this.editUserForm.userName
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/user/updateUser", param, {}
                ).then(resp => {
                    let data = resp.data;
                    loading.close();
                    if (data && data.code == 200) {
                        this.editUserVisible = false;
                        this.$message({
                            message: '修改成功！',
                            type: 'success'
                        });
                        this.getData(1, 10);
                    } else {
                        this.editUserVisible = false;
                        this.$message.error('修改失败');
                        this.getData(1, 10);
                    }
                });
            },
            submitForm() {//重置密码提交
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let param = {
                    id: this.resetPasswordForm.id,
                    userPassword: this.resetPasswordForm.password
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/user/resetPass", param, {}
                ).then(resp => {
                    let data = resp.data;
                    loading.close();
                    if (data && data.code == 200) {
                        this.resetVisible = false;
                        this.$message({
                            message: '重置成功！',
                            type: 'success'
                        });
                        this.resetForm();
                    } else {
                        this.resetVisible = false;
                        this.$message.error('重置失败');
                        this.resetForm();
                    }
                    this.getData(1, 10);
                });
            },
            resetForm() {//清空重置密码
                this.resetPasswordForm.password = ""
            },
            getDepartment() {//所属部门
                this.$axios.post("/perm/department/findDepartmentList", {}
                ).then(resp => {
                    this.departmentIdList = [];
                    let data = resp.data;
                    if (data && data.code == 200) {
                        for (let i = 0; i < data.data.length; i++) {
                            this.departmentIdList.push(data.data[i])
                        }
                    } else {
                        this.Common.GetSessionStorage(this, data.errorMessage);
                    }
                });
            },
            handleTreeData(params) {
                let childrens = [];
                for (let i = 0; i < params.length; i++) {
                    let tmpData = params[i];
                    if (tmpData.menus) {
                        let children = {
                            id: tmpData.id,
                            title: tmpData.title,
                            level: Data.level,
                            children: this.handleTreeData(tmpData.menus)
                        };
                        childrens.push(children);
                    } else {
                        let children = {
                            id: tmpData.id,
                            title: tmpData.title,
                            level: Data.level
                        };
                        childrens.push(children);
                    }
                }
                return childrens;
            },
            turnRoleToObj(list, val) { //角色转换数据
                for (let i = 0; i < val.length; i++) {
                    let role = {};
                    role.key = val[i].id;
                    role.label = val[i].name;
                    role.description = val[i].description;
                    list.push(role);
                }
            },
            getData(indexCurrent, allrowCount) {  // 用户列表
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let param = {
                    currentPage: indexCurrent,
                    pageSize: allrowCount
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/user/findPagesUser", param, {}
                ).then(resp => {
                    let data = resp.data;
                    loading.close();
                    if (data && data.code == 200) {
                        let Data = data.data;
                        this.tableData = Data.data;
                        this.Total = Data.total;
                    } else {
                        console.log(resp)
                    }
                });
            },
        },
        mounted() {
            this.getData(1, 10); //请求表格数据
            this.getDepartment(); //所属部门
        }
    }
</script>


<style scoped>
    .transfer-footer {
        margin-left: 20px;
        padding: 6px 5px;
    }

    .avatar-uploader .el-upload {
        border: 1px dashed #d9d9d9;
        border-radius: 6px;
        cursor: pointer;
        position: relative;
        overflow: hidden;
        width: 146px;
        height: 146px;
    }

    .avatar-uploader-icon {
        font-size: 28px;
        color: #8c939d;
        width: 146px;
        height: 146px;
        line-height: 146px;
        text-align: center;
    }

    .el-upload--picture-card i {
        border-radius: 6px;
        cursor: pointer;
        position: relative;
        overflow: hidden;
        width: 146px;
        height: 146px;
        font-size: 28px;
        color: #8c939d;
        text-align: center;
    }

    .el-form .el-cascader--small,
    .el-form .el-select--small {
        width: 100%;
    }

    .handle-box {
        margin-bottom: 20px;
    }

    .table {
        width: 100%;
        font-size: 14px;
    }

    .demo-input-suffix {
        display: inline-block;
        width: 27%;
        text-align: right;
    }
</style>
