<template>
  <div>
    <h4>ユーザ:{{ comment.user_id }}</h4>
    <h4>内容:{{ comment.body }}</h4>
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