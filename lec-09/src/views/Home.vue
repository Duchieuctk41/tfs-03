<template>
  <div v-if="loading">Loading...</div>
  <div v-else class="home">
    <section class="todoapp">
      <header class="header">
        <h1>Todo app</h1>
        <input
          type="text"
          class="new-todo"
          :value="newTodo"
          @change="setNewTodo"
          @keyup.enter="addTodo"
        />
      </header>
    </section>
    <section class="main" v-show="todos.length > 0">
      <ul class="todo-list">
        <li
          v-for="todo in filteredTodos"
          class="todo"
          :key="todo.id"
          :class="{
            completed: todo.completed,
            editing: editingToDo && todo.id === editingToDo.id,
          }"
        >
          <input
            v-if="editingToDo"
            class="edit"
            type="text"
            v-model="editingToDo.title"
            @keyup.enter="doneEdit(todo)"
            @keyup.esc="cancelEdit"
          />
          <div class="view">
            <input
              class="toggle"
              type="checkbox"
              :checked="todo.completed"
              @change="changeTodoComplete(todo, $event)"
            />
            <label @dblclick="editTodo(todo)">{{ todo.title }}</label>
            <button class="destroy" @click="removeTodo(todo)">X</button>
          </div>
        </li>
      </ul>
      <footer class="footer" v-show="todos.length > 0">
        <span class="todo-count">
          <strong>{{ remaining }}</strong>
          remaining
        </span>
        <ul class="filters">
          <li>
            <a
              @click.prevent="visiblity = 'all'"
              :class="{ selected: visiblity === 'all' }"
            >
              all
            </a>
          </li>
          <li>
            <a
              @click.prevent="visiblity = 'active'"
              :class="{ selected: visiblity === 'active' }"
            >
              active
            </a>
          </li>
          <li>
            <a
              @click.prevent="changeVisibility('completed')"
              :class="{ selected: visiblity === 'completed' }"
            >
              completed
            </a>
          </li>
        </ul>
      </footer>
    </section>
  </div>
</template>

<script>
const filters = {
  all: todos => {
    return todos
  },
  active: todos => {
    return todos.filter(todo => {
      return !todo.completed
    })
  },
  completed: todos => {
    return todos.filter(todo => {
      return todo.completed
    })
  },
}
export default {
  name: 'Home',
  data() {
    return {
      newTodo: '',
      visiblity: 'active',
      editingToDo: null,
    }
  },
  // filters: {
  //   pluralize: function (n) {
  //     return n === 1 ? 'item' : 'items'
  //   },
  // },
  created() {
    this.$store.dispatch('loadTodos')
  },
  computed: {
    loading() {
      return this.$store.state.loading
    },
    todos() {
      return this.$store.state.todos
    },
    filteredTodos() {
      return filters[this.visiblity](this.todos) // dữ liệu động
    },
    remaining() {
      return filters.active(this.todos).length
    },
  },
  methods: {
    setNewTodo(e) {
      this.newTodo = e.target.value
    },
    async addTodo() {
      await this.$store.dispatch('addTodo', this.newTodo)
      this.newTodo = ''
    },
    removeTodo(todo) {
      this.$store.dispatch('removeTodo', todo)
    },
    changeTodoComplete(todo, e) {
      this.$store.dispatch('updateTodo', {
        todo: todo,
        payload: {
          completed: e.target.checked,
        },
      })
    },
    changeVisibility(type) {
      this.visiblity = type
    },
    editTodo(todo) {
      this.editingToDo = { ...todo }
    },
    doneEdit(todo) {
      this.$store.dispatch('updateTodo', {
        todo: todo,
        payload: {
          title: this.editingToDo.title,
        },
      })
      this.editingToDo = null
    },
    cancelEdit() {
      this.editingToDo = null
    },
    changeAll() {
      this.$store.commit('CHANGE_ALL')
    },
  },
}
</script>
