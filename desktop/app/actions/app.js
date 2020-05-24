import { createAction } from 'redux-actions';

const ca = (n, p) => createAction(`app/${n}`, p);

export const requestInitApp = ca('REQUEST_INIT_APP');
export const setAppInfo = ca('SET_APP_INFO', (binVersion, build) => ({
  binVersion,
  build
}));
