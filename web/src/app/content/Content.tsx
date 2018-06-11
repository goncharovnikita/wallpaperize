import * as React from 'react';
import './Content.sass';
import { ITexts } from '../texts/texts.interface';
import { SelectPlatform } from './select-platform/SelectPlatform';
import { InstallCode } from './install-code/InstallCode';

export class Content extends React.Component {
    constructor(public props: {texts: ITexts}) {
        super(props);
    }

    render() {
        return(
            <div className="position-relative">
                <div className="jumbotron h-50">
                    <h1 className="display-1 text-center">Wallpaperize</h1>
                    <h1 className="font-weight-light text-center">{this.props.texts.about}</h1>
                    <hr className="my-4" />
                    <div className="row">
                        <div className="col-md-2"></div>
                        <div className="col-md-8">
                            <SelectPlatform />
                            <div className="alert alert-dark">
                                <div className="code">
                                    <InstallCode />
                                </div>
                            </div>
                        </div>
                        <div className="col-md-2"></div>
                    </div>
                </div>
                {/* <div className="overlay position-absolute h-100 w-100"></div> */}
            </div>
        );
    }
}
