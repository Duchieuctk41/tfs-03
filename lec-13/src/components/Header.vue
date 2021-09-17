<template>
  <nav class="navbar navbar-light">
    <div class="container">
      <a class="navbar-brand" href="index.html">conduit</a>
      <ul class="nav navbar-nav pull-xs-right">
        <li class="nav-item">
          <!-- Add "active" class when you're on that page" -->
          <router-link class="nav-link active" :to="{ name: 'home' }">
            Home
          </router-link>
          <router-link class="nav-link active" :to="{ name: 'login' }">
            Sign in
          </router-link>
        </li>
        <li class="nav-item">
          <router-link class="nav-link" :to="{ name: 'register' }">
            <i class="ion-compose"></i>
            &nbsp;Sign up
          </router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>
