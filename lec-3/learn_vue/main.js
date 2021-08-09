var eventBus = new Vue()
Vue.component('product', {
  props: {
    premium: {
      type: Boolean,
      required: true
    }
  },
  template: `
  <div class="product">
      <div class="product-image">
        <img :src="image" alt="socks">
      </div>
      <div class="product-info">
        <h1>{{ title }}</h1>
        <p v-if="inSocks">In Socks</p>
        <p v-else>no options</p>
        <p>Shipping: {{ shipping }} </p>
        <ul>
          <li v-for="detail in details">{{ detail }}</li>
        </ul>
        <div v-for="(variant, index) in variants" 
        :key="variant.key"
        class="color-box"
        :style="{backgroundColor: variant.variantColor }"
        @mouseover="updateProduct(index)">
        </div>
      </div>
      <button v-on:click="addToCart" 
      :disabled="!inSocks"
      :class="{ disabledButton: !inSocks }"
      >Add to Cart</button>
      <product-tabs :reviews="reviews"></product-tabs>
      </div>
  `,
  data() {
    return {
      brand: 'Vue Mastery',
      product: 'Socks',
      selectedVariant: 0,
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
          variantColor: 'blue',
          variantImage: './assets/vmSocks-gray.jpg',
          variantQuantity: 0
        },
      ],
      reviews: []
    }
  },
  methods: {
    addToCart() {
      this.$emit('add-to-cart')
    },
    updateProduct(index) {
      this.selectedVariant = index
    }
  },
  computed: {
    title() {
      return this.brand + ' ' + this.product
    },
    image() {
      return this.variants[this.selectedVariant].variantImage
    },
    inSocks() {
      return this.variants[this.selectedVariant].variantQuantity
    },
    shipping() {
      if (this.premium) {
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


Vue.component('product-review', {
  template: `
  <form @submit.prevent="onSubmit">
  <p v-if="errors.length"></p>
  <b>Please correct the following error(s):</b>
  <ul>
    <li v-for="error in errors">{{error}}</li>
  </ul>
  <input id="name" v-model="name">
  <input id="review" v-model="review">
  <select id="rating" v-model.number="rating">
    <option>1</option>
    <option>2</option>
    <option>3</option>
  </select>
  <input type="submit" value="Submit">
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
      if (this.name && this.review && this.rating) {
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
        if (!this.name) this.errors.push("Name required")
        if (!this.review) this.errors.push("review required")
        if (!this.rating) this.errors.push("rating required")
      }
    }
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
      <span class="tab"
        v-for="(tab, index) in tabs" 
        :class="{ activeTab: selectedTab === tab}"
        :key="index"
        @click="selectedTab = tab">
      {{ tab }}</span>

      <div v-show="selectedTab === 'Reviews'">
      <p v-if="!reviews.length">There are no reviews yet.</p>
      <ul>
        <li v-for="review in reviews">
        <p>{{review.name}}</p>
        <p>{{review.review}}</p>
        <p>{{review.rating}}</p>
        </li>
      </ul>
      </div>
      <product-review v-show="selectedTab === 'Make a Review'"
      ></product-review>
    
    </div>
  `,
  data() {
    return {
      tabs: ['Reviews', 'Make a Review'],
      selectedTab: 'Reviews'
    }
  }
})

var app = new Vue({
  el: '#app',
  data: {
    premium: false,
    cart: 0
  },
  methods: {
    updateCart() {
      this.cart += 1
    }
  }
})
