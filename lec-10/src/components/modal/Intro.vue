<template>
  <div class="DemoDetail__Wrapper">
    <h2 class="Style__HeadingH2">{{ state.title }}</h2>
    <p class="DemoDetail__Description">
      {{ state.description }}
    </p>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
export default {
  name: 'Intro',
  computed: mapGetters({ state: 'allModal' }),
}
</script>
<style lang="scss" scoped>
.DemoDetail__Wrapper {
  text-align: center;
  .Style__HeadingH2 {
    margin: 0;
  }
  .DemoDetail__Description {
    line-height: 2rem;
    margin: 1.2rem 0;
  }
}
</style>
