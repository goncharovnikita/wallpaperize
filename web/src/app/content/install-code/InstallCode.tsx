import * as React from 'react';
import { Observable, zip } from 'rxjs';
import { AppService, DownloadLinks } from '../../services/app.service';
import { map, switchMap } from 'rxjs/operators';
import { Platform } from '../../platform/platform';

interface InstallCodeState {
    link: string;
}

export class InstallCode extends React.Component<{}, InstallCodeState> {
    constructor(props) {
        super(props);
        this.state = {link: ''};
    }

    componentDidMount() {
        AppService.getDownloadLinks()
            .subscribe(links => {
                AppService.getSelectedPlatform()
                .subscribe(p => {
                    switch (p) {
                        case Platform.Linux:
                        this.setState({link: links.linux});
                        break;
                        case Platform.Mac:
                        this.setState({link: links.mac});
                        break;
                    }
                });
            });
    }

    render() {
        return (
            <div className="lead">
                <a href={this.state.link} target="_blank">
                    <button className="btn btn-primaty">Download</button>
                </a>
            </div>
        );
    }

    // private _getInstallCode(version: string, url: string): string {
    //     return `
    //     curl ${url}/${version} --output wallpaperize && chmod +x ./wallpaperize &&
    //     sudo mv ./wallpaperize /usr/local/bin/wallpaperize && . ~/.bashrc
    //     `;
    // }
}
