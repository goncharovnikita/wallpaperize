import { connect } from 'react-redux';
import { AppState } from '@app/state/app-state';
import { Main } from '@app/main/Main';

const mapStateToProps = (state: AppState) => {
    return {
        path: state.path
    };
};

export const MainRedux = connect(mapStateToProps)(Main);
