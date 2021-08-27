<template>
  <section class="FQA__Section">
    <div class="container">
      <h2 class="Heading_StyleH2">Challenge FAQS</h2>
      <div class="FAQ__Question">
        <details class="Question__Details" v-for="val in getState.faq" :key="getState.faq[val]">
          <summary class="Question__Summary">
            {{val.sum}}
          </summary>
          <div class="Question__Answer">
            <div class="Spaser__Answer">
              {{val.desc}}
            </div>
          </div>
        </details>
      </div>
    </div>
  </section>
</template>
<script>
import {mapGetters} from "vuex";
export default {
  name: "FAQSection",
  computed: mapGetters({getState: "allContent"})
};
</script>

<style lang="scss" scoped>
.FQA__Section {
  padding: 5rem 0px 10rem;
  .container {
    width: 95%;
    max-width: 1200px;
    margin: 0px auto;
    .Heading_StyleH2 {
      margin-bottom: 3rem;
      text-align: center;
    }
    .FAQ__Question {
      display: flex;
      flex-direction: column;
      max-width: 70rem;
      margin: 0px auto;
      .Question__Details {
        margin-bottom: 1.5rem;
        border: 1px solid rgb(219, 229, 230);
        border-radius: 1rem;
        background: rgb(255, 255, 255);
        font-family: Barlow, sans-serif;
        font-weight: 600;
        font-size: 1.8rem;
        line-height: 1.2;
        .Question__Summary {
          position: relative;
          padding: 2rem 7rem 2rem 3rem;
          cursor: pointer;
          line-height: 1.5;
          list-style: none;
        }
        .Question__Answer {
          border-top: 1px solid rgb(219, 229, 230);
          padding: 2rem 3rem;
          color: rgb(115, 115, 115);
          font-family: Heebo, sans-serif;
          font-weight: 400;
          font-size: 1.5rem;
          line-height: 1.6;
         
        }
      }
    }
  }
}
</style>