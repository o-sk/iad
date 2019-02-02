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
          v-card
            v-container(grid-list-sm fluid)
              v-layout(row wrap)
                v-flex(v-for="image in imageList" :key="image.id + image.rh" xs4 sm4 md2 d-flex)
                  v-card(flat tile class="d-flex")
                    v-img(:src="image.tu" aspect-ratio="1" class="grey lighten-2")
                      v-layout(slot="placeholder" fill-height align-center justify-center ma-0)
                        v-progress-circular(indeterminate color="grey lighten-5")
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import QueryField from '@/components/QueryField.vue';

@Component({
  components: { QueryField },
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

  get imageList() {
    const imageList: any[] = [];
    const maxLength: number = (this.resultList1.length > this.resultList2.length) ?
      this.resultList1.length : this.resultList2.length;
    for (let i = 0; i < maxLength; i++) {
      if (this.resultList1[i]) {
        imageList.push(this.resultList1[i]);
      }
      if (this.resultList2[i]) {
        imageList.push(this.resultList2[i]);
      }
    }
    return imageList;
  }
}
</script>
