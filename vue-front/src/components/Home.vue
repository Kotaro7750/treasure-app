<template>
  <div v-if="state.user === null">
    <button v-on:click="login">ログイン</button>
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

    <button v-on:click="logout">ログアウト</button>
    <button v-on:click="getArticleList">記事一覧</button>

    <ArticleDetail v-bind:article="article" v-if="state.isFocusedArticle" />

    <ArticleList v-bind:articles="article_list" v-else />
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