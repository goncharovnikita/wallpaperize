import { AppState, AppPath } from '@app/state/app-state';
import { Actions } from '@app/reducers/actions';

export const appReducer = (
  state: AppState = initState,
  action: { type: Actions; value?: any }
): AppState => {
  switch (action.type) {
    case Actions.NAVIGATE:
      return {
        ...state,
        path: action.value
      };
    default:
      return state;
  }
};

const initState: AppState = {
  path: AppPath.Init
};
