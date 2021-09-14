<template>
  <div class="home">
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
        <li v-for="todo in todos" class="todo" :key="todo.id">
          <div class="view">
            <label>{{ todo.title }}</label>
            <button class="destroy" @click="removeTodo(todo)">X</button>
          </div>
        </li>
      </ul>
    </section>
  </div>
</template>

<script>
export default {
  name: 'Home',
  computed: {
    todos() {
      return this.$store.state.todos
    },
  },
  data() {
    return {
      newTodo: '',
    }
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
  },
}
</script>
