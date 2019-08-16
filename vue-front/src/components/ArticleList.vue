<template>
  <div v-if="articles.length != 0">
    <h4 v-for="article in articles" v-bind:key="article.id">No.{{article.id}} {{article.title}}</h4>
  </div>
  <div v-else>
    <h2>記事はありません</h2>
  </div>
</template>

<script>
export default {
  props: {
    articles: {
      default: () => []
    }
  }
};
</script>