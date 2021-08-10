<template>
  <router-link class="event-link" :to="{ name: 'EventDetails', params: { id: event.id }}">
    <div class="event-card">
      <span>@ {{ event.time }} on {{ event.date }}</span>
      <h4>{{ event.title }}</h4>
    </div>
  </router-link>
</template>

<script>
export default {
  name: 'EventCard',
  props: {
    event: {
      type: Object,
      required: true
    }
  },
}
</script>

<style scoped>
.event-card {
  border: 1px solid;
  padding: 16px;
  margin-top: 16px;
}

.event-link {
  color: #2c3e50;
  text-decoration: none;
}
</style>
