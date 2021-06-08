import axios from 'axios';
// eslint-disable-next-line import/no-unresolved
import { STATUS_OK } from '@/util';

const state = {
  todoList: {},
};

const getters = {
  // eslint-disable-next-line no-shadow
  getList(state) {
    return state.todoList;
  },
};

const mutations = {
  setTodoList(_state, todoList) {
    state.todoList = todoList;
  },
};

const actions = {
  async getList(context, data) {
    context.commit('error/clearError', {}, { root: true });
    context.commit('loading/setStatus', true, { root: true });

    const response = await axios.get('/todos', { params: data });
    context.commit('loading/setStatus', false, { root: true });
    if (response.status === STATUS_OK) {
      context.commit('setTodoList', response.data.data);
      return;
    }

    context.commit('error/setError', { code: response.status, message: response.data.message }, { root: true });
  },
  async create(context, data) {
    context.commit('error/clearError', {}, { root: true });

    const response = await axios.post('/todos', data.params);
    if (response.status !== STATUS_OK) {
      context.commit('error/setError', { code: response.status, message: response.data.message }, { root: true });
    }
  },
  async modify(context, data) {
    context.commit('error/clearError', {}, { root: true });

    const response = await axios.put(`/todos/${data.id}`, data.params);
    if (response.status !== STATUS_OK) {
      context.commit('error/setError', { code: response.status, message: response.data.message }, { root: true });
    }
  },
  async remove(context, data) {
    context.commit('error/clearError', {}, { root: true });

    const response = await axios.delete(`/todos/${data.id}`);
    if (response.status !== STATUS_OK) {
      context.commit('error/setError', { code: response.status, message: response.data.message }, { root: true });
    }
  },
  async finished(context, data) {
    context.commit('error/clearError', {}, { root: true });

    const response = await axios.put(`/todos/${data.id}/finished`);
    if (response.status !== STATUS_OK) {
      context.commit('error/setError', { code: response.status, message: response.data.message }, { root: true });
    }
  },
  async unfinished(context, data) {
    context.commit('error/clearError', {}, { root: true });

    const response = await axios.put(`/todos/${data.id}/unfinished`);
    if (response.status !== STATUS_OK) {
      context.commit('error/setError', { code: response.status, message: response.data.message }, { root: true });
    }
  },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions,
};
