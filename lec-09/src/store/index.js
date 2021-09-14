import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

function randomId() {
  return Math.random().toString().substr(2, 10)
}

export default new Vuex.Store({
  state: {
    todos: [],
    loading: false,
  },
  mutations: {
    ADD_TODO(state, newTodo) {
      state.todos.push(newTodo)
    },
    REMOVE_TODO(state, todo) {
      state.todos.splice(state.todos.indexOf(todo), 1)
      // state.todos = state.todos.filter(t => t.id !== todo.id)
    },
    SET_TODOS(state, todos) {
      state.todos = todos
    },
    SET_LOADING(state, loading) {
      state.loading = loading
    },
    UPDATE_TODO(state, { todo, payload }) {
      const todos = state.todos
      todos.splice(
        todos.indexOf(todo),
        1,
        Object.assign({}, todo, { ...payload }),
      )
    },
  },
  actions: {
    async loadTodos({ commit }) {
      commit('SET_LOADING', true)
      setTimeout(async () => {
        const response = await Vue.axios.get('/todos')
        const todos = response.data
        commit('SET_TODOS', todos)
        commit('SET_LOADING', false)
      }, 500)
    },
    async addTodo({ commit }, newTodo) {
      if (!newTodo) {
        return
      }
      const todo = {
        id: randomId(),
        title: newTodo,
        completed: false,
      }
      await Vue.axios.post('todos', todo)
      commit('ADD_TODO', todo)
    },
    async removeTodo({ commit }, todo) {
      await Vue.axios.delete(`/todos/${todo.id}`)
      commit('REMOVE_TODO', todo)
    },
    async updateTodo({ commit }, { todo, payload }) {
      Vue.axios.patch(`/todos/${todo.id}`, { ...payload })
      commit('UPDATE_TODO', { todo: payload })
    },
  },
  modules: {},
})
