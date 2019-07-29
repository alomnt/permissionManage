<template>
    <div class="table">
        <div class="container">
            <div class="handle-box">
                <el-row :gutter="20">
                    <el-col :span="8">
                        <div class="demo-input-suffix">菜单名称:</div>
                        <el-input style="width:60%;" placeholder="请输入菜单名称" size="small"
                                  v-model="searchParam.resourceName"></el-input>
                    </el-col>
                    <el-col :span="8">
                        <el-button type="primary" icon="search" @click="searchBtn">搜 索</el-button>
                        <el-button icon="search" @click="resetBtn">重 置</el-button>
                    </el-col>
                </el-row>
                <el-button type="primary" icon="add" class="handle-del mr10" @click="addBtn">添加菜单</el-button>
            </div>
            <el-table :data="tableData" border class="table" ref="multipleTable"
                      @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55"></el-table-column>
                <el-table-column type="index" label=" " width="50" align="center"></el-table-column>
                <el-table-column prop="id" v-if="false"></el-table-column>
                <el-table-column prop="parentId" v-if="false"></el-table-column>
                <el-table-column prop="resourceName" label="菜单名称" sortable width="200"></el-table-column>
                <el-table-column prop="resourceLevel" label="菜单级别" width="100"></el-table-column>
                <el-table-column prop="url" label="url" width="220"></el-table-column>
                <el-table-column prop="resourceOrder" label="顺序" width="100"></el-table-column>
                <el-table-column prop="icon" label="图标" width="200"></el-table-column>
                <el-table-column fixed="right" label="操作" align="center">
                    <template slot-scope="scope">
                        <el-dropdown @command="handleCommand">
                            <el-button type="primary">操作<i class="el-icon-arrow-down el-icon--right"></i></el-button>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item :command="{name:'handleEdit',index:scope.$index}">编辑
                                </el-dropdown-item>
                                <el-dropdown-item :command="{name:'handleDelete',index:scope.$index}">删除
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
        <!-- 新增/编辑弹出框 -->
        <el-dialog :title="setTitle" :visible.sync="editVisible" width="30%" @close="editClose('permissionForm')"
                   :show-close="false">
            <el-form :model="permissionForm" :rules="rules" ref="permissionForm" label-width="110px">
                <el-form-item label="选择父菜单">
                    <el-input placeholder="输入关键字进行过滤" v-model="filterText"></el-input>
                    <el-tree class="filter-tree"
                             :data="selectTreeData"
                             :props="permissionForm.parentId"
                             @node-click="handleNodeClick"
                             @click="setCheckedKeys"
                             :filter-node-method="filterNode"
                             ref="tree2">
                    </el-tree>
                </el-form-item>
                <el-form-item label="菜单名称" prop="title">
                    <el-input v-model="permissionForm.resourceName" placeholder="请输入菜单名称"></el-input>
                </el-form-item>
                <el-form-item label="描述" prop="description">
                    <el-input v-model="permissionForm.description" placeholder="描述"></el-input>
                </el-form-item>
                <el-form-item label="url">
                    <el-input v-model="permissionForm.url" placeholder="url"></el-input>
                </el-form-item>
                <el-form-item label="图标" prop="icon">
                    <el-input v-model="permissionForm.icon" placeholder="图标"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="editVisibleBtn()">取 消</el-button>
                <el-button type="primary" @click="saveEdit()">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
    export default {
        name: 'permission-list',
        data() {
            return {

                setTitle: '新增菜单',
                tableData: [],  //表格数据
                Total: 0,    //总页数
                cur_page: 1,   //当前页
                multipleSelection: [],   //全选
                departmentIdList: [],  //部门树
                outSourceList: [],    //员工类型
                roleidList: [],     //角色
                searchParam: {
                    resourceName: '',
                    type: 0
                },
                rules: {
                    parentId: [
                        {required: true, message: '请选择父菜单', trigger: 'blur'},
                    ],
                    resourceName: [
                        {required: true, message: '请输入菜单名称', trigger: 'blur'},
                    ],
                    url: [
                        {required: true, message: '请输入url', trigger: 'blur'},
                    ]
                },
                permissionForm: {
                    id: '',
                    parentId: '',
                    resourceName: '',
                    description: '',
                    url: '',
                    icon: ''
                },
                filterText: '',
                selectTreeData: [],
                editVisible: false,   //新增弹出框
                select_word: ''   //搜索值
            }
        },
        watch: {
            filterText(val) {
                this.$refs.tree2.filter(val);
            }
        },
        methods: {
            handleSelectionChange(val) {    //全选
                this.multipleSelection = val;
            },
            // 搜索
            searchBtn() {
                this.getData(1, 10);
            },
            //重置
            resetBtn() {
                this.searchParam.resourceName = '';
                this.searchParam.current = 1;
                this.getData(1, 10);
            },
            // 清空form表单数据
            cleanFormData() {
                this.permissionForm.parentId = '';
                this.permissionForm.description = '';
                this.permissionForm.icon = '';
                this.permissionForm.resourceName = '';
                this.permissionForm.url = '';
                this.filterText = '';
            },
            // 添加
            addBtn() {
                this.setTitle = '新增菜单';
                this.cleanFormData();
                this.editVisible = true;
                this.getSelectTreeData();
            },
            handleCommand(command) {    //操作
                if (command.name == 'handleEdit') {
                    this.handleEdit(command.index);
                } else if (command.name == 'handleDelete') {
                    this.handleDelete(command.index);
                }
            },
            handleEdit(index, row) {   //编辑
                this.editVisible = true;
                this.setTitle = '编辑';
                this.getSelectTreeData();
                const item = this.tableData[index];
                // this.$refs.tree2.setCheckedKeys(item.parentId);
                // this.setCheckedKeys();
                this.filterText = item.resourceName;
                this.permissionForm = {
                    id: item.id,
                    parentId: item.parentId,
                    resourceName: item.resourceName,
                    description: item.description,
                    url: item.url,
                    icon: item.icon
                };
            },
            handleDelete(index, row) {  //删除
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let item = this.tableData[index];
                let param = {
                    ids: item.id
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/resoure/deleteResource", param, {}
                ).then(resp => {
                    let data = resp.data;
                    loading.close();
                    if (data.success) {
                        this.$message({
                            message: '删除成功！',
                            type: 'success'
                        });
                        this.cleanFormData();
                    } else {
                        this.$message.error(data.errorMessage);
                        this.cleanFormData();
                    }
                    this.getData(this.cur_page, 10);
                });
            },
            handleCurrentChange(val) {   // 分页导航
                this.cur_page = val;
                this.getData(val, 10);
            },
            saveEdit() { //编辑保存
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let param = this.permissionForm;
                this.$axios.defaults.headers['urlencoded'] = '-1';
                let url = "";
                if (this.permissionForm.id == '') {
                    url = "/perm/resoure/addResource";
                } else {
                    url = "/perm/resoure/updateResource";
                }
                this.$axios.post(url, param, {}
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
                    this.getData(this.cur_page, 10);
                });
            },
            editClose(permissionForm) {//编辑弹窗关闭
                this.$refs[permissionForm].resetFields();
            },
            editVisibleBtn() {  //新增或编辑取消
                this.cleanFormData();
                this.editVisible = false;
                // 重新加载数据
                // this.getData();
            },
            filteroutSource(row, column, cellValue) {  //员工类型
                if (cellValue == '0') {
                    return '内部用户';
                } else if (cellValue == '1') {
                    return '外包用户';
                } else if (cellValue == '3') {
                    return '渠道用户';
                }
            },
            getData(cur_page, pageSize) {
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let param = {
                    currentPage: cur_page,
                    pageSize: pageSize,
                    resourceName: this.searchParam.resourceName,
                    type: this.searchParam.type
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/resoure/findPagesResource", param, {}
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
            filterNode(value, data) {
                if (!value) return true;
                return data.label.indexOf(value) !== -1;
            },
            getSelectTreeData() {
                this.selectTreeData = [];
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/resoure/findResourceTreeAll", {}, {}
                ).then(resp => {
                    let data = resp.data;
                    if (data && data.code == 200) {
                        let Data = data.data;
                        this.selectTreeData = [];
                        this.selectTreeData.push(Data);
                    } else {
                        console.log(resp)
                    }
                });
            },
            handleNodeClick(data) {
                this.filterText = data.label;
                this.permissionForm.parentId = data.id;
            },
            setCheckedKeys() {
                this.$refs.tree.setCheckedKeys([3]);
            },
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
</style>
