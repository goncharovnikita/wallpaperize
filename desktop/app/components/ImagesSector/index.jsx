import * as React from 'react';

import styles from './index.module.scss';

export default function ImagesSector({
  title,
  afterTitle,
  getSrc,
  images,
  onImageClick
}) {
  const onImgClick = React.useCallback(
    img => {
      onImageClick(img);
    },
    [getSrc]
  );

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        {title}
        {afterTitle}
      </div>
      <div className={styles.grid}>
        {images.map(img => {
          return (
            <div key={img} className={styles.column}>
              <img
                className={styles.image}
                onClick={() => onImgClick(img)}
                src={getSrc(img)}
              />
            </div>
          );
        })}
      </div>
    </div>
  );
}
