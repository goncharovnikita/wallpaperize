import * as React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import * as ospath from 'path';

import styles from './index.module.scss';

import ImagesSector from '../../components/ImagesSector';
import {
  getInfo,
  loadRandom,
  getRandom,
  getSelected
} from '../../wallpaperize-proxy';

import { setImages } from '../../actions/pictures';

const randomPath = 'https://images.wallpaperize.goncharovnikita.com';

const useInitImages = () => {
  const dispatch = useDispatch();

  React.useEffect(() => {
    (async () => {
      const {
        random_images: localRandomImages,
        daily_images: localDailyImages
      } = await getInfo();
      const randomImages = await getRandom();
      const randomCached = localRandomImages.map(i => i.split('/').pop());
      const selected = await getSelected();

      dispatch(setImages('random', randomImages));
      dispatch(setImages('daily', localDailyImages));
    })();
  }, [dispatch]);
};

export default function App() {
  const dispatch = useDispatch();
  const dailyImages = useSelector(state => state.pictures.daily);
  const randomImages = useSelector(state => state.pictures.random);

  useInitImages();

  const getSrc = React.useCallback(path => {
    return `file://${ospath.resolve(path)}`;
  }, []);

  const handleRefresh = React.useCallback(() => {
    (async () => {
      const newRandomImages = await getRandom();
      dispatch(setImages('random', newRandomImages));
    })();
  }, []);

  const getRemoteSrc = React.useCallback(p => {
    return `${randomPath}/random_images_min/${p.replace('.jpg', '-min.jpg')}`;
  }, []);

  const afterTitle = React.useMemo(
    () => (
      <div className="lead d-flex align-items-end">
        <button
          onClick={handleRefresh}
          type="button"
          className="text-muted pointer"
        >
          refresh
        </button>
      </div>
    ),
    []
  );

  const loadRandomHandler = React.useCallback(() => {
    (async () => {
      await loadRandom();
    })();
  }, []);

  return (
    <div className={styles.container}>
      <ImagesSector
        title="Daily images"
        selected=""
        images={dailyImages}
        cachedImages={[]}
        loadHandler={loadRandomHandler}
        getSrc={getSrc}
      />
      <ImagesSector
        title="Random images"
        selected=""
        images={randomImages}
        cachedImages={[]}
        loadHandler={loadRandomHandler}
        afterTitle={afterTitle}
        getSrc={getRemoteSrc}
      />
    </div>
  );
}
