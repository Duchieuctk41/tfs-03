Vue.component('product', {
  props: {
    premium: {
      type: Boolean,
      required: true
    }
  },
  template: `
  <div class="product">
      <div class="product-item">
        <div class="product-image">
          <img :src="image" alt="socks">
        </div>
        <div class="product-info">
          <h2>{{ title }}</h2>
          <p v-if="inStocks">Out of Stock</p>
          <p v-else>no options</p>
          <p>Shipping: {{ shipping }}</p>
          <ul>
            <li v-for="detail in details">{{detail}}</li>
          </ul>
          <div>
            <span v-for="(variant, index) in variants"
            :key="variant.variantId"
            class="color-box"
            :style="{backgroundColor: variant.variantColor}"
            @mouseover="updateProduct(index)"></span>
          </div>
          <button @click="addToCart"
          :disabled="!inStocks"
          :class="{disabledButton: !inStocks}"
          >Add to cart</button>
        </div>
      </div>
      <product-tabs :reviews="reviews"></product-tabs>
    </div>
  `,
  data() {
    return {
      brand: 'Gucci ',
      product: 'Socks',
      selectedVatiant: 0,
      details: ['80% cotton', '20% denim', 'Gender-neutral'],
      variants: [
        {
          variantId: 2234,
          variantColor: 'Green',
          variantImage: './assets/vmSocks-green.jpg',
          variantQuantity: 10
        },
        {
          variantId: 3345,
          variantColor: 'Gray',
          variantImage: './assets/vmSocks-gray.jpg',
          variantQuantity: 0
        }
      ],
      reviews: []
    }
  },
  methods: {
    updateProduct(index) {
      this.selectedVatiant = index
    },
    addToCart() {
      this.$emit('add-to-cart')
    }
  },
  computed: {
    title() {
      return this.brand + ' ' + this.product
    },
    inStocks() {
      return this.variants[this.selectedVatiant].variantQuantity
    },
    image() {
      return this.variants[this.selectedVatiant].variantImage
    },
    shipping() {
      if(this.premium) {
        return "Free"
      }
      return 2.99
    }
  },
  mounted() {
    eventBus.$on('review-submitted', productReview => {
      this.reviews.push(productReview)
    })
  }
})
Vue.component('product-tabs', {
  props: {
    reviews: {
      type: Array,
      required: true
    }
  },
  template: `
  <div>
  <span class="product-tabs"
  v-for="(tab, index) in tabs"
  :class="{activeTab: selectedTab === tab}"
  :key="index"
  @click="selectedTab = tab">{{ tab }}</span>
       <div v-show="selectedTab === 'Reviews'">
       <p v-if="!reviews.length">There are no reviews yet.</p>
        <ul class="review">
        <li v-for="review in reviews">
        <p>{{ review.name }}</p>
        <p>{{ review.review }}</p>
        <p>{{ review.rating }}</p>
        </li>
      </ul>
      </div>
      <product-review v-show="selectedTab === 'Make a Review'"></product-review>
      
  </div>
  `,
  data() {
    return {
      tabs: ['Reviews', 'Make a Review'],
      selectedTab: 'Reviews'
    }
  }
})

Vue.component('product-review', {
  template: `
  <form @submit.prevent="onSubmit">
    <p v-if="errors.length">
    <b>Please correct the following error(s):</b></p>
    <ul>
    <li v-for="error in errors">{{ error }}</li>
    </ul>
      <input type="text" v-model="name" placeholder="Name:">
      <textarea type="text"v-model="review" placeholder="Review:" />
      <select v-model.number="rating">
        <option value="1">1</option>
        <option value="2">2</option>
        <option value="3">3</option>
        <option value="4">4</option>
        <option value="5">5</option>
      </select>
      <input type="submit" value="submit">
  </form>
  `,
  data() {
    return {
      name: null,
      review: null,
      rating: null,
      errors: []
    }
  },
  methods: {
    onSubmit() {
      if(this.name && this.review && this.rating) {
        let productReview = {
          name: this.name,
          review: this.review,
          rating: this.rating
        }
        eventBus.$emit('review-submitted', productReview)
        this.name = null,
        this.review = null,
        this.rating = null
      } else {
        if(!this.name) this.errors.push("Name required")
        if(!this.review) this.errors.push("review required")
        if(!this.rating) this.errors.push("rating required")
      }
    }
  }
})

var eventBus = new Vue()
var app = new Vue({
  el: '#app',
  data: {
    cart: 0,
    premium: false
  },
  methods: {
    updateCart() {
      this.cart += 1
    }
  }
})