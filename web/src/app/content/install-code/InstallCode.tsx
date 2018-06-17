import * as React from 'react';
import { AppService } from '../../services/app.service';
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
                        case Platform.Windows:
                        this.setState({link: links.windows});
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
}
