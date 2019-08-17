<template>
  <div>
    <input type="number" v-model="article.id" />
    <button v-on:click="showArticle">記事を見る</button>

    <button v-on:click="getArticleList">記事一覧</button>

    <ArticleDetail v-bind:article="article" v-if="state.isFocusedArticle" />

    <ArticleList v-bind:articles="article_list" v-else />
  </div>
</template>

<script>
import { showArticle, getArticleList } from "../api";

import ArticleDetail from "../components/ArticleDetail/ArticleDetail.vue";
import ArticleList from "../components/ArticleList.vue";
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
      article_list: []
    };
  },
  methods: {
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
  }
};
</script>