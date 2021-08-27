<template>
  <div class="home">
    <Header />
    <NavSecond />
    <main>
      <div class="container">
        <ShowSection />
      </div>
      <div class="container">
        <BriefSection />
      </div>
      <div class="container">
        <FAQSection />
      </div>
      <div class="wide-container">
        <CommunitySection />
      </div>
      <NewsletterSection />
    </main>
    <Footer />
  </div>
</template>

<script>
// @ is an alias to /src
import Header from "@/components/Header.vue";
import NavSecond from "@/components/NavSecond.vue";
import ShowSection from "@/components/ShowSection.vue";
import BriefSection from "@/components/BriefSection.vue";
import FAQSection from "@/components/FAQSection.vue";
import CommunitySection from "@/components/CommunitySection.vue";
import NewsletterSection from "@/components/NewsletterSection.vue";
import Footer from "@/components/Footer.vue";

export default {
  name: "Home",
  components: {
    Header,
    NavSecond,
    ShowSection,
    BriefSection,
    FAQSection,
    CommunitySection,
    NewsletterSection,
    Footer,
  },
};
</script>

<style lang="scss">
.container {
  width: 95%;
  max-width: 1200px;
  margin: 0px auto;
}
.wide-container {
      width: 95%;
    max-width: 1400px;
    margin: 0px auto;
}
</style>
