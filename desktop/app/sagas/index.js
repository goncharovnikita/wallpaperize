import { all, takeEvery, call, put, select } from 'redux-saga/effects';

import * as wp from '../wallpaperize-proxy';
import * as picturesActions from '../actions/pictures';
import * as appActions from '../actions/app';

function* watchRequestInitApp() {
  yield takeEvery(appActions.requestInitApp, function*() {
    const {
      app_version: binVersion,
      build,
      daily_images: dailyCachedImages,
      random_images: randomCachedImages
    } = yield call(wp.getInfo);

    yield put(appActions.setAppInfo({ binVersion, build }));
    yield put(picturesActions.setImages('dailyCached', dailyCachedImages));
    yield put(picturesActions.setImages('randomCached', randomCachedImages));
  });
}

function* watchRequestImages() {
  yield takeEvery(picturesActions.requestImages, function*({
    payload: { type }
  }) {
    if (type === 'random') {
      const randomImages = yield call(wp.getRandom);

      yield put(picturesActions.setImages('random', randomImages));
    }
  });
}

export default function*() {
  yield all([watchRequestInitApp(), watchRequestImages()]);
}
