<template>
    <div class=''>
        <el-dialog :title=title :visible.sync="delDialogVisible" width="30%" :append-to-body="true" :show-close="false">
            <span slot="footer">
                <el-button size="small" @click="delDialogVisible = false"> 取 消 </el-button>
                <el-button size="small" type="primary" @click="delSelect"> 确 定 </el-button>
            </span>
        </el-dialog>
        <el-card>
            <div class="ly-tree-container">
                <el-tree :data="treeData" :props="defaultProps" default-expand-all
                         :expand-on-click-node="false" :render-content="renderContent"></el-tree>
            </div>
        </el-card>
    </div>
</template>

<script>
    import './tree.scss';
    import {getDefaultContent, getEditContent} from './tree.utils.js';

    export default {
        name: 'ly-tree',
        data() {
            return {
                id: 1000,
                title: '',
                treeData: [],
                isEdit: false,
                edit_name: '',
                is_superuser: 'False',
                defaultProps: {
                    children: 'child',
                    label: 'name'
                },
                select_id: null,
                select_level: null,
                select_node: null,
                delDialogVisible: false,
                userIds: []
            }
        },
        methods: {
            //添加
            append(node, data, e) {
                e = event || window.event;
                e.stopPropagation();
                if (!this.isEdit) {
                    this.select_id = data.id;
                    this.edit_name = '';
                    const newChild = {
                        name: '',
                        level: data.level + 1,
                        isEdit: true
                    };
                    this.isEdit = true;
                    if (!data.child) {
                        this.$set(data, 'child', []);
                    }
                    data.child.unshift(newChild);
                } else {
                    this.$notify({
                        type: 'error', getDepartmentTree,
                        title: '操作提示',
                        message: '有正在编辑或添加的选项未完成！',
                        duration: 2000
                    });
                }
            },
            //删除
            remove(node, data, e) {
                e = event || window.event;
                e.stopPropagation();
                if (this.isEdit) {
                    this.$notify({
                        type: 'error',
                        title: '操作提示',
                        message: '有正在编辑或添加的选项未完成！',
                        duration: 2000
                    });
                    return;
                }
                this.userIds = [];
                let param = {
                    id: data.id
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/department/isDepartmentCanDelete", param, {}
                ).then(resp => {
                    let data = resp.data;
                    if (data && data.code == 200) {
                        if (data.data) {
                            this.title = "确认要删除此项吗？";
                            this.select_node = node;
                            this.delDialogVisible = true;
                        } else {
                            this.userIds = data.data;
                            this.title = "存在正在是使用的用户,确认要删除此项吗？";
                            this.select_node = node;
                            this.delDialogVisible = true;
                        }
                    }
                });
            },
            delSelect() {
                this.delItem(this.treeData, {id: this.select_node.data.id});
            },
            update(node, data, e) {
                e = event || window.event;
                e.stopPropagation();
                if (this.isEdit) {
                    this.$notify({
                        type: 'error',
                        title: '操作提示',
                        message: '有正在编辑或添加的选项未完成！',
                        duration: 2000
                    });
                    return;
                }
                this.select_id = data.id;
                this.select_level = data.level;
                this.edit_name = data.name;
                this.isEdit = true;
            },
            //编辑添加确定
            editMsg(data, node, e) {
                e = event || window.event;
                e.stopPropagation();
                if (this.edit_name.replace(/^\s+|\s+$/g, '')) {
                    if (!data.id) {
                        let virtualNode = node.parent;
                        let params = {
                            name: this.edit_name,
                            id: virtualNode.data.id
                        };
                        // console.log(params)
                        let addChild = this.addItem(this.treeData, params);
                        // console.log(addChild)
                        // 如果是用的真api,需要在添加的接口返回添加的节点
                        // 添加成功后，将返回的节点加入数据中，然后删除掉没有id的假节点
                        virtualNode.data.child.forEach((item, i) => {
                            if (!item.id) {
                                virtualNode.data.child.splice(i, 1);
                            }
                        });
                        return;
                    }
                    let params = {
                        name: this.edit_name,
                        id: data.id
                    };
                    this.updateItem(this.treeData, params);

                }
            },
            close(data, node, e) {
                e = event || window.event;
                e.stopPropagation();
                if (!data.id) {
                    node.parent.data.child.forEach((item, i) => {
                        if (!item.id) {
                            node.parent.data.child.splice(i, 1);
                        }
                    });
                }
                this.select_id = null;
                this.select_level = null;
                this.edit_name = data.name;
                this.isEdit = false;
            },
            nameChange(e) {
                e = event || window.event;
                e.stopPropagation();
                this.edit_name = e.target.value;
            },

            isSelect(data) {
                return data.id === this.select_id && data.level === this.select_level;
            },
            renderContent(h, {node, data}) {
                return (
                    <span class="ly-tree-node">
                        {
                            (this.isEdit === true && this.isSelect(data)) || data.isEdit
                                ? <input placeholder="名称不能为空" class="ly-edit__text" on-keyup={() => this.nameChange()} value={this.edit_name}/>
                                : <span>{data.name}</span>
                        }
                        {
                            (this.isEdit === true && this.isSelect(data)) || data.isEdit
                                ? getEditContent.call(this, h, data, node)
                                : getDefaultContent.call(this, h, data, node)
                        }
                    </span>
                )
            },
            //删除确定
            delItem(data, payload) {
                for (let i = 0; i < data.length; i++) {
                    if (data[i].id === payload.id) {
                        data.splice(i, 1);
                        this.deleteDepartment(payload);
                        break
                    }
                    if (data[i].child && data[i].child.length) {
                        this.delItem(data[i].child, payload);
                    }
                }
            },
            addItem(data, payload) {
                let addObj;
                for (let i = 0; i < data.length; i++) {
                    if (data[i].id === payload.id) {
                        addObj = {
                            id: this.id++,
                            name: payload.name,
                            level: data[i].level + 1,
                            child: []
                        };
                        data[i].child.unshift(addObj);
                        this.addDepartment(payload);
                        break
                    }
                    if (data[i].child && data[i].child.length) {
                        this.addItem(data[i].child, payload);
                    }
                }
            },
            updateItem(data, payload) {
                for (let i = 0; i < data.length; i++) {
                    if (data[i].id === payload.id) {
                        data[i].name = payload.name;
                        this.editDepartment(payload);
                        break
                    }
                    if (data[i].child && data[i].child.length) {
                        this.updateItem(data[i].child, payload);
                    }
                }
            },
            addDepartment(payload) { //添加部门
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                console.log(payload);
                let params = {
                    depName: payload.name,
                    parentId: payload.id
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/department/addDepartment", params, {}
                ).then(resp => {
                    let data = resp.data;
                    loading.close();
                    this.isEdit = false;
                    this.select_id = null;
                    this.select_level = null;
                    if (data && data.code == 200) {
                        this.$notify({
                            type: 'success',
                            title: '操作提示',
                            message: '添加成功！',
                            duration: 2000
                        });
                        this.getData();
                    } else {
                        this.$message.error(data.errorMessage);
                        this.getData();
                    }
                });
            },
            editDepartment(payload) {//编辑部门
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let params = {
                    title: payload.name,
                    id: payload.id
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/department/updateDepartment", params, {}
                ).then(resp => {
                    let data = resp.data;
                    loading.close();
                    this.isEdit = false;
                    this.select_id = null;
                    this.select_level = null;
                    if (data && data.code == 200) {
                        this.$notify({
                            type: 'success',
                            title: '操作提示',
                            message: '编辑成功！',
                            duration: 2000
                        });
                    } else {
                        this.$message.error(data.errorMessage);
                    }
                });
            },
            deleteDepartment(payload) { //删除部门
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                let params = {
                    ids: payload.id
                };
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/department/deleteDepartment", params, {}
                ).then(resp => {
                    let data = resp.data;
                    loading.close();
                    this.isEdit = false;
                    this.select_id = null;
                    this.select_level = null;
                    if (data && data.code == 200) {
                        this.delDialogVisible = false;
                        this.$notify({
                            type: 'success',
                            title: '操作提示',
                            message: '删除成功!',
                            duration: 2000
                        });
                    } else {
                        this.$message.error('删除失败');
                        this.getData();
                    }
                });
            },
            getData() {  //请求表格数据
                const loading = this.$loading({
                    lock: true,
                    text: 'Loading',
                    spinner: 'el-icon-loading',
                    background: 'rgba(0, 0, 0, 0.7)'
                });
                this.treeData = [];
                this.$axios.defaults.headers['urlencoded'] = '-1';
                this.$axios.post("/perm/department/findDepartmentTreeAll", {}
                ).then(resp => {
                    let data = resp.data;
                    loading.close();
                    if (data && data.code == 200) {
                        let Data = data.data;
                        this.treeData = [];
                        this.treeData.push(Data[0]);
                    } else {
                        console.log(resp)
                    }
                    console.log(this.treeData)
                });
            },
            handleTreeData(params, index) {
                let childrens = [];
                for (let i = 0; i < params.length; i++) {
                    let tmpData = params[i];
                    if (tmpData.children) {
                        let child = {
                            id: tmpData.id,
                            name: tmpData.title,
                            child: this.handleTreeData(tmpData.children, index + 1),
                            level: index + 1
                        };
                        childrens.push(child);
                    } else {
                        let child = {
                            id: tmpData.id,
                            name: tmpData.title,
                            level: index + 1
                        };
                        childrens.push(child);
                    }
                }
                return childrens;
            },
        },
        mounted() {
            this.getData(); //请求表格数据
        },
    }
</script>
<style scoped>

</style>