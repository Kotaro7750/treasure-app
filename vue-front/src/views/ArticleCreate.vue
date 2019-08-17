<template>
  <div>
    <p style="color:red;">{{errorMessage}}</p>
    <input type="text" v-model="newArticle.title" />
    <input type="text" v-model="newArticle.body" />

    <b-form-checkbox
      v-for="tag in tagList"
      :key="tag.id"
      v-model="selectedTag"
      :value="tag.id"
      button
      button-variant="warning"
    >{{ tag.name}}</b-form-checkbox>

    <b-form-select v-model="selectedJiro">
      <option value="-1">Please select one</option>
      <option v-for="jiro in jiroList" v-bind:key="jiro.id" v-bind:value="jiro.id">{{jiro.name}}</option>
    </b-form-select>

    <b-button variant="warning" v-on:click="createArticle">記事を作成する</b-button>
  </div>
</template>

<script>
import { createArticle, getTagList, getJiroList } from "../api";
export default {
  props: {
    user: {}
  },
  data() {
    return {
      newArticle: {
        title: "",
        body: ""
      },
      tagList: [],
      jiroList: [],
      errorMessage: "",
      selectedTag: [],
      selectedJiro: -1
    };
  },
  created() {
    getTagList()
      .then(resp => {
        this.tagList = resp;
      })
      .catch(error => {
        this.errorMessage = error.toString();
      });

    getJiroList()
      .then(resp => {
        this.jiroList = resp;
      })
      .catch(error => {
        this.errorMessage = error.toString();
      });
  },
  methods: {
    createArticle: function() {
      this.user
        .getIdToken()
        .then(token => {
          return createArticle(
            token,
            this.newArticle.title,
            this.newArticle.body,
            this.selectedTag,
            this.selectedJiro
          );
        })
        .then(() => {
          alert("記事の追加に成功しました!!");
        })
        .catch(error => {
          this.errorMessage = error.toString();
        });
    }
  }
};
</script>