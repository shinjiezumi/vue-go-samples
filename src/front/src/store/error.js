const state = {
  code: "",
  message: ""
};

const getters = {
  getCode: state => state.code,
  getError: state => state.message
};

const mutations = {
  setCode(state, code) {
    state.code = code
  },
  setError(state, message) {
    state.message = message
  },
};

const actions = {
  clearError(context) {
    context.commit('error/setCode', "", {root: true});
    context.commit('error/setError', "", {root: true});
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
}

