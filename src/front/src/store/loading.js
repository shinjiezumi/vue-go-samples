const state = {
  isOn: false,
};

const getters = {
  isOn(state) {
    return state.isOn
  }
};

const mutations = {
  setStatus(state, status) {
    state.isOn = status
  },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations
}