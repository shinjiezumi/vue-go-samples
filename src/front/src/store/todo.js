import axios from "axios"
import { STATUS_OK } from "../util";

const state = {
  todoList: {},
};

const getters = {
  getList(state) {
    return state.todoList
  }
};

const mutations = {
  setTodoList(state, todoList) {
    state.todoList = todoList
  }
};

const actions = {
  async getList(context, data) {
    context.commit('error/clearError', {}, {root: true});
    context.commit('loading/setStatus', true, {root: true});

    const response = await axios.get("/todo/list", {params: data});
    context.commit('loading/setStatus', false, {root: true});
    if (response.status === STATUS_OK) {
      context.commit('setTodoList', response.data.data);
      return
    }

    context.commit('error/setError', {code: response.status, message: response.data.message}, {root: true});
  },
  async create(context, data) {
    context.commit('error/clearError', {}, {root: true});

    const response = await axios.post("/todo", data.params);
    if (response.status !== STATUS_OK) {
      context.commit('error/setError', {code: response.status, message: response.data.message}, {root: true});
    }
  },
  async modify(context, data) {
    context.commit('error/clearError', {}, {root: true});

    const response = await axios.put("/todo/" + data.id, data.params);
    if (response.status !== STATUS_OK) {
      context.commit('error/setError', {code: response.status, message: response.data.message}, {root: true});
    }
  },
  async remove(context, data) {
    context.commit('error/clearError', {}, {root: true});

    const response = await axios.delete("/todo/" + data.id);
    if (response.status !== STATUS_OK) {
      context.commit('error/setError', {code: response.status, message: response.data.message}, {root: true});
    }
  },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}

