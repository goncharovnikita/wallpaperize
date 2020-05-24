import { combineReducers } from 'redux';
import { connectRouter } from 'connected-react-router';
import { History } from 'history';

import picturesReducer from './pictures';
import appReducer from './app';
import homeReducer from './home';

export default function createRootReducer(history: History) {
  return combineReducers({
    router: connectRouter(history),
    app: appReducer,
    home: homeReducer,
    pictures: picturesReducer
  });
}
