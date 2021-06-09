const state = {
  isOn: false,
};

const getters = {
  // eslint-disable-next-line no-shadow
  isOn(state) {
    return state.isOn;
  },
};

const mutations = {
  // eslint-disable-next-line no-shadow
  setStatus(state, status) {
    // eslint-disable-next-line no-param-reassign
    state.isOn = status;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
};
