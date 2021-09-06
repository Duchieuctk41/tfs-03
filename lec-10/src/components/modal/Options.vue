<template>
  <div class="DemoOptions__Wrapper">
    <div class="DemoOptions__Icon">icon</div>
    <div class="DemoOptions__Price">
      <span>{{ state.options[pos].name }}</span>
      <span>${{ state.options[pos].price }}/year</span>
    </div>
    <button @click="changeOption">Change</button>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'Price',
  data() {
    return {
      pos: 0,
    }
  },
  computed: mapGetters({ state: 'allModal' }),
  methods: {
    changeOption() {
      const len = this.state.options.length - 1
      this.pos === len ? (this.pos = 0) : this.pos++
    },
  },
}
</script>
<style lang="scss" scoped>
.DemoOptions__Wrapper {
  border-radius: 0.8rem;
  background-color: #edeeef;
  padding: 1.6rem;
  display: flex;
  align-items: center;
  .DemoOptions__Price {
    flex: 1;
    span {
      display: block;
    }
  }
}
</style>
