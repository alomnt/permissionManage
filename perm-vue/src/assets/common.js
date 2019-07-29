/*
*公共方法
*
**/
let Common = {
    trim(val) {//去除首尾空格
        return str.replace(/(^\s*)|(\s*$)/g,"");
    },
    changeNull(val) {//null显示
        return val==null ? "":val;
    },
    formatMoney(val) {//金额保留两位小数点并以逗号分割
        if(val===undefined || val===null || isNaN(val) ) {
            return '0.00';
        }
        if(String(val).indexOf("-")>-1) {
            let num = parseFloat(Number(-val)).toFixed(2);
            let arr = num.split(".");
            let left = arr[0].split("").reverse();
            let str = "";
            for(let i=0;i<left.length;i++)
            {
                str += left[i] + ((i + 1) % 3 == 0 && (i + 1) != left.length ? "," : "");
            }
            return "-"+str.split("").reverse().join("")+'.'+arr[1];
        }else if(String(val).indexOf()===-1){
            let num = parseFloat(Number(val)).toFixed(2);
            let arr = num.split(".");
            let left = arr[0].split("").reverse();
            let str = "";
            for(let i=0;i<left.length;i++)
            {
                str += left[i] + ((i + 1) % 3 == 0 && (i + 1) != left.length ? "," : "");
            }
            return str.split("").reverse().join("")+'.'+arr[1];
        }
    },
    keepTwo(val) {//保留两位小数
        return parseFloat(Number(val)).toFixed(2);
    },
    formatDates(date) {
        let d = new Date(date);
        let month = (d.getMonth() + 1) < 10 ? '0'+(d.getMonth() + 1) : (d.getMonth() + 1);
        let day = d.getDate()<10 ? '0'+d.getDate() : d.getDate();
        let hours = d.getHours()<10 ? '0'+d.getHours() : d.getHours();
        let min = d.getMinutes()<10 ? '0'+d.getMinutes() : d.getMinutes();
        let sec = d.getSeconds()<10 ? '0'+d.getSeconds() : d.getSeconds();
        let times=d.getFullYear() + '-' + month + '-' + day + ' ' + hours + ':' + min + ':' + sec;
        return times
    },
    // formatDates(data) {
    //     let time = new Date(Date.parse(data));
    //     time.setTime(time.setHours(time.getHours() + 8));
    //     let Y = time.getFullYear() + '-';
    //     let M = this.addZero(time.getMonth() + 1) + '-';
    //     let D = this.addZero(time.getDate()) + ' ';
    //     let h = this.addZero(time.getHours()) + ':';
    //     let m = this.addZero(time.getMinutes()) + ':';
    //     let s = this.addZero(time.getSeconds());
    //     return Y + M + D + h + m + s;
    // },
    // 数字补0操作
    addZero(num) {
      return num < 10 ? '0' + num : num;
    },
    formatDate(date) {//格式化日期 xxxx-xx-xx;
        if(date==="" || date == undefined){
            return "";
        }else if(typeof date == "number" || date.length==10) {
            let today = new Date(date);
            let year = today.getFullYear();
            let month = ('0' + (today.getMonth() + 1)).slice(-2);
            let day = ('0' + today.getDate()).slice(-2);
            return `${year}-${month}-${day}`;
        }else {
            let Y = date.getFullYear();
            let M = (date.getMonth() + 1) >= 10 ? date.getMonth() + 1 : "0" + (date.getMonth() + 1);
            let D = date.getDate() >= 10 ? date.getDate() : "0" + date.getDate();
            return Y + '-' + M + '-' + D;
            // return date.getFullYear()+"-"+(date.getMonth()+1)+"-"+date.getDate();
        }
    },
    checkChassisNumber(sVIN) {  //校验vin码
        var bl = false;
        /*校验是否17位*/
        if (sVIN.length != 17) {
            // mui.toast("请重新检查车架11");
            return bl;
        }

        /*校验是否有I/O/Q*/
        if (sVIN.indexOf("I") >= 0 || sVIN.indexOf("O") >= 0 || sVIN.indexOf("Q") >= 0) {
            // mui.toast('请重新检查车架22');
            return bl;
        }

        /*通过校验位判断是否正常*/
        var Arr = new Array();
        var Brr = new Array();
        Arr['A'] = 1;
        Arr['B'] = 2;
        Arr['C'] = 3;
        Arr['D'] = 4;
        Arr['E'] = 5;
        Arr['F'] = 6;
        Arr['G'] = 7;
        Arr['H'] = 8;
        Arr['J'] = 1;
        Arr['K'] = 2;
        Arr['L'] = 3;
        Arr['M'] = 4;
        Arr['N'] = 5;
        Arr['P'] = 7;
        Arr['R'] = 9;
        Arr['S'] = 2;
        Arr['T'] = 3;
        Arr['U'] = 4;
        Arr['V'] = 5;
        Arr['W'] = 6;
        Arr['X'] = 7;
        Arr['Y'] = 8;
        Arr['Z'] = 9;
        Arr['1'] = 1;
        Arr['2'] = 2;
        Arr['3'] = 3;
        Arr['4'] = 4;
        Arr['5'] = 5;
        Arr['6'] = 6;
        Arr['7'] = 7;
        Arr['8'] = 8;
        Arr['9'] = 9;
        Arr['0'] = 0;
        Brr[1]=8;
        Brr[2]=7;
        Brr[3]=6;
        Brr[4]=5;
        Brr[5]=4;
        Brr[6]=3;
        Brr[7]=2;
        Brr[8]=10;
        Brr[9]=0;
        Brr[10]=9;
        Brr[11]=8;
        Brr[12]=7;
        Brr[13]=6;
        Brr[14]=5;
        Brr[15]=4;
        Brr[16]=3;
        Brr[17]=2;

        var sKYZF="ABCDEFGHJKLMNPRSTUVWXYZ1234567890";
        var sJYW ='';
        var blKYZF = false;

        if (sVIN.length == 17) {
            var iJQS = 0;
            var intTemp = 0;
            var ht = Arr;
            var htZM = Brr;
            try {
                for (var i = 0; i <sVIN.length; i++) {
                    if (sKYZF.indexOf(sVIN.substr(i, 1)) != -1)
                    {
                        blKYZF = true;
                        iJQS = iJQS + parseInt(ht[sVIN.substr(i, 1)]) * parseInt(htZM[(i + 1)]);
                    }
                    else {
                        blKYZF = false;
                        break;
                    }
                }
                if (blKYZF) {
                    intTemp = iJQS%11;
                    if (intTemp == 10)
                    {
                        sJYW = "X";
                    }
                    else {
                        sJYW = intTemp.toString();
                    }
                    if (sJYW == sVIN.substr(8, 1)) bl = true;
                }
                else {
                    bl = false;
                }
            }
            catch(err) {
                bl = false;
            }
        }
        // mui.toast("正则校验通过");
        return bl;
    },
    AmountMoneyToWY(money) {  //金额转换 元转万元
        if(money == undefined || money == "") {
            return "0";
        }else if(money == 0 || typeof money === "object") {
            return "0";
        }else {
            return (money/10000).toFixed(2);
        }
    },
    AmountMoneyToY(money) {  //金额转换 万元转元
        if(money == undefined || money == "") {
            return "0";
        }else if(money == "0") {
            return "0";
        }else {
            return money*10000;
        }
    },
    getdate(s) {//日期转换
        var now = new Date(),
            y = now.getFullYear(),
            m = now.getMonth() + 1,
            d = now.getDate();
        return y + "-" + (m < 10 ? "0" + m : m) + "-" + (d < 10 ? "0" + d : d);
        //return y + "-" + (m < 10 ? "0" + m : m) + "-" + (d < 10 ? "0" + d : d) + " " + now.toTimeString().substr(0, 8);
    },
    downLoadImgArr(imgArr){/*imgArr为图片批量数组格式[{url:''}]*/
        imgArr.map(function(i){
            var a = document.createElement('a');
            a.setAttribute('download','');
            a.href = i.url;
            document.body.appendChild(a);
            a.click();
        })
    },
    /*身份证校验*/
    IdentityCodeValid(code) {
        var city={11:"北京",12:"天津",13:"河北",14:"山西",15:"内蒙古",21:"辽宁",22:"吉林",23:"黑龙江 ",31:"上海",32:"江苏",33:"浙江",34:"安徽",35:"福建",36:"江西",37:"山东",41:"河南",42:"湖北 ",43:"湖南",44:"广东",45:"广西",46:"海南",50:"重庆",51:"四川",52:"贵州",53:"云南",54:"西藏 ",61:"陕西",62:"甘肃",63:"青海",64:"宁夏",65:"新疆",71:"台湾",81:"香港",82:"澳门",91:"国外 "};
        var tip = "";
        var pass= true;
        if(!code || !/^\d{6}(18|19|20)?\d{2}(0[1-9]|1[12])(0[1-9]|[12]\d|3[01])\d{3}(\d|X)$/i.test(code)){
            tip = "身份证号格式错误";
            pass = false;
        }else if(!city[code.substr(0,2)]){
            tip = "地址编码错误";
            pass = false;
        }else{
            //18位身份证需要验证最后一位校验位
            if(code.length == 18){
                code = code.split('');
                //∑(ai×Wi)(mod 11)
                //加权因子
                var factor = [ 7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2 ];
                //校验位
                var parity = [ 1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2 ];
                var sum = 0;
                var ai = 0;
                var wi = 0;
                for (var i = 0; i < 17; i++)
                {
                    ai = code[i];
                    wi = factor[i];
                    sum += ai * wi;
                }
                var last = parity[sum % 11];
                if(parity[sum % 11] != code[17]){
                    tip = "校验位错误";
                    pass =false;
                }
            }
        }
        if(!pass) alert(tip);
        return pass;
    },
    //权限校验
    authority(h,name,callback) {
        return h('a', {
            style: {
                marginRight: '12px',
                fontSize: '12px'
            },
            on: {
                click: () => {
                    callback();
                }
            }
        }, name);
    },
    commonSlice(str,n){
        var str=str+'';
        if(str == 'undefined' || str == "" || str=='null') {
            return "--";
        }else if(str.indexOf('00:00:00') != -1){
            return str.replace('00:00:00','')
        }else{
            if(n){
                return str.slice(0,n)
            }else{
                return str;
            }

        }
    },
    /*15位的身份证编码首先把出生年扩展为4位，简单的就是增加一个19或18,这样就包含了所有1800-1999年出生的人;
    正则表达式:
    出生日期1800-2099  (18|19|20)?\d{2}(0[1-9]|1[12])(0[1-9]|[12]\d|3[01])
    身份证正则表达式 /^\d{6}(18|19|20)?\d{2}(0[1-9]|1[12])(0[1-9]|[12]\d|3[01])\d{3}(\d|X)$/i            
    15位校验规则 6位地址编码+6位出生日期+3位顺序号
    18位校验规则 6位地址编码+8位出生日期+3位顺序号+1位校验位

    校验位规则     公式:∑(ai×Wi)(mod 11)……………………………………(1)
                公式(1)中： 
                i----表示号码字符从由至左包括校验码在内的位置序号； 
                ai----表示第i位置上的号码字符值； 
                Wi----示第i位置上的加权因子，其数值依据公式Wi=2^(n-1）(mod 11)计算得出。
                i 18 17 16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1
                Wi 7 9 10 5 8 4 2 1 6 3 7 9 10 5 8 4 2 1

        身份证号合法性验证 
        支持15位和18位身份证号
        支持地址编码、出生日期、校验位验证*/
};
export default Common;
// module.exports = common;




// export var formatDate = function (date) {
// 	let today = new Date(date)
// 	let year = today.getFullYear()
// 	let month = ('0' + (today.getMonth() + 1)).slice(-2)
// 	let day = ('0' + today.getDate()).slice(-2)
// 	let hour = today.getHours()
// 	let minute = today.getMinutes()
// 	let second = today.getSeconds()
// 	return `${year}-${month}-${day} ${hour}:${minute}:${second}`
// };
// export var validators = {
//     required(val) {
//         return /^[^ ]+$/.test(val);
//     }
// }