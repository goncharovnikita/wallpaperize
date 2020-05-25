import { createAction } from 'redux-actions';

const ca = (n, p) => createAction(`app/${n}`, p);

export const requestInitApp = ca('REQUEST_INIT_APP');
export const setAppInfo = ca('SET_APP_INFO', (binVersion, build) => ({
  binVersion,
  build
}));
export const setImageInstallStatus = ca(
  'SET_IMAGE_INSTALL_STATUS',
  imageInstallStatus => ({
    imageInstallStatus
  })
);
export const requestInstallImage = ca('REQUEST_INSTALL_IMAGE', src => ({
  src
}));
