<template>
  <div>
    <h2>タイトル</h2>
    <div>{{ article.title }}</div>

    <h3>タグ</h3>
    <li v-for="tag in tags" v-bind:key="tag.id">{{ tag.name }}</li>

    <h3>二郎</h3>

    {{ jiro.name }}
    <h2>内容</h2>
    <div>{{ article.body }}</div>
  </div>
</template>

<script>
export default {
  props: {
    article: {
      default: () => ({ title: "default title", body: "default body" })
    },
    tags: {
      default: () => [{ id: 0, name: "なし" }]
    },
    jiro: {
      default: () => {}
    }
  }
};
</script>