<template>
    <div class="table">
        <div class="container">
            <div class="handle-box">
                <el-row :gutter="20">
                    <el-col :span="8">
                        <div class="demo-input-suffix">角色名称:</div>
                        <el-input style="width:60%;" placeholder="请输入角色名称" size="small" v-model="searchParam.roleName"></el-input>
                    </el-col>
                    <el-col :span="8">
                        <div class="demo-input-suffix">角色编码:</div>
                        <el-input style="width:60%;" placeholder="请选择角色编码" size="small" v-model="searchParam.roleCode"></el-input>
                    </el-col>
                    <el-col :span="8">
                        <el-button type="primary" icon="search" @click="searchBtn">搜 索</el-button>
                        <el-button icon="search" @click="resetBtn">重 置</el-button>
                    </el-col>
                </el-row>
                <el-button type="primary" icon="add" class="handle-del mr10" @click="addBtn">添加角色</el-button>
            </div>
            <el-table :data="tableData" border class="table" ref="multipleTable" @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55"></el-table-column>
                <el-table-column type="index" label=" " width="50" align="center"></el-table-column>
                <el-table-column prop="id" v-if="false"></el-table-column>
                <el-table-column prop="roleName" label="角色名称"></el-table-column>
                <el-table-column prop="roleCode" label="角色编码" width="300"></el-table-column>
                <el-table-column prop="description" label="描述" width="300"></el-table-column>
                <el-table-column label="操作" align="center">
                    <template slot-scope="scope">
                        <el-button type="text" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                        <el-button type="text" @click="handleDistributeResource(scope.$index, scope.row)">分配资源</el-button>
                        <el-button type="text" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination background @current-change="handleCurrentChange" layout="prev, pager, next" :total="Total"></el-pagination>
            </div>
        </div>
        <!-- 新增弹出框 -->
        <el-dialog :title="setTitle" :visible.sync="addVisible" width="30%" @close="addClose('roleData')" :show-close="false">
            <el-form :model="roleData" :rules="rules" ref="roleData" label-width="110px">
                <el-form-item label="角色名称" prop="roleName">
                    <el-input v-model="roleData.roleName" placeholder="请输入角色名称"></el-input>
                </el-form-item>
                <el-form-item label="角色编码" prop="roleName">
                    <el-input v-model="roleData.roleCode" placeholder="请输入角色编码"></el-input>
                </el-form-item>
                <el-form-item label="描述" prop="description">
                    <el-input v-model="roleData.description" placeholder="描述"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="addCancel('roleData')">取 消</el-button>
                <el-button type="primary" @click="addConfirm('roleData')">确 定</el-button>
            </span>
        </el-dialog>
        <!--编辑弹出框-->
        <el-dialog :title="setTitle" :visible.sync="editVisible" width="30%" @close="editClose('roleData')" :show-close="false">
            <el-form :model="roleData" :rules="rules" ref="roleData" label-width="110px">
                <el-form-item label="角色名称" prop="roleName">
                    <el-input v-model="roleData.roleName" placeholder="请输入角色名称"></el-input>
                </el-form-item>
                <el-form-item label="角色名称" prop="roleName">
                    <el-input v-model="roleData.roleCode" placeholder="请输入角色名称"></el-input>
                </el-form-item>
                <el-form-item label="描述" prop="description">
                    <el-input v-model="roleData.description" placeholder="描述"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="editCancel('roleData')">取 消</el-button>
                <el-button type="primary" @click="editConfirm('roleData')">确 定</el-button>
            </span>
        </el-dialog>
        <!-- 删除提示框 -->
        <el-dialog title="提示" :visible.sync="deleteDialogFlag" width="30%" center :show-close="false">
            <span class="delTip">是否确认要删除?</span>
            <span slot="footer" class="dialog-footer">
                <el-button @click="deleteCancel()">取 消</el-button>
                <el-button type="primary" @click="deleteConfrim()">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
    export default {
        name: 'permission-list',
        data() {
            return {

                setTitle: '',
                tableData: [],  //表格数据
                Total: 0,       //总页数
                cur_page: 1,   //当前页
                multipleSelection: [],   //全选
                searchParam: { //搜索
                    current: 1,
                    rowCount: 10,
                    roleName: '',
                    roleCode: ''
                },
                roleData: {//新增或编辑数据
                    id: '',
                    roleName: '',
                    roleCode: '',
                    description: ''
                },
                isRoleExist: false,
                roleList: [], //角色类型
                addVisible: false,         //添加弹出框
                editVisible: false,       //编辑弹出框
                deleteDialogFlag: false, //删除弹出框
                rules: {
                    roleName: [
                        {required: true, message: '请输入角色名称', trigger: 'blur'},
                    ],
                    roleCode: [
                        {required: true, message: '请输入角色编码', trigger: 'blur'},
                    ],
                    description: [
                        {required: true, message: '请输入描述', trigger: 'blur'},
                    ]
                }
            }
        },
        methods: {
            handleSelectionChange(val) {//全选
                this.multipleSelection = val;
            },
            handleCurrentChange(val) {// 分页导航
                this.cur_page = val;
                this.getData(val, 10);
            },
            searchBtn() {// 搜索
                this.getData(1, 10);
            },
            resetBtn() {//重置
                this.searchParam.roleName = '';
                this.searchParam.roleCode = '';
                this.getData(1, 10);
            },
            cleanFormData() {// 清空form表单数据
                this.roleData.id = '';
                this.roleData.roleName = '';
                this.roleData.roleCode = '';
                this.roleData.description = '';
            },
            addBtn() { // 添加角色按钮
                this.setTitle = '新增角色';
                this.addVisible = true;
            },
            addClose(roleData) { //新增弹窗关闭
                this.$refs[roleData].resetFields();
            },
            addCancel(roleData) {//新增弹窗取消
                // this.$refs[roleData].resetFields();
                this.addVisible = false;
                this.cleanFormData();
                // this.getData(1,10);
            },
            addConfirm(roleData) {//添加确认
                this.$refs[roleData].validate((valid) => {
                    if (valid) {
                        const loading = this.$loading({
                            lock: true,
                            text: 'Loading',
                            spinner: 'el-icon-loading',
                            background: 'rgba(0, 0, 0, 0.7)'
                        });
                        let param = {
                            roleName: this.roleData.roleName,
                            roleCode: this.roleData.roleCode,
                            description: this.roleData.description
                        };
                        this.$axios.defaults.headers['urlencoded'] = '-1';
                        this.$axios.post("/perm/role/addRole", param, {}
                        ).then(resp => {
                            let data = resp.data;
                            loading.close();
                            if (data && data.code == 200) {
                                this.addVisible = false;
                                this.$message({
                                    message: '保存成功！',
                                    type: 'success'
                                });
                                // this.$router.push('./role-edit?id=' + data.data.id);
                                this.cleanFormData();
                            } else {
                                this.addVisible = false;
                                this.$message.error('保存失败');
                                this.cleanFormData();
                            }
                            this.getData(1, 10);
                        });
                    }
                });
            },
            handleEdit(index, row) {//编辑角色
                this.setTitle = '编辑角色';
                this.editVisible = true;
                let item = this.tableData[index];
                this.roleData.id = item.id;
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/role/findRoleById", {id: item.id}, {}
                ).then(resp => {
                    let data = resp.data;
                    if (data && data.code == 200) {
                        let roleDetail = data.data;
                        // 初始化反显数据
                        this.roleData.id = roleDetail.id;
                        this.roleData.roleName = roleDetail.roleName;
                        this.roleData.roleCode = roleDetail.roleCode;
                        this.roleData.description = roleDetail.description;
                    } else {
                        this.$message.error('数据加载失败');
                        this.isRoleExist = true;
                    }
                });
            },
            checkRoleName() {//检查角色名称是否存在
                if (this.roleData.roleCode != "" && this.roleData.roleCode != null) {
                    let param = {
                        id: this.roleData.id,
                        roleCode: this.roleData.roleCode
                    };
                    this.$axios.defaults.headers['urlencoded'] = '-1';
                    this.$axios.post("/perm/role/isRoleCodeExist", param, {}
                    ).then(resp => {
                        let data = resp.data;
                        if (data.success) {
                            this.isRoleExist = false;
                        } else {
                            this.$message.error(data.errorMessage);
                            this.isRoleExist = true;
                        }
                    });
                }
            },
            editClose(roleData) {//编辑弹窗关闭
                this.$refs[roleData].resetFields();
                this.getData(1, 10);
            },
            editCancel(roleData) {//编辑弹窗取消
                // this.$refs[roleData].resetFields();
                this.cleanFormData();
                this.editVisible = false;
                // this.getData(1,10);
            },
            editConfirm(roleData) {//编辑弹窗确认
                this.$refs[roleData].validate((valid) => {
                    if (valid && !this.isRoleExist) {
                        const loading = this.$loading({
                            lock: true,
                            text: 'Loading',
                            spinner: 'el-icon-loading',
                            background: 'rgba(0, 0, 0, 0.7)'
                        });
                        let param = {
                            id: this.roleData.id,
                            roleName: this.roleData.roleName,
                            roleCode: this.roleData.roleCode,
                            description: this.roleData.description
                        };
                        this.$axios.defaults.headers['urlencoded'] = '-1';
                        this.$axios.post("/perm/role/updateRole", param, {}
                        ).then(resp => {
                            let data = resp.data;
                            loading.close();
                            if (data && data.code == 200) {
                                this.editVisible = false;
                                this.$message({
                                    message: '保存成功！',
                                    type: 'success'
                                });
                                this.cleanFormData();
                            } else {
                                this.editVisible = false;
                                this.$message.error('保存失败');
                                this.cleanFormData();
                            }
                            this.getData(1, 10);
                        });
                    } else {
                        this.$message.error('角色名称已存在');
                    }
                });
            },
            handleDistributeResource(index, row) {// 分配资源
                let item = this.tableData[index];
                this.$router.push('./role-edit?id=' + item.id);
            },
            handleDelete(index, row) {//删除角色
                this.deleteDialogFlag = true;
                let item = this.tableData[index];
                this.roleData.id = item.id;
            },
            deleteCancel() {//删除取消
                this.deleteDialogFlag = false;
                // this.getData();
            },
            deleteConfrim() {//删除角色确认
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let param = {
                    ids: this.roleData.id
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/role/deleteRole", param, {}
                ).then(resp => {
                    let data = resp.data;
                    loading.close();
                    if (data && data.code == 200) {
                        this.$message({
                            message: '删除成功！',
                            type: 'success'
                        });
                        this.cleanFormData();
                    } else {
                        this.$message.error('删除失败');
                        this.cleanFormData();
                    }
                    this.deleteDialogFlag = false;
                    this.getData(1, 10);
                });
            },
            getData(indexCurrent, allrowCount) {//请求表格数据
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let param = {
                    currentPage: indexCurrent,
                    pageSize: allrowCount,
                    roleName: this.searchParam.roleName,
                    roleCode: this.searchParam.roleCode
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/role/findRolePages", param, {}
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
            }
        },
        mounted() {
            this.getData(1, 10); //请求表格数据
        }
    }
</script>


<style scoped>
    .handle-box {
        margin-bottom: 20px;
    }

    .handle-select {
        width: 120px;
    }

    .handle-input {
        width: 300px;
        display: inline-block;
    }

    .del-dialog-cnt {
        font-size: 16px;
        text-align: center
    }

    .table {
        width: 100%;
        font-size: 14px;
    }

    .red {
        color: #ff0000;
    }

    .demo-input-suffix {
        display: inline-block;
        width: 27%;
        text-align: right;
    }

    .delTip {
        text-align: center;
    }
</style>
