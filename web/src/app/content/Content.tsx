import * as React from 'react';
import './Content.sass';
import { ITexts } from '../texts/texts.interface';
import { SelectPlatform } from './select-platform/SelectPlatform';
import { InstallCode } from './install-code/InstallCode';

export class Content extends React.Component {
    constructor(public props: {texts: ITexts}) {
        super(props);
    }

    content(): JSX.Element {
        return (
            <div className="jumbotron">
                <h1 className="display-1 text-center">Wallpaperize</h1>
                <h1 className="font-weight-light text-center">{this.props.texts.about}</h1>
                <hr className="my-4" />
                <div className="row">
                    <div className="col-md-1"></div>
                    <div className="col-md-10">
                        <SelectPlatform />
                        <div className="alert alert-dark">
                            <div className="code">
                                <InstallCode />
                            </div>
                        </div>
                    </div>
                    <div className="col-md-1"></div>
                </div>
            </div>
        );
    }

    render() {
        return(
            <div className="position-relative">
                <div className="row">
                <div className="col-md-1 col-sm-0 col-lg-2"></div>
                <div className="col-md-10 col-sm-12 col-lg-8">
                    {this.content()}
                </div>
                <div className="col-md-1 col-sm-0 col-lg-2"></div>
                </div>
            </div>
        );
    }
}
