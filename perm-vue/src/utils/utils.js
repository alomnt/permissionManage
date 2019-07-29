export const getToken = () => {
    let token = '';
    let user = JSON.parse(window.sessionStorage.getItem('user'));
    if (user) {
        token = user.token
    }
    return token
};

export const fmtToken = () => {
    let fmttoken = 'bearer ';
    let user = window.sessionStorage.getItem('user');
    if (user) {
        fmttoken += getToken()
    }
    return fmttoken
};

export const getUserId = () => {
    let userId = '';
    let user = JSON.parse(window.sessionStorage.getItem('user'));
    if (user) {
        userId = user.id
    }
    return userId
};