import * as React from 'react';
import { WallpaperizeInfo, getInfo } from '@app/wallpaperize-proxy/get-info';

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
            <h1>Wallpaperize version {this.state.about.app_version}</h1>
        );
    }
}
