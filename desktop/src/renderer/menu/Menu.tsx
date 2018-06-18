import * as React from 'react';
import './Menu.scss';
import { MenuButton } from '@app/menu/menu-button/MenuButton';

interface MenuState {
    opened: boolean;
}

export class Menu extends React.Component<{}, MenuState> {
    constructor(props: any) {
        super(props);
        this.state = {opened: false};
    }

    menuClass = (): string => this.state.opened ? "menu opened" : "menu";

    triggerMenu = (): void => this.setState({opened: !this.state.opened});

    render(): JSX.Element {
        return (
            <div className={this.menuClass()}>
                <MenuButton onclick={this.triggerMenu} opened={this.state.opened}/>
            </div>
        );
    }
}
