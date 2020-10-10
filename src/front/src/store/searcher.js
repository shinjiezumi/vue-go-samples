import axios from "axios"
import { STATUS_OK } from "@/util";

const state = {
  searchResult: {},
};

const getters = {
  getResult(state) {
    return state.searchResult
  }
};

const mutations = {
  setResult(state, searchResult) {
    state.searchResult = searchResult
  }
};

const actions = {
  async search(context, data) {
    context.commit('error/clearError', {}, {root: true});
    context.commit('loading/setStatus', true, {root: true});

    const response = await axios.get("/searcher/search", {params: data});
    context.commit('loading/setStatus', false, {root: true});
    if (response.status === STATUS_OK) {
      context.commit('setResult', response.data.data);
      return
    }

    context.commit('error/setError', {code: response.status, message: response.data.message}, {root: true});
  },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}

