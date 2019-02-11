import { mount } from '@vue/test-utils';
import Vue from 'vue';
import Vuetify from 'vuetify';
import App from '@/App.vue';

Vue.use(Vuetify);

describe('App', () => {
  test('is a Vue instance', () => {
    const wrapper = mount(App);
    expect(wrapper.isVueInstance()).toBeTruthy();
  })
})
