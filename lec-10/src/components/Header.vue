<template>
  <header class="Header__Wrapper">
    <div class="Header__Content">
      <a href="#" class="Header__Logo">
        <img src="@/assets/logo-desktop.svg" alt="logo here" />
      </a>
      <div class="Header__Nav">
        <nav>
          <ul class="MainMenu__List">
            <li class="MainMenu__Item">
              <a href="#">
                Unlock
                <span>pro</span>
              </a>
            </li>
            <li class="MainMenu__Item"><a href="#">Challenges</a></li>
            <li class="MainMenu__Item"><a href="#">Solutions</a></li>
            <li class="MainMenu__Item"><a href="#">Resources</a></li>
          </ul>
          <a href="#" class="Notification"><span>ring</span></a>
        </nav>
        <button class="ProfileMenu__Button">github</button>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: 'Header',
}
</script>

<style lang="scss" scoped>
.Header__Wrapper {
  padding: 1.5rem;
  height: 71px;
  box-sizing: border-box;
  .Header__Content {
    margin: 0 auto;
    width: 95%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    .Header__Nav {
      display: flex;
      align-items: center;
      nav {
        display: flex;
        align-items: center;
        .MainMenu__List {
          list-style-type: none;
          .MainMenu__Item {
            display: inline-block;
            border: none;
            padding: 0px 1.5rem;
          }
        }
        .Notification {
          margin: 0 1.5rem;
        }
      }
    }
  }
}
</style>
