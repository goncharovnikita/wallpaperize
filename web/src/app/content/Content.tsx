import * as React from 'react';
import './Content.sass';
import { ITexts } from '../texts/texts.interface';

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
                    <p>It uses utility classNamees for typography and spacing to space content out within the larger container.</p>
                    <a className="btn btn-primary btn-lg" href="#" role="button">Learn more</a>
                </div>
                {/* <div className="overlay position-absolute h-100 w-100"></div> */}
            </div>
        );
    }
}
