import * as React from 'react';
import {Component} from 'react';
import { Slide } from './slide/Slide';
import { Content } from './content/Content';
import { ITexts } from './texts/texts.interface';
import { texts } from './texts/texts';
import { AppService } from './services/app.service';

export class AppComponent extends Component {
    private _texts: ITexts;
    constructor(props) {
        super(props);
        this._texts = texts;
    }

    render() {
        return (
            <div className="content h-100">
            <Slide children={<Content {...{texts: this._texts}} />} />
            </div>
        );
    }
}
