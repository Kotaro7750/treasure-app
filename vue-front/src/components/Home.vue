<template>
  <div>
    <b-navbar toggleable="lg" type="light" variant="warning">
      <b-navbar-brand to="/">ラーメン二郎 Treasure店</b-navbar-brand>
      <b-navbar-toggle target="nav-collapse" right></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav v-if="state.user === null">
          <b-nav-item v-on:click="login">ログイン</b-nav-item>
        </b-navbar-nav>
        <b-navbar-nav v-else>
          <b-nav-item to="/articles/list">記事一覧</b-nav-item>
          <b-nav-item :to="{ name: 'articleCreate', params: {state: state.user } }">記事作成</b-nav-item>
          <b-nav-item v-on:click="logout">ログアウト</b-nav-item>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>

    <div v-if="state.user === null">
      <p></p>
      <img
        src="../../public/jiro.jpg"
        alt="ラーメン296"
        title="ラーメン296"
        width="30%"
        height="30%"
        class="logo offset-lg-4"
      />
      <h1 class="offset-lg-4">ラーメン二郎 is God!!</h1>
    </div>

    <div v-else>
      <div>{{state.message}}</div>

      <p style="color:red;">{{state.errorMessage}}</p>

      <router-view :authstate="state.user"></router-view>
    </div>
  </div>
</template>

<script>
import firebase from "../firebase";
import { getPrivateMessage } from "../api";

export default {
  data() {
    return {
      state: {
        user: null,
        message: "",
        errorMessage: "",
        isFocusedArticle: false
      }
    };
  },
  created() {
    firebase.auth().onAuthStateChanged(user => {
      if (user) {
        this.state.user = user;
      } else {
        this.state.user = null;
      }
    });
  },
  methods: {
    //auth
    getPrivateMessage: function() {
      this.state.user
        .getIdToken()
        .then(token => {
          return getPrivateMessage(token);
        })
        .then(resp => {
          this.state.message = resp.message;
        })
        .catch(error => {
          this.state.errorMessage = error.toString();
        });
    },
    login: function() {
      firebase.login();
    },
    logout: function() {
      firebase.logout();
    }
  }
};
</script>