import Vue from "vue";
import Vuex from "vuex";
import { user } from "./modules/user/index";
Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        counter : 0,
        user: null,
    },
    mutations: {
        increase:(state) => {
            state.counter++
        },
        decrease: (state)=> {
            state.counter--
        } 
    },
    actions: {
        increase: ({ commit }) => {
            commit("increase");
        },
        decrease: ({ commit }) =>{
            commit("decrease");
        }
    },
    modules: {
        user
    }
})