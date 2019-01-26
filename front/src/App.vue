<template>
  <v-app>
    <v-container>
      <v-layout row wrap>
        <v-flex xs12 sm6 md3 offset-md3>
          <v-text-field
            label="query1"
            v-model="query1"
            @keypress.enter="submitQuery1"
            @change="submitQuery1"
            :loading="loading1"
            outline
          ></v-text-field>
        </v-flex>
        <v-flex xs12 sm6 md3>
          <v-text-field
            label="query2"
            v-model="query2"
            @keypress.enter="submitQuery2"
            @change="submitQuery2"
            :loading="loading2"
            outline
          ></v-text-field>
        </v-flex>
      </v-layout>

      <v-layout>
        <v-flex xs12 sm12 md12>
          <v-card>
            <v-container grid-list-sm fluid>
              <v-layout row wrap>
                <v-flex
                  v-for="image in imageList"
                  :key="image.id + image.rh"
                  xs4 sm4 md2
                  d-flex
                >
                  <v-card flat tile class="d-flex">
                    <v-img
                      :src="image.tu"
                      aspect-ratio="1"
                      class="grey lighten-2"
                    >
                      <v-layout
                        slot="placeholder"
                        fill-height
                        align-center
                        justify-center
                        ma-0
                      >
                        <v-progress-circular indeterminate color="grey lighten-5"></v-progress-circular>
                      </v-layout>
                    </v-img>
                  </v-card>
                </v-flex>
              </v-layout>
            </v-container>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import axios from 'axios';

@Component({})
export default class App extends Vue {
  private query1!: string;
  private query2!: string;
  private resultList1: any[];
  private resultList2: any[];
  private loading1: boolean;
  private loading2: boolean;
  private imageList: any[];

  public constructor() {
    super();
    this.query1 = '';
    this.query2 = '';
    this.resultList1 = [];
    this.resultList2 = [];
    this.loading1 = false;
    this.loading2 = false;
    this.imageList = [];
  }

  private submitQuery1() {
    this.loading1 = true;
    axios
      .get(`/api/?q=${this.query1}`)
      .then((res) => {
        this.resultList1 = [];
        this.resultList1 = res.data;
        this.loading1 = false;
      });
  }

  private submitQuery2() {
    this.loading2 = true;
    axios
      .get(`/api/?q=${this.query2}`)
      .then((res) => {
        this.resultList2 = [];
        this.resultList2 = res.data;
        this.loading2 = false;
      });
  }

  @Watch('resultList1')
  private onChangeResultList1() {
    this.setImageList();
  }

  @Watch('resultList2')
  private onChangeResultList2() {
    this.setImageList();
  }

  private setImageList() {
    this.imageList = [];
    const maxLength: number = (this.resultList1.length > this.resultList2.length) ?
      this.resultList1.length : this.resultList2.length;
    for (let i = 0; i < maxLength; i++) {
      if (this.resultList1[i]) {
        this.imageList.push(this.resultList1[i]);
      }
      if (this.resultList2[i]) {
        this.imageList.push(this.resultList2[i]);
      }
    }
  }
}
</script>
