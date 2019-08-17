<template>
  <div>
    <input type="number" v-model="article.id" />
    <b-button variant="warning" v-on:click="showArticle">記事を見る</b-button>

    <b-button variant="warning" v-on:click="getArticleList">記事一覧</b-button>

    <ArticleDetail v-bind:article="article" v-if="state.isFocusedArticle" />

    <div v-if="state.isFocusedArticle">
      <h3>コメント</h3>
      <input type="text" v-model="comment_body" />
      <b-button variant="warning" v-on:click="createComment">コメントする</b-button>
    </div>

    <ArticleList v-bind:articles="article_list" v-else />
  </div>
</template>

<script>
import { showArticle, getArticleList, createComment } from "../api";

import ArticleDetail from "../components/ArticleDetail/ArticleDetail.vue";
import ArticleList from "../components/ArticleList.vue";
export default {
  components: { ArticleDetail, ArticleList },
  props: {
    user: {}
  },
  data() {
    return {
      state: {
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
      comment_body: ""
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
          this.article.jiro = resp.jiro;
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
    },
    createComment: function() {
      this.user
        .getIdToken()
        .then(token => {
          return createComment(token, this.article.id, this.comment_body);
        })
        .then(() => {
          this.showArticle(Number(this.article_id));
        })
        .catch(error => {
          this.state.errorMessage = error.toString();
        });
    }
  }
};
</script>