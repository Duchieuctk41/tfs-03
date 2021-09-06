<template>
  <div class="NavSecond__Wrapper">
    <div class="wide-container">
      <h1 class="NavSecond__Heading">Challenge</h1>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NavSecond',
}
</script>

<style lang="scss" scoped>
h1 {
  margin: 0;
}
.NavSecond__Wrapper {
  position: relative;
  margin-bottom: 3rem;
  border-bottom: 1px solid rgb(219, 229, 230);
  border-top: 1px solid #dbe5e6;
  .wide-container {
    margin: 0 auto;
    width: 95%;
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 54px;

    .NavSecond__Heading {
      padding: 1.5rem 3rem;
      border-right: 1px solid rgb(219, 229, 230);
      border-left: 1px solid rgb(219, 229, 230);
      font-weight: 700;
      letter-spacing: 1px;
      text-transform: uppercase;
      font-size: 1.4rem;
      line-height: 2.3rem;
    }
  }
}
</style>
