<template lang="pug">
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
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component({})
export default class ShowCase extends Vue {
  @Prop() private imageList1!: any[];
  @Prop() private imageList2!: any[];

  get imageList() {
    const imageList: any[] = [];
    const maxLength: number = (this.imageList1.length > this.imageList2.length) ?
      this.imageList1.length : this.imageList2.length;
    for (let i = 0; i < maxLength; i++) {
      if (this.imageList1[i]) {
        imageList.push(this.imageList1[i]);
      }
      if (this.imageList2[i]) {
        imageList.push(this.imageList2[i]);
      }
    }
    return imageList;
  }
}
</script>
