import axios from "axios"
import { setToken, STATUS_OK } from "../util";

const state = {
  user: "",
};

const getters = {
  check(state) {
    return state.user !== "" && state.user.name !== ""
  }
};

const mutations = {
  setToken(state, token) {
    setToken(token)
  },
  setUser(state, user) {
    state.user = user
  }
};

const actions = {
  async login(context, data) {
    context.commit('error/setCode', "", {root: true});
    context.commit('error/setError', "", {root: true});

    const response = await axios.post("/login", data);
    if (response.status === STATUS_OK) {
      context.commit('setToken', response.data.token);
      return
    }

    context.commit('error/setCode', response.status, {root: true});
    context.commit('error/setError', response.data.message, {root: true});
  },

  async register(context, data) {
    context.commit('error/setCode', "", {root: true});
    context.commit('error/setError', "", {root: true});

    const response = await axios.post("/register", data);
    if (response.status === STATUS_OK) {
      context.commit('setToken', response.data.token);
      return
    }

    context.commit('error/setCode', response.status, {root: true});
    context.commit('error/setError', response.data.message, {root: true});
  },

  async currentUser(context) {
    context.commit('error/setCode', "", {root: true});
    context.commit('error/setError', "", {root: true});

    const response = await axios.get("/user",);
    if (response.status === STATUS_OK) {
      context.commit('setUser', response.data);
      return
    }

    context.commit('setUser', "");
    context.commit('error/setCode', response.status, {root: true});
    context.commit('error/setError', response.data.message, {root: true});
  },

};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}

