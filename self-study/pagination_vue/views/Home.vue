<template>
  <div class="home">
    <h1>hello world</h1>
    <Articles :current-posts="currentPosts" loading />
    <Pagination :total-page="totalPage" @send="receive" />
  </div>
</template>
<script>
import axios from "axios";
import Articles from "../components/Articles.vue";
import Pagination from "../components/Pagination.vue";
export default {
  data() {
    return {
      posts: [],
      loading: false,
      currentPage: 1,
      postsPerPage: 10,
    };
  },
  components: {
    Articles,
    Pagination,
  },
  async beforeCreate() {
    const res = await axios.get("https://jsonplaceholder.typicode.com/posts");
    this.posts = res.data;
  },
  computed: {
    indexOfLastPost() {
      return this.currentPage * this.postsPerPage;
    },
    indexOfFirsPost() {
      return this.indexOfLastPost - this.postsPerPage;
    },
    currentPosts() {
      return this.posts.slice(this.indexOfFirsPost, this.indexOfLastPost);
    },
    totalPage() {
      return Math.ceil(this.posts.length / this.postsPerPage);
    },
  },
  methods: {
    receive(number) {
      this.currentPage = number;
    },
  },
};
</script>
