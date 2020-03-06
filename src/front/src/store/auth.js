import axios from "axios"
import { setToken, STATUS_OK } from "../util";
import router from "../router";

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
    context.commit('error/clearError', {}, {root: true});

    const response = await axios.post("/login", data);
    if (response.status === STATUS_OK) {
      context.commit('setToken', response.data.token);
      return
    }

    context.commit('error/setError', {code: response.status, message: response.data.message}, {root: true});
  },

  async register(context, data) {
    context.commit('error/clearError', {}, {root: true});

    const response = await axios.post("/register", data);
    if (response.status === STATUS_OK) {
      context.commit('setToken', response.data.token);
      return
    }

    context.commit('error/setError', {code: response.status, message: response.data.message}, {root: true});
  },

  async currentUser(context) {
    context.commit('error/clearError', {}, {root: true});

    const response = await axios.get("/user",);
    if (response.status === STATUS_OK) {
      context.commit('setUser', response.data);
      return
    }

    context.commit('setUser', "");
  },

  logout(context) {
    context.commit("setToken", "");
    context.commit("setUser", "");

    router.push('/login');
  }
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}

