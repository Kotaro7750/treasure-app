<template>
  <div>
    <h6>{{article.user_id}}さんが投稿</h6>
    <h1>
      <b>
        <u>{{ article.title }}</u>
      </b>
    </h1>

    <li v-for="tag in tags" v-bind:key="tag.id" class="tag">{{ tag.name }}</li>

    <p></p>

    <h4>{{ jiro.name }}</h4>
    <p></p>

    <div>
      <span id="body">{{ article.body }}</span>
    </div>
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

<style>
li {
  display: inline-block;
}

.tag {
  padding: 0.5em 0.5em;
  margin: 0em 0.25em;
  color: #000000; /*文字色*/
  background: #fff;
  border: solid 2px #ffc107; /*線*/
  border-radius: 10px; /*角の丸み*/
}

#body {
  display: block;
  background-color: #ededed;
  color: #000000; /*文字色*/
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
}
</style>