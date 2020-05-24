import { createSelector } from 'reselect';

export const selectDailyImages = createSelector(
  state => state.pictures.daily,
  state => state.pictures.dailyCached,
  (daily, dailyCached) => Array.from(new Set([...daily, ...dailyCached]))
);

export const selectRandomImages = createSelector(
  state => state.pictures.random,
  random => random
);
