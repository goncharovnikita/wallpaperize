import { Dispatch, connect } from "react-redux";
import { InitApp } from "@app/init/Init";
import { Actions } from "@app/reducers/actions";
import { AppPath } from "@app/state/app-state";

const mapDispatchToProps = (disp: Dispatch) => {
    return {
        initialize: () => {
            disp({type: Actions.NAVIGATE, value: AppPath.Main});
        }
    };
};

export const InitRedux = connect(null, mapDispatchToProps)(InitApp);
