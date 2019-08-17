<template>
  <div>
    <input type="number" v-model="jiro.id" />
    <b-button variant="warning" v-on:click="showJiro">二郎を見る</b-button>

    <b-button variant="warning" v-on:click="getJiroList">二郎一覧</b-button>

    <JiroDetail v-bind:jiro="jiro" v-if="state.isFocusedJiro" />

    <JiroList v-bind:jiros="jiro_list" v-else />
  </div>
</template>

<script>
import { showJiro, getJiroList } from "../api";

import JiroDetail from "../components/JiroDetail/JiroDetail.vue";
import JiroList from "../components/JiroList.vue";
export default {
  components: { JiroDetail, JiroList },
  props: {
    user: {}
  },
  data() {
    return {
      state: {
        message: "",
        errorMessage: "",
        isFocusedJiro: false
      },
      jiro: {
        id: 0,
        name: "nothing"
      },
      jiro_list: []
    };
  },
  methods: {
    //jiro
    showJiro: function() {
      showJiro(Number(this.jiro.id))
        .then(resp => {
          this.state.isFocusedJiro = true;
          this.jiro = resp;
        })
        .catch(error => {
          this.state.errorMessage = error.toString();
        });
    },
    getJiroList: function() {
      getJiroList()
        .then(resp => {
          this.state.isFocusedJiro = false;
          this.jiro_list = resp;
        })
        .catch(error => {
          this.state.errorMessage = error.toString();
        });
    }
  }
};
</script>