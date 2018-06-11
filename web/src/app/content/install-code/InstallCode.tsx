import * as React from 'react';
import { Observable, zip } from 'rxjs';
import { AppService } from '../../services/app.service';
import { map, switchMap } from 'rxjs/operators';
import { Platform } from '../../platform/platform';

export class InstallCode extends React.Component {
    private _downloadURL: string;
    state: {
        code: string;
    };
    constructor(props) {
        super(props);
        this.state = {
            code: '',
        };
        this._downloadURL = 'https://wallpaperize.goncharovnikita.com/builds';
    }

    componentDidMount() {
        this._getSelectedCode()
            .subscribe(code => {
                this.setState({code});
            });
    }

    render() {
        return (
            <div className="lead">
                {this.state.code}
            </div>
        );
    }

    private _getSelectedCode(): Observable<string> {
        return AppService.getSelectedVersion()
            .pipe(
                switchMap(version => {
                    return AppService.getSelectedPlatform()
                        .pipe(
                            map(platform => {
                                switch (platform) {
                                    case Platform.Mac:
                                    // return this._getInstallCode(`darwin-amd64-${version}`, this._downloadURL);
                                    return this._getInstallCode(`darwin-amd64-1.0.0`, this._downloadURL);
                                    default:
                                    return this._getInstallCode(`linux-amd64-${version}`, this._downloadURL);
                                }
                            })
                        );
                })
            );
    }

    private _getInstallCode(version: string, url: string): string {
        return `
        curl ${url}/${version} --output wallpaperize -D && chmod +x ./wallpaperize &&
        cp ./wallpaperize /usr/local/bin/wallpaperize | rm ./wallpaperize &&
        . ~/.bashrc
        `;
    }
}
