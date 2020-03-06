const state = {
  code: "",
  message: ""
};

const getters = {
  getCode: state => state.code,
  getError: state => state.message
};

const mutations = {
  setError(context, data) {
    state.code = data.code;
    state.message = data.message;
  },
  clearError() {
    state.code = "";
    state.message = "";
  }

};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
}

