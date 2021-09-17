import AuthService from '../common/auth.service'
import { REGISTER } from './actions.type'
import { SET_ERRORS } from './mutations.type'

const state = {
  isAuthenticated: false,
  errors: null,
  user: {},
}

const getters = {
  errors(state) {
    if (!state.errors) {
      return []
    }
    return Object.keys(state.errors).map(key => {
      return `${key} ${state.errors[key].join(' ')}`
    })
  },
}

const actions = {
  async [REGISTER]({ commit }, credentials) {
    try {
      const response = await AuthService.register(credentials)
      console.log(response)
      console.log('vao day')
      return true
    } catch (err) {
      commit(SET_ERRORS, err.response.data.errors)
      return false
    }
  },
}

const mutations = {
  [SET_ERRORS](state, errors) {
    state.errors = errors
  },
}

export default {
  state,
  getters,
  actions,
  mutations,
}
