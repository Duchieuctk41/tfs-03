import { FETCH_POSTS } from '@/store/actions.type'
import { SET_POSTS, ADD_MSG } from './mutations.type'
import HomeService from '../common/home.service'

const state = {
  posts: [],
  msg: '',
}

const actions = {
  async [FETCH_POSTS]({ commit }) {
    try {
      console.log('fetch1')
      const res = await HomeService.fetchPosts()
      commit(SET_POSTS, res.data)
    } catch (err) {
      commit(ADD_MSG, err)
    }
  },
}

const mutations = {
  [SET_POSTS](state, posts) {
    state.posts = posts
  },
  [ADD_MSG](state, err) {
    state.msg = err
  },
}

export default {
  state,
  actions,
  mutations,
}
