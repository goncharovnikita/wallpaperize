import * as React from 'react';
import { Menu } from '@app/menu/Menu';
import { App } from '@app/app/App';
import { SavedImages } from '@app/saved-images/SavedImages';
import { AppPath } from '@app/state/app-state';

export class Main extends React.Component<{path: AppPath}> {
    constructor(props: any) {
        super(props);
    }

    component = (): JSX.Element => {
        switch (this.props.path) {
            case AppPath.Saved:
            return <SavedImages />;
            default:
            return <App />;
        }
    }

    render() {
        return (
            <div className="d-flex h-100">
                <Menu />
                {this.component()}
            </div>
        );
    }
}
