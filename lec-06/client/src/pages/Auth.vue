<template>
  <div>
    <h2>Đăng nhập</h2>
    <form @submit.prevent="submit">
      <input type="email" v-model="email" placeholder="example@gmail.com" />
      <input type="password" v-model="password" placeholder="password" />
      <button type="submit">login</button>
    </form>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "Auth",
  data() {
    return {
      email: "",
      password: "",
    };
  },
  methods: {
    submit: async function () {
      const d = JSON.stringify({
        email: this.email,
        password: this.password,
      })
      console.log(d)
      const { data } = await axios.post("http://localhost:3000/api/login", d);
      console.log(data);
    },
  },
};
</script>