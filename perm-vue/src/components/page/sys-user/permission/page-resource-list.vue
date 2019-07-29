<template>
    <div class="table">
        <div class="container">
            <div class="handle-box">
                <el-row :gutter="20">
                    <el-col :span="8">
                        <div class="demo-input-suffix">页面元素名称:</div>
                        <el-input style="width:60%;" placeholder="请输入页面元素名称" size="small"
                                  v-model="searchParam.name"></el-input>
                    </el-col>
                    <el-col :span="8">
                        <div class="demo-input-suffix">权限:</div>
                        <el-input style="width:60%;" placeholder="请输入权限" size="small"
                                  v-model="searchParam.buttonType"></el-input>
                    </el-col>
                    <el-col :span="8">
                        <el-button type="primary" icon="search" @click="searchBtn">搜 索</el-button>
                        <el-button icon="search" @click="resetBtn">重 置</el-button>
                    </el-col>
                </el-row>
                <el-button type="primary" icon="add" class="handle-del mr10" @click="addBtn">添加页面元素</el-button>
            </div>
            <el-table :data="tableData" border class="table" ref="multipleTable"
                      @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55"></el-table-column>
                <el-table-column type="index" label=" " width="50" align="center"></el-table-column>
                <el-table-column prop="id" v-if="false"></el-table-column>
                <el-table-column prop="parentId" v-if="false"></el-table-column>
                <el-table-column prop="belongPage" label="所属页面"></el-table-column>
                <el-table-column prop="name" label="页面元素名称"></el-table-column>
                <el-table-column prop="buttonType" label="权限控制" width="120"></el-table-column>
                <el-table-column prop="description" label="描述" width="120"></el-table-column>
                <el-table-column label="操作" width="180" align="center">
                    <template slot-scope="scope">
                        <el-button type="text" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                        <el-button type="text" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination background @current-change="handleCurrentChange" layout="prev, pager, next"
                               :total="Total"></el-pagination>
            </div>
        </div>
        <!-- 编辑弹出框 -->
        <el-dialog :title="setTitle" :visible.sync="editVisible" width="30%" :show-close="false">
            <el-form :model="resourceForm" :rules="rules" ref="resourceForm" label-width="110px">
                <el-form-item label="所属页面" prop="parentId">
                    <el-input placeholder="输入关键字进行过滤" v-model="filterText"></el-input>
                    <el-tree class="filter-tree"
                             :data="selectTreeData"
                             show-checkbox
                             node-key="id"
                             :default-expanded-keys="expandedKeys"
                             :default-checked-keys="checkedKeys"
                             @node-click="handleNodeClick"
                             @click="setCheckedKeys"
                             :filter-node-method="filterNode"
                             ref="resourceTree">
                    </el-tree>
                </el-form-item>
                <el-form-item label="页面元素名称" prop="name">
                    <el-input v-model="resourceForm.name" placeholder="请输入页面元素名称"></el-input>
                </el-form-item>
                <el-form-item label="权限控制" prop="buttonType">
                    <el-input v-model="resourceForm.buttonType" placeholder="请输入权限控制"
                              @blur="checkButtonType()"></el-input>
                </el-form-item>
                <el-form-item label="描述" prop="description">
                    <el-input v-model="resourceForm.description" placeholder="描述"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="handleAddCancel()">取 消</el-button>
                <el-button type="primary" @click="saveEdit('resourceForm')">确 定</el-button>
            </span>
        </el-dialog>

        <!-- 删除提示框 -->
        <el-dialog title="提示" :visible.sync="deleteDialogFlag" width="30%" center :show-close="false">
            <span class="delTip">是否确认要删除?</span>
            <span slot="footer" class="dialog-footer">
                <el-button @click="cancelDelete()">取 消</el-button>
                <el-button type="primary" @click="deleteResource()">确 定</el-button>
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
                searchParam: {
                    name: '',
                    buttonType: '',
                    type: 1
                },
                rules: {
                    parentId: [
                        {required: true, message: '请选择所属页面', trigger: 'blur'},
                    ],
                    name: [
                        {required: true, message: '请输入菜单名称', trigger: 'blur'},
                    ],
                    buttonType: [
                        {required: true, message: '请输入url', trigger: 'blur'},
                    ]
                },
                resourceForm: {
                    id: '',
                    parentId: '',
                    name: '',
                    buttonType: '',
                    description: '',
                    type: 1
                },
                checkType: '',
                isResourceExist: false,
                deleteDialogFlag: false,
                editVisible: false,   //新增弹出框
                filterText: '',
                selectTreeData: [],
                expandedKeys: [],
                checkedKeys: [],
                select_word: ''   //搜索值
            }
        },
        watch: {
            filterText(val) {
                this.$refs.resourceTree.filter(val);
            }
        },
        methods: {
            handleCurrentChange(val) {   // 分页导航
                this.searchParam.current = val;
                this.getData();
            },
            handleSelectionChange(val) {    //全选
                this.multipleSelection = val;
            },
            // 搜索
            searchBtn() {
                this.getData(1, 10);
            },
            //重置
            resetBtn() {
                this.searchParam.name = '';
                this.searchParam.buttonType = '';
                this.getData(1, 10);
            },
            // 清空form表单数据
            cleanFormData() {
                this.resourceForm.id = '';
                this.resourceForm.parentId = '';
                this.resourceForm.name = '';
                this.resourceForm.buttonType = '';
                this.resourceForm.description = '';
            },
            filterNode(value, data) {
                if (!value) return true;
                return data.label.indexOf(value) !== -1;
            },
            getSelectTreeData() {
                this.selectTreeData = [];
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/permission/getAllList", {}, {}
                ).then(resp => {
                    let data = resp.data;
                    if (data.success) {
                        let Data = data.data;
                        for (let i = 0; i < Data.length; i++) {
                            if (Data[i].menus) {
                                let children = this.handleTreeData(Data[i].menus);
                                let tmp = {
                                    id: Data[i].id,
                                    label: Data[i].title,
                                    children: children
                                };
                                this.selectTreeData.push(tmp);
                            } else {
                                let tmp = {
                                    id: Data[i].id,
                                    label: Data[i].title,
                                };
                                this.selectTreeData.push(tmp);
                            }
                        }
                    } else {
                        console.log(resp)
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
                            label: tmpData.title,
                            children: this.handleTreeData(tmpData.menus)
                        };
                        childrens.push(children);
                    } else {
                        let children = {
                            id: tmpData.id,
                            label: tmpData.title
                        };
                        childrens.push(children);
                    }
                }
                return childrens;
            },
            handleNodeClick(data) {
                this.resourceForm.parentId = data.id;
                this.filterText = data.label;
            },
            setCheckedKeys() {
                this.$refs.resourceTree.setCheckedKeys();
            },
            // 添加
            addBtn() {
                this.checkType = 0;
                this.setTitle = '新增页面元素';
                this.editVisible = true;
                this.getSelectTreeData();
            },
            handleAddCancel() {  //新增或编辑取消
                this.cleanFormData();
                this.editVisible = false;
                // 重新加载数据
                this.getData();
            },
            handleEdit(index, row) {   //编辑
                this.checkType = 1;
                this.editVisible = true;
                this.setTitle = '编辑';
                this.$nextTick(() => {
                    const item = this.tableData[index];
                    this.getSelectTreeData();
                    this.resourceForm.id = item.id;
                    this.resourceForm.parentId = item.parentId;
                    this.resourceForm.name = item.name;
                    this.resourceForm.buttonType = item.buttonType;
                    this.resourceForm.description = item.description;
                    // 反显选中的值
                    this.checkedKeys = [];
                    this.expandedKeys = [];
                    this.checkedKeys.push(item.parentId);
                    this.expandedKeys.push(item.parentId);
                    this.$refs.resourceTree.setCheckedKeys(this.checkedKeys);
                });
            },
            checkButtonType() {
                if (this.resourceForm.buttonType != "" && this.resourceForm.buttonType != null) {
                    let param = {
                        buttonType: this.resourceForm.buttonType,
                        checkType: this.checkType,
                        id: this.resourceForm.id
                    };
                    this.$axios.defaults.headers['urlencoded'] = '-1';
                    this.$axios.post("/perm/permission/isResourcePermissionExist", param, {}
                    ).then(resp => {
                        let data = resp.data;
                        if (!data.data) {
                            this.$message.error('页面元素已经存在');
                            this.isResourceExist = true;
                        } else {
                            this.isResourceExist = false;
                        }
                    });
                }
            },
            handleDelete(index, row) {
                this.deleteDialogFlag = true;
                let item = this.tableData[index];
                this.resourceForm.id = item.id;
            },
            deleteResource() {  //删除
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let param = {
                    id: this.resourceForm.id
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/permission/delByIds", param, {}
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
                    this.getData();
                });
            },
            cancelDelete() {
                this.deleteDialogFlag = false;
                this.getData();
            },
            saveEdit(resourceForm) { //编辑保存
                this.$refs[resourceForm].validate((valid) => {
                    if (valid && !this.isResourceExist) {
                        const loading = this.$loading({
                            lock: true,
                            text: 'Loading',
                            spinner: 'el-icon-loading',
                            background: 'rgba(0, 0, 0, 0.7)'
                        });
                        let dataParam = {
                            id: this.resourceForm.id,
                            parentId: this.resourceForm.parentId,
                            name: this.resourceForm.name,
                            description: this.resourceForm.description,
                            buttonType: this.resourceForm.buttonType,
                            type: this.resourceForm.type,
                            checkType: this.checkType
                        };
                        this.$axios.defaults.headers['urlencoded'] = '-1';
                        this.$axios.post("/perm/permission/createOrUpdatePageResource", dataParam, {}
                        ).then(resp => {
                            let data = resp.data;
                            loading.close();
                            if (data.success) {
                                this.editVisible = false;
                                this.$message({
                                    message: '保存成功！',
                                    type: 'success'
                                });
                            } else {
                                this.editVisible = false;
                                this.$message.error('保存失败');
                            }
                            this.getData();
                            this.cleanFormData();
                        });
                    }
                });
            },
            getData(cur_page, pageSize) {
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let param = {
                    current: cur_page,
                    rowCount: pageSize,
                    name: this.searchParam.name,
                    buttonType: this.searchParam.buttonType,
                    type: this.searchParam.type
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/permission/pages", param, {}
                ).then(resp => {
                    let data = resp.data;
                    if (data.success) {
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
</style>
