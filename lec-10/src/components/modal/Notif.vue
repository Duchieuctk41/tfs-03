<template>
  <h2>{{ titlePost }}</h2>
</template>
<script>
export default {
  name: 'Notif',
  props: {
    titlePost: {
      type: String,
      default: 'title post',
    },
  },
}
</script>
