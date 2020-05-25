import React from 'react';
import cls from 'classnames';

import styles from './index.module.scss';
import Button from '../Button';

const Status = ({ children }) => (
  <small className={styles.status}>{children}</small>
);

function ImageStatus({ status }) {
  if (status === 'installing') {
    return <Status>Image installing...</Status>;
  }

  return null;
}

export default function ImageOverlay({
  visible,
  status,
  src,
  cached,
  setWallpaper,
  onClose
}) {
  const imageRef = React.useRef();

  const handleOverlayClick = React.useCallback(() => {
    onClose();
  }, [onClose]);

  const handleSetWallpaper = React.useCallback(() => {
    setWallpaper(src);
  }, [src, setWallpaper]);

  return (
    <div className={cls(styles.container, { [styles.invisible]: !visible })}>
      <div className={styles.imageContainer}>
        {src ? (
          <img ref={imageRef} src={src} alt="Global" className={styles.image} />
        ) : null}
      </div>
      <div className={styles.toolbar}>
        <Button
          type="button"
          onClick={handleSetWallpaper}
          className={styles.button}
          disabled={status !== 'idle'}
        >
          Set as wallpaper
        </Button>
        <ImageStatus status={status} />
        {cached ? <Status>Image cached</Status> : null}
      </div>
      <div
        onClick={handleOverlayClick}
        role="overlay"
        className={styles.overlay}
      />
    </div>
  );
}
