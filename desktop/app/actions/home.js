import { createAction } from 'redux-actions';

const ca = (n, p) => createAction(`home/${n}`, p);

export const setSelectedImage = ca(
  'SET_SELECTED_IMAGE',
  (selectedImageType, selectedImage) => ({
    selectedImage,
    selectedImageType
  })
);
