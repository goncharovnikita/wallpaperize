import * as React from 'react';
import { init } from '@app/wallpaperize-proxy/init';
import { Redirect } from 'react-router';

interface InitState {
    initialized: boolean;
}

export class InitApp extends React.Component<InitState, InitState> {
    constructor(props: any) {
        super(props);
        this.state = {initialized: false};
        this.main();
    }
    render() {
        return(
            this.state.initialized ? <Redirect to="/main" />
            : <div>Wallpaperize initializing</div>
        );
    }

    async main() {
        const inited = await init();
        this.setState({initialized: inited});
    }
}
