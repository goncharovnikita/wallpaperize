import * as React from 'react';
import './MenuContent.sass';
import { MenuItemProps, MenuItem } from '@app/menu/menu-content/MenuItem';
import { AppPath } from '@app/state/app-state';

interface State {
    items: MenuItemProps[];
}

interface Props {
    path: AppPath;
    selectItem: (p: AppPath) => void;
}

export class MenuContent extends React.Component<Props, State> {
    constructor(props: any) {
        super(props);
        this.state = this.initState();
    }

    items = (): JSX.Element[] => {
        return this.state.items.map(item => {
            return <MenuItem key={item.text} {...item} />;
        });
    }

    render() {
        return (
            <div className="menu-content ml-2">
            {this.items()}
            </div>
        );
    }

    private initState(): State {
        return {
            items: [
                {
                    text: 'Main',
                    selected: this.props.path === AppPath.Main,
                    onclick: () => this.props.selectItem(AppPath.Main)
                },
                {
                    text: 'Saved images',
                    selected: this.props.path === AppPath.Saved,
                    onclick: () => this.props.selectItem(AppPath.Saved)
                }
            ]
        };
    }
}
