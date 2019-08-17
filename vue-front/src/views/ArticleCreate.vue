<template>
  <p>
    <input type="text" v-model="newArticle.title" />
    <input type="text" v-model="newArticle.body" />
    <button v-on:click="createArticle">記事を作成する</button>
  </p>
</template>

<script>
import { createArticle } from "../api";
export default {
  props: {
    state: {}
  },
  data() {
    return {
      newArticle: {
        title: "",
        body: ""
      }
    };
  },
  methods: {
    createArticle: function() {
      this.state
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
  }
};
</script>