<template>
  <section class="Newsletter__Section">
    <div class="Newsletter__Content">
      <div class="Newsletter__TextContent">
        <h2 class="Heading__StyleH2">SUBSCRIBE TO OUR NEWSLETTER</h2>
        <p class="Newsletter__Strapline">
          Stay up-to-date with new challenges, featured solutions, selected
          articles and Frontend Mentor latest news
        </p>
        <div>
          <form action="" class="NewsletterForm__Form">
            <div class="NewsletterForm__Checkbox">
              <input type="text" class="sr-only" />
              <label for=""
                >I am happy for Frontend Mentor to contact me by email</label
              >
            </div>
            <div class="NewsletterForm__Email">
              <label for="" class="newsletter-email"
                >Please enter your email address to subscribe to the newsletter
              </label>
              <input type="text" placeholder="email@example.com" />
              <button class="CTA__StyledLink NewsletterForm__Button">
                Subscribe
              </button>
            </div>
          </form>
        </div>
      </div>
      <div class="Newsletter__ImageWrapper">
        <img src="@/assets/image.svg" alt="illustratate" />
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: "NewsletterSection",
};
</script>

<style lang="scss" scoped>
.Newsletter__Section {
  padding: 10rem 0px 15rem;
  background-repeat: no-repeat;
  background-position: 150% bottom;
  .Newsletter__Content {
    width: 95%;
    max-width: 80%;
    margin: 0px auto;
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 5rem;
    -webkit-box-align: center;
    align-items: center;
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 5rem;
    -webkit-box-align: center;
    align-items: center;
    .Newsletter__TextContent {
      .Heading__StyleH2 {
        font-family: Heebo, sans-serif;
        font-size: 1.5rem;
        font-weight: 900;
        letter-spacing: 3px;
        text-transform: uppercase;
        margin: 0;
      }
      .Newsletter__Strapline {
        font-family: Barlow, sans-serif;
        font-size: 2.8rem;
        margin: 2rem 0px;
      }

      .NewsletterForm__Form {
        .NewsletterForm__Checkbox {
          margin-bottom: 2rem;
          .sr-only {
            border: 0;
            clip: rect(0, 0, 0, 0);
            height: 1px;
            margin: -1px;
            overflow: hidden;
            padding: 0;
            position: absolute;
            width: 1px;
          }
        }
        .NewsletterForm__Email {
          display: flex;
          -webkit-box-pack: justify;
          justify-content: space-between;
            margin-bottom: 2rem;

          label {
            border: 0;
            clip: rect(0, 0, 0, 0);
            height: 1px;
            margin: -1px;
            overflow: hidden;
            padding: 0;
            position: absolute;
            width: 1px;
          }
          input {
            flex-basis: 60%;
            margin-bottom: 0px;
            display: block;
            width: 100%;
            border: 1px solid rgb(28, 32, 34);
            padding: 1.3rem;
            background-color: rgb(250, 250, 250);
            font-size: 2rem;
          }
          .CTA__StyledLink {
            flex-basis: 30%;
            cursor: pointer;
            font-size: 1.6rem;
          }
          .NewsletterForm__Button {
            display: inline-flex;
            -webkit-box-align: center;
            align-items: center;
            -webkit-box-pack: center;
            justify-content: center;
            border: 2px solid rgb(237, 44, 73);
            border-radius: 10rem;
            padding: 1.2rem 3rem;
            background-color: rgb(237, 44, 73);
            font-family: Barlow, sans-serif;
            font-size: 1.4rem;
            font-style: italic;
            font-weight: 700;
            letter-spacing: 0.1rem;
            line-height: 1.2;
            text-align: center;
            text-decoration: none;
            text-transform: uppercase;
            backface-visibility: hidden;
            transition: transform 0.3s ease 0s, box-shadow 0.3s ease 0s;
            color: rgb(255, 255, 255) !important;
          }
        }
      }
    }
    .Newsletter__ImageWrapper {
      display: inline-block;
    max-width: 100%;
    overflow: hidden;
    position: relative;
    box-sizing: border-box;
    margin: 0;
    }
  }
}
</style>