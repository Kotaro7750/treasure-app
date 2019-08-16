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
      <button v-on:click="getPrivateMessage">認証情報</button>

      <div>
        <input type="number" v-model="article.id" />
        <button v-on:click="showArticle">記事を見る</button>
      </div>

      <p>
        <input type="text" v-model="newArticle.title" />
        <input type="text" v-model="newArticle.body" />
        <button v-on:click="createArticle">記事を作成する</button>
      </p>

      <button v-on:click="getArticleList">記事一覧</button>

      <ArticleDetail v-bind:article="article" v-if="state.isFocusedArticle" />

      <ArticleList v-bind:articles="article_list" v-else />
    </div>
  </div>
</template>

<script>
import firebase from "../firebase";
import {
  getPrivateMessage,
  showArticle,
  createArticle,
  getArticleList
} from "../api";

import ArticleDetail from "./ArticleDetail/ArticleDetail.vue";
import ArticleList from "./ArticleList.vue";

export default {
  components: { ArticleDetail, ArticleList },
  data() {
    return {
      state: {
        user: null,
        message: "",
        errorMessage: "",
        isFocusedArticle: false
      },
      article: {
        content: {
          id: 0,
          title: "nothing",
          body: "nothing"
        },
        comments: [],
        tags: []
      },
      article_list: [],
      newArticle: {
        title: "",
        body: ""
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
    },
    //article
    showArticle: function() {
      showArticle(Number(this.article.id))
        .then(resp => {
          this.state.isFocusedArticle = true;
          this.article.content = resp.article;
          this.article.tags = resp.tag;
          this.article.comments = resp.comment;
        })
        .catch(error => {
          this.state.errorMessage = error.toString();
        });
    },
    createArticle: function() {
      this.state.user
        .getIdToken()
        .then(token => {
          return createArticle(
            token,
            this.newArticle.title,
            this.newArticle.body
          );
        })
        .then(resp => {
          this.state.message = resp;
        })
        .catch(error => {
          this.state.errorMessage = error.toString();
        });
    },
    getArticleList: function() {
      getArticleList()
        .then(resp => {
          this.state.isFocusedArticle = false;
          this.article_list = resp;
        })
        .catch(error => {
          this.state.errorMessage = error.toString();
        });
    }
    //comment
  }
};
</script>