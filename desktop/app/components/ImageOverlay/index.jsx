import React from 'react';
import cls from 'classnames';

import styles from './index.module.scss';

export default function ImageOverlay({ visible, onClose, src }) {
  const imageRef = React.useRef();

  const handleOverlayClick = React.useCallback(() => {
    onClose();
  }, [onClose]);

  return (
    <div className={cls(styles.container, { [styles.invisible]: !visible })}>
      <div className={styles.imageContainer}>
        {src ? (
          <img ref={imageRef} src={src} alt="Global" className={styles.image} />
        ) : null}
      </div>
      <div
        onClick={handleOverlayClick}
        role="overlay"
        className={styles.overlay}
      />
    </div>
  );
}
