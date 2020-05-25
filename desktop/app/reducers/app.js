import { handleActions } from 'redux-actions';
import produce from 'immer';

import { setAppInfo, setImageInstallStatus } from '../actions/app';

const initialState = {
  binVersion: '',
  build: '',
  imageInstallStatus: 'idle'
};

export default handleActions(
  {
    [setAppInfo]: (state, { payload }) =>
      produce(state, draft => {
        Object.assign(draft, payload);
      }),
    [setImageInstallStatus]: (state, { payload }) =>
      produce(state, draft => {
        Object.assign(draft, payload);
      })
  },
  initialState
);
