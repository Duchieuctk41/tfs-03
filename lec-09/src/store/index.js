import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

function randomId() {
  return Math.random().toString().substr(2, 10)
}

export default new Vuex.Store({
  state: {
    todos: [],
  },
  mutations: {
    ADD_TODO(state, newTodo) {
      state.todos.push(newTodo)
    },
    REMOVE_TODO(state, todo) {
      state.todos.splice(state.todos.indexOf(todo), 1)
      // state.todos = state.todos.filter(t => t.id !== todo.id)
    },
  },
  actions: {
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
  },
  modules: {},
})
