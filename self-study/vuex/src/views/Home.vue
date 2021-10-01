<template>
  <div class="home">
    <ul>
      <li v-for="post in posts" :key="post.id">{{ post.title }}</li>
    </ul>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import { FETCH_POSTS } from '@/store/actions.type'
export default {
  name: 'Home',
  computed: {
    ...mapState({
      posts: state => state.home.posts,
    }),
  },
  async beforeCreate() {
    console.log('fetch2')
    await this.$store.dispatch(FETCH_POSTS)
  },
}
</script>
