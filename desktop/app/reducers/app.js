import { handleActions } from 'redux-actions';
import produce from 'immer';

import { setAppInfo } from '../actions/app';

const initialState = {
  binVersion: '',
  build: ''
};

export default handleActions(
  {
    [setAppInfo]: (state, { payload }) =>
      produce(state, draft => {
        Object.assign(draft, payload);
      })
  },
  initialState
);
