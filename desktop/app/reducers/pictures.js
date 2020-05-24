import { handleActions } from 'redux-actions';
import produce from 'immer';

import { setImages } from '../actions/pictures';

const initialState = {
  daily: [],
  dailyCached: [],
  random: [],
  randomCached: []
};

export default handleActions(
  {
    [setImages]: (state, { payload: { type, ids } }) =>
      produce(state, draft => {
        draft[type] = ids;
      })
  },
  initialState
);
