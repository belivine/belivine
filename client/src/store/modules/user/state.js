import cookie from 'vue-cookies'
const token = cookie.get("belivine"); 
const state = {
    user : null,
    token: token,
    isLoggedin : false,
};

export {state};