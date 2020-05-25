import React from 'react';
import cls from 'classnames';

import styles from './index.module.scss';

export default function Button({ className, type, children, ...props }) {
  return (
    <button
      {...props}
      type={type || 'button'}
      className={cls(styles.button, className)}
    >
      {children}
    </button>
  );
}
