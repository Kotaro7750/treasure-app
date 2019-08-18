<template>
  <div class="offset-lg-4">
    <div v-if="jiros.length != 0">
      <h4 v-for="jiro in jiros" v-bind:key="jiro.id">No.{{jiro.id}} {{jiro.name}}</h4>
    </div>
    <div v-else>
      <h2>二郎はありません</h2>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    jiros: {
      default: () => []
    }
  }
};
</script>