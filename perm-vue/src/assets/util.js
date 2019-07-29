export default {
    GetSessionStorage: function (index, data) {
        if (data == '没有登录') {
            index.$router.push({path: '/login'});
        } else {
            index.$message.error(data);
        }
    },
}
