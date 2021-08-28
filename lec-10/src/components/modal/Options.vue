<template>
  <div class="DemoOptions__Wrapper">
    <div class="DemoOptions__Icon">icon</div>
    <div class="DemoOptions__Price">
      <span>Annual plan</span>
      <span>$59.99/year</span>
    </div>
    <a href="">Change</a>
  </div>
</template>

<script>
export default {
  name: "Price",
};
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