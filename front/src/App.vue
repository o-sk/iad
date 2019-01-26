<template>
  <v-app>
    <v-container>
      <v-layout row wrap>
        <v-flex xs12 sm6 md3 offset-md3>
          <v-text-field
            label="query"
            v-model="query"
            @keypress.enter="submitQuery"
            @change="submitQuery"
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
                  :key="image.id"
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
import { Component, Vue } from 'vue-property-decorator';
import axios from 'axios';

@Component({})
export default class App extends Vue {
  private query!: string | null;
  private imageList: any[] | null;

  public constructor() {
    super();
    this.query = null;
    this.imageList = null;
  }

  private submitQuery() {
    axios
      .get(`/api/?q=${this.query}`)
      .then((res) => (
        this.imageList = res.data
      ));
  }
}
</script>
