<template>
   <div class="container">
    <Screen :message="msg"/>
    <Keyboard @update="updateMessage"/>
  </div>
</template>

<script>
import Screen from './components/Screen.vue'
import Keyboard from './components/Keyboard.vue'

export default {
  name: "App",
  components: {
    Screen,
    Keyboard,
  },
  data() {
    return {
      msg: {
        input: '0',
        output: 'Result'
      },
    }
  },
  methods: {
    updateMessage(val) {
      this.msg = val
    }
  }
};
</script>

<style scoped>
h1 {
  text-align: center;
}
.container {
  background-color: #000;
  border: 1px solid #afafaf;
  max-width: 390px;
  margin: 50px auto;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-bottom: 5px;
  color: #fff;
}
</style>
