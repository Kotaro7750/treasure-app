<template>
  <div v-if="state.user === null">
    <button v-on:click="login">Please login</button>
  </div>

  <div v-else>
    <div>{{state.message}}</div>

    <p style="color:red;">{{state.errorMessage}}</p>
    <button v-on:click="getPrivateMessage">Get Private Message</button>

    <div>
      <input type="number" v-model="article.id" />
      <button v-on:click="showArticle">Show Article</button>
    </div>

    <div>
      <input type="text" v-model="newArticle.title" />
      <input type="text" v-model="newArticle.body" />
      <button v-on:click="createArticle">Create Article</button>
    </div>

    <button v-on:click="logout">Logout</button>

    <h1>記事</h1>
    <Article v-bind:article="article" v-bind:tags="tags" />

    <dir v-if="comments.length != 0">
      <h1>コメント</h1>
      <Comment v-for="comment in comments" v-bind:key="comment.id" v-bind:comment="comment" />
    </dir>
  </div>
</template>

<script>
import firebase from "../firebase";
import { getPrivateMessage, showArticle, createArticle } from "../api";

import Article from "./Article.vue";
import Comment from "./Comment.vue";

export default {
  components: { Article, Comment },
  data() {
    return {
      state: {
        user: null,
        message: "",
        errorMessage: ""
      },
      article: {
        id: 0,
        title: "nothing",
        body: "nothing"
      },
      newArticle: {
        title: "",
        body: ""
      },
      comments: [],
      tags: [{ id: 0, name: "なし" }]
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
          this.article.title = resp.article.title;
          this.article.body = resp.article.body;
          this.tags = resp.tag;
          this.comments = resp.comment;
          //this.appendComment(resp.comment);
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
    }
    //comment
  }
};
</script>