import Vue from "vue";
import Vuex from "vuex";
Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        counter : 0,
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

    }
})