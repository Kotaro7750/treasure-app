<template>
  <div>
    <h2>ユーザ:{{ comment.user_id }}</h2>
    <h2>内容:{{ comment.body }}</h2>
  </div>
</template>


<script>
export default {
  props: {
    comment: {
      default: () => [{ user_id: 0, body: "default body" }]
    }
  }
};
</script>