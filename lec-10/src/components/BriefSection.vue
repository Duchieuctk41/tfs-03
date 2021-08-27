<template>
  <section class="Brief__Section">
    <h2 class="Heading__StyleH2">Chalenge</h2>
    <div class="Brief__Content">
      <div>
        <h3 class="Brief__Subheading">Brief</h3>
        <div class="Spacer__Wrapper">
          <p class="Text__Wrapper">
            Your challenge is to build out this order summary card component and
            get it looking as close to the design as possible.
          </p>
          <p class="Text__Wrapper">
            You can use any tools you like to help you complete the challenge.
            So if you've got something you'd like to practice, feel free to give
            it a go.
          </p>
          <p class="Text__Wrapper">Your users should be able to:</p>
          <ul class="Text__Wrapper">
            <li class="Item__Wrapper">Your users should be able to:</li>
          </ul>
          <p class="Text__Wrapper">
            Download the project and go through the README.md file. This will
            provide further details about the project and help you get set up.
          </p>
          <p class="Text__Wrapper">
            Want some support on the challenge? Join our Slack community and ask
            questions in the help channel.
          </p>
        </div>
      </div>
      <div>
        <h3 class="Brief__Subheading">Getting started</h3>
        <ol class="OrderedList__List">
          <li class="OrderedList__ListItem">
            <div class="Spacer__Wrapper">
              <p class="Text__Wrapper">Download the starter code</p>
            </div>
          </li>
          <li class="OrderedList__ListItem">
            <div class="Spacer__Wrapper">
              <p class="Text__Wrapper">
                Set up the project with version control (e.g. Git)
              </p>
            </div>
          </li>
          <li class="OrderedList__ListItem">
            <div class="Spacer__Wrapper">
              <p class="Text__Wrapper">
                Read the README.md file and have a look around the project
              </p>
            </div>
          </li>
          <li class="OrderedList__ListItem">
            <div class="Spacer__Wrapper">
              <p class="Text__Wrapper">
                Set up the project with version control (e.g. Git)
              </p>
            </div>
            Get colors, fonts etc from the style-guide.md file
          </li>
          <li class="OrderedList__ListItem">
            <div class="Spacer__Wrapper">
              <p class="Text__Wrapper">
                Set up your project/file architecture however you want
              </p>
            </div>
          </li>
          <li class="OrderedList__ListItem">
            <div class="Spacer__Wrapper">
              <p class="Text__Wrapper">Start coding!</p>
            </div>
          </li>
        </ol>
        <h3 class="Brief__Subheading">Ideas to test yourself</h3>
        <ul class="BulletedList__List">
          <li class="BulletedList__ListItem">
            <div class="Spacer__Wrapper">
              <p class="Text__Wrapper">
                Write your styles using a pre-processor, such as Sass, Less or
                Stylus
              </p>
            </div>
          </li>
          <li class="BulletedList__ListItem">
            <div class="Spacer__Wrapper">
              <p class="Text__Wrapper">
                Train your eye for detail by getting your solution as close to
                the design as you can
              </p>
            </div>
          </li>
          <li class="BulletedList__ListItem">
            <div class="Spacer__Wrapper">
              <p class="Text__Wrapper">
                Try estimating the time it will take for you to build the
                project. Then see if the time taken matches up to your estimate.
                Project estimations are a skill that is often overlooked, but is
                important for professional developers
              </p>
            </div>
          </li>
        </ul>
      </div>
    </div>
    <a href="#" class="CTA__StyledLink">Visit challenge hub</a>
  </section>
</template>

<script>
export default {
  name: "BriefSection",
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