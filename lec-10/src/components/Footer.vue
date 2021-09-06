<template>
  <footer class="Footer__Wrapper">
    <div class="Footer__Upper">
      <div class="container">
        <div class="Footer__Column">
          <h2 class="Heading__StyleH4">Frontend Mentor</h2>
          <p class="Footer__Description">
            Gain real experience of building websites and providing code
            reviews. Build your portfolio and help others achieve their goals.
          </p>
        </div>
        <div class="Footer__Column">
          <h2 class="Heading__StyleH4">Expore</h2>
          <ul>
            <li><a href="">Unlock PRO</a></li>
            <li><a href="">Challenges</a></li>
            <li><a href="">Solutions</a></li>
            <li><a href="">Resources</a></li>
          </ul>
        </div>
        <div class="Footer__Column">
          <h2 class="Heading__StyleH4">COMMUNITY</h2>
          <ul>
            <li><a href="">Contact Us</a></li>
            <li><a href="">Slack</a></li>
            <li><a href="">FAQs</a></li>
            <li><a href="">Blog</a></li>
          </ul>
        </div>
        <div class="Footer__Column">
          <ul class="Footer__SocialLinks">
            <li><a href="#">Tw</a></li>
            <li><a href="">Li</a></li>
            <li><a href="">Fa</a></li>
          </ul>
        </div>
      </div>
    </div>
    <div class="Footer__Lower">
      <div class="container">
        <span>© Frontend Mentor 2021</span>
        <ul class="Footer__List">
          <li><a href="">Terms</a></li>
          <li>
            <a href=""></a>
            Cookie Policy
          </li>
          <li>
            <a href=""></a>
            Private Policy
          </li>
          <li>
            <a href=""></a>
            License
          </li>
        </ul>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'Footer',
}
</script>

<style lang="scss" scoped>
.Footer__Wrapper {
  border-top: 1px solid rgb(28, 32, 34);
  .Footer__Upper {
    padding: 7rem 0px;
    .container {
      display: grid;
      gap: 5rem;
      grid-template-columns: 4fr 1fr 2fr 5fr;
      .Footer__Column {
        &:first-child {
          padding-right: 4rem;
        }
        .Heading__StyleH4 {
          display: block;
          margin-block-start: 0.83em;
          margin-block-end: 0.83em;
          margin-inline-start: 0px;
          margin-inline-end: 0px;
          font-weight: bold;
          font-size: 1.3rem;
          font-weight: 900;
          text-transform: uppercase;
        }
        ul {
          list-style-type: none;
          margin: 2rem 0px;
        }
        &:last-child {
          display: flex;
          -webkit-box-pack: end;
          justify-content: flex-end;
          .Footer__SocialLinks {
            li {
              display: inline-block;
              margin: 0px 0.5rem;
            }
          }
        }
      }
    }
  }
  .Footer__Lower {
    border-top: 1px solid rgb(115, 115, 115);
    padding: 3rem 0px;
    color: rgb(115, 115, 115);
    font-size: 1.2rem;
    .container {
      width: 95%;
      max-width: 1200px;
      margin: 0px auto;
      display: flex;
      -webkit-box-pack: justify;
      justify-content: space-between;
      ul {
        li {
          margin-left: 3rem;
          display: inline-block;
        }
      }
    }
  }
}
</style>
