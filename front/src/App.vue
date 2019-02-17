<template lang="pug">
  v-app
    v-container
      v-layout(row wrap)
        v-flex(xs12 sm6 md3 offset-md3)
          query-field(label="query1" @get="get1" @fail="fail1")
        v-flex(xs12 sm6 md3)
          query-field(label="query2" @get="get2" @fail="fail2")

      v-layout
        v-flex(xs12 sm12 md12)
          show-case(:imageList1="resultList1" :imageList2="resultList2")
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import QueryField from "@/components/QueryField.vue";
import ShowCase from "@/components/ShowCase.vue";

@Component({
  components: { QueryField, ShowCase }
})
export default class App extends Vue {
  private resultList1: any[];
  private resultList2: any[];

  public constructor() {
    super();
    this.resultList1 = [];
    this.resultList2 = [];
  }

  private get1(result: any) {
    this.resultList1 = [];
    this.resultList1 = result;
  }

  private get2(result: any) {
    this.resultList2 = [];
    this.resultList2 = result;
  }

  private fail1() {
    this.resultList1 = [];
  }

  private fail2() {
    this.resultList2 = [];
  }
}
</script>
