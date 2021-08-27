<template>
  <section class="Community__Section">
    <div class="container">
      <img src="@/assets/logo-slack.svg" alt="Slack" class="Community__Image" />
      <h2 class="Heading_StyleH2">Join our slack community</h2>
      <p class="Community__Strapline">
        Join over 50,000 people taking the challenges, talking about their code,
        helping each other, and chatting about all things front-end!
      </p>
      <div>
        <form action="" class="CommunityForm__Form">
          <div class="CommunityForm__InputWrapper">
            <label for="slack-email" class="sr-only"
              >Please enter your email address to sign up for the flack
              community
            </label>
            <input
              id="slack-email"
              type="text"
              placeholder="example@gmail.com"
            />
          </div>
          <button class="CTA__StyledLink CommunityForm__Button">
            Request invite
          </button>
        </form>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: "CommunitySection",
};
</script>

<style lang="scss" scoped>
.Community__Section {
  position: relative;
  border-radius: 1rem;
  padding: 6rem 0px 20rem;
  background-color: rgb(62, 84, 163);
  color: rgb(255, 255, 255);
  text-align: center;
  overflow: hidden;
  .container {
    width: 95%;
    max-width: 1200px;
    margin: 0px auto;
    img {
      max-width: 100%;
      height: auto;
    }
    .Community__Image {
      display: inline-block;
      max-width: 20rem;
      margin-bottom: 2rem;
    }
    .Heading_StyleH2 {
      font-family: Heebo, sans-serif;
      font-size: 1.5rem;
      font-weight: 900;
      letter-spacing: 3px;
      text-transform: uppercase;
      margin-bottom: 4rem;
    }
    .Community__Strapline {
      margin: 0px auto 5rem;
      font-family: Barlow, sans-serif;
      font-size: 2.8rem;
      width: 90%;
    }

    .CommunityForm__Form {
      width: 65%;
      display: flex;
      -webkit-box-pack: justify;
      justify-content: space-between;
      margin: 0 auto;
      .CommunityForm__InputWrapper {
        flex-basis: 65%;
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
        input {
          display: block;
          width: 100%;
          border: 1px solid rgb(255, 255, 255);
          padding: 1.3rem;
          background-color: transparent;
          color: rgb(255, 255, 255);
          font-size: 2rem;
        }
      }
      .CTA__StyledLink {
        flex-basis: 30%;
        border: 2px solid rgb(255, 255, 255);
        background-color: rgb(255, 255, 255);
        color: rgb(62, 84, 163);
        cursor: pointer;
      }
      .CommunityForm__Button {
        display: inline-flex;
        -webkit-box-align: center;
        align-items: center;
        -webkit-box-pack: center;
        justify-content: center;
        border: 2px solid;
        border-radius: 10rem;
        padding: 1.2rem 3rem;
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
      }
    }
  }
}
</style>