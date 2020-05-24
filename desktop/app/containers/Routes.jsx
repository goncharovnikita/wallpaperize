import React from 'react';
import { Switch, Route, Redirect } from 'react-router-dom';
import { useDispatch } from 'react-redux';

import { requestInitApp } from '../actions/app';

import Home from './Home';

export default function Routes() {
  const dispatch = useDispatch();

  React.useEffect(() => {
    dispatch(requestInitApp());
  }, []);

  return (
    <Switch>
      <Route path="/" component={Home} />
      <Redirect to="/" />
    </Switch>
  );
}
