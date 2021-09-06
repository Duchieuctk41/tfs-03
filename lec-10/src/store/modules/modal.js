import { v4 as uuidv4 } from 'uuid'

const state = {
  title: 'Order summary',
  description:
    'You can now listen to millions of songs. audioblocks, and podcasts on any device anywhere you like!',
  options: [
    { name: 'Annual Plan', price: '59.99' },
    { name: 'Pro Plan', price: '99.99' },
  ],
  notif: [],
}

const getters = {
  allModal: state => state,
}

const actions = {
  addNotif(context, content) {
    let id = uuidv4()
    console.log(typeof id)
    context.commit('ADD_NOTIF', { content, id })
    setTimeout(() => {
      context.commit('DELETE_NOTIF', id)
    }, 2000)
  },
  deleteNotif({ commit }, id) {
    commit('DELETE_NOTIF', id)
  },
}

const mutations = {
  ADD_NOTIF: (state, { content, id }) => {
    console.log(content)
    state.notif.push({
      id: id,
      content: content,
    })
  },
  DELETE_NOTIF: (state, id) => {
    let tmp = state.notif.filter(val => val.id !== id)
    state.notif = [...tmp]
  },
}

export default {
  state,
  getters,
  actions,
  mutations,
}
