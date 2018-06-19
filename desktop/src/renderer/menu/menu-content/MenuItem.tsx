import * as React from 'react';

export interface MenuItemProps {
    selected: boolean;
    text: string;
    onclick: () => void;
}

export class MenuItem extends React.Component<MenuItemProps> {
    baseClass = 'menu-item text-secondary pointer text-center mt-1 ';
    constructor(props: MenuItemProps) {
        super(props);
    }

    mainClass = (): string => this.props.selected ? this.baseClass + 'text-info'
        : this.baseClass

    render(): JSX.Element {
        return (
            <h3 onClick={this.props.onclick} className={this.mainClass()}>{this.props.text}</h3>
        );
    }
}
