<template>
  <div>
    <p style="color:red;">{{errorMessage}}</p>

    <b-input-group prepend="記事タイトル" class="col-lg-5 offset-lg-4">
      <b-form-input type="text" v-model="newArticle.title"></b-form-input>
    </b-input-group>
    <p></p>
    <b-input-group prepend="記事内容" class="col-lg-5 offset-lg-4">
      <b-form-input type="text" v-model="newArticle.body"></b-form-input>
    </b-input-group>
    <p></p>

    <b-input-group prepend="タグ" class="col-lg-5 offset-lg-4">
      <b-form-checkbox
        v-for="tag in tagList"
        :key="tag.id"
        v-model="selectedTag"
        :value="tag.id"
        button
        button-variant="warning"
      >{{ tag.name}}</b-form-checkbox>
    </b-input-group>

    <p></p>
    <b-input-group prepend="二郎" class="col-lg-5 offset-lg-4">
      <b-form-select v-model="selectedJiro">
        <option value="-1">なし</option>
        <option v-for="jiro in jiroList" v-bind:key="jiro.id" v-bind:value="jiro.id">{{jiro.name}}</option>
      </b-form-select>
    </b-input-group>
    <p></p>

    <b-input-group class="col-lg-5 offset-lg-4">
      <b-button variant="warning" v-on:click="createArticle">記事を作成する</b-button>
    </b-input-group>
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