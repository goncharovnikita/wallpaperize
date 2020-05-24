import * as React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import * as ospath from 'path';

import ImagesSector from '../components/ImagesSector';
import ImageOverlay from '../components/ImageOverlay';
import { loadRandom } from '../wallpaperize-proxy';

import { requestImages } from '../actions/pictures';
import { setSelectedImage } from '../actions/home';
import { selectDailyImages, selectRandomImages } from '../selectors/pictures';

const randomPath = 'https://images.wallpaperize.goncharovnikita.com';

const useInitImages = () => {
  const dispatch = useDispatch();

  React.useEffect(() => {
    dispatch(requestImages('random'));
    dispatch(requestImages('daily'));
  }, [dispatch]);
};

const getSrc = path => {
  return `file://${ospath.resolve(path)}`;
};

const getRemoteSrc = p => {
  return `${randomPath}/random_images_min/${p.replace('.jpg', '-min.jpg')}`;
};

export default function Home() {
  const dispatch = useDispatch();
  const dailyImages = useSelector(selectDailyImages);
  const randomImages = useSelector(selectRandomImages);
  const selectedImage = useSelector(state => state.home.selectedImage);
  const selectedImageType = useSelector(state => state.home.selectedImageType);

  useInitImages();

  const handleRefresh = React.useCallback(() => {
    dispatch(requestImages('random'));
  }, []);

  const handleDailyImageClick = React.useCallback(
    id => {
      dispatch(setSelectedImage('daily', id));
    },
    [dispatch]
  );

  const handleRandomImageClick = React.useCallback(
    id => {
      dispatch(setSelectedImage('random', id));
    },
    [dispatch]
  );

  const selectedImageUrl = React.useMemo(() => {
    if (!selectedImage) return;

    const fn = selectedImageType === 'daily' ? getSrc : getRemoteSrc;

    return fn(selectedImage);
  }, [selectedImage, selectedImageType]);

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

  const handleCloseSelectedImage = React.useCallback(() => {
    dispatch(setSelectedImage(null, null));
  }, [dispatch]);

  return (
    <div>
      <ImagesSector
        title="Daily images"
        selected=""
        images={dailyImages}
        cachedImages={[]}
        loadHandler={loadRandomHandler}
        getSrc={getSrc}
        onImageClick={handleDailyImageClick}
      />
      <ImagesSector
        title="Random images"
        selected=""
        images={randomImages}
        cachedImages={[]}
        loadHandler={loadRandomHandler}
        afterTitle={afterTitle}
        getSrc={getRemoteSrc}
        onImageClick={handleRandomImageClick}
      />
      <ImageOverlay
        visible={Boolean(selectedImageUrl)}
        src={selectedImageUrl}
        onClose={handleCloseSelectedImage}
      />
    </div>
  );
}
