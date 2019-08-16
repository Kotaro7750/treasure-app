<template>
  <div>
    <h1>コンテンツ</h1>
    <h2>タイトル</h2>
    <div>{{ article.title }}</div>
    <h2>内容</h2>
    <div>{{ article.body }}</div>
  </div>
</template>


<script>
export default {
  props: {
    article: {
      default: () => ({ title: "default title", body: "default body" })
    }
  }
};
</script>