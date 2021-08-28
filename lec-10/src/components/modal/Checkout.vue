<template>
  <div class="DemoCheckout__Wrapper">
    <button @click="paymentHandle" class="DemoCheckout__Button">
      Proceed to Payment
    </button>
    <button>Cancel Order</button>
    <div v-if="state.notif.length" class="DemoCheckout__Notif">
      <Notif v-for="val in state.notif" :key="val.id" :title-post="val.content"/>
    </div>
  </div>
</template>
<script>
import Notif from "@/components/modal/Notif.vue";
import { mapGetters, mapActions } from "vuex";
export default {
  name: "Checkout",
 
  components: {
    Notif,
  },
  computed: mapGetters({ state: "allModal" }),
  methods: {
    ...mapActions(["addNotif"]),
    paymentHandle() {
      this.addNotif("payment success!")
    },
  },
};
</script>
<style lang="scss" scoped>
.DemoCheckout__Wrapper {
  .DemoCheckout__Button {
    display: block;
  }
}
</style>