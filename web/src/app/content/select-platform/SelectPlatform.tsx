import * as React from 'react';
import { Platform } from '../../platform/platform';
import { Observable } from 'rxjs';
import { AppService } from '../../services/app.service';
import { map } from 'rxjs/operators';

interface IPlatformButton {
    text: string;
    platform: Platform;
}

export class SelectPlatform extends React.Component {

    state: {
        buttons: JSX.Element[],
    };
    constructor(props) {
        super(props);
        this.state = {
            buttons: [],
        };

    }

    componentDidMount() {
        this._getButtons(this._getPlatformButtons())
            .subscribe(b => {
                this.setState({
                    buttons: b,
                });
            });
    }

    render() {
        return (
            <div className="m-3 d-flex justify-content-around">
                <div className="btn-group" role="group" aria-label="Basic example">
                    {...this.state.buttons}
                </div>
            </div>
        );
    }

    private _getButtons(btns: IPlatformButton[]): Observable<JSX.Element[]> {
        return AppService.getSelectedPlatform()
            .pipe(
                map((v: Platform) => {
                    return btns.map(b => {
                        return <button onClick={(_) => this._selectPlatform(b.platform)} key={b.platform} className={this._getButtonClassName(v, b.platform)}>{b.text}</button>;
                    });
                })
            );
    }

    private _getButtonClassName(sp: Platform, p: Platform): string {
        return "btn " + (p === sp ? "btn-primary" : "btn-secondary");
    }

    private _getPlatformButtons(): IPlatformButton[] {
        return [
            {
                text: 'Mac OS',
                platform: Platform.Mac
            },
            {
                text: 'Linux',
                platform: Platform.Linux
            },
            {
                text: 'Windows',
                platform: Platform.Windows
            }
        ];
    }

    private _selectPlatform(p: Platform): void {
        AppService.selectPlatform(p);
    }
}
