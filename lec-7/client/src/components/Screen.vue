<template>
  <div class="screen">
    <h4 id="inputscr">{{message.input}}</h4>
    <h3 id="ouputscr">{{message.output}}</h3>
  </div>
</template>
<script>
export default {
  name: "Screen",
  props: ['message'],
}
</script>

<style scoped>
.screen {
  width: 100%;
  padding: 0 5px;
  box-sizing: border-box;
  margin-bottom: 15px;
}

h4 {
  text-align: end;
  font-size: 1.5rem;
  margin: 5px 5px 0 0;
}
h3 {
  font-size: 1.7rem;
  margin: 0px 0px 10px 5px;
}

</style>

