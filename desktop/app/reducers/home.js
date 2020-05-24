import { handleActions } from 'redux-actions';
import produce from 'immer';

import { setSelectedImage } from '../actions/home';

const initialState = {
  selectedImage: null,
  selectedImageType: null
};

export default handleActions(
  {
    [setSelectedImage]: (state, { payload }) =>
      produce(state, draft => {
        Object.assign(draft, payload);
      })
  },
  initialState
);
