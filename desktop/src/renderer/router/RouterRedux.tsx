import { connect } from 'react-redux';
import { AppState } from '@app/state/app-state';
import { Router } from '@app/router/Router';

const mapStateToProps = (state: AppState) => {
    return {
        path: state.path
    };
};

export const RouterRedux = connect(mapStateToProps)(Router);
