import { shallowMount } from "@vue/test-utils";
import Vue from "vue";
import Vuetify from "vuetify";
import ShowCase from "@/components/ShowCase.vue";

Vue.use(Vuetify);

describe("ShowCase", () => {
  test("propsData", () => {
    const wrapper = shallowMount(ShowCase, {
      propsData: {
        imageList1: [1, 2],
        imageList2: [11, 12]
      }
    });
    expect(wrapper.vm.$props.imageList1).toEqual([1, 2]);
    expect(wrapper.vm.$props.imageList2).toEqual([11, 12]);
  });

  xtest("computed", () => {
  });
});
