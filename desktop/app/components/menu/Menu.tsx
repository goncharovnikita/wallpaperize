import * as React from 'react';
import './Menu.css';
import { MenuButton } from './menu-button/MenuButton';
import { MenuContentRedux } from './menu-content/MenuContentRedux';

interface MenuState {
  opened: boolean;
}

export class Menu extends React.Component<{}, MenuState> {
  constructor(props: any) {
    super(props);
    this.state = { opened: false };
  }

  menuClass = (): string =>
    this.state.opened ? 'menu-sidebar opened' : 'menu-sidebar';

  contentClass = (): string => (this.state.opened ? '' : 'hidden');

  triggerMenu = (): void =>
    this.setState(({ opened }) => ({ opened: !opened }));

  menuContent = (): JSX.Element => <MenuContentRedux />;

  render(): JSX.Element {
    return (
      <div className="menu d-flex">
        <div className={this.menuClass()}>
          <MenuButton onclick={this.triggerMenu} opened={this.state.opened} />
        </div>
        <div className={this.contentClass()}>{this.menuContent()}</div>
      </div>
    );
  }
}
