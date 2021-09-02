// src/plugins/vuetify.js
import '@mdi/font/css/materialdesignicons.css'
import Vue from 'vue'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css';
// import colors from 'vuetify/lib/util/colors';

Vue.use(Vuetify)

const opts = {
    icons: {
        iconfont: 'mdi',
    },
    theme: {
        themes: {
          light: {
            soft: '#f8f8fa',
          },
        },
    },
}

export default new Vuetify(opts)