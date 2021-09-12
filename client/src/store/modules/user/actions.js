import cookie from 'vue-cookies'
import { login, register } from "../../../graphql/mutations";
import { apolloClient } from "../../../vue-apollo";
import { get_user } from "@/graphql/queries";

const actions = { 
    async register ({ commit, dispatch }, authDetails) {
        try {
          const { data } = await apolloClient.mutate({ mutation: register, variables: { ...authDetails } })
          commit('set_token', data.createUser)
          cookie.set('belivine', data.createUser)
          dispatch('setUser')
        } catch (e) {
          return e.graphQLErrors[0].message;
        }
    },
    async setUser ({ commit }) {
        const { data } = await apolloClient.query({ query: get_user });
        commit('initUser', data.get_profile)
    },
    async logOut ({ commit }) {
        commit('logout_user');
    },
    async login({ commit, dispatch }, authDetails) {
        try{
            const { data } = await apolloClient.mutate({mutation: login, variables: {...authDetails}})
            commit('set_token', data.login)
            cookie.set('belivine', data.login)
            dispatch('setUser')
        } catch(e) {
            return e.graphQLErrors[0].message;
        }
    },
}

export {actions}