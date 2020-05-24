import * as React from 'react';

export interface MenuItemProps {
  selected: boolean;
  text: string;
  onclick: () => void;
}

export class MenuItem extends React.Component<MenuItemProps> {
  baseClass = 'menu-item text-secondary pointer text-center mt-1 ';

  mainClass = (): string =>
    this.props.selected ? `${this.baseClass}text-info` : this.baseClass;

  render(): JSX.Element {
    return <h3 className={this.mainClass()}>{this.props.text}</h3>;
  }
}
