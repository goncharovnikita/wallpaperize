import * as React from 'react';
import { Route } from 'react-router';
import { App } from '../app/App';
import { MenuAbout } from '../global-menu/about/About';
import { InitApp } from '@app/init/Init';

export class Router extends React.Component {
    render() {
        return (
            <div>
                <Route path="/main" component={App} />
                <Route path="/init" component={InitApp} />
                <Route path="/menu/about" component={MenuAbout} />
                {/* <Redirect from="/" to="/init" /> */}
            </div>
        );
    }
}
