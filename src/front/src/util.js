import { APP_NAME } from './constants';

export const STATUS_OK = 200;
export const STATUS_UNAUTHORIZED = 401;

export const getToken = () => localStorage.getItem('token');

export const setToken = (token) => {
  localStorage.setItem('token', token);
};

export const formatDate = (date) => {
  if (!date) return null;

  return date.replace(/-/g, '/');
};

export const parseDate = (date) => {
  if (!date) return null;

  return date.replace(/\//g, '-');
};

export const generateTitle = (pageTitle) => `${pageTitle}｜${APP_NAME}`;
