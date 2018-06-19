import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { HashRouter } from 'react-router-dom';
import { createStore } from 'redux';
import { appReducer } from '@app/reducers/reducer';
import { Provider } from 'react-redux';
import { RouterRedux } from '@app/router/RouterRedux';
// tslint:disable:no-var-requires
const fixPath = require('fix-path');
fixPath();

const store = createStore(appReducer);

ReactDOM.render(
    <Provider store={store}>
      <HashRouter>
        < RouterRedux />
      </HashRouter>
    </Provider> ,
    document.getElementById('app')
);
