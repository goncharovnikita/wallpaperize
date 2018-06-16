import * as React from 'react';
import { Route, Redirect } from 'react-router';
import { App } from '../app/App';
import { MenuAbout } from '../global-menu/about/About';

export class Router extends React.Component {
    render() {
        return (
            <div>
                <Route path="/main" component={App} />
                <Route path="/menu/about" component={MenuAbout} />
                <Redirect from="/" to="/main" />
            </div>
        );
    }
}
