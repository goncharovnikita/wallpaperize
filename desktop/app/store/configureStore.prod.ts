import { createStore, applyMiddleware } from 'redux';
import { createHashHistory } from 'history';
import createSagaMiddleware from 'redux-saga';
import { routerMiddleware } from 'connected-react-router';
import createRootReducer from '../reducers';
import { Store, AppState } from '../reducers/types';
import sagas from '../sagas';

const history = createHashHistory();
const rootReducer = createRootReducer(history);
const router = routerMiddleware(history);
const sagaMiddleware = createSagaMiddleware();
const enhancer = applyMiddleware(sagaMiddleware, router);

function configureStore(initialState?: AppState): Store {
  const store = createStore(rootReducer, initialState, enhancer);

  sagaMiddleware.run(sagas);

  return store;
}

export default { configureStore, history };
