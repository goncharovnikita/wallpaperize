import * as React from 'react';
import { WallpaperizeInfo, getInfo } from '@app/wallpaperize-proxy/get-info';
import * as electron from 'electron';

export class MenuAbout extends React.Component {
    state: {about: WallpaperizeInfo};
    constructor(props: any) {
        super(props);
        this.state = {about: {} as WallpaperizeInfo};
        this.setAbout();
    }

    async setAbout(): Promise<void> {
        const version = await getInfo();
        this.setState({about: version});
    }

    render() {
        return (
            <div>
                <h1>Wallpaperize version {this.state.about.app_version}</h1>
                <h2>Desktop version {electron.remote.app.getVersion()}</h2>
            </div>
        );
    }
}
