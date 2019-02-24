<template lang="pug">
  v-text-field(:label="label" v-model="query" @change="submitQuery" :loading="loading" outline)
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import axios from "axios";

@Component({})
export default class QueryField extends Vue {
  @Prop() private label!: string;
  private query: string = "";
  private loading: boolean = false;

  private submitQuery() {
    this.loading = true;
    axios
      .get(`/api/?q=${this.query}`)
      .then(res => {
        this.loading = false;
        this.$emit("get", res.data);
      })
      .catch(e => {
        this.loading = false;
        this.$emit("fail");
      });
  }
}
</script>
