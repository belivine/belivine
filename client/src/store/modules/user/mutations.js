import cookie from 'vue-cookies'

const mutations = {
    set_token (state, token) {
        state.token = token
      },
    initUser (state, user) {
        state.isLoggedin = true
        console.log("cek isi user", user);
        state.user = { ...user }
    },
    logout_user (state) {
        state.isLoggedin = ''
        state.token = '' && cookie.remove('belivine')
    }
};

export {mutations}