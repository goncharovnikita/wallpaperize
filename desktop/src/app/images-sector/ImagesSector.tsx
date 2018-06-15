import * as React from 'react';
import { Image } from '@app/app/images-sector/image/Image';

export interface ImagesSectorProps {
    title: string;
    images: string[];
    loadHandler: () => Promise<void>;
}

export class ImagesSector extends React.Component {
    private _btnDisabled: boolean;
    state: ImagesSectorProps;
    constructor(props: ImagesSectorProps) {
        super(props);
        this.state = {
            images: props.images,
            title: props.title,
            loadHandler: () => new Promise(r => r())
        };
        this._btnDisabled = false;
    }

    shouldComponentUpdate(nextProps: ImagesSectorProps, _: any): boolean {
        const shouldUpdate = nextProps.images !== this.state.images;
        if (shouldUpdate) {
            this.setState(nextProps);
        }
        return shouldUpdate;
    }

    getImages(): JSX.Element[] {
        return this.state.images.map((img, index) => {
            return <Image key={index} {...{image: img}} />;
        });
    }

    async onBtnClick(): Promise<void> {
        this._btnDisabled = true;
        await this.state.loadHandler();
        this._btnDisabled = false;
    }

    render() {
        return (
            <div>
                <h3 className="display-3">
                {this.state.title}
                    <button onClick={(_) => this.onBtnClick()}
                        disabled={this._btnDisabled} className="ml-2 btn btn-secondary">Load more</button>
                </h3>
                <hr/>
                <div className="row">
                    {this.getImages()}
                </div>
            </div>
        );
    }
}
