<template>
  <section class="Show__Section">
    <div class="DesignPreview__Wrapper">
      <div class="DesignPreview_Preview">
        <button>
          <img
            :src="require(`@/assets/${imgURL}`)"
            alt="Desktop design for order summary"
          />
        </button>
      </div>
      <ul>
        <li
          class="DesignPreview__PreviewTab"
          v-for="val in getState.image"
          :key="val.name"
        >
          <button @click="setImage(val.img)">{{ val.name }}</button>
        </li>
      </ul>
    </div>
    <div class="ChallengeDetails__Wrapper">
      <div class="ChallengeDetails__Tags">
        <span class="ChallengeDetails__PriceTag">{{ getState.price }}</span>
        <ul class="Languages__List">
          <li v-for="lang in getState.lang" :key="lang" class="Languages__Item">
            {{ lang }}
          </li>
        </ul>
        <div class="Level_Wrapper">
          <span class="Level__Number">{{ getState.level.number }}</span>
          <span class="Level_Name">{{ getState.level.name }}</span>
        </div>
      </div>
      <h2 class="Heading_StyleH1">{{ getState.title }}</h2>
      <p class="ChallengeDetails__Description">
        A perfect project for newbies who are starting to build confidence with
        layouts!
      </p>
      <div class="ChallengeDetails__Assets">
        <h2 class="Heading__StyleH2">ASSETS PROVIDED</h2>
        <ul class="ChallengeDetails__List">
          <!-- <li class="ChallengeDetails__ListItem">
            Sketch and Figma design file access -
            <a href="#">
              Unlock access with <span class="uppercased">PRO</span></a
            >
          </li> -->
          <li
            v-for="val in getState.asset"
            :key="val.id"
            class="ChallengeDetails__ListItem"
          >
            {{ val.name }}
          </li>
        </ul>
      </div>
      <div class="ChallengeDetails__CTASection">
        <div class="ChallengeDetails__Info">
          <div>icon</div>
          <span>It looks like you've already started this challenge.</span>
        </div>
        <a href="#" class="CTA__StyledLink">Visit challenge hub</a>
      </div>
    </div>
  </section>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "ShowSection",
  data() {
    return {
      imgURL: "desktop.jpg",
    };
  },
  computed: mapGetters({ getState: "allContent" }),
  methods: {
    setImage(img) {
      return this.imgURL = img
    }
  }
};
</script>

<style lang="scss" scoped>
.Show__Section {
  grid-template-columns: 6fr 5fr;
  display: grid;
  gap: 5rem;

  .DesignPreview__Wrapper {
    .DesignPreview_Preview {
      height: 35vw;
      max-height: 48rem;
      margin-bottom: 0px;
      button {
        width: 100%;
        height: 100%;
        img {
          object-fit: cover;
          max-width: 100%;
          height: auto;
        }
      }
    }

    ul {
      padding: 0;
      margin: 1.1rem 0px;
      .DesignPreview__PreviewTab {
        display: inline-block;
        margin-right: 0.5rem;
      }
    }
  }
  .ChallengeDetails__Wrapper {
    .ChallengeDetails__Tags {
      display: flex;
      flex-wrap: wrap;
      margin-bottom: 1.5rem;

      .ChallengeDetails__PriceTag {
        margin-right: 1rem;
      }
      .Languages__List {
        margin: 0;
        margin-right: auto;
        .Languages__Item {
          display: inline-block;
        }
      }
    }
    .Heading_StyleH1 {
      max-width: 80%;
      margin: 0.5rem 0px 3rem;
      font-size: 3rem;
      line-height: 1.3;
      font-weight: 500;
    }
    .ChallengeDetails__Description {
      color: #737373;
      font-size: 1.4rem;
      line-height: 1.6;
    }
    .ChallengeDetails__Assets {
      background-color: rgb(240, 244, 244);
      margin: 2.5rem 0px 0px;
      border-radius: 0.75rem;
      padding: 2.5rem;
      .Heading__StyleH2 {
        font-family: Heebo, sans-serif;
        font-size: 1.5rem;
        font-weight: 900;
        letter-spacing: 3px;
        text-transform: uppercase;
        margin: 0;
      }
      .ChallengeDetails__List {
        margin: 1.7rem 0px 0px;
        padding-left: 1.5rem;
        .ChallengeDetails__ListItem {
          position: relative;
          padding-left: 1rem;
          font-size: 1.4rem;
          margin-bottom: 0.5rem;
        }
      }
    }
    .ChallengeDetails__CTASection {
      display: flex;
      flex-wrap: wrap;
      align-items: center;
      justify-content: space-between;
      margin: 2.5rem 2.5rem 0 0;
      .ChallengeDetails__Info {
        display: flex;
        -webkit-box-align: center;
        align-items: center;
        margin-right: 2.5rem;
        color: rgb(115, 115, 115);
        font-size: 1.6rem;
        font-weight: 500;
        line-height: 1.5;
      }
      .CTA__StyledLink {
        white-space: nowrap;
        margin-top: 2.5rem;
      }
    }
  }
}
</style>
