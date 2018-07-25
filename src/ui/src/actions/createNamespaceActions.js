import { getNamespaces } from './kubernetes-actions';

export const NAMESPACE_MODAL_SHOW = 'namespace-modal:show';
export const NAMESPACE_MODAL_HIDE = 'namespace-modal:hide';
export const NAMESPACE_ADD_ERROR = 'namespace-modal:error';
export const CREATE_NAMESPACE = 'namespace:create';

const showNamespaceModal = () => ({
  type: NAMESPACE_MODAL_SHOW,
  payload: {
    show: true,
    errorMessage: '',
  },
});

const hideNamespaceModal = () => ({
  type: NAMESPACE_MODAL_HIDE,
  payload: {
    show: false,
    errorMessage: '',
  },
});

const onAddError = (message) => ({
  type: NAMESPACE_ADD_ERROR,
  payload: {
    show: true,
    errorMessage: message,
  }
});

const createNamespace = (name) => (dispatch, getState, api) => {
  api
    .postData('/namespace?name=' + name.toLowerCase())
    .then(response => {
      if (!response.ok) {
        return response.text()
      }
      return ''
    })
    .then(text => {
      //check for error
      if (text !== '') {
        throw new Error(text);
      }
    })
    .then(() => {
      dispatch(hideNamespaceModal());
      dispatch(getNamespaces());
    })
    .catch(e => {
      dispatch(onAddError(e.message));
    });
};

export { createNamespace, showNamespaceModal, hideNamespaceModal };


