<template>
  <ul class="ul">
    <li
      class="li"
      v-for="number in pageNumbers"
      :key="number"
      @click="$emit('send', number)"
    >
      {{ number }}
    </li>
  </ul>
</template>
<script>
export default {
  name: 'pagination',
  emits: ['send'],
  props: {
    postsPerPage: Number,
    totalPosts: Number,
    totalPage: Number,
  },
  data() {
    return {
      pageNumbers: [],
    }
  },
  watch: {
    totalPage: function () {
      for (let i = 1; i <= this.totalPage; i++) {
        this.pageNumbers.push(i)
      }
    },
  },
}
</script>
<style scope>
.ul {
  list-style-type: none;
}
.li {
  cursor: pointer;
  padding: 5px;
  background: #acacac;
  display: inline-block;
  margin-right: 5px;
}
</style>
