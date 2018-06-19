import * as React from 'react';
import { MainRedux } from '@app/main/MainRedux';
import { InitRedux } from '@app/init/InitRedux';
import { AppPath } from '@app/state/app-state';

export class Router extends React.Component<{path: AppPath}> {
    constructor(props) {
        super(props);
    }

    shouldComponentUpdate(p, v): boolean {
        return true;
    }

    component = (): JSX.Element => {
        switch (this.props.path) {
            case AppPath.Saved:
            case AppPath.Main:
            return <MainRedux />;
            case AppPath.Init:
            default:
            return <InitRedux />;
        }
    }

    render() {
        return (
            <div>
                {this.component()}
            </div>
        );
    }
}
