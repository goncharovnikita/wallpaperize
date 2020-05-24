import { combineReducers } from 'redux';
import { connectRouter } from 'connected-react-router';
import { History } from 'history';

import picturesReducer from './pictures';

export default function createRootReducer(history: History) {
  return combineReducers({
    router: connectRouter(history),
    pictures: picturesReducer
  });
}
