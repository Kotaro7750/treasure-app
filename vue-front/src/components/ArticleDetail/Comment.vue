<template>
  <div>
    <h6>{{ comment.user_id }}さん</h6>
    <h4 class="comment">{{ comment.body }}</h4>
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

<style>
.comment {
  position: relative;
  display: inline-block;
  margin: 0.5em 0;
  padding: 7px 10px;
  min-width: 120px;
  max-width: 100%;
  color: #555;
  font-size: 16px;
  background: #fff;
  border: solid 2px rgb(255, 193, 7);
  box-sizing: border-box;
}

.comment:before {
  content: "";
  position: absolute;
  top: -24px;
  left: 10%;
  margin-left: -15px;
  border: 12px solid transparent;
  border-bottom: 12px solid #fff;
  z-index: 2;
}

.comment:after {
  content: "";
  position: absolute;
  top: -29px;
  left: 10%;
  margin-left: -17px;
  border: 14px solid transparent;
  border-bottom: 14px solid rgb(255, 193, 7);
  z-index: 1;
}
</style>