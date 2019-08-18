<template>
  <div class="offset-lg-4">
    <h2>店名</h2>
    <div>{{ jiro.name }}</div>
    {{ jiro.status }}
    <h3>住所</h3>
    {{ jiro.name }}
  </div>
</template>

<script>
export default {
  props: {
    jiro: {
      default: () => ({})
    }
  }
};
</script>