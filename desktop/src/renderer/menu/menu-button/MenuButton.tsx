import * as React from 'react';

export interface MenuButtonProps {
    onclick: () => void;
    opened: boolean;
}

export class MenuButton extends React.Component<MenuButtonProps> {
    private baseClass = "menu-but d-flex mt-2 align-items-center ";
    mainClass = (): string => this.baseClass + 'justify-content-around';

    render(): JSX.Element {
        return (
            <div className={this.mainClass()}>
                <i onClick={this.props.onclick} className="fas text-light pointer fa-bars"></i>
            </div>
        );
    }
}
