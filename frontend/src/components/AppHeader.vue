<template>
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item" href="https://filedgr.com/">
          <img src="../assets/filedgr-logo.png">
        </a>

        <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false"
           data-target="navbarBasicExample">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>
      <div id="navbarBasicExample" class="navbar-menu">

        <div class="navbar-end">
          <div class="navbar-item">
            <div v-if="!$auth.loading">
              <b-button icon-left="login" v-if="!$auth.isAuthenticated" @click="login">  Log in </b-button>
              <b-dropdown
                    v-if="$auth.isAuthenticated"
                    v-model="navigation"
                    position="is-bottom-left"
                    append-to-body
                    aria-role="menu">
                    <b-button class="navbar-item"
                        slot="trigger"
                        role="button" icon-right="menu-down">Menu</b-button>
                    <b-dropdown-item custom aria-role="menuitem">
                        <img :src="$auth.user.picture">
                        Logged as <b>{{ $auth.user.name }}</b>
                    </b-dropdown-item>
                    <hr class="dropdown-divider">
                    <b-dropdown-item @click="logout" value="logout" aria-role="menuitem" >
                        <b-icon icon="logout"></b-icon>
                        Logout
                    </b-dropdown-item>
                </b-dropdown>
            </div>
          </div>
        </div>
      </div>
    </nav>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

@Component
export default class AppHeader extends Vue {
 private isPublic: true
  logo = require('@/assets/filedgr-logo.png')

  login () {
    // eslint-disable-next-line @typescript-eslint/camelcase
    this.$auth.loginWithRedirect({ appState: { targetUrl: '/dashboard' } })
  }

  logout () {
    this.$auth.logout({
      returnTo: window.location.origin
    })
  }
}
</script>

<style scoped>

.navbar {
  background-color: #ffffff;
}

.navbar img {
  float: left;
  height: 50px;
  margin-left: 10px;
  display: inline;
}

.navbar-right {
  float: right;
  margin-right: 20px;
  vertical-align: middle;
  display: inline;
}

#nav {
  /*width: 50%;*/
  /*padding: 20px;*/
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

</style>
