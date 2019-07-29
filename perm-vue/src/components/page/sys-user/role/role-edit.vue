<template>
    <div class="">
        <div class="container">
            `
            <el-tabs v-model="message" @tab-click="handleClick()">
                <el-tab-pane :label="'分配菜单'" name="menu">
                    <template slot-scope="scope">
                        <el-input placeholder="输入关键字进行过滤" v-model="filterText"></el-input>
                        <el-tree class="filter-tree"
                                 :data="selectTreeData"
                                 node-key="id"
                                 show-checkbox
                                 :default-expanded-keys="expandedKeys"
                                 :default-checked-keys="checkedKeys"
                                 @node-click="handleNodeClick"
                                 @click="setCheckedKeys"
                                 :filter-node-method="filterNode"
                                 ref="tree2">
                        </el-tree>
                    </template>
                </el-tab-pane>

                <el-tab-pane :label="`分配页面元素`" name="resource">
                    <template v-if="message === 'resource'">
                        <div class="container">
                            <div class="handle-box">
                                <el-row :gutter="20">
                                    <el-col :span="8">
                                        <div class="demo-input-suffix">页面元素名称:</div>
                                        <el-input style="width:40%;" placeholder="请输入页面元素名称" size="small" v-model="searchParam.name"></el-input>
                                    </el-col>
                                    <el-col :span="8">
                                        <el-button type="primary" icon="search" @click="searchBtn">搜 索</el-button>
                                        <el-button icon="search" @click="resetBtn">重 置</el-button>
                                    </el-col>
                                </el-row>
                            </div>
                            <el-table :data="tableData" border class="table" ref="multipleTable" @selection-change="handleSelectionChange">
                                <el-table-column type="selection" width="55"></el-table-column>
                                <el-table-column type="index" label=" " width="50" align="center"></el-table-column>
                                <el-table-column prop="id" v-if="false"></el-table-column>
                                <el-table-column prop="parentId" v-if="false"></el-table-column>
                                <!--<el-table-column prop="belongPage" label="所属页面"></el-table-column>-->
                                <el-table-column prop="resourceName" label="页面元素名称"></el-table-column>
                                <!--<el-table-column prop="buttonType" label="权限控制" width="120"></el-table-column>-->
                                <el-table-column prop="description" label="描述" width="120"></el-table-column>
                            </el-table>
                            <div class="pagination">
                                <el-pagination background @current-change="handleCurrentChange" layout="prev, pager, next" :total="Total"></el-pagination>
                            </div>
                        </div>
                    </template>
                </el-tab-pane>
            </el-tabs>
            <div class="handle-row">
                <el-button type="primary" @click="handleSaveEdit()">保 存</el-button>
                <el-button type="primary" @click="handleBack()">返 回</el-button>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'tabs',
        data() {
            return {

                message: 'menu', //默认页面
                tableData: [],  //表格数据
                Total: 0,       //总页数
                cur_page: 1,   //当前页
                multipleSelection: [],// 选中的行
                searchParam: {
                    current: 1,
                    rowCount: 10,
                    name: '',
                    buttonType: '',
                    type: 1
                },
                roleId: '',
                authMenus: [], //以授权菜单
                filterText: '',
                selectTreeData: [],
                expandedKeys: [],
                checkedKeys: [],
                select_word: '',   //搜索值
            }
        },
        watch: {
            filterText(val) {
                this.$refs.tree2.filter(val);
            }
        },
        methods: {
            handleClick() {//切换页面
                if (this.message != 'menu') {
                    this.getResourcePermissionsByRoleIds();
                }
            },
            // 返回
            handleBack() {
                this.$router.push('./role-list');
            },
            handleSelectionChange(val) {//全选
                this.multipleSelection = val;
            },
            handleCurrentChange(val) {// 分页导航
                this.searchParam.current = val;
                this.getResourceData(val, 10);
            },
            // 搜索
            searchBtn() {
                this.getResourceData(1, 10);
            },
            //重置
            resetBtn() {
                this.searchParam.name = '';
                this.searchParam.current = 1;
                this.getResourceData(1, 10);
            },
            filterNode(value, data) {
                if (!value) return true;
                return data.label.indexOf(value) !== -1;
            },
            handleSaveEdit() {//编辑保存
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let election = this.$refs.tree2.getCheckedKeys().toString()
                let halfElection = this.$refs.tree2.getHalfCheckedKeys().toString();
                let menuPermission = '';
                if (halfElection != null) {
                    menuPermission = election + "," + halfElection;
                } else {
                    menuPermission = election;
                }
                let resourcePermission = '';
                if (this.multipleSelection) {
                    for (let index in this.multipleSelection) {
                        resourcePermission = resourcePermission + "," + this.multipleSelection[index]['id'];
                    }
                }
                let permission = menuPermission + resourcePermission;
                let param = {
                    roleId: this.roleId,
                    resourceIds: permission
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/role/createRoleResource", param, {}
                ).then(resp => {
                    let data = resp.data;
                    loading.close();
                    if (data && data.code == 200) {
                        this.$message({
                            message: '保存成功！',
                            type: 'success'
                        });
                        this.handleBack();
                    } else {
                        this.$message.error('保存失败');
                    }
                });
            },
            getSelectTreeData() {//获取菜单树
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                this.selectTreeData = [];
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/resoure/findResourceTreeAll", {}, {}
                ).then(resp => {
                    let data = resp.data;
                    loading.close();
                    if (data && data.code == 200) {
                        let Data = data.data;
                        this.selectTreeData = [];
                        this.selectTreeData.push(Data);
                        this.getRolePermissionsByRoleIds();//获取已授权菜单
                    } else {
                        console.log(resp)
                    }
                });
            },
            getRolePermissionsByRoleIds() {//获取已授权菜单
                this.authMenus = [];
                let param = {
                    id: this.roleId,
                    resourceType: '0'
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/role/getRoleResourceByRoleIds", param, {}
                ).then(resp => {
                    let data = resp.data.data;
                    for (let i in data) {
                        this.checkedKeys.push(data[i]);
                        this.expandedKeys.push(data[i]);
                    }
                    this.$refs.tree2.setCheckedKeys(this.checkedKeys);
                    this.authMenus = data;      //已授权菜单
                    this.getResourceData(1, 10);//获取已授菜单的页面元素
                })
            },
            handleNodeClick(data) {
                this.filterText = data.label;
            },
            setCheckedKeys() {
                this.$refs.tree2.setCheckedKeys();
            },
            getResourceData(indexCurrent, allrowCount) {//获取已授权菜单页面元素
                let param = {
                    currentPage: indexCurrent,
                    pageSize: allrowCount,
                    resourceType: '1'
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/resoure/findPagesResource", param, {}
                ).then(resp => {
                    let data = resp.data;
                    if (data && data.code == 200) {
                        let Data = data.data;
                        this.tableData = Data.data;
                        this.Total = Data.total;
                    } else {
                        console.log(resp)
                    }
                });
            },
            getResourcePermissionsByRoleIds() {//加载已授权的页面元素
                let param = {
                    id: this.roleId,
                    resourceType: '1'
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/role/getRoleResourceByRoleIds", param, {}
                ).then(resp => {
                    let data = resp.data;
                    this.$nextTick(() => {
                        this.tableData.forEach(row => {
                            if (data.data.indexOf(row.id) >= 0) {
                                this.$refs.multipleTable.toggleRowSelection(row, true);
                            }
                        });
                    })
                })
            }
        },
        mounted() {
            this.roleId = this.$route.query.id;
            this.getSelectTreeData(); //获取菜单树
        }
    }

</script>

<style>
    .message-title {
        cursor: pointer;
    }

    .handle-row {
        margin-top: 30px;
    }

    .table {
        width: 80%;
        font-size: 14px;
    }

    .demo-input-suffix {
        display: inline-block;
        width: 40%;
        text-align: right;
    }
</style>

