<template>
  <div>
    <input type="number" v-model="jiro.id" />
    <b-button variant="warning" v-on:click="showJiro">二郎を見る</b-button>

    <b-button variant="warning" v-on:click="getJiroList">二郎一覧</b-button>

    <b-button variant="warning" v-on:click="nearestJiro">最寄りの二郎</b-button>
    <div v-if="nearest_route != ''">
      <h3>ルート</h3>
      {{this.nearest_route}}
    </div>

    <JiroDetail v-bind:jiro="jiro" v-if="state.isFocusedJiro" />

    <JiroList v-bind:jiros="jiro_list" v-else />
  </div>
</template>

<script>
import { showJiro, getJiroList, nearestJiro } from "../api";

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
      jiro_list: [],
      nearest_route: ""
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
    },
    nearestJiro: function() {
      this.user
        .getIdToken()
        .then(token => {
          return nearestJiro(token, "今の位置");
        })
        .then(resp => {
          this.jiro.id = resp.jiro.id;
          this.showJiro();
          this.nearest_route = resp.route;
        })
        .catch(error => {
          this.state.errorMessage = error.toString();
        });
    }
  }
};
</script>