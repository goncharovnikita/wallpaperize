import { AppState, AppPath } from "@app/state/app-state";
import { connect, Dispatch } from "react-redux";
import { MenuContent } from "@app/menu/menu-content/MenuContent";
import { Actions } from "@app/reducers/actions";

const mapStateToProps = (state: AppState) => {
    return {
        path: state.path
    };
};

const mapDispatchToProps = (disp: Dispatch) => {
    return {
        selectItem: (p: AppPath): void => {
            switch (p) {
                case AppPath.Saved:
                disp({type: Actions.NAVIGATE, value: AppPath.Saved});
                break;
                default:
                disp({type: Actions.NAVIGATE, value: AppPath.Main});
                break;
            }
        }
    };
};

export const MenuContentRedux = connect(mapStateToProps, mapDispatchToProps)(MenuContent);
