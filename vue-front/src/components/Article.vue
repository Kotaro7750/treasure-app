<template>
  <div>
    <h1>title:{{ article.title }}</h1>
    <h1>body:{{ article.body }}</h1>
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