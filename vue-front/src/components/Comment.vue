<template>
  <div>
    <h1>user_id:{{ comment.user_id }}</h1>
    <h1>body:{{ comment.body }}</h1>
  </div>
</template>


<script>
export default {
  props: {
    comment: {
      default: () => ({ user_id: 0, body: "default body" })
    }
  }
};
</script>