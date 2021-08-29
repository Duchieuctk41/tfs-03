<template>
  <div class="Modal__Demo" @click="$emit('closeModal')">
    <button>close</button>
    <div class="Demo__Wrapper" @click.stop="stopCloseModal(event)">
      <header class="Demo__Illustrate">
        <img src="@/assets/illustration-hero.svg" alt="illustrate_image" />
      </header>
      <main class="DemoMain__Wrapper">
        <div class="container">
          <Intro />
          <Options />
          <Checkout />
        </div>
      </main>
    </div>
  </div>
</template>

<script>
import Intro from "@/components/modal/Intro.vue";
import Options from "@/components/modal/Options.vue";
import Checkout from "@/components/modal/Checkout.vue";
export default {
  name: "Modal",
  components: {
    Intro,
    Options,
    Checkout,
  },
  methods: {
    stopCloseModal(event) {
      event.stopPropagation();
    },
  },
};
</script>

<style lang="scss" scoped>
.Modal__Demo {
  position: fixed;
  z-index: 10;
  left: 0;
  top: 0;
  width: 100%; /* Full width */
  height: 100%;
  overflow: hidden;
  background-color: rgba(58, 58, 58, 0.9);

  .Demo__Wrapper {
    position: absolute;
    width: 40rem;
    max-width: 40rem;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    border-radius: 1.6rem;
    overflow: hidden;
    background: #fff;

    .Demo__Illustrate {
      img {
        width: 100%;
        height: auto;
      }
    }
    .DemoMain__Wrapper {
      .container {
        padding: 3.2rem;
        box-sizing: border-box;
      }
    }
  }
}
</style>