<template>
  <div>
    <Article
      v-bind:article="article.content"
      v-bind:tags="article.tags"
      v-bind:jiro="article.jiro"
      class="offset-lg-3 col-lg-6"
    />

    <p></p>

    <div v-if="article.comments.length != 0" class="offset-lg-3 col-lg-6">
      <h3>コメント</h3>
      <Comment
        v-for="comment in article.comments"
        v-bind:key="comment.id"
        v-bind:comment="comment"
      />
    </div>
  </div>
</template>

<script>
import Article from "./Article.vue";
import Comment from "./Comment.vue";

export default {
  components: { Article, Comment },
  props: {
    article: {
      default: () => ({ title: "default title", body: "default body" })
    }
  }
};
</script>