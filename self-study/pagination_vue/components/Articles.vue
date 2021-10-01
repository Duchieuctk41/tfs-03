<template>
  <ul>
    <li v-for="post in currentPosts" :key="post.id">{{ post.title }}</li>
  </ul>
</template>
<script>
export default {
  name: 'Articles',
  props: {
    currentPosts: Array,
    loading: {
      type: Boolean,
      default: false,
    },
  },
}
</script>
