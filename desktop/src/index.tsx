import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { HashRouter } from 'react-router-dom';
import { Router } from './router/Router';
// tslint:disable:no-var-requires
const fixPath = require('fix-path');
fixPath();

ReactDOM.render(
    <HashRouter>
      < Router />
    </HashRouter>,
    document.getElementById('app-root')
);
