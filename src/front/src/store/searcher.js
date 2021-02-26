import axios from "axios"
import {STATUS_OK} from "@/util";

const state = {
  searchResult: {
    Qiita: [],
    SlideShare: [],
    Feedly: [],
  },
};

const getters = {
  getResult(state) {
    return state.searchResult
  }
};

const mutations = {
  setResult(state, searchResult) {
    state.searchResult.Qiita = searchResult.Qiita.List
    state.searchResult.SlideShare = searchResult.SlideShare.List
    state.searchResult.Feedly = searchResult.Feedly.List
  }
};

const actions = {
  async init(context) {
    context.commit('setResult', {
      Qiita: [],
      SlideShare: [],
      Feedly: [],
    });
    context.commit('error/clearError', {}, {root: true});
  },
  async search(context, data) {
    context.commit('error/clearError', {}, {root: true});
    context.commit('loading/setStatus', true, {root: true});

    const response = await axios.get("/searcher/search", {params: data});
    context.commit('loading/setStatus', false, {root: true});
    if (response.status === STATUS_OK) {
      context.commit('setResult', response.data);

      // エラーがあれば表示
      let errorMessages = []
      for (let v of ['Qiita', 'SlideShare', 'Feedly']) {
        if (response.data[v].Error.Message !== '') {
          errorMessages.push(v + ':' + response.data[v].Error.Message)
        }
      }
      if (errorMessages.length > 0) {
        context.commit('error/setError', {message: errorMessages.join("\n")}, {root: true});
      }

      return
    }

    context.commit('error/setError', {code: response.status, message: response.message}, {root: true});
  },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}

