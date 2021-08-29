<template>
  <section class="Brief__Section">
    <h2 class="Heading__StyleH2">Chalenge</h2>
    <div class="Brief__Content">
      <div>
        <h3 class="Brief__Subheading">Brief</h3>
        <div class="Spacer__Wrapper">
          <p
            v-for="val in getState.brief"
            :key="getState.brief[val]"
            class="Text__Wrapper"
          >
            {{ val.content }}
          </p>
          <!--           
          <ul class="Text__Wrapper">
            <li class="Item__Wrapper">Your users should be able to:</li>
          </ul>
           -->
        </div>
      </div>
      <div>
        <h3 class="Brief__Subheading">Getting started</h3>
        <ol class="OrderedList__List">
          <li
            class="OrderedList__ListItem"
            v-for="val in getState.guide"
            :key="getState.guide[val]"
          >
            <div class="Spacer__Wrapper">
              <p class="Text__Wrapper">{{ val.content }}</p>
            </div>
          </li>
        </ol>
        <h3 class="Brief__Subheading">Ideas to test yourself</h3>
        <ul class="BulletedList__List">
          <li
            class="BulletedList__ListItem"
            v-for="val in getState.idea"
            :key="getState.idea[val]"
          >
            <div class="Spacer__Wrapper">
              <p class="Text__Wrapper">{{ val.content }}</p>
            </div>
          </li>
        </ul>
      </div>
    </div>
    <a @click="visitChallenHub" class="CTA__StyledLink">Visit challenge hub</a>
  </section>
</template>

<script>
import { mapGetters } from "vuex";
import {visitMixin} from "@/mixins/visitMixin";
export default {
  name: "BriefSection",
  mixins: [visitMixin],
  computed: mapGetters({ getState: "allContent" }),
};
</script>

<style lang="scss" scoped>
.Brief__Section {
  padding: 5rem 0;
  .Heading__StyleH2 {
    margin-bottom: 8rem;
    text-align: center;
  }
  .Brief__Content {
    display: grid;
    gap: 10rem;
    grid-template-columns: repeat(2, 1fr);
    .Brief__Subheading {
      font-family: Barlow, sans-serif;
      font-size: 2.8rem;
      font-weight: 700;
      margin-bottom: 2rem;
    }
    .Spacer__Wrapper {
      .Text__Wrapper {
        margin-bottom: 2rem;
        .Item__Wrapper {
          margin-bottom: 0.5rem;
        }
      }
    }
    .OrderedList__List {
      padding-left: 2rem;
      margin: 2rem 0px;
      .OrderedList__ListItem {
        position: relative;
        margin: 0.5rem 0px;
        list-style-type: decimal;
      }
    }
  }
}
</style>