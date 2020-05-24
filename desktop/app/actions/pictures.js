import { createAction } from 'redux-actions';

const ca = (n, p) => createAction(`pictures/${n}`, p);

export const setImages = ca('SET_IMAGES', (type, ids) => ({ type, ids }));
