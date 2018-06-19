import * as React from 'react';
import { init } from '@app/wallpaperize-proxy/init';

export class InitApp extends React.Component<{initialize: () => void}> {
    constructor(props) {
        super(props);
        this.state = {initialized: false};
        this.main();
    }

    render() {
        return(
            <div>Wallpaperize initializing</div>
        );
    }

    async main() {
        await init();
        this.props.initialize();
    }
}
