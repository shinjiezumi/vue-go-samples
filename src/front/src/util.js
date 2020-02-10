export const STATUS_OK = 200;

export const getToken = () => {
  return localStorage.getItem("token")
};

export const setToken = (token) => {
  localStorage.setItem("token", token)
};

